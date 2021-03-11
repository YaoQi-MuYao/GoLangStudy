package main

import "fmt"

/* 切片 len() 获取长度函数 ；cap() 获取容量函数 */
func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
