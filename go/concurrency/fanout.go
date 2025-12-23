package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    wg := sync.WaitGroup{}
    jobs := make(chan int)

    wg.Add(3)
    go worker(1, jobs, &wg)
    go worker(2, jobs, &wg)
    go worker(3, jobs, &wg)

    for i := 0; i < 10; i++ {
        jobs <- i
    }

    close(jobs)
    wg.Wait() // without it, the last job probably won't run
}

func worker(workerId int, jobs <-chan int, wg *sync.WaitGroup) {
    for job := range jobs {
        fmt.Printf("Worker %v started, job %v\n", workerId, job)
        time.Sleep(time.Second * 3)
        fmt.Printf("Worker %v ended, job %v\n", workerId, job)
    }

    wg.Done()
}
