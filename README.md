# Ибраимов Дамир Эльдарович ПИМО-01-25
# Практическое задание 3 - HTTP сервер 

## 🎯 Цели
- Освоить работу со стандартной библиотекой net/http
- Реализовать HTTP-сервер с маршрутизацией через http.ServeMux
- Обрабатывать GET/POST запросы, query-параметры и JSON
- Добавить middleware для логирования
- Реализовать CRUD-операции для задач в памяти

## Запуск проекта
```bash
go run ./cmd/server
```
# API Endpoints
- GET /health
```bash
curl -i http://localhost:8080/health
```

- GET /tasks
```bash
curl -i http://localhost:8080/tasks
```

- GET /tasks?q=фильтр
```bash
curl -i "http://localhost:8080/tasks?q=молоко"
```

- POST /tasks
```bash
curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Новая задача"}'
```

- GET /tasks/{id}
```bash
curl -i http://localhost:8080/tasks/1
```

- POST без title
```bash
curl -i -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{}'
```

- GET с невалидным ID
```bash
curl -i http://localhost:8080/tasks/abc
```

# Структура проекта
```bash
pz3-http/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers.go
│   │   ├── middleware.go
│   │   └── responses.go
│   └── storage/
│       └── memory.go
├── go.mod
└── requests.md
```



