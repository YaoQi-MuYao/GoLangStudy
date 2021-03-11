package main

import "fmt"

func main() {

	/* 创建二维数组 */
	animals := [][]string{}

	/* 创建三个一维数组，各数组长度不同 */
	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	/* 使用append()函数将一维数组添加到二维数组中 */
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	/* 循环输出 */
	for i := range animals {
		fmt.Printf("Row: %v\n", 1)
		fmt.Println(animals[i])
	}
}
