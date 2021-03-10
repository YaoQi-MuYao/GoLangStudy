package main

import "fmt"

func main() {

	/* 定义局部变量 */
	a := 100
	b := 200
	var ret int

	/* 调用函数，传入参数 */
	ret = max(a, b)
	fmt.Printf("最大值是 ：%d\n", ret)

	/* 调用函数 */
	s, s1 := swap("muyao", "wenruo")
	fmt.Println(s, s1)
}

/*
	自定义函数
*/
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

/*
	返回多个值
*/
func swap(x, y string) (string, string) {
	return x, y
}
