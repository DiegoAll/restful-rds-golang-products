# terraform/variables.tf

variable "aws_region" {
  description = "La región de AWS donde se desplegarán los recursos."
  type        = string
  default     = "us-east-1"
}

variable "vpc_cidr_block" {
  description = "Bloque CIDR para la VPC."
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidr_block" {
  description = "Bloque CIDR para la subred pública."
  type        = string
  default     = "10.0.1.0/24"
}

variable "private_subnet_cidr_block" {
  description = "Bloque CIDR para la subred privada."
  type        = string
  default     = "10.0.2.0/24"
}

variable "availability_zone_public" {
  description = "Zona de disponibilidad para la subred pública."
  type        = string
  default     = "us-east-1a"
}

variable "availability_zone_private" {
  description = "Zona de disponibilidad para la subred privada."
  type        = string
  default     = "us-east-1b"
}

variable "db_instance_identifier" {
  description = "Identificador para la instancia RDS."
  type        = string
}

variable "db_master_username" {
  description = "Nombre de usuario maestro para la instancia RDS."
  type        = string
}

variable "db_master_password" {
  description = "Contraseña maestra para la instancia RDS."
  type        = string
  sensitive   = true # Marca como sensible para no mostrar en logs de Terraform
}

variable "db_name" {
  description = "Nombre de la base de datos dentro de la instancia RDS."
  type        = string
  default     = "products"
}

variable "rds_port" {
  description = "Puerto para la instancia RDS."
  type        = number
  default     = 5432
}

variable "my_public_ip" {
  description = "Tu dirección IP pública para acceso al Security Group de RDS y SSH a EC2. Formato CIDR (ej. 181.142.86.4/32)."
  type        = string
}

variable "kms_key_alias" {
  description = "Alias de la clave KMS para el cifrado de RDS. Usa 'alias/aws/rds' por defecto para la clave gestionada por AWS."
  type        = string
  default     = "alias/aws/rds"
}

# --- Variables para la instancia EC2 ---
variable "instance_type" {
  description = "Tipo de instancia EC2 para la API. Debe coincidir con la arquitectura de la AMI y el binario de Go."
  type        = string
  default     = "t3.micro" # CAMBIO: Vuelve a t4g.micro para arquitectura ARM64 (Graviton)
}

variable "key_pair_name" {
  description = "Nombre del par de claves EC2 existente para SSH."
  type        = string
}

variable "api_port" {
  description = "Puerto en el que la API Go escuchará."
  type        = number
  default     = 9090 # Puerto común para APIs
}

variable "github_repo_url" {
  description = "URL del repositorio Git de tu aplicación Go."
  type        = string
}

variable "go_version" {
  description = "Versión de Go a instalar en la instancia EC2."
  type        = string
  default     = "1.23.2" # Mantiene la versión de tu log
}

variable "go_arch" {
  description = "Arquitectura de Go a instalar (ej. arm64, amd64). Debe coincidir con la AMI y el tipo de instancia."
  type        = string
  default     = "amd64" # CAMBIO: Se mantiene arm64 para t4g.micro
}

variable "rds_secret_name" {
  description = "Nombre del secreto en Secrets Manager que contiene las credenciales de RDS."
  type        = string
  # Ejemplo: "my-rds-credentials"
}

# --- Variables para Cognito (usadas en el user_data de EC2, aunque cognito.tf sea placeholder) ---
# variable "cognito_user_pool_id" {
#   description = "ID del User Pool de Cognito."
#   type        = string
# #   default     = "YOUR_COGNITO_USER_POOL_ID" # Reemplazar con el ID real
# }

# variable "cognito_client_id" {
#   description = "ID del Cliente de la aplicación Cognito."
#   type        = string
# #   default     = "YOUR_COGNITO_CLIENT_ID" # Reemplazar con el ID real
# }
