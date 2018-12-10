package main

import "fmt"

func main() {
	defer function1() //多个defer的情况类似于压栈，function1最先入栈，最后执行；
	defer function2()
	defer function3()
	cc := aaf()
	fmt.Printf("%d \n", cc)
}
func function1() {
	fmt.Printf("test 11 11 1 1 1 1\n")
}
func function2() {
	fmt.Printf("test 22 2 2 2  2 2 \n")
}
func function3() {
	fmt.Printf("test 3 3 3 3  3 3  3 \n")
}

func aaf() int {
	i := 0
	defer fmt.Printf("%d \n", i) //defer 会先于return执行
	i++
	return i
}
