package main

import "fmt"

// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

	//FOR

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value :", i)
	// }

	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }

	// pow := make([]int, 10)
	// for i := range pow {
	// 	pow[i] = 1 << uint(i)
	// }
	// for _, value := range pow {
	// 	fmt.Printf("%d\n", value)
	// }

	// a := [...]string{"a", "b,", "c", "d"}
	// for i := range a {
	// 	fmt.Println("Array item, ", i, "is ", a[i])
	// }

	baskentler := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}
	for key, val := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", val)
	}

}
