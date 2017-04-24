package result

import "sync"

//Структура, подсчитывающая результат
type result struct {
	wg sync.WaitGroup
	result_channel *chan int
	sum int
}

//Конструктор, в который передается канал, в который поступают результаты
func MakeResultWaiter(ch *chan int) (*result){
	r := &result{
		sync.WaitGroup{},
		ch,
		0,
	}
	go r.lifecycle()
	return r
}

//Говорим, чтобы ждал на один урл больше
func (r *result) WaitForUrl(){
	r.wg.Add(1)
}

//Запускается в конструкторе, суммирует все вхождения
func (r *result) lifecycle(){
	for {
		select {
			case result, ok := <- *(r.result_channel):
				if !ok {
					return
				} else {
					r.sum += result
					r.wg.Done()
				}
		}
	}
}

//Ждет все оставшиеся урлы и возвращает сумму
func (r *result) GetResult() int{
	r.wg.Wait()
	return r.sum
}

//Убивает процесс и закрывает канал
func (r *result) Close() {
	close(*r.result_channel)
}
