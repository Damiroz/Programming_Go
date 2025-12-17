package core

import "time"

// Note - основная модель заметки
type Note struct {
	ID        int       `json:"id" example:"1"`
	Title     string    `json:"title" example:"Список покупок"`
	Content   string    `json:"content" example:"Молоко, хлеб, яйца"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-01T12:00:00Z"`
}

// NoteCreate - DTO для создания заметки
// Используется в POST запросах
type NoteCreate struct {
	Title   string `json:"title" example:"Новая заметка"`
	Content string `json:"content" example:"Текст заметки"`
}

// NoteUpdate - DTO для обновления заметки
// Используется в PATCH запросах (поля опциональны)
type NoteUpdate struct {
	Title   *string `json:"title,omitempty" example:"Обновленный заголовок"`
	Content *string `json:"content,omitempty" example:"Обновленный текст"`
}