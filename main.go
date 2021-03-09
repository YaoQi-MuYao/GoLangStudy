package main

import "fmt"

var (
	c int = 3
	d int = 4
)
func main() {
	var a int = 1
	var b int
	b = a
	b = 2
	println(a)
	println(b)
	println(c)
	println(d)
	fmt.Println(c, d)
}
