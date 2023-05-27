package db

import (
	"context"
)

type Worker struct {
	ID     int
	Pool   *Pool
	Input  chan *Task
	Output chan *Task
}

func NewWorker(id int, pool *Pool) *Worker {
	return &Worker{
		ID:   id,
		Pool: pool,
	}
}
func (w *Worker) Start(ctx context.Context) {
	defer w.Pool.WG.Done()
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
			err = insertAccount(subject, w.Pool.DB)
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
			err = deleteAccountByID(subject.ID, w.Pool.DB)
			if err != nil {
				return nil, err
			}
			// err = deleteAccountByEmail(subject.Email, w.Pool.DB)
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
			data, err := getAccountsByID(subject.ID, w.Pool.DB)
			if err != nil {
				return nil, err
			}
			// data, err = getAccountsByEmail(subject.Email, w.Pool.DB)
			// if err != nil {
			// 	return nil, err
			// }
			task.Data = data
		}
		//
	//
	case TASK_TYPE_UPDATE:
		switch subject := task.Data.(type) {
		case Account:
			err = updateAccountByID(subject, w.Pool.DB)
			if err != nil {
				return nil, err
			}
			// err = updateAccountByEmail(subject, w.Pool.DB)
			// if err != nil {
			// 	return nil, err
			// }
		}
		//
	}
	return task, nil
}
