# terraform/outputs.tf

output "vpc_id" {
  description = "ID de la VPC creada."
  value       = aws_vpc.practicelab_vpc.id # Referencia actualizada
}

output "public_subnet_id" {
  description = "ID de la subred pública."
  value       = aws_subnet.practicelab_public_subnet.id # Referencia actualizada
}

output "private_subnet_id" {
  description = "ID de la subred privada."
  value       = aws_subnet.practicelab_private_subnet.id # Referencia actualizada
}

output "rds_endpoint" {
  description = "Endpoint de la instancia RDS."
  value       = aws_db_instance.db_instance_rds.address
}

output "rds_security_group_id" {
  description = "ID del Security Group de RDS."
  value       = aws_security_group.rds_sg.id
}

output "db_subnet_group_name" {
  description = "Nombre del DB Subnet Group."
  value       = aws_db_subnet_group.practicelab_rds_subnet_group.name # Referencia actualizada
}

# --- Salidas para la instancia EC2 ---
output "ec2_public_ip" {
  description = "Dirección IP pública de la instancia EC2 de la API."
  value       = aws_instance.go_api_instance.public_ip
}

output "ec2_public_dns" {
  description = "Nombre DNS público de la instancia EC2 de la API."
  value       = aws_instance.go_api_instance.public_dns
}

output "ec2_instance_id" {
  description = "ID de la instancia EC2 de la API."
  value       = aws_instance.go_api_instance.id
}

output "cognito_user_pool_id" {
  description = "ID del User Pool de Cognito."
  value       = aws_cognito_user_pool.go_user_pool.id
}

output "cognito_client_id" {
  description = "ID del Cliente de la aplicación Cognito."
  value       = aws_cognito_user_pool_client.go_app_client.id
}
