# crud-app
Sample crud app based on golang to check robot framework API testing


# Getting start

**Prerequisite**
1. Golang (version >=1.20)
2. MySQL (version >=8.0)

Prepare `.env` file

```
# The web app port
ADDR=0.0.0.0:8000
# Database infomation
DB_HOST=127.0.0.1
DB_USER=root
DB_PASSWORD=root
DB_NAME=fuelt
DB_PORT=3306
```

To export `.env` file to environment variable

```
export $(grep -v '^#' .env | xargs)
```


**Start server**

```
go run cmd/http-server/*.go
```

Sample API:

```
#### Run migration, create table `users` in database

GET /migrate

#### Create new user

POST /users
Body:
{ "Username": "McKinsey", "Password": "ABC123def" }

#### List all user

GET /users
```
