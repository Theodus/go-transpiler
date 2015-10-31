package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Name  string
	Right int
	Left  int
}

func (p *Philosopher) Eat(t []sync.Mutex) {
	t[p.Left].Lock()
	defer t[p.Left].Unlock()
	t[p.Right].Lock()
	defer t[p.Right].Unlock()
	fmt.Println(p.Name, "is eating.")
	time.Sleep(time.Second)
	fmt.Println(p.Name, "is done eating.")
}

func main() {
	phils := []*Philosopher{
		&Philosopher{"Judith Butler", 0, 1},
		&Philosopher{"Grilles Deleuze", 1, 2},
		&Philosopher{"Karl Marx", 2, 3},
		&Philosopher{"Emma Goldman", 3, 4},
		&Philosopher{"Michel Faucault", 4, 0},
	}
	table := make([]sync.Mutex, len(phils))
	var wg sync.WaitGroup
	for _, p := range phils {
		wg.Add(1)
		go func(p *Philosopher) {
			defer wg.Done()
			p.Eat(table)
		}(p)
	}
	wg.Wait()
}
