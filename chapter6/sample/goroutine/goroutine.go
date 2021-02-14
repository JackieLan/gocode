package goroutine

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGoroutine() {
	test1()
	test2()
}

func test1() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start goroutine!")
	//go printAlphabet('a', &wg)
	//go printAlphabet('A', &wg)
	go printPrime("A", &wg)
	go printPrime("B", &wg)
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Set runtime.GOMAXPROCS to ", runtime.NumCPU())
}

func printAlphabet(initByte byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for count := 0; count < 300; count++ {
		for char := initByte; char < initByte+26; char++ {
			fmt.Printf("%c ", char)
		}
	}
}

func printPrime(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	isPrime := true
	for outer := 900000; outer < 900100; outer++ {
		isPrime = true
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%s:%d  ", prefix, outer)
		}
	}
	fmt.Println("Complet3ed", prefix)
}

func test2() {
	var wg sync.WaitGroup
	court := make(chan int)
	wg.Add(2)
	go player("Nadal", court, &wg)
	go player("Djokovic", court, &wg)
	court <- 1
	wg.Wait()
	// Player Djokovic Hit 1
	// Player Nadal Hit 2
	// Player Djokovic Hit 3
	// Player Nadal Hit 4
	// Player Djokovic Hit 5
	// Player Nadal Hit 6
	// Player Djokovic Hit 7
	// Player Nadal Hit 8
	// Player Djokovic Missed
	// Player Nadal Won
}

func player(name string, court chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
