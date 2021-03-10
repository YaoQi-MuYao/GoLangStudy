package main

import "fmt"

/*
	GoLang的switch条件控制
	fallthrough 关键字可以继续执行下一个case
	不管下一个case是否符合标准
*/
func main() {

	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "E":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差")
	}

	fmt.Printf("您的等级 %s\n", grade)
}
