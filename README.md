### JWT-аутентификация: создание и проверка токенов. Middleware для авторизации
### Ибраимов Дамир Эльдарович, ПИМО-01-25.

#### Суть практической 
   JWT‑аутентификация HS256, выдача access и refresh токенов, middleware AuthN и AuthZ, RBAC и ABAC правило для `/api/v1/users/{id}`, in‑memory репозиторий с bcrypt‑хэшами, endpoint `/api/v1/refresh` с blacklist для отозванных refresh.


#### Структура проекта

```bash
practica10/
├── cmd
│   └── server
│       └── main.go           # Точка входа, загружает конфиг и запускает HTTP сервер.
├── internal
│   ├── core
│   │   ├── user.go           # Доменная модель User.
│   │   └── service.go        # Бизнес-логика: LoginHandler, RefreshHandler, MeHandler, UserHandler, blacklist для refresh.
│   ├── http
│   │   ├── middleware
│   │   │   ├── authn.go      # Middleware для JWT-аутентификации: извлечение токена, валидация, кладёт клеймы в context.
│   │   │   └── authz.go      # Middleware для авторизации: проверяет роль из клеймов (RBAC).
│   │   └── router.go         # Настройка chi роутера, подключение middleware и DI.
│   ├── platform
│   │   ├── config
│   │   │   └── config.go     # Загрузка конфигурации: JWT_SECRET, TTL, PORT.
│   │   └── jwt
│   │       └── jwt.go        # Пакет для работы с токенами: Sign (access), SignRefresh (с jti), Parse.
│   └── repo
│       └── user_mem.go       # In-memory "база данных": хранит пользователей и проверяет bcrypt.
├── go.mod
└── README.md
```
---

#### Инструкция запуска

**Переменные окружения**  
- `JWT_SECRET` — секрет для HS256 (обязательно).  
- `APP_PORT` — порт сервера (по умолчанию `8080`).  
- (опционально) `JWT_TTL` — общий TTL если используется; в реализации используются явные TTL для access и refresh.


```bash
export JWT_SECRET=dev-secret
export APP_PORT=8080
go run ./cmd/server
```
![screen](image2.png)

**Скриншоты**

```
# export JWT_SECRET=dev-secret; export JWT_TTL=2h; export APP_PORT=8080; go run ./cmd/server
```
![screen](image1.png)
```
curl -s -X POST http://localhost:8080/api/v1/login \
 -H "Content-Type: application/json" \
 -d '{"Email":"admin@example.com","Password":"secret"}'
```
![screen](image3.png)
```
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJwejEwLWNsaWVudHMiLCJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzY0OTU4NzIxLCJpYXQiOjE3NjQ5NTE1MjEsImlzcyI6InB6MTAtYXV0aCIsInJvbGUiOiJhZG1pbiIsInN1YiI6MX0.DjpOSiVxOHZHPZCuKU0OvP-U5I2hDo53jPfHPIO9264
curl -s http://localhost:8080/api/v1/me -H "Authorization: Bearer $TOKEN"
curl -s http://localhost:8080/api/v1/admin/stats -H "Authorization: Bearer $TOKEN"
```
![screen](image4.png)
```
TOKEN_USER=$(curl -s -X POST http://localhost:8080/api/v1/login \
 -H "Content-Type: application/json" -d '{"Email":"user@example.com","Password":"secret"}' | jq -r .token)
curl -i http://localhost:8080/api/v1/admin/stats -H "Authorization: Bearer $TOKEN_USER"
```
![screen](image5.png)


#### Контрольные вопросы и ответы

1. **Что такое клеймы JWT и чем отличаются registered public private Почему важно exp**  
   **Клеймы** — это утверждения (claims) в payload JWT: стандартные (registered), публичные и приватные. **Registered** — заранее определённые имена (например, `iss`, `sub`, `aud`, `exp`, `iat`) с рекомендованным смыслом. **Public** — общедоступные имена, которые можно зарегистрировать в IANA или договориться о них между сервисами. **Private** — произвольные пользовательские клеймы для конкретного приложения (например, `role`). Поле **exp** важно потому, что ограничивает срок действия токена; без него токен действовал бы бесконечно, что повышает риск компрометации. Проверка `exp` предотвращает использование старых токенов.

2. **Чем stateless-аутентификация на JWT отличается от сессионных cookie на сервере Плюсы минусы**  
   **JWT (stateless)**: сервер не хранит сессии — валидируется подпись и клеймы. **Плюсы**: простота масштабирования, меньше нагрузки на сервер, удобство для микросервисов. **Минусы**: сложнее отзыв токенов (нужен blacklist/refresh), токен может содержать чувствительные данные, риск хранения на клиенте.  
   **Сессионные cookie (stateful)**: сервер хранит сессии (в памяти/БД/Redis). **Плюсы**: лёгкий отзыв и управление сессиями, меньший риск утечки данных в токене. **Минусы**: требуется хранение состояния и синхронизация между инстансами, дополнительная нагрузка.

3. **Как устроена цепочка middleware и почему AuthZ должна идти после AuthN**  
   Middleware — цепочка функций, каждая получает `request` и `next`. **AuthN** (аутентификация) должна идти первой, потому что она определяет **кто** делает запрос и кладёт клеймы в `context`. **AuthZ** (авторизация) использует эти клеймы (роль, sub и т.д.) для принятия решения о доступе. Если AuthZ выполняется до AuthN, у неё нет данных о пользователе и она не сможет корректно проверить права.

4. **RBAC vs ABAC когда что выбирать Примеры**  
   **RBAC (Role Based Access Control)** — права назначаются ролям (admin, user). Подходит для простых систем с фиксированными ролями. Пример: только `admin` может просматривать статистику.  
   **ABAC (Attribute Based Access Control)** — решение на основе атрибутов (роль, owner id, время, IP). Подходит для гибких политик. Пример: пользователь с ролью `user` может читать `/users/{id}` только если `{id} == sub` (владелец ресурса). Часто комбинируют RBAC и ABAC.

5. **Как безопасно хранить пароль и почему нужен bcrypt argon2 вместо SHA-256 соль pepper**  
   Пароли нужно хранить в виде адаптивных хешей (bcrypt, Argon2). Эти алгоритмы медленнее и устойчивы к брутфорсу, поддерживают настройку стоимости (work factor). **SHA‑256** — быстрый хеш, не предназначен для паролей: атаки перебором выполняются очень быстро. **Соль** — уникальная случайная строка для каждого пароля, предотвращает использование радужных таблиц. **Pepper** — секрет, общий для всех паролей и хранящийся отдельно (например, в конфиге/секретном хранилище), добавляет дополнительный уровень защиты при компрометации БД. Всегда используйте проверенные библиотеки и не храните пароли в открытом виде.
