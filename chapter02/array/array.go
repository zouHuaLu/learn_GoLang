package main

import "fmt"

func main() {
	r := [...]int{99: -1}
	fmt.Printf("%T\n", r)

	// 数组的比较
	a := [2]int{1, 2}
	aa := &a
	b := [...]int{1, 2}
	bb := &b
	fmt.Println(a == b)
	fmt.Println(aa == bb)
	fmt.Printf("aa的地址是：%p/n，bb的地址是：%p/n ", aa, bb)

	// 以指针的形式将数组传入函数中
	aaa := [32]byte{31: 23}
	fmt.Println(aaa)
	zero(&aaa)
	fmt.Println(aaa)
}

// 将数组里的每一项修改为0
func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
