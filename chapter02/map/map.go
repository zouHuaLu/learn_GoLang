package main

import "fmt"

/*
哈希表是一种巧妙并且实用的数据结构。
它是一个无序的key/value对的集合，其中所有的key都是不同的，然后通过给定的key可以在常数时间复杂度内检索、更新或删除对应的value。
*/

/*
一个map就是一个哈希表的引用，map类型可以写为map[K]V，
其中K和V分别对应key和value。map中所有的key都有相同的类型，所有的value也有着相同的类型，
但是key和value之间可以是不同的数据类型。
其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。
*/
func main() {
	//	内置的make函数可以创建一个map：
	ages := make(map[string]int)
	//	也可以用map字面值的语法创建map，同时还可以指定一些最初的key/value
	people := map[string]int{
		"alice": 23,
		"jay":   27,
	}
	fmt.Println(ages, people)
	fmt.Println(people["jay"])
	delete(people, "alice")
	fmt.Println(ages, people)

	//	要想遍历map中全部的key/value对的话，可以使用range风格的for循环实现，和之前的slice遍历语法类似。
	for name, age := range people {
		fmt.Println("--------")
		fmt.Printf("%s\t %d\n", name, age)
	}
}
