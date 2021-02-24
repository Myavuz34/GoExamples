package main

import (
	"runtime"
	"time"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}

func main() {

	// go say("World")
	// say("Hello")
	// xFunc()

	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)
}
