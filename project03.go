package main

//编写一个函数，接受一个整数参数n，返回n的阶乘（factorial）。在main函数中调用该函数并打印结果。
import "fmt"
//编写一个函数，接受一个整数参数n，返回n的阶乘（factorial）。在main函数中调用该函数并打印结果
func add(n int) int {
	if n == 0 {
		return 1
	}
	return n * add(n-1)
}

func main() {
	fmt.Println(add(7))
}
