package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"example.com/notes-api/internal/http/handlers"
	"example.com/notes-api/internal/repo"
	// _ "example.com/notes-api/docs" // Импорт доков обычно делается в main.go, но можно и тут
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Инициализация зависимостей
	repo := repo.NewNoteRepoMem()
	h := handlers.NewHandler(repo)

	// Роуты API
	r.Route("/notes", func(r chi.Router) {
		r.Get("/", h.ListNotes)
		r.Post("/", h.CreateNote)
		r.Get("/{id}", h.GetNote)
		// Добавьте Patch и Delete при необходимости
	})

	// Роут для документации Swagger
	// URL: http://localhost:8080/docs/index.html
	r.Get("/docs/*", httpSwagger.WrapHandler)

	return r
}