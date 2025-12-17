package repo

import (
	"errors"
	"sync"
	"time"

	"example.com/notes-api/internal/core"
)

type NoteRepoMem struct {
	mu    sync.RWMutex
	data  map[int]core.Note
	seqID int
}

func NewNoteRepoMem() *NoteRepoMem {
	return &NoteRepoMem{
		data:  make(map[int]core.Note),
		seqID: 1,
	}
}

func (r *NoteRepoMem) Create(n core.Note) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	n.ID = r.seqID
	n.CreatedAt = time.Now()
	r.data[n.ID] = n
	r.seqID++
	return n.ID, nil
}

func (r *NoteRepoMem) List() ([]core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	notes := make([]core.Note, 0, len(r.data))
	for _, n := range r.data {
		notes = append(notes, n)
	}
	return notes, nil
}

func (r *NoteRepoMem) Get(id int) (core.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	n, ok := r.data[id]
	if !ok {
		return core.Note{}, errors.New("note not found")
	}
	return n, nil
}
