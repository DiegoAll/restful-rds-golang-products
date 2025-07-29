# terraform/ec2.tf

# Security Group para la API (si tu API se ejecuta en una EC2)
# Permite tráfico saliente a cualquier lugar y tráfico entrante desde tu IP para SSH, y tráfico de API.
resource "aws_security_group" "api_sg" {
  name        = "go-api-ec2-sg"
  description = "Allow SSH and HTTP/API traffic to Go API, allow outbound to RDS"
  vpc_id      = aws_vpc.practicelab_vpc.id # Referencia a la VPC creada en vpc.tf

  # Regla de entrada para SSH desde tu IP
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.my_public_ip]
    description = "Allow SSH from specific IP"
  }

  # Regla de entrada para el puerto de la API (desde cualquier origen, para pruebas)
  # Para producción, considera restringirlo a un Load Balancer o API Gateway.
  ingress {
    from_port   = var.api_port
    to_port     = var.api_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Permite acceso público a tu API. ¡Revisar para producción!
    description = "Allow API traffic on port ${var.api_port}"
  }

  # Regla de salida general (para actualizaciones, Git, etc.)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1" # Todos los protocolos
    cidr_blocks = ["0.0.0.0/0"] # Permitir todo el tráfico saliente a Internet
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "GoAPISecurityGroup"
  }
}

# --- RECURSOS IAM PARA LA INSTANCIA EC2 (Acceso a Secrets Manager) ---

# Consulta el ID de la cuenta de AWS (necesario para construir el ARN del secreto)
data "aws_caller_identity" "current" {}

# Define la política IAM para permitir que la EC2 lea secretos de Secrets Manager
resource "aws_iam_policy" "secrets_manager_read_policy" {
  name        = "GoAPISecurityManagerReadPolicy"
  description = "Allows EC2 instance to read specific secrets from AWS Secrets Manager."

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret" # DescribeSecret es útil para depuración
        ],
        Resource = "arn:aws:secretsmanager:${var.aws_region}:${data.aws_caller_identity.current.account_id}:secret:${var.rds_secret_name}*"
        # El comodín * es importante si Secrets Manager añade un sufijo aleatorio al ARN del secreto.
      },
      {
        Effect = "Allow",
        Action = [
          "kms:Decrypt" # Necesario si el secreto está cifrado con una clave KMS (que no sea la default)
        ],
        Resource = "*" # Idealmente, esto debería ser el ARN de la clave KMS si usas una CMK específica
                       # Si usas la clave default de Secrets Manager, esta acción a menudo es implícita
                       # pero es buena práctica incluirla si hay dudas.
      }
    ]
  })
}

# Define el rol IAM que la instancia EC2 asumirá
resource "aws_iam_role" "ec2_api_role" {
  name               = "GoAPIEc2Role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ec2.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })

  tags = {
    Name = "GoAPIEc2Role"
  }
}

# Adjunta la política al rol IAM
resource "aws_iam_role_policy_attachment" "ec2_api_policy_attach" {
  role       = aws_iam_role.ec2_api_role.name
  policy_arn = aws_iam_policy.secrets_manager_read_policy.arn
}

# Crea un perfil de instancia IAM para adjuntar el rol a la EC2
resource "aws_iam_instance_profile" "ec2_api_profile" {
  name = "GoAPIEc2Profile"
  role = aws_iam_role.ec2_api_role.name
}

