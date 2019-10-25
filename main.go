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

}
