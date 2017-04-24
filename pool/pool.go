package pool

import (
	"sync"
)

type pool struct {
	m  sync.Mutex
	wg sync.WaitGroup

	stop chan struct{}
	task chan Task

	worker_count int
}

func NewPool(size int, buf_size int) *pool {
	p := &pool{
		stop: make(chan struct{}),
		task: make(chan Task, buf_size),
	}
	p.Resize(size)
	return p
}

func (p *pool) GetSize() int {
	return p.worker_count
}

func (p *pool) Resize(new_size int) {
	p.m.Lock()
	defer p.m.Unlock()
	for p.worker_count < new_size {
		p.worker_count++
		p.wg.Add(1)
		go p.worker()
	}
	for p.worker_count > new_size {
		p.worker_count--
		p.stop <- struct{}{}
	}
}

func (p *pool) worker() {
	defer p.wg.Done()
	for {
		select {
		case new_task, ok := <-p.task:
			if !ok {
				return
			}
			//fmt.Println("executing")
			new_task.Execute()
		case <-p.stop:
			return
		}
	}
}

func (p *pool) SendTask(task Task) {
	p.task <- task
}

func (p *pool) Close() {
	close(p.task)
}

func (p *pool) Wait() {
	p.wg.Wait()
}
