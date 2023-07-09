package server

import "github.com/gammazero/workerpool"

type Pool struct {
	Volium int
	Pool   workerpool.WorkerPool
}

func New(volium int) *Pool {
	return &Pool{
		Volium: volium,
		Pool:   *workerpool.New(volium),
	}
}

func Submit(p *Pool, f func()) {
}
