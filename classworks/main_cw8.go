package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

}

func main2() {
	c := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		<-c
		fmt.Println("Hello go 2 after cg")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		<-c
		fmt.Println("Hello go 3 after cg")
		wg.Done()
	}()

	go func() {
		c <- struct{}{}
		c <- struct{}{}
	}()

	fmt.Println("Hello")

	wg.Wait()

	fmt.Println("Hello2")

}

func main1() {
	ctxBg := context.Background()
	ctxValue := context.WithValue(ctxBg, "hello", "world")
	ctx, cancel := context.WithCancel(ctxValue)
	ctxValueAndCancel := context.WithValue(ctx, "bye", "cancel")

	go printHelloAndBye(ctxValueAndCancel)
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("END")

}

func printHelloAndBye(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context is done")
			return
		default:
			fmt.Println(ctx.Value("hello"), ctx.Value("bye"))
		}
	}
}
