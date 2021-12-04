# PUPGO BACKEND

PUPGO Backend is a Go backend package to serve Pupgo



## 1. Build 

### Component

MySQL Database

Go 1.17.2

Docker



## 2. Function and Structure

This repository contains several functions, such as communicate with db



- cmd/ directory is for run main.go to serve the backend database.

- configs/ contains server port which by defalut is 8080
- internals/ store several package
  - db/  migration scripts for database
  - gorm/ go orm framework related implemenation for controlling database and
  - firebase/ contains basic testing for firebase and the config file for firebase key
  - graph/  graphql resolver get user request and parse them.
  - handler/ for graphql hander and login handler.
  - notification/ sending notification functions.
  - test/ unit test for coverage in the report.





Go is a easily to build and reliable language with high concurrecy that fit for our projects backend needed.



## 3. Build

### 3.1 Build Database



migration script need the up.sh in internal/db/

To use that shell scripts , first need to install go-migrate(https://github.com/golang-migrate/migrate) 

And also the .env file as below

it should be put in ./internal/db/

```bash
export DB_RESOURCE_NAME="mysql"
export DB_USERNAME=<username>
export DB_PASSWORD=<password>
export DB_URL=""
export DB_PORT="3306"
export DB_TABLENAME="pupgo"
```

run command should be 

```sh
./up.sh
```



### 3.2 Docker
We build our backend server into Docker image. It can make us deploy PupGo API server effortlessly.
Please build Docker image on the root folder.
```bash
docker build . -t IMAGE_NAME:IMAGE_TAG
```


### 3.3 Go

To install go : here's the link. https://go.dev/doc/install



Run this command under same directory with go.mod. Will auto install dependicies.

```shell
go get
```



### 3.4 AWS S3 or image database

To build the service, first need a AWS environment for frontend to upload S3 file. Those can served or used AWS.






