package main

import (
	"fmt"
	"time"
)

func main() {

	//Konsol Tarih işlemleri

	t := time.Date(2021, time.February, 03, 14, 30, 0, 0, time.UTC)

	fmt.Printf("Date at : %s\n", t)

	fmt.Println("------")

	now := time.Now()
	fmt.Printf("Time is : %s", now)

	fmt.Println("-----")

	fmt.Println("The Month is ", t.Month())
	fmt.Println("The Day is ", t.Day())
	fmt.Println("The weekday is", t.Weekday())

	fmt.Println("-----")

	//tarih ekleme
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	fmt.Println("------")

	longFormat := "Wendesday, February 3,2006"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))
	fmt.Println("------")

	shortFormat := "03/02/2021"
	fmt.Println("Today is ", tomorrow.Format(shortFormat))
}
