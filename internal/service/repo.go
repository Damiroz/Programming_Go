package service

import "fmt"

type User struct {
	ID    int64
	Email string
}

var ErrNotFound = fmt.Errorf("not found")

// UserRepo — интерфейс репозитория для изоляции зависимостей
type UserRepo interface {
	ByEmail(email string) (User, error)
}