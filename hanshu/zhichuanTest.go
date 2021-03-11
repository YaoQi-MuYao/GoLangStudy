package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf(" 交换前 a 的值为 ：%d\n", a)
	fmt.Printf(" 交换前 b 的值为 ：%d\n", b)

	swap1(a, b)

	fmt.Printf(" 交换后 a 的值为 ：%d\n", a)
	fmt.Printf(" 交换后 b 的值为 ：%d\n", b)
}

/*
	值传函数，不会改变参数，不会对实际参数有影响
*/
func swap1(x, y int) int {

	var temp int

	temp = x
	x = y
	y = temp

	return temp
}
