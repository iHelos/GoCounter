package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/iHelos/GoCounter/pool"
	"github.com/iHelos/GoCounter/task"
	"os"
)

func main() {
	k := flag.Int("k", 5, "max goroutines count")
	b := flag.Int("b", 128, "max buffer tasks")

	flag.Parse()

	p := pool.NewPool(1, *b)
	defer p.Close()

	results := make(chan int, *b)

	url_count := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if p.GetSize() < *k {
			p.Resize(p.GetSize() + 1)
		}
		p.SendTask(&task.Task{line, &results})
		url_count++
	}

	sum := 0
	for i := 0; i < url_count; i++ {
		sum += <-results
	}
	fmt.Printf("Total: %d\n", sum)
}
