# Todo List Restful Api Using Golang

Documentation on application design

## Tool/Dependencies

- Gin for https server and routing
- PostgreSQL for database
- Render as hosting server
- Github to deploy code to Render

## Request

Request will be accepted in `json` format, pass `Content-Type` as `application/json`.

## Response

### Message Response

Message response return a object of `success` and `message` of types, `bool` and `string` respectively

#### MessageResponse

```json
{
  "success": true,
  "message": "some action occurred successfully"
}
```

### Data Response

Data response returns an object of `success`, `message` and `data` of types, `bool`, `string` and a `map[string]any` respectively. This is will returned when fetching data and the action was successful. For cases where we have to read a list of data, pagination property will be added as part of the `data`.

#### Single row

**DataResponse**

```json
{
   "success": true,
   "message": "some resource was read successfully",
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

- A field with `?` before it is optional
- `true`: Authentication is required before authorization
- `false`: No authentication required and is opened

### User

Table name = "users"

#### Field table

| Field     | Type     | Properties              | Description                                                                       |
| --------- | -------- | ----------------------- | --------------------------------------------------------------------------------- |
| id        | `int`    | NOT NULL, AUTOINCREMENT | user id that uniquely identifies user                                             |
| username  | `string` | NOT NULL, UNIQUE        | username is required                                                              |
| password  | `string` | NOT NULL                | hashed password before insert                                                     |
| createdAt | `Date`   | NOT NULL                | At creation, is the the current timestamp                                         |
| updatedAt | `Date`   | NOT NULL                | At creation, is the the current timestamp and on update, is the current timestamp |

#### Behaviour table

| Action         | Argument             | Return                              | Description                          |
| -------------- | -------------------- | ----------------------------------- | ------------------------------------ |
| SignUp         | `username, password` | `MessageResponse`                   | Create a user                        |
| Login          | `username, password` | `MessageResponse` \| `DataResponse` | Authenticate and authorize user      |
| UpdateUsername | `username`           | `MessageResponse`                   | Update a user's username             |
| UpdatePassword | `password`           | `MessageResponse`                   | Update a user's password             |
| ReadUser       |                      | `DataResponse`                      | Read this authenticated user details |
| DeleteUser     |                      | `MessageResponse`                   | Delete user details                  |
| Logout         |                      | `MessageResponse`                   | Logout user                          |

#### Validation

- id: `int`, positive
- username: `string`, minimum of 5, maximum of 10
- password: `string`, minimum of 5, maximum of 10

#### API Endpoints

| Action         | Auth    | Method   | Endpoint          | Description                          |
| -------------- | ------- | -------- | ----------------- | ------------------------------------ |
| SignUp         | `false` | `POST`   | `/users`          | to create a new user                 |
| Login          | `false` | `POST`   | `/users/auth`     | to login a user                      |
| UpdateUsername | `true`  | `PUT`    | `/users/username` | to update username                   |
| UpdatePassword | `true`  | `PUT`    | `/users/password` | to update password                   |
| ReadUser       | `true`  | `GET`    | `/users`          | to create a new task                 |
| DeleteUser     | `true`  | `DELETE` | `/users`          | to delete a user and all their todos |
| Logout         | `true`  | `GET`    | `/users/auth`     | to logout a user                     |

### Todo

Table name = "todos"

#### Field table

| Field     | Type     | Properties              | Description                                                                       |
| --------- | -------- | ----------------------- | --------------------------------------------------------------------------------- |
| id        | `int`    | NOT NULL, AUTOINCREMENT | todo id that uniquely identifies task                                             |
| task      | `string` | NOT NULL                | task is required                                                                  |
| completed | `bool`   | NOT NULL                | marks the status of the task (todo) as done or not                                |
| createdAt | `Date`   | NOT NULL                | At creation, is the the current timestamp                                         |
| updatedAt | `Date`   | NOT NULL                | At creation, is the the current timestamp and on update, is the current timestamp |

#### Behaviour table

| Action          | Argument                   | Return                  | Description                          |
| --------------- | -------------------------- | ----------------------- | ------------------------------------ |
| CreateTodo      | `task`                     | `MessageResponse`       | Create a todo                        |
| ReadTodo        | `id`                       | `DataResponse`          | Read a user's todo by id             |
| ReadTodos       | `?pageNumber`, `?pageSize` | `PaginatedDataResponse` | Read a section of a user's todos     |
| UpdateTask      | `id`, `task`               | `MessageResponse`       | Update a user's todo task by id      |
| UpdateCompleted | `id`, `completed`          | `MessageResponse`       | Update a user's todo completed state |
| DeleteTodo      | `id`                       | `MessageResponse`       | Delete a user's todo by id           |

#### Validation

- id: `int`, positive
- task: `string`, min of 5, maximum 255
- completed: `bool`, true or false
- ?pageNumber: `int`, positive
- ?pageSize: `int`, positive

#### API Endpoints

| Action          | Auth    | Method   | Endpoint                         | Description                    |
| --------------- | ------- | -------- | -------------------------------- | ------------------------------ |
| CreateTodo      | `false` | `POST`   | `/todos`                         | to create a new todo           |
| ReadTodo        | `true`  | `GET`    | `/todos/:id`                     | to read a todo                 |
| ReadTodos       | `true`  | `GET`    | `/todos?pageNumber=1&pageSize=4` | to read todos                  |
| UpdateTask      | `true`  | `PUT`    | `/todos/:id`                     | to update task                 |
| UpdateCompleted | `true`  | `GET`    | `/todos/:id/state`               | to update the completion state |
| DeleteTodo      | `true`  | `DELETE` | `/todos/:id`                     | to delete a todo               |

### Entity Relationship

A user may zero or more todos

## Authentication

- `GetAuthenticatedToken(userId, username string) string`: get authentication token that last for 30 minutes using username nad password by creating a new user or logging in

## Middleware

- `IsAuthenticated(authToken: string) bool`: pass the authenticated token to get authorization to access resources
- set the authenticated user to request.user
