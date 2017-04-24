package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/iHelos/GoCounter/pool"
	"github.com/iHelos/GoCounter/task"
	"github.com/iHelos/GoCounter/result"
	"os"
	"regexp"
)

func main() {
	k := flag.Int("k", 5, "max goroutines count")
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
		if p.GetSize() < *k {
			p.Resize(p.GetSize() + 1)
		}
		p.SendTask(&task.Task{line, &results, rp})
		result_waiter.WaitForUrl()
	}
	fmt.Printf("Total: %d\n", result_waiter.GetResult())
}
