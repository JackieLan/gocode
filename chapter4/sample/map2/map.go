package map2

import "fmt"

func TestMap() {
	dict1 := make(map[string]int)
	printMap(dict1) // map[] 0
	dict1["E"] = 500
	printMap(dict1) // map[E:500] 1
	dict2 := map[string]int{"A": 100, "B": 200, "C": 300, "D": 400}
	printMap(dict2) // map[A:100 B:200 C:300 D:400] 4
	dict2["E"] = 500
	printMap(dict2) // map[A:100 B:200 C:300 D:400 E:500] 5
	//dict3 := map[int][]string{}
	var dict4 map[string]int
	// dict4["E"] = 500 // panic: assignment to entry in nil map
	dict4 = map[string]int{}
	dict4["E"] = 500
	valueE, exists := dict4["E"]
	fmt.Println(valueE, exists) // 500 true
	valuee, exists := dict4["e"]
	fmt.Println(valuee, exists)     // 0 false
	for key, value := range dict4 { // E 500
		fmt.Println(key, value)
	}
	delete(dict4, "E")
	delete(dict4, "F")
}

func printMap(dict map[string]int) {
	fmt.Println(dict, len(dict))
}
