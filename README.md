
# Todo List API

An roadmap.sh challenge who consist at build RESTful API for manage todo list.


## Goals

- User authentication
- Schema design and Databases
- RESTful API design
- CRUD operations
- Error handling
- Security


## Requirements

- User registration to create a new user
- Login endpoint to authenticate the user and generate a token
- CRUD operations for managing the to-do list
- Implement user authentication to allow only authorized users to access the to-do list
- Implement error handling and security measures
- Use a database to store the user and to-do list data (you can use any database of your choice)
- Implement proper data validation
- Implement pagination and filtering for the to-do list


## API Reference

#### Register

```http
  POST /api/register
```

| Payload |          Type     | Description                       |
| :---------------|:-------- | :--------------------------------- |
| `Name`          | `string` | **Required**. Name of your account |
| `Email`         | `string` | **Required**. Put unique email     |
| `Password`      | `string` | **Required**. Enter string password|


```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvd25lcl9pZCI6IjciLCJlbWFpbCI6ImplYW5kdXBvbnRAZ21haWwuY29tIiwiZXhwIjoxNzY5MDE2OTIxfQ.pv8B3rsf80QgN1KZKlPUqSG2J3GAlYIcYseXwh30G7E"
    },
    "message": "user registered successfully"
}
```

---

#### Login

```http
  POST /api/login
```

| Payload |          Type     | Description                       |
| :---------------|:-------- | :--------------------------------- |
| `Email`         | `string` | **Required**. Put unique email     |
| `Password`      | `string` | **Required**. Enter your password  |


```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvd25lcl9pZCI6IjciLCJlbWFpbCI6ImplYW5kdXBvbnRAZ21haWwuY29tIiwiZXhwIjoxNzY5MDE3MjAxfQ.uTKdTj8g2dVc2ep-GzIBGVCuu6lPTFx23btSGcaAlSY"
    },
    "message": "login successful"
}
```

---

#### New tasks

```http
  POST /api/todos
```

| Payload            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Title`            | `string` | **Required**. Title less than 50 char   |
| `Description`      | `string` | **Required**. Description less 200 char |


| Headers            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Authorization`    | `string` | **Required**. JWT authentication token  |

```json
{
    "data": {
        "id": "50",
        "title": "Réunion équipe",
        "description": "Brainstorming sur l'IA",
        "status": "",
        "created_at": "2026-01-21T13:41:13.044558Z",
        "updated_at": "2026-01-21T13:41:13.044558Z"
    },
    "message": "task created successfully"
}
```

---

#### Update tasks

```http
  PUT /api/todos/:id
```

| Payload            | Type     | Description                               |
| :---------------   |:-------- | :---------------------------------        |
| `Title`            | `string` | **Optional**. Title less than 50 char     |
| `Description`      | `string` | **Optional**. Description less 200 char   |
| `Status`           | `string` | **Optional**. Change status of your tasks |


| Headers            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Authorization`    | `string` | **Required**. JWT authentication token  |

```json
{
    "data": {
        "id": "50",
        "title": "Réunion équipe",
        "description": "Annuler",
        "status": "",
        "created_at": "2026-01-21T13:41:13.044558Z",
        "updated_at": "2026-01-21T13:50:39.573056Z"
    },
    "message": "task updated successfully"
}
```
---

#### Update tasks

```http
  PUT /api/todos/:id
```

| Payload            | Type     | Description                               |
| :---------------   |:-------- | :---------------------------------        |
| `Title`            | `string` | **Optional**. Title less than 50 char     |
| `Description`      | `string` | **Optional**. Description less 200 char   |
| `Status`           | `string` | **Optional**. Change status of your tasks |


| Headers            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Authorization`    | `string` | **Required**. JWT authentication token  |

```json
{
    "data": {
        "id": "50",
        "title": "Réunion équipe",
        "description": "Annuler",
        "status": "",
        "created_at": "2026-01-21T13:41:13.044558Z",
        "updated_at": "2026-01-21T13:50:39.573056Z"
    },
    "message": "task updated successfully"
}
```

---

#### Delete tasks

```http
  DELETE /api/todos/:id
```

| Headers            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Authorization`    | `string` | **Required**. JWT authentication token  |

---

#### Get tasks

```http
  GET /api/todos
```

| Params             | Type     | Description                                  |
| :---------------   |:-------- | :---------------------------------           |
| `page`             | `int` | **Required**. Pick the page you wanna see       |
| `limit`            | `int` | **Required**. Put the limit of elements per page|


| Headers            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Authorization`    | `string` | **Required**. JWT authentication token  |

```json
{
    "data": {
        "results": [
            {
                "id": 51,
                "title": "Développement API",
                "description": "Implémentation des endpoints REST pour la gestion des utilisateurs"
            },
            {
                "id": 52,
                "title": "Correction de bugs",
                "description": "Résolution des erreurs de validation lors de la création des tâches"
            },
            {
                "id": 53,
                "title": "Sécurité JWT",
                "description": "Mise en place et vérification de l’authentification par token JWT"
            }
        ],
        "page": 1,
        "limit": 3,
        "total": 5
    },
    "message": "tasks retrieved successfully"
}
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`JWT_SECRET_KEY` = Secret key used to sign and verify JWT tokens.
It must be long, random, and kept strictly confidential.

`POSTGRES_DB` =  Name of the PostgreSQL database used by the application.

`POSTGRES_USR` = PostgreSQL username with permissions to access and modify the database.

`POSTGRES_PWD` = Password associated with the PostgreSQL user.

`POSTGRES_HOST` = ddress of the PostgreSQL server (e.g. localhost, 127.0.0.1, or a Docker service name).



## Roadmap

- Additional browser support

- Add more integrations

