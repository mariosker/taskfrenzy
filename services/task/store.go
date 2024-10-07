package task

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mariosker/taskfrenzy/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) CreateTask(task types.Task) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO tasks (title, description) VALUES ($1, $2)", task.Title, task.Description)
	if err != nil {
		return err
	}

	return nil
}
