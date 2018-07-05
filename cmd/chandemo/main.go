package main

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
		fmt.Println("generate:", i)

	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		fmt.Println("filter:", i)
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for i := 0; i < 5; i++ {
		prime := <-ch
		fmt.Println(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	// ch := make(chan int)
	// go func() {
	// 	for i := 0; ; i++ {
	// 		ch <- i
	// 	}
	// 	//close(ch)
	// }()
	// //ch1 := ch

	// for {
	// 	ch1 := make(chan int)
	// 	ch = ch1
	// 	fmt.Println("ch:" + strconv.Itoa(<-ch))
	// 	fmt.Println("ch1:" + strconv.Itoa(<-ch1))
	// }
	// for v := range ch {
	// 	fmt.Println("ch:" + strconv.Itoa(v))
	// }
	// for v := range ch1 {
	// 	fmt.Println("ch1:" + strconv.Itoa(v))
	// }

	// for {
	// 	//ch1 := make(chan int)
	// 	//ch = ch1
	// 	fmt.Println("ch:" + strconv.Itoa(<-ch))
	// 	//fmt.Println("ch1:" + strconv.Itoa(<-ch))
	// }
}
