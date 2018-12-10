package main

import "fmt"

func main() {
	var min, max int
	min, max = minMax(78, 66)
	fmt.Printf("max is %d ,mim is %d", max, min)
	n := 0
	relay := &n
	change(60, 51, relay)
}
func minMax(a int, b int) (min int, max int) {
	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}
	return

}
func change(a int, b int, relay *int) {
	*relay = a * b

}
