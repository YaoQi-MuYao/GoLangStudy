package main

import "fmt"

/* 空指针 */
func main() {
	var ptr *int

	fmt.Printf("ptr 的值为：%x\n", ptr)

	/* 判断是否为空指针 */
	if nil == ptr {
		fmt.Println("是空指针")
	}
}
