package main

import "fmt"

func main() {

	/* 初始化二维数组 */
	a := [3][4]int{
		{0, 1, 2, 3},
		{3, 4, 5, 6},
		{7, 8, 9, 10},
	}
	sites := [2][2]string{}

	fmt.Println(a)
	/* 像二维数组添加元素 */
	sites[0][0] = "google"
	sites[0][1] = "Runoob"
	sites[1][0] = "taobao"
	sites[1][1] = "weibo"

	fmt.Println(sites)
}
