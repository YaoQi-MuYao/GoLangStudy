package main

import "fmt"

func main() {

	sum := 1

	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for1()
	for2()
}

func for1() {

	/* 第二种形式 */
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	/* 第三种形式 */
	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}

/*
	foreach 循环结构
*/
func for2() {

	strings := []string{"muyao", "wenruo"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/*
	无限循环
*/
func for3() {
	for true {
		fmt.Printf("这是无限循环 \n")
	}
}
