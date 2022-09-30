package main

import (
	"reflect"
)

// var x []int
// x = make([]int, 219_000_000_000)

func main() {
	x := 1
	reflect.ValueOf(x)
	_ = x
	y := x
	_ = y
}

// func square(x *int) *int {
// 	val := *x * *x
// 	return &val
// }
