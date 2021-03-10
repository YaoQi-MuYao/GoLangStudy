package main

import "fmt"

func main() {
	var a = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - a 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - a 变量类型为 = %T\n", c)

	/* &和* 运算符实例 */
	ptr = &a
	fmt.Printf("a 的地址为  %d\n", &a)
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("ptr的值为 %d\n", ptr)
	fmt.Printf("ptr所指的值为 %d\n", *ptr)
}
