# restful-rds-golang-products

restful-rds-golang-products


Amazon Virtual Private Cloud Public IPv4 Addresses - $0.32
$0.005 por dirección IPv4 pública en uso por hora × 63.019 horas

Este cobro NO es por la RDS directamente, sino por el uso de una dirección IP pública asignada a un recurso dentro de tu VPC (como una instancia EC2, una interfaz de red, etc.).

Aunque tengas una RDS en el Free Tier, el uso de direcciones IP públicas tiene un costo si:

Tienes una Elastic IP (IP pública fija) asignada pero no asociada a una instancia activa

O si tienes una interfaz de red (ENI) que tenga una IP pública en uso (por ejemplo, una EC2 en ejecución)

📌 RDS NO utiliza una IP pública por defecto, a menos que la configures como accesible públicamente ("Publicly Accessible = true")

    aws ec2 modify-subnet-attribute \
    --subnet-id subnet-07ab674ef6e9292c6 \
    --map-public-ip-on-launch

**0.005 por hora = 30 COP (La Ip publica de la RDS) 0.005 * 24 H = 483.870 COP (24 H)**

Lo que hace es habilitar que todas las instancias lanzadas en esa subred reciban automáticamente una IP pública dinámica (no Elastic IP, pero igual tiene costo si está en uso).

Aunque la RDS no se lanza directamente en esa subnet pública, si configuraste tu RDS como “publicly accessible” y le asignaste una IP pública al momento de crearla, entonces AWS te cobra $0.005 por hora mientras esa IP está en uso.



---------------------------------------


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

