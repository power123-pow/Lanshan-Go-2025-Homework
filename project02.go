package main

import "fmt"

// 使用for循环计算1到1000的和，并打印结果。
func main() {
	var sum = 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
