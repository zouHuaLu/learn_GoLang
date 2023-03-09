package main

import "fmt"

func main() {
	sum := add(1, 3)
	fmt.Println(sum)

	a := 10
	chvalue(a) //实参传递给形参是值拷贝
	fmt.Println(a)

	chpointer(&a) // 值拷贝，复制的是a的地址值
	fmt.Println(a)
}

//func add(a, b int) byte {
//	sum := a + b
//
//	return byte(sum)
//}

func add(a, b int) (sum byte) {
	sum = byte(a + b)
	return
}

//Go 函数实参到形参的传递永远是值拷贝

func chvalue(a int) int {
	a = a + 1
	return a
}

func chpointer(a *int) {
	*a = *a + 1
	return
}
