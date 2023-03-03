package main

import "fmt"

func init() {
	fmt.Println("init函数先调用")
}

func main() {
	fmt.Println("main函数后调用")
	//var num1 int = 99

	//var num2 float64 = 99.99
	var b bool = true
	//var myChar byte = 'h'
	var str string

	//	使用fmt.Sprintf方法
	str = fmt.Sprintf("%t", b)
	fmt.Printf("类型是%T str=%q", str, str)

	// strconv函数
}
