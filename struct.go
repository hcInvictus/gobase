package main

import (
	"encoding/json"
	"fmt"
)

type foo int

type Person struct {
	name string
	age  int
}

type MyName struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var myage foo
	myage = 100
	fmt.Printf("%T %v \n", myage, myage) // main.foo 100

	var age int
	age = 100
	fmt.Printf("%T %v \n", age, age) // int 100

	// 此处报错，提示类型不同，非法的操作符号
	// fmt.Println(age + myage)

	// 使用类型强制转换
	fmt.Println(int(myage) + age)

	p1 := &Person{"James", 20}
	p2 := Person{"ch", 100}
	fmt.Println(p2)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)

	// struct Marshal
	pMarsh, _ := json.Marshal(p2)
	fmt.Println(string(pMarsh))

	name1 := MyName{"chen", 100}
	bs, _ := json.Marshal(name1)
	fmt.Println(string(bs))
}
