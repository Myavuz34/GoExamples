package main

import (
	"runtime"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum

}

func xFunc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string) {
	for {
		println(<-ch)
	}
}
func main() {

	// s := []int{1, 2, 3, 4, 5, 6}

	// c := make(chan int)
	// go sum(s[:len(s)], c)
	// go sum(s[:len(s)/2], c)
	// x, y := <-c, <-c

	// fmt.Println(x, y)

	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	go xFunc(ch)
	go printer(ch)
	time.Sleep(100 * time.Millisecond)

}
