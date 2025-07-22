# restful-rds-golang-products

restful-rds-golang-products


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

