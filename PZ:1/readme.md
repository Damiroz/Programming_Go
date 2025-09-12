# Hello API
Первый Микросервис на Go с REST API эндпоинтами.

## Требования
- Go 1.25+
- Модуль `github.com/google/uuid`

## Запуск 

```bash
#Создание структуры 
mkdir helloapi
cd helloapi

#go.mod — «паспорт» проекта с именем модуля и версией Go.
go mod init example.com/helloapi

mkdir -p cmd/server.

```
## Создание самого кода 

-Создайте файл cmd/server/main.go

-Подтянем пакет для генерации UUID и заменим «temp» на реальный идентификатор.

```bash
go get github.com/google/uuid@latest
go mod tidy
```

## Запуск сервера и быстрая проверка
```bash
go run ./cmd/server
#В другом окне PowerShell проверьте эндпоинты:
curl http://localhost:8080/hello
curl http://localhost:8080/user
curl http://localhost:8080/health
```

## Сборка бинарника и форматирование 
```bash
go build -o helloapi.exe ./cmd/server
.\helloapi.exe
```

-Отформатируйте и проверьте код стандартными инструментами Go
```bash
go fmt ./...
go vet ./...
```

## Создание переменной окружения 
-Для Macbook 
```bash
export APP_PORT="8081"
# Чтобы сервер мог слушать порт :8081
```

## Проблемы, которые могут возникнуть 

-go: command not found  переоткройте PowerShell после установки Go, проверьте PATH, переустановите Go-installer.

-no Go files in ... при go run — запускайте из корня helloapi командой go run ./cmd/server.

-Порт уже занят (listen tcp :8080: bind: ...) — запустите на другом порту (export APP_PORT="8081"), либо завершите процесс, который держит порт.

-Неверный JSON/заголовки — проверяйте Content-Type: application/json и используйте json.NewEncoder(w).Encode(...).

-go: module github.com/google/uuid found (vX.Y.Z), but does not contain package ... — проверьте правильность import и повторите go get/go mod tidy.
