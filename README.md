# Microservices: Auth + Tasks

## Описание проекта

Проект реализует два независимых микросервиса:

- Auth Service (порт 8081) — отвечает за аутентификацию и проверку токена.
- Tasks Service (порт 8082) — выполняет CRUD‑операции над задачами и требует валидный токен.

Сервисы общаются между собой через HTTP: Tasks вызывает Auth для проверки токена.

---

## Структура проекта

```bash
.
├── README.md
├── docs
│   └── pz18_grpc.md
├── go.mod
├── go.sum
├── proto
│   └── auth.proto
├── screens
│   ├── image1.png
│   ├── image2.png
│   ├── image3.png
│   ├── image4.png
│   ├── image5.png
│   ├── image6.png
│   ├── image7.png
│   └── image8.png
├── services
│   ├── auth
│   │   ├── cmd
│   │   │   └── auth
│   │   │       └── main.go
│   │   ├── internal
│   │   │   ├── grpc
│   │   │   │   └── server.go
│   │   │   ├── http
│   │   │   │   ├── handlers.go
│   │   │   │   └── router.go
│   │   │   └── service
│   │   │       └── auth.go
│   │   └── pb
│   │       └── proto
│   │           ├── auth.pb.go
│   │           └── auth_grpc.pb.go
│   └── tasks
│       ├── cmd
│       │   └── tasks
│       │       └── main.go
│       └── internal
│           ├── grpcclient
│           │   └── auth_client.go
│           ├── http
│           │   ├── handlers.go
│           │   └── router.go
│           └── service
│               └── tasks.go
└── shared
    ├── httpx
    │   └── client.go
    └── middleware
        ├── logging.go
        └── requestid.go
```
## Установка зависимостей

```bash
cd tech-ip-sem2
go get github.com/google/uuid
go clean -cache
go clean -modcache
go mod tidy
```


## Auth service

- Проверяет токен через gRPC метод AuthService.Verify.

- Использует фиксированный токен: demo-token

##  Tasks service
- HTTP API для задач.
- Перед каждой операцией вызывает Auth по gRPC:
- - AuthService.Verify(token)
- На основе ответа Auth возвращает:
- - 401 Unauthorized — неверный токен
- - 503 Service Unavailable — Auth недоступен
- - 200 OK — успешная операция

### Переменные окружения

- AUTH_GRPC_PORT — порт gRPC Auth (по умолчанию 50052)
- TASKS_PORT — порт HTTP Tasks (по умолчанию 8082)
- AUTH_GRPC_ADDR — адрес Auth для Tasks (по умолчанию localhost:50052)

### gRPC‑проверка Verify

Команда генерации 
```bash
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  services/auth/pb/proto/auth.proto
```


## Запуск сервисов

# 1. Auth (gRPC)

```bash
export AUTH_GRPC_PORT=50051
export AUTH_HTTP_PORT=8081
go run ./services/auth/cmd/auth
```
![image](screens/image1.png)

## 2. Tasks

```bash
export TASKS_PORT=8082
export AUTH_GRPC_ADDR=127.0.0.1:50052
go run ./services/tasks/cmd/tasks
```
![image](screens/image3.png)


# 3. Сценарий проверки

## Валидный токен + Проверка, что Tasks вызывает Auth по gRPC

```bash
curl -s -X POST http://localhost:8082/v1/tasks \
-H "Content-Type: application/json" \
-H "Authorization: Bearer demo-token" \
-d '{"title":"Test","description":"Check","due_date":"2026-01-10"}'
```

![image](screens/image4.png)
![image](screens/image2.png)
![image](screens/image6.png)
![image](screens/image7.png)

## Неверный токен

```bash
curl -X POST http://localhost:8082/v1/tasks \
  -H "Authorization: WRONG" \
  -H "Content-Type: application/json" \
  -d '{"title":"Test"}'                       
```

![image](screens/image5.png)

## Auth выключен
Остановить Auth и выполнить команду

```bash
curl -X GET http://localhost:8082/v1/tasks \
  -H "Authorization: demo-token"
```
![image](screens/image8.png)

## Маппинг ошибок

| **gRPC статус**       | **Смысл**                                 | **HTTP код** | **Комментарий / пример**                                 |
|-----------------------|-------------------------------------------|--------------|---------------------------------------------------------|
| `Unauthenticated`     | Неверный или отсутствующий токен          | `401`        | Токен невалиден → `unauthorized`                        |
| `PermissionDenied`    | Нет прав доступа                          | `403`        | Доступ запрещён                                          |
| `NotFound`            | Ресурс не найден                          | `404`        | Задача с таким id не найдена                             |
| `InvalidArgument`     | Некорректный запрос / тело                | `400`        | Неправильный JSON или отсутствуют обязательные поля      |
| `DeadlineExceeded`    | Превышен дедлайн (таймаут)                | `503`        | Клиент установил deadline; считать как недоступность Auth |
| `Unavailable`         | Сервис недоступен / сетевые ошибки        | `503`        | Auth не отвечает / соединение не установлено             |
| `Internal`            | Внутренняя ошибка сервера                 | `503`        | Ошибка на стороне Auth или gRPC                          |
| `OK`                  | Успех                                     | `200`        | Валидный токен, операция выполнена                       |


## Контрольные вопросы
1. Что такое .proto и почему он контракт?
- .proto — это декларативный файл, в котором описываются структуры сообщений и интерфейсы сервисов (RPC‑методы) в формате Protocol Buffers. На его основе автоматически генерируется код клиента и сервера для разных языков, поэтому обе стороны получают одинаковое представление о формате данных и сигнатурах вызовов. Именно это делает .proto формальным контрактом: изменения в нём влияют на совместимость и требуют согласованного обновления клиентов и серверов.

2. Что такое deadline в gRPC? 
- Deadline — это ограничение времени для выполнения конкретного RPC‑вызова, передаваемое клиентом в контексте запроса. Если сервер не успевает ответить до дедлайна, вызов отменяется и клиент получает ошибку (например, DeadlineExceeded), что предотвращает «висящие» запросы и освобождает ресурсы. Дедлайны полезны для управления ресурсами, обеспечения предсказуемости поведения и реализации отказоустойчивых стратегий (retry, fallback).

3. Почему “exactly-once” не даётся просто так даже в RPC?
- Гарантия exactly‑once затруднена из‑за сетевых сбоев, потери или дублирования сообщений и перезапусков компонентов: клиент может не получить ответ и повторить запрос, хотя сервер уже выполнил операцию. Для надёжного exactly‑once нужны дополнительные механизмы: идемпотентность операций, уникальные идентификаторы запросов, согласованное хранение состояния и дедупликация на стороне сервера. Эти механизмы усложняют логику и требуют согласованного проектирования как клиента, так и сервера.

4. Как обеспечивать совместимость при расширении .proto?
- Чтобы не ломать совместимость при расширении .proto, соблюдают правила: не переиспользовать и не удалять номера полей, добавлять новые поля с новыми номерами и делать их опциональными (или с дефолтными значениями). Старые клиенты будут игнорировать неизвестные поля, а новые — корректно обрабатывать старые сообщения. При серьёзных изменениях используют версионирование сервисов или отдельные RPC, чтобы избежать несовместимых изменений в продакшене.