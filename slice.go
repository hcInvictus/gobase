package main

import "fmt"

func main() {

	var ints []int = make([]int, 2, 5) // 初始化前2个元素为0,从第三个开始是没有地址的，如果访问第三个元素会地址越界
	fmt.Println(ints)                  // 0 0
	ints = append(ints, 1)
	//fmt.Println(ints[3]) // 越界，panic

	slices := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}
	fmt.Println(len(slices), cap(slices))
	s1 := slices[1:4] // 前闭后开
	s2 := slices[7:]
	// cap  = 从当前起始位置到原来切片最后的大小
	fmt.Println(len(s1), cap(s1)) // 22 33 44 len= 3 cap=9
	fmt.Println(len(s2), cap(s2)) // 88 99 100 len = 3 cap = 3
	fmt.Println(&s1, &s2)

	// 操作原切片，s1 ,s2
	s1[0] = 2
	fmt.Println(slices) // 会改变slices的值 11 2 33 44 ....

	// 给s2添加元素会怎么样 , slices是否会改变len，cap呢,发生了什么
	// 原s2 = 88 99 100
	s2 = append(s2, 111) // 当添加新的111 时s2底层的空间已经不够存储新元素，
	// 此时把s2的空间扩容到2倍大，然后将原来s2的数据拷贝到新空间，在进行操作
	fmt.Println(slices, len(slices), cap(slices)) //
	fmt.Println(s2, len(s2), cap(s2))             // 88 99 100 111  len=4,cap=6

	// 扩容规则实验
	arr := []int{1, 2}         // oldlen=2,oldcap=2
	arr = append(arr, 3, 4, 5) // len=5,cap=5
	// oldcap*2 < cap   ====> newcap = cap
	// oldlen < 1024    =====> newcap=oldcap*2
	// oldlen >= 1024   =====> newcap = oldcap*1.25
	fmt.Println(len(arr), cap(arr))

	// 字符串操作
	str := new([]string)
	fmt.Println(*str)
	*str = append(*str, "chen")
	fmt.Println(*str, len(*str), cap(*str)) // chen 1 1
}
