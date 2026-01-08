package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to file")

// Globaler Slice für Memory Leak Simulation
var leakyData [][]byte

// http://localhost:6060/debug/pprof
func main() {
	flag.Parse()

	// CPU Profiling starten
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Profiling Server auf Port 6060
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	go http.ListenAndServe(":6060", nil)

	// Simulierte Workload mit Performance-Problemen
	for i := 0; i < 100; i++ {
		cpuIntensiveTask()
		memoryLeakTask()
		blockingTask()
	}

	// Memory Profile schreiben
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

	fmt.Println("Fertig!")
}

// Problem 1: CPU-intensive Operation
func cpuIntensiveTask() {
	primes := findPrimes(100000)
	fmt.Printf("Gefunden: %d Primzahlen\n", len(primes))
}

func findPrimes(max int) []int {
	var primes []int
	for n := 2; n < max; n++ {
		isPrime := true
		for i := 2; i < n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, n)
		}
	}
	return primes
}

// Problem 2: Memory Leak
func memoryLeakTask() {
	data := make([]byte, 1024*1024) // 1 MB
	leakyData = append(leakyData, data)
}

// Problem 3: Blocking Operation
func blockingTask() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {
		// Simuliere langsame Verarbeitung
		// TODO
		time.Sleep(1 * time.Millisecond)
		// Option
		/*
			// Simuliere CPU-intensive Verarbeitung
			        result := 0
			        for j := 0; j < 10000; j++ {
			            result += val * j
			        }
			        _ = result
		*/
	}
}

/*
func blockingTask2() {
    ch := make(chan int) // unbuffered!
    done := make(chan bool)

    go func() {
        for i := 0; i < 1000; i++ {
            ch <- i // blockiert bis Empfänger liest
        }
        close(ch)
    }()

    go func() {
        for val := range ch {
            // Simuliere langsame Verarbeitung
            time.Sleep(100 * time.Microsecond)
            _ = val
        }
        done <- true
    }()

    <-done
}
*/
