package pool

import (
	"sync/atomic"
)

//Счетчик
type count struct {
	count int32
}

//Увеличиваем число
func (c *count) Add() {
	atomic.AddInt32(&c.count, 1)
}

//Уменьшаем число
func (c *count) Done() {
	atomic.AddInt32(&c.count, -1)
}

//Получаем число
func (c *count) Get() int {
	return int(atomic.LoadInt32(&c.count))
}
