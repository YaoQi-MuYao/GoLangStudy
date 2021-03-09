package main

import "fmt"
import "unsafe"

/*
	常量用作枚举类
 */
const (
	UNKNOWN = "wenruo"
	FEMALE = len(UNKNOWN)
	MALE = unsafe.Sizeof(UNKNOWN)
)

func main() {
	fmt.Println(UNKNOWN, FEMALE, MALE)
}
