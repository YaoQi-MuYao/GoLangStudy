package main

import "fmt"

/*
	GoLang 嵌套循环 求素数
*/
func main() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break //如果没有因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}