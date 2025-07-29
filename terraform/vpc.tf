# terraform/vpc.tf

provider "aws" {
  region = var.aws_region
}

# 1. Crear una custom VPC (practicelab)
resource "aws_vpc" "practicelab_vpc" { # Nombre del recurso interno actualizado
  cidr_block           = var.vpc_cidr_block
  enable_dns_support   = true  # Habilitar resolución DNS
  enable_dns_hostnames = true  # Habilitar nombres DNS

  tags = {
    Name = "practicelab-vpc" # Nombre de la VPC actualizado
  }
}

# 2. Crear subred pública en la AZ us-east-1a
resource "aws_subnet" "practicelab_public_subnet" { # Nombre del recurso interno actualizado
  vpc_id                  = aws_vpc.practicelab_vpc.id # Referencia a la VPC actualizada
  cidr_block              = var.public_subnet_cidr_block
  availability_zone       = var.availability_zone_public
  map_public_ip_on_launch = true # Habilitar autoasignación de IP pública

  tags = {
    Name = "practicelab-subnet1a" # Nombre de la subred actualizado
  }
}

# 2.1. Crear subred privada en la AZ us-east-1b (para cumplir requisito de DB Subnet Group)
resource "aws_subnet" "practicelab_private_subnet" { # Nombre del recurso interno actualizado
  vpc_id            = aws_vpc.practicelab_vpc.id # Referencia a la VPC actualizada
  cidr_block        = var.private_subnet_cidr_block
  availability_zone = var.availability_zone_private

  tags = {
    Name = "practicelab-subnet1b" # Nombre de la subred actualizado
  }
}

# 3. Crear una Internet Gateway y asociarla
resource "aws_internet_gateway" "practicelab_igw" { # Nombre del recurso interno actualizado
  vpc_id = aws_vpc.practicelab_vpc.id # Referencia a la VPC actualizada

  tags = {
    Name = "practicelab-igw" # Nombre del Internet Gateway actualizado
  }
}

# 4. Crear tabla de ruteo pública
resource "aws_route_table" "practicelab_public_route_table" { # Nombre del recurso interno actualizado
  vpc_id = aws_vpc.practicelab_vpc.id # Referencia a la VPC actualizada

  tags = {
    Name = "practicelab-rt" # Nombre de la tabla de ruteo actualizado
  }
}

# 4.1. Crear la ruta a Internet en la tabla de ruteo pública
resource "aws_route" "public_internet_route" {
  route_table_id         = aws_route_table.practicelab_public_route_table.id # Referencia a la tabla de ruteo actualizada
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.practicelab_igw.id # Referencia al IGW actualizada
}

# 4.2 Asociar la route table pública a la subred pública
resource "aws_route_table_association" "public_subnet_association" {
  subnet_id      = aws_subnet.practicelab_public_subnet.id # Referencia a la subred actualizada
  route_table_id = aws_route_table.practicelab_public_route_table.id # Referencia a la tabla de ruteo actualizada
}

# Opcional: Para la subnet privada, no necesitamos una ruta a Internet.
# Utilizará la tabla de ruteo principal de la VPC por defecto,
# lo que permite la comunicación interna con otras subredes de la VPC.
# Si quisieras asegurar que use la tabla principal, podrías asociarla explícitamente:
/*
resource "aws_route_table_association" "private_subnet_association" {
  subnet_id      = aws_subnet.practicelab_private_subnet.id # Referencia a la subred actualizada
  route_table_id = aws_vpc.practicelab_vpc.default_route_table_id # Referencia a la VPC actualizada
}
*/