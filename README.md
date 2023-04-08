# Todo API in GO

Todo API allows the client to create, read, update and delete todo items.

This project is sample project that shows how to use 
DDD(Domain-Driven Design) with hexagonal architecture in Golang and Gin.

## Persistence

1. Using In-Memory
1. Cassandra (Todo)

## Requirements

1. Go 1.18
1. Gin
1. Testify
1. Swagger

## Usage

`make run`

## Build

`make build`

## Using Docker

To build the app with docker, run

`docker build -t todoapi .`

`docker run -d -p 8080:80 --name todoapiapp todoapi`

`./run.sh`


`docker exec cass1 nodetool status`

`docker exec -it cass1  cqlsh  -e "describe keyspaces"`

## Mocks

Generate mocks with mockery: 

```
$ mockery --dir=pkg/domain/task --name=AddTaskRepository --filename=service_mock.go --output=pkg/domain/task/mocks --outpkg=mocks
```

## Documentation

URL: http://localhost:8081/docs/index.html