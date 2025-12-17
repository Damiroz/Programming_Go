package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"example.com/notes-api/internal/core"
	"example.com/notes-api/internal/repo"
)

type Handler struct {
	Repo *repo.NoteRepoMem
}

func NewHandler(repo *repo.NoteRepoMem) *Handler {
	return &Handler{Repo: repo}
}

// ListNotes godoc
// @Summary      Список заметок
// @Description  Возвращает список заметок
// @Tags         notes
// @Param        page   query  int     false  "Номер страницы"
// @Param        limit  query  int     false  "Размер страницы"
// @Param        q      query  string  false  "Поиск по title"
// @Success      200    {array}  core.Note
// @Header       200    {integer}  X-Total-Count  "Общее количество"
// @Failure      500    {object}  map[string]string
// @Router       /notes [get]
func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.Repo.List()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// CreateNote godoc
// @Summary      Создать заметку
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        input  body     core.NoteCreate  true  "Данные новой заметки"
// @Success      201    {object} core.Note
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /notes [post]
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var input core.NoteCreate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	note := core.Note{
		Title:   input.Title,
		Content: input.Content,
	}

	id, err := h.Repo.Create(note)
	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}
	note.ID = id
	// В реальном коде CreatedAt ставится в репозитории, здесь для ответа заглушка:
	// note.CreatedAt = time.Now() 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// GetNote godoc
// @Summary      Получить заметку
// @Tags         notes
// @Param        id   path   int  true  "ID"
// @Success      200  {object}  core.Note
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [get]
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	note, err := h.Repo.Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Note not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}