# terraform/rds.tf

# Consulta la clave KMS por su alias para obtener su ARN
data "aws_kms_alias" "rds_kms_alias" {
  name = var.kms_key_alias # Usará "alias/aws/rds" por defecto
}

# 6. Crear el Security Group para RDS
resource "aws_security_group" "rds_sg" {
  name        = "practicelab-rds-sg" # Nombre del Security Group actualizado
  description = "Acceso RDS desde IP local y desde SG de API"
  vpc_id      = aws_vpc.practicelab_vpc.id # Usa la VPC creada en vpc.tf (CORREGIDO)

  # Regla de entrada para permitir el acceso desde tu IP pública
  ingress {
    from_port   = var.rds_port
    to_port     = var.rds_port
    protocol    = "tcp"
    cidr_blocks = [var.my_public_ip] # Tu IP pública
    description = "Allow PostgreSQL from my IP"
  }

  # Regla de entrada para permitir el acceso desde el Security Group de la API
  ingress {
    from_port       = var.rds_port
    to_port         = var.rds_port
    protocol        = "tcp"
    security_groups = [aws_security_group.api_sg.id] # Permite acceso desde el SG de la API (definido en ec2.tf)
    description     = "Allow PostgreSQL from API Security Group"
  }

  # Regla de salida para permitir todo el tráfico saliente (RDS necesita esto)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "practicelab-rds-sg" # Tag actualizado
  }
}

# 7. Crear un DB Subnet Group
# Incluye ambas subredes (pública y privada) para cumplir con el requisito de AZs.
resource "aws_db_subnet_group" "practicelab_rds_subnet_group" { # Nombre del recurso interno actualizado (CORREGIDO)
  name        = "practicelab-rds-subnet-group" # Nombre del DB Subnet Group actualizado
  description = "Subred para RDS en AZ us-east-1a y us-east-1b"
  subnet_ids  = [aws_subnet.practicelab_public_subnet.id, aws_subnet.practicelab_private_subnet.id] # Referencia a las subredes en vpc.tf (CORREGIDO)

  tags = {
    Name = "practicelab" # Tag actualizado
  }
}

# 8. Crear la instancia RDS PostgreSQL
resource "aws_db_instance" "db_instance_rds" {
  identifier                = var.db_instance_identifier
  engine                    = "postgres"
  engine_version            = "17.4"
  instance_class            = "db.t4g.micro"
  allocated_storage         = 20
  storage_type              = "gp2"
  storage_encrypted         = true
  # ¡CORRECCIÓN AQUÍ! Usamos el ARN de la clave KMS obtenido de la fuente de datos
  kms_key_id                = data.aws_kms_alias.rds_kms_alias.target_key_arn
  username                  = var.db_master_username
  password                  = var.db_master_password
  db_name                   = var.db_name
  vpc_security_group_ids    = [aws_security_group.rds_sg.id] # Referencia al SG de RDS
  db_subnet_group_name      = aws_db_subnet_group.practicelab_rds_subnet_group.name # Referencia al DB Subnet Group (nombre actualizado)
  availability_zone         = var.availability_zone_public # Aunque el subnet group tenga 2 AZs, la instancia se lanza en una específica.
  publicly_accessible       = true # Como tu lo definiste en el CLI
  backup_retention_period   = 1
  port                      = var.rds_port
  monitoring_interval       = 0 # Deshabilitado el monitoreo mejorado para costos
  enabled_cloudwatch_logs_exports = ["postgresql"]
  multi_az                  = false # Single AZ como en tu CLI
  skip_final_snapshot = true

  # Puedes añadir esto si quieres deshabilitar Performance Insights para ahorrar costos
  # performance_insights_enabled = false
  # performance_insights_retention_period = 0 # O un valor bajo si lo habilitas

  tags = {
    Name = "practicelab" # Tag actualizado
  }

  # Evita que Terraform borre la base de datos si el recurso es destruido
  # Considera cuidadosamente si quieres esto en producción.
  # deletion_protection = true
}