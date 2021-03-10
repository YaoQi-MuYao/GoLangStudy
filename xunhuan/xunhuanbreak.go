package main

import "fmt"

/* break 在循环中的使用 */
func main() {

	var a int = 10

	/* break 标准使用 */
	for a < 20 {
		fmt.Printf("a 的值为 ：%d\n", a)
		a++
		if a > 15 {
			/* 使用break 跳出循环 */
			break
		}
	}

	forBreak()
}

func forBreak() {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}
