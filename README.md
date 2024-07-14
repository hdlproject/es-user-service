# ES User Service
This project contains a user service. 
Through this project, I want to illustrate how I leverage some technologies to build a user service. 
This may not be practical and inefficient since the purpose is solely for showing-off. 
This project implements Clean Architecture. 
This project is supposed to be a part of a bigger project which establish a complete microservice system.

## About the Project
### Config
Config
This project needs some configurations. 
The configurations can be supplied through several methods. 
Each method will have a priority and the higher one will take precedence. 
Here is the list of the methods ordered from the highest priority:
- environment variable
- config file
- default value

This feature is possible by the help of the [viper package](https://github.com/spf13/viper).

### External Dependencies
This project works with some storages to function properly. 
Here is the list:
- PostgreSQL to store the user data (e.g. balance, auth, and location).
- Redis to store the user location. 
  Even though this storage is volatile, it offers fast read write performance, so it is considered suitable to store rapid changing data such as location.
- MongoDB to store the events (e.g. top-up). 
  This storage do not need a predefined schema, so it is suitable to store event data with variate structure.
- RabbitMQ to broadcast the events to the other system members (e.g. top-up).
- Centrifuge to broadcast the events as well as an alternative to RabbitMQ.

## How to Run
### DB Migration
```shell script
$ migrate create -ext sql -dir db/migrations -seq <migration_name>
$ migrate -database "postgresql://postgres:postgres@localhost:5433/es-user-service?sslmode=disable" -path db/migrations up
```

## TODO
- Create common interface for pubsub (RabbitMQ and Centrifuge)
