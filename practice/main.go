package main

import "fmt"

func main() {
	// If both functions have go the program exits
	// Why? because the main function has then finished
	// NOTE: This is bad practice
	// go count("sheep")
	// go count("fish")
	// time.Sleep(time.Second * 2)
	// fmt.Scanln()

	// NOTE: Useful when the goroutine just needs to perform an action
	// and we just need to wait for the action to be done
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// NOTE: Useful when the routines need to send data back
	// c := make(chan string)
	// go count("sheep", c)
	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// NOTE: Here the loop for c2 would be blocked by c1
	// so you would use a select
	// c1 := make(chan string)
	// c2 := make(chan string)
	// go count("sheep", c1)
	// for msg1 := range c1 {
	// 	fmt.Println(msg1)
	// }
	// go count("cow", c2)
	// for msg2 := range c2 {
	// 	fmt.Println(msg2)
	// }

	// c := make(chan string, 2)
	// c <- "hello"
	// c <- "world"
	//
	// msg := <-c
	// fmt.Println(msg)
	//
	// msg = <-c
	// fmt.Println(msg)

	// c1 := make(chan string)
	// c2 := make(chan string)
	//
	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()
	//
	// go func() {
	// 	for {
	// 		c2 <- "Every two seconds"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()
	//
	// for {
	// 	// Doing it like this the first reading c1 will block,
	// 	// the second reading of c2
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	//
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

// The <-chan and chan<- forces only read or write to the channels
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		// for i := 1; true; i++ {
// 		// fmt.Println(i, thing)
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
//
// 	close(c)
// }
