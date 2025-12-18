package repo

import (
	"context"
	"database/sql"
	"example.com/pract16/internal/models"
)

type NoteRepo struct {
	DB *sql.DB
}

func (r *NoteRepo) Create(ctx context.Context, n *models.Note) error {
	return r.DB.QueryRowContext(ctx, "INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id", n.Title, n.Content).Scan(&n.ID)
}

func (r *NoteRepo) Get(ctx context.Context, id int64) (models.Note, error) {
	var n models.Note
	err := r.DB.QueryRowContext(ctx, "SELECT id, title, content FROM notes WHERE id = $1", id).Scan(&n.ID, &n.Title, &n.Content)
	return n, err
}

func (r *NoteRepo) Update(ctx context.Context, id int64, title, content string) error {
	_, err := r.DB.ExecContext(ctx, "UPDATE notes SET title = $1, content = $2 WHERE id = $3", title, content, id)
	return err
}

func (r *NoteRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM notes WHERE id = $1", id)
	return err
}

func (r *NoteRepo) List(ctx context.Context, limit, offset int64) ([]models.Note, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, title, content FROM notes LIMIT $1 OFFSET $2", limit, offset)
	if err != nil { return nil, err }
	defer rows.Close()
	var res []models.Note
	for rows.Next() {
		var n models.Note
		rows.Scan(&n.ID, &n.Title, &n.Content)
		res = append(res, n)
	}
	return res, nil
}

func (r *NoteRepo) ListAll(ctx context.Context) ([]models.Note, error) {
	return r.List(ctx, 100, 0)
}