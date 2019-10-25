package main

import (
	"fmt"
)

func add(i, j int) int {
	return i + j
}

func main() {
	fmt.Println(add(1, 2))

	foo := func() string {
		return "Anonymous Func"
	}

	fmt.Println(foo())

	bar := func() {
		fmt.Println("Anonymous Func2")
	}

	bar()

	go func(i, j int) { //Go routine
		fmt.Println(i + j)
	}(1, 2)

}
