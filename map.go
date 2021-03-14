package main

import "fmt"

func main() {
	maps := map[string]string{}
	maps["a"] = "aaa"
	maps["b"] = "bbb"
	fmt.Println(len(maps),maps)

	// 判断map中某个key是否存在
	if _ ,ok := maps["a"];ok {
		fmt.Println("---")
	}

	m := make(map[string]int,2)
	fmt.Println(m)


	//  实现安全的map

}
