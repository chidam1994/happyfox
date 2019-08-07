## Setup
run the following to clone this repo in your go path
```
go get github.com/chidam1994/happyfox
```
cd into the repo and run the following to install all dependencies
```
go get ./...
```
open config/config.json and update the db connection details in this file

## Migrations
run the following to install goose which will be used to run the migrations
```
go get -u github.com/pressly/goose/cmd/goose
```
after installing goose cd into migrations directory and run the following command after replacing the db connection details
```
goose postgres "user=postgres password=postgres dbname=postgres sslmode=disable" up
```

## Start Server
To start the app run
cd back to the root folder happyfox and run
```
go run .\server\main.go
```
## Api endpoints
check api_reference file 
