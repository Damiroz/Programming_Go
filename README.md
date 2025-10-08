# Ибраимов Дамир Эльдарович ПИМО-01-25
# Практическое задание 4 -  CRUD Service

Простой CRUD-сервис для управления списком задач с использованием роутера chi.

## Описание

REST API сервис для создания, чтения, обновления и удаления задач. Хранилище - in-memory (в памяти), подходит для демонстрационных целей и обучения.

## Технологии

- **Go** 1.21+
- **Chi** - легкий и быстрый HTTP роутер
- **Standard Library** - для работы с HTTP и JSON

## Структура проекта
```bash
pz4-todo/
├── go.mod
├── main.go
├── internal/
│   └── task/
│       ├── model.go
│       ├── repo.go
│       └── handler.go
└── pkg/
    └── middleware/
        ├── logger.go
        └── cors.go
```


## Установка и запуск

1. **Клонирование и настройка:**
```bash
mkdir pz4-todo
cd pz4-todo

# API Endpoints
- GET /health
```bash
curl -i http://localhost:8080/health
```

2. **Создание структуры папок:**
```bash
mkdir -p internal/task
mkdir -p pkg/middleware
```

3. **Инициализация модуля:**
```bash
go mod init example.com/pz4-todo
go get github.com/go-chi/chi/v5
```

4. **Запуск сервера:**
```bash
go run .
```

# API Endpoints
- Health Check
```bash
GET /health - проверка работоспособности сервиса
```

- GET /tasks/{id}
```bash
curl -i http://localhost:8080/tasks/1
```

- Создание задачи
```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Изучить Go"}'
```

- Получение списка задач
```bash
curl http://localhost:8080/api/tasks
```

- Получение задачи по ID
```bash
curl http://localhost:8080/api/tasks/1
```

- Обновление задачи
```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Изучить Go и chi","done":true}'
```

- Удаление задачи
```bash
curl -X DELETE http://localhost:8080/api/tasks/1
```


# Middleware

- Logger - логирование всех HTTP запросов
- CORS - разрешение кросс-доменных запросов
- RequestID - добавление ID к каждому запросу
- Recoverer - обработка паник

# Scrins 
Папка со всеми скриншотами запросов

# Итог

В ходе практической работы был успешно разработан CRUD-сервис для управления задачами на языке Go с использованием роутера chi. Сервис предоставляет REST API с полным набором операций создания, чтения, обновления и удаления записей.
- Реализована модульная архитектура проекта с четким разделением ответственности между компонентами
- Все CRUD-операции работают корректно и возвращают соответствующие HTTP-статусы
- Интегрированы middleware для логирования запросов и обработки CORS
- Сервис успешно протестирован через curl и готов к использованию
- Обеспечена потокобезопасность операций с данными через sync.RWMutex


# Сложности и решения:
- Организация корректной валидации входных данных (обработка пустых полей, неверных форматов)
- Реализация обработки preflight OPTIONS-запросов для CORS
- Оптимизация структуры хранения данных в памяти с использованием map и счетчика






