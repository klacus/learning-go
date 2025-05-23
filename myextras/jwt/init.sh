#!/bin/bash

cd ./src

go mod init jwt
go mod tidy

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u github.com/gin-gonic/gin
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4
go get -u github.com/joho/godotenv
go get -u github.com/rs/zerolog
go get -u github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
