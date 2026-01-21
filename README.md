
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


**Response**
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


**Response**
```json
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvd25lcl9pZCI6IjciLCJlbWFpbCI6ImplYW5kdXBvbnRAZ21haWwuY29tIiwiZXhwIjoxNzY5MDE3MjAxfQ.uTKdTj8g2dVc2ep-GzIBGVCuu6lPTFx23btSGcaAlSY"
    },
    "message": "login successful"
}
```

---

#### 

```http
  POST /api/todos
```

| Payload            | Type     | Description                             |
| :---------------   |:-------- | :---------------------------------      |
| `Title`            | `string` | **Required**. Title less than 50 char   |
| `Description`      | `string` | **Required**. Description less 200 char |



**Response**
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
