package result

import "sync"

type result struct {
	wg sync.WaitGroup
	result_channel *chan int
	sum int
}

func MakeResultWaiter(ch *chan int) (*result){
	r := &result{
		sync.WaitGroup{},
		ch,
		0,
	}
	go r.lifecycle()
	return r
}

func (r *result) WaitForUrl(){
	r.wg.Add(1)
}

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

func (r *result) GetResult() int{
	r.wg.Wait()
	return r.sum
}

func (r *result) Close() {
	close(*r.result_channel)
}
