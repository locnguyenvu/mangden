# MDN
Sample boilerplat code CRUD app based on golang 

**Table of contents**
* Start application
    - [Migration](#migration)
    - [Api server](#api-server)
    - [Grpc server](#grpc-server)
* [Setup](#setup)
* [Development](#development)

**Prerequisite**
1. Golang (version >=1.20)
2. MySQL (version >=8.0)

<details>
    <summary>Start MySQL local with docker</summary>

```bash
docker run -dit -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=mdn --name mdn-mysql-db mysql:latest
```
</details>

---
## Migration

Set environment variable to run the app

Database infomation
```bash
export DB_HOST=127.0.0.1 DB_USER=root DB_PASSWORD=root DB_NAME=mdn DB_PORT=3306 
```

**Run database migration**

```bash
go run cmd/migrate/*.go
```

## API server 

RESTful API serever via HTTP requests (`GET`, `POST`, `PUT`, `DELETE`)

Set environment variable to run the app

Web server port
```bash
export ADDR=0.0.0.0:8000 # Default 8000
```

Database infomation
```bash
export DB_HOST=127.0.0.1 DB_USER=root DB_PASSWORD=root DB_NAME=mdn DB_PORT=3306 
```

Logger config
```bash
export LOG_FORMAT=json LOG_LEVEL=info
```


**Start server**

```bash
go run cmd/apiserver/*.go
```

Routes

|method|path|body|description|
|-|-|-|-|
|`GET`|`/users`||List ten records of newly created users|
|`POST`|`/users`|```{"userName":"loc_00", "password":"passwd", "firstName":"nguyen", "lastName":"loc", "yob":1992} ```|Create new user|
|`GET`|`/users/{id}`||Get info of user with id 10|
|`PUT`|`/users/{id}`|```{"firstName":"nguyen", "lastName":"loc", "yob":1992} ```|Update info of user with id 10|
|`DELETE`|`/users/{id}`||Delete user with id 10|


## GRPC server

Generate code from `user.proto` file

```bash
protoc --go-grpc_out=module:./ --go_out=module:./ proto/user.proto
```

gRPC server port
```bash
export ADDR=0.0.0.0:50051 # Default 50051
```

Database infomation
```bash
export DB_HOST=127.0.0.1 DB_USER=root DB_PASSWORD=root DB_NAME=mdn DB_PORT=3306 
```

**Start server**

```bash
go run cmd/grpcserver/*.go
```

gRPC apis
|methods|description|
|-|-|
|`/grpc.UserService/List`|List ten records of newly created users|
|`/grpc.UserService/Get`|Get info of a specific user with provided id|
|`/grpc.UserService/Create`|Create new user|
|`/grpc.UserService/Delete`|Delete a specific user|
|`/grpc.UserService/Update`|Updae info of a specific user with provided id|

## Setup

Rename all package with your module name

```bash
./scripts/setup
```

## Development

Follow the instruction of the command `./srcipts/dev`

```bash
Usage: dev <command> [options]

Commands:
  hapiserver      Start apiserver with hotreload container
                  Options:
                      -d              start debug server at port 2345
                      -e <env-file>   path to environment file, default is `.env`

```