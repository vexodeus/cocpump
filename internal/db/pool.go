package db

import (
	"context"
	"database/sql"
	"sync"
)

type Pool struct {
	Name    string
	DB      *sql.DB
	WG      *sync.WaitGroup
	Workers []*Worker
	Inputs  chan *Task
}

func NewPool(name string, db *sql.DB) *Pool {
	return &Pool{
		Name: name,
		DB:   db,
		WG:   &sync.WaitGroup{},
	}
}
func (p *Pool) addWorkers(n int) {
	last := len(p.Workers) - 1
	for i := 1; i <= n; i++ {
		w := NewWorker(last+i, p)
		p.Workers = append(p.Workers, w)
	}
}

func (p *Pool) Start(ctx context.Context) {
	for _, w := range p.Workers {
		go w.Start(ctx)
	}
	p.WG.Wait()
}
