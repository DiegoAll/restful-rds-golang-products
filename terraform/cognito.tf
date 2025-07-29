# terraform/cognito.tf

# Si decides añadir un User Pool de Cognito y su cliente, el código iría aquí.
# Ejemplo básico (sin implementar completamente):
/*
resource "aws_cognito_user_pool" "my_user_pool" {
  name = "my-app-user-pool"

  # ... otras configuraciones del User Pool
}

resource "aws_cognito_user_pool_client" "my_app_client" {
  name         = "my-app-client"
  user_pool_id = aws_cognito_user_pool.my_user_pool.id

  # ... otras configuraciones del cliente
}
*/

# terraform/cognito.tf

# --- User Pool ---
resource "aws_cognito_user_pool" "go_user_pool" {
  name = "MyNewGoUserPool"

  password_policy {
    minimum_length    = 8
    require_lowercase = true
    require_numbers   = true
    require_symbols   = true
    require_uppercase = true
  }

  email_configuration {
    email_sending_account = "COGNITO_DEFAULT"
  }

  username_attributes       = ["email"]
  auto_verified_attributes  = ["email"]
  mfa_configuration         = "OFF"

  tags = {
    Name = "MyNewGoUserPool"
  }
}

# --- App Client ---
resource "aws_cognito_user_pool_client" "go_app_client" {
  name         = "MyNewGoAppClient"
  user_pool_id = aws_cognito_user_pool.go_user_pool.id

  generate_secret = false

  # Flujos de autenticación habilitados (coincide con backend Go)
  explicit_auth_flows = [
    "ADMIN_NO_SRP_AUTH",    # Para autenticación admin sin SRP
    "USER_PASSWORD_AUTH"    # Para autenticación directa con user/pass
  ]

  # Atributos accesibles por el cliente
  read_attributes  = ["email", "profile"]
  write_attributes = ["email", "profile"]

  # Puedes descomentar estos si deseas personalizar duración de tokens
  # refresh_token_validity = 30
  # access_token_validity  = 1
  # id_token_validity      = 1
}


