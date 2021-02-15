package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func doLogin(writer http.ResponseWriter, req *http.Request) {
	_, err := writer.Write([]byte("do login"))
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
}

func HttpServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", doLogin)
	server := &http.Server{
		Addr:         ":8081",
		WriteTimeout: time.Second * 2,
		Handler:      mux,
	}
	log.Fatal(server.ListenAndServe())
}

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int, 10)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateIntA(send):
			case ch <- <-GenerateIntB(send):
			case <-done:
				send <- struct{}{}
				send <- struct{}{}
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func testChan() {
	done := make(chan struct{})

	ch := GenerateInt(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	done <- struct{}{}
	fmt.Println("stop gen")
}

/*
管道使用
*/
func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v + 10
		}
		close(out)
	}()
	return out
}

func testChain() {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := chain(chain(chain(in)))
	out2 := out
	fmt.Printf("out = %+v\n", out)
	fmt.Printf("out2 = %+v\n", out2)
	for v := range out {
		fmt.Println("out v = ", v)
	}
}
func main() {
	// HttpServer()
	// testChan()
	// testChain()
	// TestTask()
	// ctxTest()
	// relfectTest()
	// InjectTest()
	// DeferTest()
	// ArrayTest()
	// FuncTest()
	// RunTimeTest()
	ChanTest()
}
