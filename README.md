﻿# Dating APP - Technical Test
Build with golang and gin with jwt authentication. It features a simple and better performance, and customize with requirements needed.

### Required

 * GO 1.19.5 - [go1.19.5](https://go.dev/doc/devel/release#go1.19).

### Using

- Gin Web Framework 1.3.0 - [Gin-Gionic](https://github.com/gin-gonic/gin)
- MySQL 5.7.26 - [MySQL](https://dev.mysql.com/doc/relnotes/mysql/5.7/en/news-5-7-26.html)
- Go Validator v10 - [go-validator](https://github.com/go-playground/validator)

<br>

<h2>Installation</h2>

* Init workdir
```sh
git@github.com:mnurhads/dating-technical-test.git
cd dating-technical-test
```

* Setup Database
```sh
# database in folder db > datingapp.sql
#setting configuration connection database on folder config

# install golang package in folder vendor
go mod tidy
```

* Copy .env.example to .env
```sh
cp .env.example .env
# change default config .env with your local config
```

* Database Note
```sh
# restAPI will automatically migrate when there is no table in you database
```

* Start restAPI 
```sh
# start with default
go run main.go
```

