package main

import "fmt"

/* 结构体 */
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {

	/* 创建一个新的结构体 */
	fmt.Println(Book{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	/* 也可以使用 key => value 的形式 */
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})

	/* 忽略的字段为0或空 */
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})
}
