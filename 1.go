package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/iHelos/GoCounter/pool"
	"github.com/iHelos/GoCounter/result"
	"github.com/iHelos/GoCounter/task"
	"os"
	"regexp"
)

func main() {
	//Максимальное количество горутин-воркеров
	k := flag.Int("k", 5, "max goroutines count")
	//Максимальный размер буфера тасков
	b := flag.Int("b", 128, "max buffer tasks")

	flag.Parse()

	p := pool.NewPool(0, *b)
	defer p.Close()

	results := make(chan int, *b)
	result_waiter := result.MakeResultWaiter(&results)
	defer result_waiter.Close()

	rp, err := regexp.Compile("Go")
	if err != nil {
		fmt.Printf("Bad regexp\n%s", err.Error())
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		size := p.GetSize()
		if size < *k {
			p.Resize(size + 1)
		}
		p.SendTask(&task.Task{line, &results, rp})
		result_waiter.WaitForUrl()
	}
	fmt.Printf("Total: %d\n", result_waiter.GetResult())
}
