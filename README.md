# Atmail Task

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project.

### Prerequisites

Create config file named config.yaml with the values below

```
{
  DB_DRIVER: mysql,
  DB_USERNAME: root,
  DB_PASSWORD: opensesame,
  DB_HOST: 127.0.0.1,
  DB_REPLICA_HOST: localhost,
  DB_PORT: 3306,
  DB_NAME: atmail
}
```

### Migration

Create a database schema according to your DB_NAME config

Then run migration script below

```
migrate -path storage/mysql/migrations -database "mysql://root:opensesame@tcp(localhost:3306)/atmail" -verbose up
```

Run the application with the command below

```
go run ./cmd serve --config config.yaml
```

## Running the tests

To be continue

### API documentation

Please check the file named "Testing.postman_collection.json"