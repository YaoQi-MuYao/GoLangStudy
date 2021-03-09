package main

import "fmt"

/*
	iota初始值为0
	i 1左移0 1 为1
	j 3左移1 110 为6
	k 3左移2 1100 为12
	l 3左移3 11000 为24
*/
const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
	fmt.Println("k = ", k)
	fmt.Println("l = ", l)
}
