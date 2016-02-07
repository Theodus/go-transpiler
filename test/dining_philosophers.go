package main

import (
	"fmt"
	"sync"
	"time"
)

type philosopher struct {
	name  string
	right int
	left  int
}

func (p *philosopher) eat(t []sync.Mutex) {
	t[p.left].Lock()
	defer t[p.left].Unlock()
	t[p.right].Lock()
	defer t[p.right].Unlock()
	fmt.Println(p.name, "is eating.")
	time.Sleep(time.Second)
	fmt.Println(p.name, "is done eating.")
}

func main() {
	phils := []*philosopher{
		&philosopher{"Judith Butler", 0, 1},
		&philosopher{"Grilles Deleuze", 1, 2},
		&philosopher{"Karl Marx", 2, 3},
		&philosopher{"Emma Goldman", 3, 4},
		&philosopher{"Michel Faucault", 4, 0},
	}
	table := make([]sync.Mutex, len(phils))
	var wg sync.WaitGroup
	for _, p := range phils {
		wg.Add(1)
		go func(p *philosopher) {
			defer wg.Done()
			p.eat(table)
		}(p)
	}
	wg.Wait()
}
