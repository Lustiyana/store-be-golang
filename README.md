<!-- @format -->

# Final Project Sanbercode Golang

Ini merupakan API untuk aplikasi sosial media sederhana. API ini dibuat untuk memenuhi project akhir pada Bootcamp Golang Sanbercode.

# Base

`https://final-project-sanbercode-golang-production.up.railway.app`

# Authorization

Authorization using JWT token

### API

| Method | Endpoint    | Protected |
| ------ | ----------- | --------- |
| POST   | `/register` | FALSE     |
| POST   | `/login`    | false     |

# Endpoints

## Users

### Register

- URL
  - `/register`
- Method
  - POST
- Request Body

| Key                | Data Type | Required | Body type |
| ------------------ | --------- | -------- | --------- |
| `fullname`         | `string`  | TRUE     | JSON      |
| `email`            | `string`  | TRUE     | JSON      |
| `password`         | `string`  | TRUE     | JSON      |
| `confirm_password` | `string`  | TRUE     | JSON      |
| `address`          | `string`  | TRUE     | JSON      |
| `profince`         | `string`  | TRUE     | JSON      |
| `city`             | `string`  | TRUE     | JSON      |
| `postal_code`      | `string`  | TRUE     | JSON      |
| `country`          | `string`  | TRUE     | JSON      |
| `phone_number`     | `string`  | TRUE     | JSON      |

- Response

```
{
    "success": true,
    "message": "Akun berhasil dibuat",
    "data": {}
}
```

### Login

- URL
  - `/login`
- Method
  - POST
- Request Body

| Key        | Data Type | Required | Body type |
| ---------- | --------- | -------- | --------- |
| `email`    | `string`  | TRUE     | JSON      |
| `password` | `string`  | TRUE     | JSON      |

- Response

```
{
    "success": true,
    "message": "Berhasil Masuk",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJsdXN0aXlhbmExODAxQGdtYWlsLmNvbSIsInBhc3N3b3JkIjoidXRpY2FudGlrIiwiZXhwIjoxNzI0MzMwOTkxLCJpYXQiOjE3MjQyNDQ1OTEsImlzcyI6ImdvbGFuZ19zYW5iZXJjb2RlIn0.DuvSz2sL5pBM0G6nx-TlUWvOVt-BIF9Daxd7WniwM_Q"
    }
}
```
