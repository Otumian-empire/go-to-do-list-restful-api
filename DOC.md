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

#### Validation

- id: `int`, positive
- username: `string`, minimum of 5, maximum of 10
- password: `string`, minimum of 5, maximum of 10

#### API Endpoints

| Action         | Auth    | Method   | Endpoint          | Description                     |
| -------------- | ------- | -------- | ----------------- | ------------------------------- |
| SignUp         | `false` | `POST`   | `/users`          | to create a new user            |
| Login          | `true`  | `POST`   | `/users/auth`     | to login a user                 |
| UpdateUsername | `true`  | `PUT`    | `/users/username` | to update username              |
| UpdatePassword | `true`  | `PUT`    | `/users/password` | to update password              |
| ReadUser       | `true`  | `GET`    | `/users`          | to create a new task            |
| DeleteUser     | `true`  | `DELETE` | `/users`          | to retrieve a list of all tasks |

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

| Action          | Auth    | Method   | Endpoint          | Description                     |
| --------------- | ------- | -------- | ----------------- | ------------------------------- |
| CreateTodo      | `false` | `POST`   | `/users`          | to create a new user            |
| ReadTodo        | `true`  | `POST`   | `/users/auth`     | to login a user                 |
| ReadTodos       | `true`  | `PUT`    | `/users/username` | to update username              |
| UpdateTask      | `true`  | `PUT`    | `/users/password` | to update password              |
| UpdateCompleted | `true`  | `GET`    | `/users`          | to create a new task            |
| DeleteTodo      | `true`  | `DELETE` | `/users`          | to retrieve a list of all tasks |

## Tool/Dependencies

- Gin for https server and routing
- PostgreSQL for database
- Render as hosting server
- Github to deploy code to Render
