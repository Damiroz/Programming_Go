package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stubRepo — тестовая реализация (стаб) репозитория
type stubRepo struct {
	users map[string]User
}

func (r stubRepo) ByEmail(email string) (User, error) {
	u, ok := r.users[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

func TestService_FindIDByEmail(t *testing.T) {
	// Подготовка данных
	repo := stubRepo{
		users: map[string]User{
			"test@example.com": {ID: 101, Email: "test@example.com"},
		},
	}
	srv := New(repo)

	t.Run("user found", func(t *testing.T) {
		id, err := srv.FindIDByEmail("test@example.com")
		require.NoError(t, err)
		assert.Equal(t, int64(101), id)
	})

	t.Run("user not found", func(t *testing.T) {
		_, err := srv.FindIDByEmail("unknown@mail.com")
		assert.ErrorIs(t, err, ErrNotFound)
	})
}