package array

import "fmt"

func TestArray() {
	var array1 [5]int
	var array2 [5]int = [5]int{10, 20, 30}
	array3 := [5]int{10, 20, 30, 40}
	array4 := [...]int{10, 20, 30, 40, 50}
	array5 := [...]int{1: 10, 2: 20}
	array6 := [5]*int{1: new(int), 2: new(int), 3: &array2[2]}
	*array6[1] = 1
	*array6[2] = 2
	fmt.Println(array1)                                     //[0 0 0 0 0]
	fmt.Println(array2)                                     //[10 20 30 0 0]
	fmt.Println(array3)                                     //[10 20 30 40 0]
	fmt.Println(array4)                                     //[10 20 30 40 50]
	fmt.Println(array5)                                     //[0 10 20]
	fmt.Println(array6, *array6[1], *array6[2], *array6[3]) //[<nil> 0xc000016098 0xc0000160a0 0xc00001e0a0 <nil>] 1 2 30
	*array6[3] = 300
	fmt.Println(array2)                                     //[10 20 300 0 0]
	fmt.Println(array6, *array6[1], *array6[2], *array6[3]) //[<nil> 0xc000016098 0xc0000160a0 0xc00001e0a0 <nil>] 1 2 300
}
