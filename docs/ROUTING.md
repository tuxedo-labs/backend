
## API Reference

#### Register

```http
  POST /api/auth/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Your name |
| `email` | `string` | **Required**. Your email unique |
| `password` | `string` | **Required**. Your password |

response

```json
{
  "message": "Success register"
}
```

#### Login

```http
  POST /api/auth/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. Your email |
| `password`      | `string` | **Required**. Your password |


