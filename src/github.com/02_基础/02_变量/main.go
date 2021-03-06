package main

import "fmt"

func variableZeroValue() {
	var a int
	var b string
	var c []float64
	var d func() bool
	var e struct{
		name string
		age int
	}
	fmt.Println(a,b,c,d,e) // 0 "" [] <nil> { 0}
}

func variableInitialValue() {
	var (
		a int
		b string
		c []float64
		d func() bool
		e struct{
			name string
			age int
		}
	)
	fmt.Println(a,b,c,d,e) // // 0 "" [] <nil> { 0}
}

func variableTypeDeduction() {
	var num, name, gender = 10, "jack", true
	fmt.Println(num, name, gender) // 10 jack true
}

func variableShorter() {
	// 短变量
	num, name, gender := 10, "jack", true
	fmt.Println(num, name, gender) // 10 jack true
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
