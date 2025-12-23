package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerWithCtx(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancellation from context")
			return
		default:
			z := rand.Intn(100)
			userId := ctx.Value("userID").(string)
			fmt.Printf("Hello from %v zahl: %v \n", userId, z)
		}
	}
}

func worker2(wg *sync.WaitGroup, cancellation chan int) {
	defer wg.Done()

	for {
		select {
		case <-cancellation:
			fmt.Println("Cancelled")
			return
		default:
			z := rand.Intn(100)
			fmt.Println("Hello ", z)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	//cancellation := make(chan int)
	wg.Add(1)
	//go worker2(&wg, cancellation)

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", "1234")
	ctx, cancel := context.WithCancel(ctx)
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	go workerWithCtx(&wg, ctx)

	time.Sleep(10 * time.Second)
	//abbruch <- 1
	cancel()
	wg.Wait()
}
