package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("minFloat:%f \n",math.MaxFloat32)
	fmt.Printf("maxFloat:%f \n",math.MaxFloat64)


	// 浮点数据默认fload64类型
	var num3 = 1.1
	fmt.Printf("num3 的数据类型%T", num3) // float64

	fmt.Println()
	num4 := .123 //==>0.123
	fmt.Println(num4)

	//科学计数法
	num5 := 5.1234e+2  //==>5.1234 * 10^2
	num5_2 := 5.1234e2 //==>5.1234 * 10^2
	num5_3 := 5.1234E2 //==>5.1234 * 10^2
	num6 := 5.1234e+2  //==>5.1234 * 10^2
	num7 := 5.1234e-2  //==>5.1234 / 10^2
	num8 := 5.1234E-2  //==>5.1234 / 10^2
	fmt.Println(num5, num5_2,num5_3, num6, num7,num8)

	
}
