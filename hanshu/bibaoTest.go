package main

import (
	"fmt"
)

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

func getSequence() func() int {
	i := 0
	//fmt.Println(i)
	return func() int {
		i += 1
		return i
	}
}
