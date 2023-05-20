# gorm

## Overview
Sample code using [gorm](https://github.com/go-gorm/gorm) to access database.

## Getting Started

1. Start mysql container
```
docker-compose up -d mysql
```
2. Run the code
```
go run gorm/main.go
```

### If you want to access database

1. Run exec command
```
docker-compose exec mysql bash
```
2. Connect to the server
```
mysql -h 127.0.0.1 -P 3306 -u root -p
```
