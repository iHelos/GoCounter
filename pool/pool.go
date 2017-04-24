package pool

import (
	"sync"
)

//Пул заданий
//m - мьютек для ресайза
//wg - группа, учитывающая все процессы
//stop - канал, который слушают Воркеры, чтобы самоубиться
//task - канал, который слушают Воркеры, чтобы выполнить задание
//worker_count - количество сущестующих Воркеров
type pool struct {
	m  sync.Mutex
	wg sync.WaitGroup

	stop chan struct{}
	task chan Task

	worker_count int
}

//Создание нового пула
//size - исходный размер
//buf_size - размер буфера канала заданий
func NewPool(size int, buf_size int) *pool {
	p := &pool{
		stop: make(chan struct{}),
		task: make(chan Task, buf_size),
	}
	p.Resize(size)
	return p
}

//Получение текущего размера пула
func (p *pool) GetSize() int {
	return p.worker_count
}

//Изменение размера пула (добавление или уничтожение Воркеров)
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

//Жизненный цикл Воркера
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

//Отправить задание в пул
func (p *pool) SendTask(task Task) {
	p.task <- task
}

//Закрыть пул Воркеров
func (p *pool) Close() {
	close(p.task)
}
