package main

import "fmt"

func main() {
	a := 100
	b := 200

	fmt.Printf("交换前 a 的值 ：%d\n", a)
	fmt.Printf("交换前 b 的值 ：%d\n", b)
	fmt.Println(&a, &b)
	swap2(&a, &b)
	fmt.Printf("交换后 a 的值 ：%d\n", a)
	fmt.Printf("交换后 b 的值 ：%d\n", b)
}

/*
	引用传递，会改变参数的值
*/
func swap2(x *int, y *int) {
	var temp int

	fmt.Println(x, *x, y, *y)

	temp = *x
	*x = *y
	*y = temp
}
