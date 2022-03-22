package main

import (
	"fmt"

	"github.com/multipackage_1/package_2"
)

func main() {
	fmt.Println(package_2.OneIntPlusAnother(3, 4))
	stack := makeStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
