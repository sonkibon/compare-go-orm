# gorm

## Overview
Sample code using [ent](https://github.com/ent/ent) to access database.

## Getting Started

1. Start mysql container
```
docker-compose up -d mysql
```
1. hoge
```
cd ent
```
1. hoge
```
go install entgo.io/ent/cmd/ent@v0.12.3
```
1. hoge
```
ent new Employee Department DeptManager
```
1. hoge
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
