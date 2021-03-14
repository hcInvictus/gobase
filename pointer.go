package main

import "fmt"

// 如何使用指针

// 函数值传递,会把参数重新拷贝一份传过来，对参数的修改是对拷贝值的修改
func ShowValue(value int) {
	value = 300
	fmt.Println(value, &value)
}

// 函数传递指针，是对原地址值的修改
func ShowValue1(value *int) {
	*value = 300
	fmt.Println(value, &value)
}
func main() {
	a := 100
	fmt.Println(&a)

	// b  指向a的地址，对b做解引用操作就能取出地址的值
	var b = &a
	fmt.Println(*b)

	// 由于b指向a的地址，对b进行赋值操作，会直接改变a的值
	*b = 200
	fmt.Println(a, b)

	fmt.Println("------")
	// 函数值传递
	ShowValue(a)
	fmt.Println(a)

	// 函数传递指针
	ShowValue1(&a)
	fmt.Println(a)
}
