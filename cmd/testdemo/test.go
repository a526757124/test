package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/**
 *FIF0,
 */

func sample1(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- "I'm samplel num:" + strconv.Itoa(i)
		time.Sleep(1 * time.Second)
	}
}
func sample2(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}

func abc() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type test1 struct {
	_ int
	s string
}

func main() {
	abc()

	ch1 := make(chan string, 3)
	ch2 := make(chan int, 5)

	for i := 0; i < 10; i++ {
		go sample1(ch1)
		go sample2(ch2)
	}

	time.Sleep(10 * time.Second)

	var a Integer = 1
	var b LessAdder = &a
	b.Add(5)
	fmt.Println(a)
	fmt.Println(b.Less(7))

}

//定义一个接口，接口内有eat和run两个方法
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
