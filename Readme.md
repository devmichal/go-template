## For better your work install:

### This is a free template to create amazing project in go

### nodemon is for rebuilt your project in real time

    npm install -g nodemon

### dynamodb-admin is to monitor your DynamoDB

    npm install -g dynamodb-admin

    dynamodb-admin

## Run this project

### Copy envs

    make local-setup

### Install packages

    go install

### Create a vendor
    
    make vendor

### Run docker image

    docker-compose up -d

### Run your server

    make hot

### Read the docs.

    api/swagger/swagger.json