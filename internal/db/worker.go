package db

import (
	"context"
	"database/sql"
)

type Worker struct {
	ID     int
	DB     *sql.DB
	Input  chan *Task
	Output chan *Task
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
			result, _ := w.Process(task)
			w.Output <- result
		}
		select {
		case <-ctx.Done():
			break
		}
	}
}
func (w *Worker) Process(task *Task) (*Task, error) {
	var err error
	switch task.Type {
	//
	case TASK_TYPE_INSERT:
		switch subject := task.Data.(type) {
		//
		case Account:
			err = insertAccount(subject, w.DB)
			if err != nil {
				return nil, err
			}
		}
		//
	//
	case TASK_TYPE_DELETE:
		switch subject := task.Data.(type) {
		//
		case Account:
			err = deleteAccountByID(subject.ID, w.DB)
			if err != nil {
				return nil, err
			}
			// err = deleteAccountByEmail(subject.Email, w.DB)
			// if err != nil {
			// 	return nil, err
			// }
		}
		//
	//
	case TASK_TYPE_GET:
		switch subject := task.Data.(type) {
		//
		case Account:
			data, err := getAccountsByID(subject.ID, w.DB)
			if err != nil {
				return nil, err
			}
			// data, err = getAccountsByEmail(subject.Email, w.DB)
			// if err != nil {
			// 	return nil, err
			// }
			task.Data = data
		}
		//
	//
	case TASK_TYPE_UPDATE:
		//
	}
	return task, nil
}
