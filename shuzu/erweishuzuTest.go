package main

import "fmt"

func main() {

	/* 创建数组 */
	values := [][]int{}

	/* 使用append() 函数向空的二维数组添加两行一为数组 */
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	/* 显示两行数据 */
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	/* 访问第一个元素 */
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])

}
