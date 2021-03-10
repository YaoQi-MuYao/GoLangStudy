package main

import "fmt"

func main() {
	/* 定义局部变量 */
	a := 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过本次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为：%d\n", a)
		a++
	}

	label()
}

func label() {
	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re
		}
	}
}
