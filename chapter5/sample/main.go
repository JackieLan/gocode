package main

import (
	"fmt"
	"github.com/JackieLan/gocode/chapter5/sample/interface2"
	"github.com/JackieLan/gocode/chapter5/sample/method"
	"github.com/JackieLan/gocode/chapter5/sample/type2"
)

func main() {
	type2.TestType()
	method.TestMethod()
	interface2.TestInterface()
	lisa := interface2.NewUser("Lisa", "lisa@example.com")
	fmt.Println(lisa) // {Lisa lisa@example.com}
}
