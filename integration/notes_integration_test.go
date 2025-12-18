package integration

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"example.com/pract16/internal/db"
	"example.com/pract16/internal/httpapi"
	"example.com/pract16/internal/repo"
	"example.com/pract16/internal/service"
)

func TestNoteFullFlow(t *testing.T) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" { t.Skip("DB_DSN not set") }

	database, _ := sql.Open("postgres", dsn)
	db.MustApplyMigrations(database)

	repository := repo.NoteRepo{DB: database}
	srv := service.Service{Notes: repository}
	h := &httpapi.Handler{Service: srv}

	r := gin.Default()
	r.POST("/notes", h.CreateNote)
	r.GET("/notes/:id", h.GetNote)
	r.PUT("/notes/:id", h.UpdateNote)
	r.DELETE("/notes/:id", h.DeleteNote)
	r.GET("/notes", h.ListNotes)

	t.Run("1. Get Non-existent 404", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/notes/9999", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	var createdID int64
	t.Run("2. Create and Update", func(t *testing.T) {
		// Create
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/notes", bytes.NewBufferString(`{"title":"Initial","content":"Text"}`))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		// Нам нужно вытащить ID для следующих шагов (упрощенно возьмем 1, если БД чистая)
		createdID = 1 

		// Update (PUT)
		w = httptest.NewRecorder()
		req, _ = http.NewRequest("PUT", fmt.Sprintf("/notes/%d", createdID), bytes.NewBufferString(`{"title":"Updated","content":"New"}`))
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("3. List with Pagination", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/notes?limit=1&offset=0", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("4. Delete", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/notes/%d", createdID), nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}