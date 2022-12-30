<h3 align="center">GO - JWT Authentication Tool</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/localhost-8000/go-jwt-auth.svg)](https://github.com/localhost-8000/go-jwt-auth/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/localhost-8000/go-jwt-auth.svg)](https://github.com/localhost-8000/go-jwt-auth/pulls)
<!-- [![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE) -->

</div>

---

<p align="center"> A simple tool written in GO to generate JWT tokens for authentication purposes. <br>Useful for testing APIs. 
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Built Using](#built_using)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

User can register and login to the system. Once logged in, user can generate JWT tokens for authentication purposes.
Afterwards, user also will be able to logout from the system.

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

You need to have GO installed on your machine. Also, you need to have a PostgreSQL database running on your machine. 



### Installing

1. Create a database in PostgreSQL.
2. Create a .env file in the root directory of the project and add the following environment variables:

```sh
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
```

3. Run the following command to install the dependencies:

```sh
go get
```

4. Run the following command to start the server:

```sh
go run main.go
```


## 🎈 Usage <a name="usage"></a>

1. Register a new user by sending a POST request to the following endpoint:

```sh
http://localhost:3000/api/register
``` 

2. Login to the system by sending a POST request to the following endpoint. JWT token will be saved in the cookies:

```sh
http://localhost:3000/api/login
```

3. Get the user data by sending a GET request to the following endpoint:

```sh
http://localhost:3000/api/user
```

4. Logout from the system by sending a POST request to the following endpoint:

```sh
http://localhost:3000/api/logout
```


## ⛏️ Built Using <a name = "built_using"></a>

- [PostgreSQL](https://www.postgresql.org/) - Database
- [Go Fiber](https://gofiber.io/) - Web Framework
- [GO](https://go.dev/) - Programming Language

## ✍️ Authors <a name = "authors"></a>

- [@localhost](https://github.com/localhost-8000) 


