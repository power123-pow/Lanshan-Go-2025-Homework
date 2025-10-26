package main

/*
编写一个 Go 程序，实现以下功能：

1. 不断让用户输入整数（输入 `0` 表示结束输入）；
2. 使用 `for` 循环统计输入数字的总和与个数；
3. 定义一个函数 `average(sum int, count int) float64` 用于计算平均值；
4. 根据平均值输出结果：
   - 若平均值 ≥ 60，输出“平均成绩为 xx.xx，成绩合格”；
   - 否则输出“平均成绩为 xx.xx，成绩不合格”。*/
import (
	"fmt"
)

func average(sum int, count int) float64 {
	if count == 0 {
		return 0.0
	}
	return float64(sum) / float64(count)
}

func main() {
	var sum int
	var count int
	var num int
	for {
		fmt.Print("请输入一个整数(输入0结束):")
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
		count++
	}
	avg := average(sum, count)
	if count == 0 {
		fmt.Println("未输入有效值")
	} else {
		if avg >= 60 {
			fmt.Printf("平均成绩为%.2f，成绩合格\n", avg)
		} else {
			fmt.Printf("平均成绩为%.2f，成绩不合格\n", avg)
		}
	}
}
