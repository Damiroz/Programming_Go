package httpapi

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func BuildRouter(d *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	h := NewHandlers(d) // Создаем обработчики, передавая DB-соединение

	r.Get("/health", h.Health) // Проверка здоровья

	// Пользователи
	r.Post("/users", h.CreateUser)

	// Заметки
	r.Post("/notes", h.CreateNote)      // Создание заметки с тегами
	r.Get("/notes/{id}", h.GetNoteByID) // Получение заметки с автором и тегами

	return r
}