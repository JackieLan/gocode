package slice

import "fmt"

func TestSlice() {
	slice1 := make([]string, 5)
	fmt.Println(slice1, len(slice1), cap(slice1)) // [    ] 5 5
	slice2 := make([]string, 3, 5)
	fmt.Println(slice2, len(slice2), cap(slice2)) // [  ] 3 5
	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	printSlice(slice3)      // [Red Blue Green Yellow Pink] 5 5
	printSlice(slice3[1:3]) // [Blue Green] 2 4
	var slice4 []int
	printSlice2(slice4)        // [] 0 0, nil slice
	fmt.Println(slice4 == nil) // true
	slice5 := make([]int, 0)   // same with slice5 := []int{}
	printSlice2(slice5)        // [] 0 0, empty slice
	fmt.Println(slice5 == nil) // false
	slice5 = append(slice5, 0)
	printSlice2(slice5) // [0] 1 1
	slice5 = append(slice5, 1)
	printSlice2(slice5) // [0 1] 2 2
	slice5 = append(slice5, 2)
	printSlice2(slice5) // [0 1 2] 3 4
	slice5 = append(slice5, 3)
	printSlice2(slice5) // [0 1 2 3] 4 4
	slice5 = append(slice5, 4)
	printSlice2(slice5) // [0 1 2 3 4] 5 8
	slice5 = append(slice5, 5)
	printSlice2(slice5) // [0 1 2 3 4 5] 6 8
	// slice5[7] = 7 // panic: runtime error: index out of range [7] with length 6
	source := []string{"A0", "B1", "C2", "D3", "E4", "F5"}
	slice6 := source[2:3:4]
	printSlice(slice6) // [C2] 1 2
	slice6 = append(slice6, "newD3")
	slice6 = append(slice6, "newE4")
	printSlice(slice6)                                  // [C2 newD3 newE4] 3 4
	printSlice(source)                                  // [A0 B1 C2 newD3 E4 F5] 6 6
	printSlice2(append([]int{1, 2}, []int{3, 4, 5}...)) // append(s1, s2...)  [1 2 3 4 5] 5 6
	testSliceRange()
	testMultiDimensionSlice()
	testPassSliceInFunc()
}

func printSlice(slice []string) {
	fmt.Println(slice, len(slice), cap(slice))
}

func printSlice2(slice []int) {
	fmt.Println(slice, len(slice), cap(slice))
}

func testSliceRange() {
	slice := []int{0, 10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Value: %d, Value-Addr: %X, ElemAddr: %X\n", value, &value, &slice[index])
	}
	//Value: 0, Value-Addr: C0000161C0, ElemAddr: C00001E1E0
	//Value: 10, Value-Addr: C0000161C0, ElemAddr: C00001E1E8
	//Value: 20, Value-Addr: C0000161C0, ElemAddr: C00001E1F0
	//Value: 30, Value-Addr: C0000161C0, ElemAddr: C00001E1F8
	//Value: 40, Value-Addr: C0000161C0, ElemAddr: C00001E200
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d, Value: %d\n", index, slice[index])
	}
	//Index: 2, Value: 20
	//Index: 3, Value: 30
	//Index: 4, Value: 40
}

func testMultiDimensionSlice() {
	slice := [][]int{{10}, {100, 200}}
	fmt.Printf("%v %v\n", slice[0], slice[1]) //[10] [100 200]
	slice[0] = append(slice[0], 20)
	fmt.Printf("%v %v\n", slice[0], slice[1]) //[10 20] [100 200]
}

func testPassSliceInFunc() {
	slice := make([]int, 1e6)
	slice = foo(slice)
	fmt.Println(slice[:10]) // [100 0 0 0 0 0 0 0 0 0]
	bar(slice)
	fmt.Println(slice[:10]) // [200 0 0 0 0 0 0 0 0 0]
	slice2 := make([]*int, 1e6)
	bar2(slice2)
	fmt.Println(slice2[:10]) // [0xc00008c090 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]
}

func foo(slice []int) []int { // only 24 bytes are needed for a slice
	slice[0] = 100
	return slice
}
func bar(slice []int) {
	slice[0] = 200
}

func bar2(slice []*int) {
	slice[0] = new(int)
	*slice[0] = 100
}
