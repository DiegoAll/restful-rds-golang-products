# restful-rds-golang-products


RESTful microservice that allows the management of soccer tournament data. It uses AWS cloud resources and
is provisioned and deployed using Terraform.


<p align="center">
  <img src="diagram.jpeg" width="600"/>
</p>


    aws ec2 modify-subnet-attribute \
    --subnet-id subnet-07ab674ef6e9292c6 \
    --map-public-ip-on-launch


## Run application locally

    docker-compose down -v --rmi all
    docker-compose up --build -d


## Accesss remote database

    psql -h db-instance-rds-tf.c8le640i0kbl.us-east-1.rds.amazonaws.com -U p0stgr3s -d products -p 5432


## Provisioning Infrastructure

    terraform init
    terraform plan
    terraform apply

## Removing Provisioned Infrastructure

    terrform destroy






