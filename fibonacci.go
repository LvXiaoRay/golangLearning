package main

import "fmt"

func main() {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("当 i 为 %d 时，结果为 %d \n", i, result)
	}
	printff(10)
	casee := scase(86)
	fmt.Printf(casee)
}
func fibonacci(i int) (res int) {
	if i <= 0 {
		res = 1
	} else {
		res = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}
func printff(i int) {
	fmt.Printf("%d\n", i)
	if i > 1 {
		printff(i - 1)
	}
	//panic("数据范围不合法");非法操作提示

}
func switchcase(i int) string {
	switch {
	case i < 60:

	}
	return "a"
}
func scase(i int) string {
	switch { //switch 有两种写法；这种是 switch true 写法；
	case i < 60:
		fmt.Printf("pass60")
		return "D"
	case i < 70:
		fmt.Printf("pass60")
		return "C"
	case i < 80:
		return "B"
	case i < 90:
		return "A"
	default:
		return "Error"

	}

}
