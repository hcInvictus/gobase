package main

import "fmt"

func greet(namea, nameb string) (string, string) {
	return fmt.Sprint(namea, nameb), namea
}

// 返回一个函数，该函数的返回值是int
func makeGreeter() func() int {
	return func() int {
		return 100
	}
}

// 回调函数
func rangeNums(nums []int, callback func(int)) {
	for _, value := range nums {
		callback(value)
	}
}
func showNum(value int) {
	fmt.Println("value is", value)
}

func main() {

	a, b := greet("aaa", "bbb")
	fmt.Println(a, b)

	mG := makeGreeter()
	fmt.Println(mG())
	fmt.Printf("%T \n", mG)

	nums := []int{11, 22, 33, 44}
	rangeNums(nums, showNum)
}
