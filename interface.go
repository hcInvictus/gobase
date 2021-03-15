package main

import (
	"fmt"
	"sort"
)

//  问题1 interface能做什么
// 问题2  interface底层实现原理

//  什么是interface ?
// interface是定义了一组函数签名的集合，但是没有实现这些函数
// 如果某个数据类型(struct)实现了interfa中定义的那些函数，则称该数据类型实现了这个interface


// 为啥要用interface ?
// 1 为了实现泛型编程
// 2  为了

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []person //自定义

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int)  {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// ByAge 数据类型实现了 Sort interface的三个函数,此处就可以直接调用该interface
	// 该interface一共定义了 less len swap 三个函数
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
