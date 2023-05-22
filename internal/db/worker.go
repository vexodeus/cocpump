package db

import (
	"context"
	"database/sql"
)

type Worker struct {
	ID     int
	DB     *sql.DB
	Input  chan Task
	Output chan Task
}

func NewWorker(id int, db *sql.DB) *Worker {
	return &Worker{
		ID: id,
		DB: db,
	}
}
func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case task := <-w.Input:
			result := w.Process(task)
			w.Output <- result
		}
		select {
		case <-ctx.Done():
			break
		}
	}
}
func (w *Worker) Process(task Task) Task {
	switch task.Type {
	case int(TASK_TYPE_INSERT):
		//
	case int(TASK_TYPE_DELETE):
		//
	case int(TASK_TYPE_GET):
		//
	case int(TASK_TYPE_UPDATE):
		//
	}
	return task
}
