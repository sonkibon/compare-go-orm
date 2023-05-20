# sqlboiler

## Overview
Sample code using [sqlboiler](https://github.com/volatiletech/sqlboiler) to access database.

## Getting Started

1. Uncomment

Check this [comment](https://github.com/sonkibon/compare-go-orm/blob/main/docker-compose.yml#L5)

2. Start mysql container
```
docker-compose up -d mysql
```
3. Run the code
```
go run sqlboiler/main.go
```

### If you want to generate code

1. Install packages
```
go install github.com/volatiletech/sqlboiler/v4@v4.14.2
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2
```
2. Change directory
```
cd sqlboiler
```
3. Generate code
```
sqlboiler mysql --output entity --pkgname entity --add-soft-deletes
```