# Define la instancia EC2
resource "aws_instance" "go_api_instance" {
  # IMPORTANTE: Esta AMI es de Ubuntu 22.04 LTS para arquitectura ARM64 (Graviton) en us-east-1.
  # Siempre verifica la AMI más reciente y adecuada en la consola de AWS.
  # Ejemplo: ami-053b0d53c279e2c60 (Ubuntu Server 22.04 LTS (HVM), SSD Volume Type - arm64)
  ami           = "ami-020cba7c55df1f615" # CAMBIO: AMI para ARM64 (Graviton)
  instance_type = var.instance_type # Usará t4g.micro por defecto (ARM64)
  key_name      = var.key_pair_name
  vpc_security_group_ids = [aws_security_group.api_sg.id]
  subnet_id     = aws_subnet.practicelab_public_subnet.id

  # Asignar una IP pública para poder accederla (si la subred es pública)
  associate_public_ip_address = true

  # Adjunta el perfil de instancia IAM a la EC2
  iam_instance_profile = aws_iam_instance_profile.ec2_api_profile.name

  # --- User Data para instalar Go y configurar la aplicación ---
  user_data = <<-EOF
              #!/bin/bash
              set -eux # Salir inmediatamente si un comando falla, y mostrar los comandos

              # Actualizar el sistema e instalar herramientas básicas (ej. git, wget)
              apt update -y
              apt install -y git wget

              # Instalar Go
              GO_VERSION="${var.go_version}"
              GO_ARCH="${var.go_arch}"
              GO_FILENAME="go$${GO_VERSION}.linux-$${GO_ARCH}.tar.gz"
              GO_URL="https://golang.org/dl/$${GO_FILENAME}"

              wget "$GO_URL"
              tar -C /usr/local -xzf "$GO_FILENAME"
              rm "$GO_FILENAME"

              # Configurar el PATH para Go para todos los usuarios y para esta sesión
              echo 'export PATH=$PATH:/usr/local/go/bin' | tee /etc/profile.d/go.sh
              export PATH=$PATH:/usr/local/go/bin # Asegura que esté disponible en este script inmediatamente

              # Crear un directorio para la aplicación
              mkdir -p /opt/app/go_api

              # Clonar el repositorio de Git
              git clone ${var.github_repo_url} /opt/app/go_api/

              # Cambiar la propiedad de los archivos clonados al usuario ubuntu
              chown -R ubuntu:ubuntu /opt/app/go_api/

              # Navegar al directorio de la aplicación (asumiendo que main.go está en cmd/api)
              cd /opt/app/go_api/cmd/api

              # Configurar el entorno para el usuario 'ubuntu' antes de ejecutar Go
              # y ejecutar 'go mod tidy' y 'go build' como el usuario 'ubuntu'
              sudo -u ubuntu bash -c '
                export HOME=/home/ubuntu
                export GOPATH=$HOME/go
                export GOMODCACHE=$GOPATH/pkg/mod
                export PATH=$PATH:/usr/local/go/bin

                go mod tidy || { echo "Error en go mod tidy."; exit 1; }
                go build -o /opt/app/go_api/main . || { echo "Error en go build."; exit 1; }
              ' || { echo "Error en los comandos de Go ejecutados como usuario ubuntu."; exit 1; }

              # --- AJUSTE CLAVE AQUÍ: Configurar variables de entorno en un archivo .env para systemd ---
              # Se escribirá en el directorio de la aplicación, que es gestionado por 'ubuntu'.
              cat << 'EOT_APP_ENV' > /opt/app/go_api/.env
              APP_ENV=production
              RDS_SECRET_NAME=${var.rds_secret_name}
              UserPoolID=${aws_cognito_user_pool.go_user_pool.id}
              ClientID=${aws_cognito_user_pool_client.go_app_client.id}
              Region=${var.aws_region}
              EOT_APP_ENV

              # Asegurarse de que el archivo .env tenga permisos de lectura para el usuario del servicio
              chown ubuntu:ubuntu /opt/app/go_api/.env
              chmod 640 /opt/app/go_api/.env

              # Crear un servicio systemd para tu aplicación Go
              cat << 'EOT_SERVICE' > /etc/systemd/system/go-api.service
              [Unit]
              Description=My Go API Service
              After=network.target

              [Service]
              User=ubuntu
              WorkingDirectory=/opt/app/go_api
              # --- CAMBIO: Cargar variables de entorno desde el nuevo archivo .env ---
              EnvironmentFile=/opt/app/go_api/.env
              # -------------------------------------------------------------------
              ExecStart=/opt/app/go_api/main
              Restart=always
              StandardOutput=syslog
              StandardError=syslog
              SyslogIdentifier=my-go-api

              [Install]
              WantedBy=multi-user.target
              EOT_SERVICE

              # Habilitar y arrancar el servicio
              systemctl daemon-reload
              systemctl start go-api.service
              systemctl enable go-api.service

              echo "Go API setup complete!"
              EOF

  tags = {
    Name        = "GoApiInstance"
    Environment = "Production"
    Project     = "restful-rds-golang"
  }
}