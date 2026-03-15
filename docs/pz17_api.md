# Примеры запросов и ответов (PZ17 API)

## Auth service


### POST /v1/auth/login

#### Пример запроса

```bash
curl -i -X POST http://localhost:8081/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"student","password":"student"}'
```

- Пример успешного ответа (200)

```json
{
  "access_token": "demo-token",
  "token_type": "Bearer"
}
```

- Ошибка 400

```json
{
  "error": "invalid_request"
}
```

- Ошибка 401

```json
{
  "error": "invalid_credentials"
}
```

### GET /v1/auth/verify

#### Пример запроса

```bash
curl -i http://localhost:8081/v1/auth/verify \
  -H "Authorization: Bearer demo-token" \
  -H "X-Request-ID: req-001"
```
- Пример успешного ответа (200)

```json
{
  "valid": true,
  "subject": "student"
}
```

- Пример ошибки 401

```json
{
  "valid": false,
  "error": "unauthorized"
}
```

## Tasks service

### POST /v1/tasks

#### Пример запроса

```bash
curl -i -X POST http://localhost:8082/v1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer demo-token" \
  -d '{"title":"Read lecture","description":"Prepare notes","due_date":"2026-01-10"}'
```

- Пример успешного ответа (201)

```json
{
  "id": "t_001",
  "title": "Read lecture",
  "description": "Prepare notes",
  "due_date": "2026-01-10",
  "done": false
}
```
- Ошибка 400

```json
{
  "error": "invalid_payload"
}
```

- Ошибка 401

```json
{
  "error": "unauthorized"
}
```

### GET /v1/tasks

#### Пример запроса

```bash
curl -i http://localhost:8082/v1/tasks \
  -H "Authorization: Bearer demo-token"
```

- Пример успешного ответа (200)

```json
[
  {"id":"t_001","title":"Read lecture","done":false},
  {"id":"t_002","title":"Do practice","done":true}
]
```

### GET /v1/tasks/{id}

#### Пример запроса

```bash
curl -i http://localhost:8082/v1/tasks/t_001 \
  -H "Authorization: Bearer demo-token"
```

- Пример успешного ответа (200)

```json
{
  "id": "t_001",
  "title": "Read lecture",
  "description": "Prepare notes",
  "done": false
}
```
- Oшибка 404

```json
{
  "error": "task_not_found"
}

```

### PATCH /v1/tasks/{id}

#### Пример запроса

```bash
curl -i -X PATCH http://localhost:8082/v1/tasks/t_001 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer demo-token" \
  -d '{"title":"Read lecture (updated)","done":true}'
```

-  Пример успешного ответа (200)

```json
{
  "id": "t_001",
  "title": "Read lecture (updated)",
  "description": "Prepare notes",
  "due_date": "2026-01-10",
  "done": true
}
```

### DELETE /v1/tasks/{id}

#### Пример запроса 
```bash
curl -i -X DELETE http://localhost:8082/v1/tasks/t_001 \
  -H "Authorization: Bearer demo-token"
```

- Ошибка 404

```json
{
  "error": "task_not_found"
}
```

- Пример запроса без токена (должно быть 401)

```bash
curl -i http://localhost:8082/v1/tasks
```

- Ответ 

```json
{
  "error": "unauthorized"
}
```
