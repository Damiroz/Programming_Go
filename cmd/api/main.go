package main

//go:generate swag init -g cmd/api/main.go -o docs
import (
	"fmt"
	"net/http"
	"github.com/swaggo/http-swagger"
	_ "example.com/notes-api/docs" // ОБЯЗАТЕЛЬНО: Импорт сгенерированных доков
	internalHttp "example.com/notes-api/internal/http"
	
)

// @title           Notes API
// @version         1.0
// @description     Учебный REST API для заметок (CRUD).
// @contact.name    Backend Course
// @contact.email   example@university.ru
// @BasePath        /

func main() {
	r := internalHttp.NewRouter()

	fmt.Println("Server starting on :8080...")
	fmt.Println("Swagger UI: http://localhost:8080/docs/index.html")
	r.Get("/docs/*", httpSwagger.WrapHandler) // для chi: r.Get

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
