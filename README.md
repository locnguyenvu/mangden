# Learn golang project

## Create new project with this template

```
$ curl https://raw.githubusercontent.com/locnguyenvu/mangden/main/setup.sh | sh
```

## Setup local environment

### Prerequisites
1. Google Protobuf 
    
    * Binary: https://github.com/protocolbuffers/protobuf
    * Golang add-ons: https://github.com/protocolbuffers/protobuf-go

2. Docker

3. MySQL server

### Preparation

Create `.env` file like example below

```
ADDR=0.0.0.0:8081

LOG_LEVEL=debug
LOG_FORMAT=text

DB_HOST=host.docker.local
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=mangden
```

Install dependencies:
```
$ go mod download 

$ go mod tidy
```

### Run development server

```
$ ./mdn hotreload-http-server [options]

    -d <port>      Debugger connection port
```

## Documents

Development use [cosmtrek/air](https://github.com/cosmtrek/air) for hot-reload, configuration file is `http-api.air.toml`

# References

* https://rakyll.org/style-packages/
* https://blog.golang.org/laws-of-reflection
