package main

import "fmt"

/* 定义结构体 */
type Circle1 struct {
	radius float64
}

func main() {
	var c Circle1
	fmt.Println(c.radius)
	c.radius = 10.00
	fmt.Println(c.getArea())
	c.changeRadius(20)
	fmt.Println(c.radius)
	change(&c, 30)
	fmt.Println(c.radius)
}

/* 函数方法 */
func (c Circle1) getArea() float64 {
	return c.radius * c.radius
}

/* 函数方法 通过指针修改c的值 */
func (c *Circle1) changeRadius(radius float64) {
	c.radius = radius
}

/* 普通函数 通过传递指针修改c的值 */
func change(c *Circle1, radius float64) {
	c.radius = radius
}
