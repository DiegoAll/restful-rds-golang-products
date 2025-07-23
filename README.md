# restful-rds-golang-products

restful-rds-golang-products


No, no es estrictamente obligatorio usar el SDK de AWS Go para interactuar con Cognito, pero es altamente recomendable y la forma más práctica y segura de hacerlo.


Aunque técnicamente podrías construir todas las llamadas HTTP y el manejo de seguridad desde cero, la cantidad de tiempo, esfuerzo y potencial para introducir errores de seguridad es muy alta. Utilizar el SDK de AWS Go para Cognito te ahorra una enorme cantidad de trabajo, te proporciona un código más robusto y seguro, y te permite centrarte en la lógica de negocio de tu aplicación en lugar de en los detalles de la integración con la API de AWS.


db.t3.micro

## Run application

    docker-compose down -v --rmi all
    docker-compose up --build -d


## Accesss database

    psql -h db-instance-rds.c8le640i0kbl.us-east-1.rds.amazonaws.com -U p0stgr3s -d products -p 5432


## API Endpoints

    curl -X GET http://localhost:9090/v1/products

    curl -X POST http://localhost:9090/v1/products \
    -H "Content-Type: application/json" \
    -H "X-API-Key: super_secreto_api_key_valida" \
    -d '{
        "name": "Nuevo Producto de Prueba",
        "description": "Una descripción de este producto",
        "price": 50.00
    }'


##  Database Querys

    docker run -p 8081:8081 \
    -e PORT=8081 \
    -e POSTGRES_DSN="host=db-instance-rds.c8le640i0kbl.us-east-1.rds.amazonaws.com port=5432 user=p0stgr3s password=p4rc3r02025! dbname=fixture sslmode=require" \
    your-api-image-name

