package main

import "fmt"

func main() {
	switch i := 'y'; i {
	case 'y':
		fmt.Println("yes")
		fallthrough //会跳过接下来的case条件表达式，直接执行下一个case语句

	case 'n':
		fmt.Println("no")
		goto L1
	}

	//L1:
	//	for i := 0; ; i++ {
	//		for j := 0; ; j++ {
	//			if i >= 5 {
	//				fmt.Println(i)
	//				break L1 // 跳出L1标签所在的for循环
	//			}
	//
	//			if j > 10 {
	//				fmt.Println("跳出j", j)
	//				break // 默认仅跳出离break最近的内层循环
	//			}
	//		}
	//	}
	//}

L1:
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if i >= 5 {
				fmt.Println(i)
				continue L1 // 跳到L1标签所在的for循环i++处执行
			}

			if j > 10 {
				//默认仅跳到离break最近的内层循环j++处执行
				fmt.Println("跳出j", j)
				continue
			}
		}
	}
}
