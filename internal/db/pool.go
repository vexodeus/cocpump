package db

import "database/sql"

type Pool struct {
	Name    string
	DB      *sql.DB
	Workers []*Worker
	Inputs  chan *Task
}

func NewPool(name string, db *sql.DB) *Pool {
	return &Pool{
		Name: name,
		DB:   db,
	}
}
func (p *Pool) addWorkers(n int) {
	last := len(p.Workers) - 1
	for i := 0; i < n; i++ {
		w := NewWorker(last+i, p)
		p.Workers = append(p.Workers, w)
	}
}
