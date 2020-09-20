# Go Demo API

This project is a proof of concept to understand the Go (Golang) ecosystem.

## Endpoints

### Show User

`GET /users/:id`

Return a user from a given ID.

#### Response

```json
{
    "ID": 1233,
    "FirstName": "John",
    "LastName": "Smith",
    "Email": "john.smith@acme.com",
    "RoleID": 1,
    "CreatedAt": "2020-10-11T17:57:50Z",
    "UpdatedAt": "2020-10-11T17:57:50Z",
    "Role": {
        "ID": 1,
        "Name": "User"
    }
}
```

If a user is not found the following response is returned:

```json
{
    "Errors": {
        "NotFound": [
            "User not found."
        ]
    },
    "Message": "There was a problem processing your request"
}
```

### Create Contact Request

Send an email based on environment variable. The email is placed on a queue to be send by a worker.

#### Response

```json
{
    "Success": true
}
```

## Technologies Used

This module uses the following Go packages:

- [Fiber](https://github.com/gofiber/fiber): for HTTP routing and middlewares
- [GORM](https://github.com/go-gorm/gorm): ORM
- [taskq](https://github.com/vmihailenco/taskq): for async job processing via Redis
- [GoDotEnv](https://github.com/joho/godotenv): for loading env variables from an `.env` file
- [email](https://github.com/jordan-wright/email): to send email via SMTP
