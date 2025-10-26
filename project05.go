package main

//定义一个常量 pi = 3.14，再定义变量 r = 5，计算圆的面积（area = pi * r * r）并打印结果。
import "fmt"

func main() {
	const pi = 3.14
	const r = 5
	area := pi * r * r
	fmt.Printf("area=%.2f", area)
}
