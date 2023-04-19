package main

import "fmt"

func main() {
	//go1.18版本开始出现了泛型
	strs := []string{"giao1","giao2"}
	printArray2(strs)
}
//传统的方法 使用的是断言的转换
func printArray(arr []interface{}) {
	for _,a := range arr {
		fmt.Println(a)
	}
}

func printArray2[T string|int](arr []T) {
	for _,a := range arr {
		fmt.Println(a)
	}
}