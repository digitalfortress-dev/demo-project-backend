# Golang Demo-Project-Backend using Echo Framework
This project ships following features
* Submit name, date of birth, bhone number, email, address, photo, and appointment time of patient
* View all the registered patient from website
* Authentication Using Jwt
* CORS 
* Easy database migration 

## Configured with
- [gorm](https://github.com/jinzhu/gorm):
- [jwt-go](https://github.com/dgrijalva/jwt-go):
- [echo](https://github.com/labstack/echo):
- [pusher-http-go](github.com/pusher/pusher-http-go):
- [postgres]("github.com/jinzhu/gorm/dialects/postgres"):
- Go modules
- Built-in **CORS Middleware**
- Built-in **JWT Middleware**
- Feature **PostgreSQL 14**
- Environment support **macOS Monterey 12.1**

## My local information PostgreSQL
Using PostgreSQL 14, please follow my local information postgres to create the same information and connect database in my project
```
(
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mydb"
)
```

### Creating My Database in PostgreSQL
```
$ CREATE DATABASE mydb;
```

### Installation
This project ships following step by step below to install
- Go to the [golang](https://go.dev/dl/).Then, pick up `go1.17.5.darwin-amd64.pkg` and dowload. Finally, install as the instructions appear
- Make sure you installed [zshrc](https://ohmyz.sh/).Then, Gopath, Path and Goroot are set to `~/.zshrc`:
```
$ export GO111MODULE=on
  $GOPATH $HOME/go
  export PATH=$PATH:$GOPATH/bin
  export GOROOT=/usr/local/go
  export PATH=$PATH:$GOROOT/bin
```
- Restart your `zshrc`
```
$ source ~/.zshrc
```
- Verify that you've installed Go by opening a command prompt and type the following command:
```
$ go version
```
- After the installation, launch VSCode. Flow the instructions carefully to setup `Go` with VSCODE. For more information, please visit this [link](https://dev.to/ko31/how-to-setup-golang-with-vscode-1i4i).

## Building My Application
- Enable dependency tracking for my project, just use following command:
```
go mod init main
```
- Download essential library from your terminal in my project
```
go get github.com/dgrijalva/jwt-go
go get github.com/google/uuid
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/pusher/pusher-http-go
go get gopkg.in/go-playground/validator.v9
go get gopkg.in/gormigrate.v1
```
- Execute the fllowing command in VSCode to build my project
```
$ go build -v
```
- Run my binary file by type following command:
```
$ ./main
```
## Payload of API to test Postman
Please see it in folder `Test` in my project
### Download and Set Up Postgres 14:
Download [Postgres](https://www.sqlshack.com/setting-up-a-postgresql-database-on-mac/)

### Gorm migration
For detailed guide, please visit this [link](https://gorm.io/docs/index.html).

## Dowload Postman
Download [Postman](https://www.getpostman.com/)
