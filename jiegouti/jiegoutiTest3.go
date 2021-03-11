package main

import "fmt"

type Book1 struct {
	title   string
	author  string
	subject string
	book_id int
}

/* 结构体作为函数参数 */
func main() {
	var book1 Book1 /* 声明 Book1 为 Books 类型 */
	var book2 Book1 /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	book1.title = "Go 语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go 语言教程"
	book1.book_id = 6495407

	/* book 2 描述 */
	book2.title = "Python 教程"
	book2.author = "www.runoob.com"
	book2.subject = "Python 语言教程"
	book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(book1)

	/* 打印 Book2 信息 */
	printBook(book2)
}

func printBook(book Book1) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
