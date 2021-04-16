package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	//常数表达式可以执行任意精度的运算
	const n = 50000000
	const d = 3e20 / n

	//数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(d)
	fmt.Println(int64(d))

	//math.Sin函数需要一个 float64 的参数。
	fmt.Println(math.Sin(n))
}
