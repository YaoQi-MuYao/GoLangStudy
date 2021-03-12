package main

import "fmt"

/* GoLang 接口 */
/* 自定义类型 必须全部实现接口方法 才算实现了接口 */
type Phone interface {
	call()
	talk()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I'm nokia Phone")
}

func (nokiaPhone NokiaPhone) talk() {

}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I'm iphone Phone")
}

func main() {

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	//phone = new(IPhone)
	//phone.call()

}
