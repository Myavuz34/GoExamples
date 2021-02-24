package main

import (
	"fmt"
	"time"
)

func main() {

	// fmt.Printf("Şuanki Unix zamnı: %v\n", time.Now().Unix())
	// time.Sleep(2 * time.Second)
	// fmt.Printf("Şuanki Unix zamnı: %v\n", time.Now().Unix())

	// t := time.Date(2020, time.February, 24, 15, 0, 0, 0, time.UTC)
	// fmt.Println(t)

	// now := time.Now()
	// fmt.Println(now)

	// tomorrow := now.AddDate(0, 0, 1)
	// fmt.Println(tomorrow)

	x := fmt.Println

	now := time.Now()
	x(now)

	x(now.Year())
	x(now.Month())
	x(now.Day())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())
	x(now.Weekday())

	xTime := time.Date(1071, 10, 11, 20, 0, 0, 0, time.UTC)
	x(xTime)

	//Date Compare
	x(xTime.Before(now))
	x(xTime.After(now))
	x(xTime.Equal(now))

	//Date different

	diff := now.Sub(xTime)
	x(diff)
	x(diff.Hours())
	x(diff.Seconds())
}
