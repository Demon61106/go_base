package main

import (
	"context"
	"fmt"
	"time"
)

func subSub(ctx context.Context) {
	fmt.Println("subsub work")
	select {
	case <-ctx.Done():
		fmt.Println("subsub get cancel sig, time = ", time.Now())
	}
}

func sub(ctx context.Context) {
	fmt.Println("sub work")
	go subSub(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("sub get cancel sig, time = ", time.Now())
	}
}

func ctxWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go sub(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("start to cancel")
	cancel()
	time.Sleep(3 * time.Second)
}

func ctxWithDeadline() {
	fmt.Println("ctxWithDeadline start ", time.Now())
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	go sub(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("ctxWithDeadline end ", time.Now())
}

func ctxTest() {
	// ctxWithCancel()
	ctxWithDeadline()
}
