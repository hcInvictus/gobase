package main

import (
	"fmt"
	"sync"
)

// 使用channel，waitgroup打印奇偶数
func print_1(ch chan int,group *sync.WaitGroup){
	defer group.Done()
	for i:=1 ;i<=100;i++ {
		ch <- i
		if(i % 2 == 1){
			fmt.Println("routine1",i)
			//runtime.Gosched()
		}
	}
}
func print_2(ch chan int,group *sync.WaitGroup){
	defer group.Done()
	for i:=1 ;i<=100;i++ {
		<-ch
		if(i % 2 == 0){
			fmt.Println("routine2",i)
			//runtime.Gosched()
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go print_1(ch,&wg)
	go print_2(ch,&wg)

	wg.Wait()

}
