package main

import (
	"fmt"
	"sync"
)

func producter(ch chan int,i int,group *sync.WaitGroup) {
	ch <- i
	fmt.Println("producter",i)
	group.Done()
}

func cousmer(ch chan int,i int,group *sync.WaitGroup){
	 x := <-ch
	fmt.Println("cousmer",x)
	group.Done()
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	for i:=0;i<10;i++ {
		go producter(ch,i,&wg)
		go cousmer(ch,i,&wg)
		wg.Add(1)
	}

	//for j:=0;j<10;j++{
	//
	//	wg.Add(1)
	//}
	wg.Wait()
}