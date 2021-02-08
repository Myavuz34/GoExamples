package main

import "fmt"

func main() {

	// myMap := make(map[int]string)
	// fmt.Println(myMap)
	// myMap[43] = "Foo"
	// myMap[11] = "Bar"
	// fmt.Println(myMap)

	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"

	antalya := states["ANT"]
	fmt.Println(antalya)
	delete(states, "ANK")
	fmt.Println(states)

	states["RZ"] = "RİZE"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}
