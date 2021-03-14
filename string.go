package main

import "fmt"

func main() {
	s := "strings"

	for i:=0;i<len(s);i++{
		fmt.Printf("%c \n",s[i])
	}

	for k,v := range s{
		fmt.Printf("%d %c \n",k,v)
	}

	// 将字符串转化为切片方式遍历
	str := []rune(s)
	for i:=0;i<len(str);i++{
		fmt.Printf("%c \n",str[i])
	}

	// 字符串切片强制转换为字符串  == >  string(str)
}