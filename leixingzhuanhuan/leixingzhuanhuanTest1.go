package main

import "fmt"

/* GoLang 类型转换 PS：不支持隐式转换 */
func main() {

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Println(mean)
}
