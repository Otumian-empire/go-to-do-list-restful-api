# Todo List Restful Api Using Golang

## Request

Request will be accepted in `json` format, pass `Content-Type` as `application/json`.

## Response

### Message Response

Message response return a object of `success` and `message` of types, `bool` and `string` respectively

#### MessageResponse

```json
{
  "success": true,
  "message": "some action occurred successful"
}
```

### Data Response

Data response returns an object of `success`, `message` and `data` of types, `bool`, `string` and a `map[string]any` respectively. This is will returned when fetching data and the action was successful. For cases where we have to read a list of data, pagination property will be added as part of the `data`.

#### Single row

**DataResponse**

```json
{
   "success": true,
   "message": "some resource was read successful",
   "data" {
      "id":1,
      "username":"john123",
      ...
   }
}
```

#### Multiple rows

**PaginatedDataResponse**

```json
{
  "success": true,
  "message": "some resource was read successful",
  "data" {
      "rows": [
         {
            "id":1,
            "username":"row1",
            ...
         },
         {
            "id":2,
            "username":"row2",
            ...
         },
      ],
      "pagination": {
         "pageNumber": 1,
         "pageSize": 10,
         "count": 43
      }
  }
}
```

## Entities

### User

#### Field table

| Field     | Type     | Properties              | Description                                                                       |
| --------- | -------- | ----------------------- | --------------------------------------------------------------------------------- |
| id        | `int`    | NOT NULL, AUTOINCREMENT | user id that uniquely identifies user                                             |
| username  | `string` | NOT NULL, UNIQUE        | username is required                                                              |
| password  | `string` | NOT NULL                | hashed password before insert                                                     |
| createdAt | `Date`   |                         | At creation, is the the current timestamp                                         |
| updatedAt | `Date`   |                         | At creation, is the the current timestamp and on update, is the current timestamp |

#### Behaviour table

| Action         | Argument             | Return                              | Description                                       |
| -------------- | -------------------- | ----------------------------------- | ------------------------------------------------- |
| SignUp         | `username, password` | `MessageResponse`                   | Create a user                                     |
| Login          | `username, password` | `MessageResponse` \| `DataResponse` | Authenticate and authorize user                   |
| UpdateUsername | `username`           | `MessageResponse`                   | Update a user's username, authentication required |
| UpdatePassword | `password`           | `MessageResponse`                   | Update a user's password, authentication required |
| ReadUser       |                      | `DataResponse`                      | Read user details, authentication required        |
| DeleteUser     |                      | `DataResponse`                      | Delete user details, authentication required      |

#### Validation

- id: int, database generated
- username: string, minimum of 5, maximum of 10
- password: string, minimum of 5, maximum of 10

#### API Endpoints

| Action         | Method   | Endpoint          | Description                     |
| -------------- | -------- | ----------------- | ------------------------------- |
| SignUp         | `POST`   | `/users`          | to create a new user            |
| Login          | `POST`   | `/users/auth`     | to login a user                 |
| UpdateUsername | `PUT`    | `/users/username` | to update username              |
| UpdatePassword | `PUT`    | `/users/password` | to update password              |
| ReadUser       | `GET`    | `/users`          | to create a new task            |
| DeleteUser     | `DELETE` | `/users`          | to retrieve a list of all tasks |

## Tool/Dependencies

- Gin for https server and routing
- PostgreSQL for database
- Render as hosting server
- Github to deploy code to Render
