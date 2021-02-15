package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func initTask(tasks chan<- task, results chan<- int, num int) {
	ps := num / 10
	for i := 0; i < ps; i++ {
		b := 10*i + 1
		e := 10 * (i + 1)
		fmt.Println("--> ", b, e)
		t := task{
			begin:  b,
			end:    e,
			result: results,
		}
		fmt.Printf("t--> %+v\n", t)
		tasks <- t
		fmt.Printf("tsk--> %+v\n", tasks)
	}
	close(tasks)
}

func dispathTasks(tasks <-chan task, results chan<- int, wait *sync.WaitGroup) {
	for t := range tasks {
		wait.Add(1)
		fmt.Printf("1pocessTask task = %#v\n", t)
		fmt.Printf("2pocessTask task = %#v\n", &t)
		go pocessTask(&t, wait)
	}
	wait.Wait()
	close(results)
}

func pocessTask(t *task, wait *sync.WaitGroup) {
	fmt.Printf("3pocessTask task = %#v\n", t)
	// fmt.Printf("pocessTask task = %#v\n", &t)
	t.do()
	wait.Done()
}

func processResult(results <-chan int) int {
	sum := 0
	for v := range results {
		sum += v
	}
	return sum
}

func TestTask() {
	tasks := make(chan task, 10)
	results := make(chan int, 10)
	waitGroups := &sync.WaitGroup{}

	initTask(tasks, results, 100)
	dispathTasks(tasks, results, waitGroups)
	sum := processResult(results)

	fmt.Println("sum = ", sum)
}
