# gorm

## Overview
Sample code using [ent](https://github.com/ent/ent) to access database.

## Getting Started

1. Start mysql container
```
docker-compose up -d mysql
```
2. Run the code
```
go run gorm/main.go
```

### If you want to generate code

1. Install packages
```
go install entgo.io/ent/cmd/ent@v0.12.3
```
2. Create schemage
```
ent new Employee Department DeptManager
```
3. Edit schemage
4. Generate code
```
go generate ./ent
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
