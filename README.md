# restful-rds-golang-products


restful-rds-golang-products

consentscreen authentication only using Goog,le Credentials service.
OAuth + OIDC using OIDC fot obtain identity.


OAuth 2.0 con OpenID Connect, el cual requiere una interacción del usuario mediante el navegador web.

API de usuarios

🔒 Esto está pensado para:
Aplicaciones web.
Aplicaciones móviles.
Aplicaciones desktop con navegador embebido.

🔑 2. Service accounts (para GCP)
Permiten autenticación automática sin intervención del usuario.


## Run application

    docker-compose down -v --rmi all
    docker-compose up --build -d


## API Endpoints



##  Database Querys


    docker run -p 8081:8081 \
    -e PORT=8081 \
    -e POSTGRES_DSN="host=db-instance-rds.c8le640i0kbl.us-east-1.rds.amazonaws.com port=5432 user=p0stgr3s password=p4rc3r02025! dbname=fixture sslmode=require" \
    your-api-image-name

