package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	name string
}

var sing  *singleton
var mu sync.Mutex

func getInstance() *singleton{
	if sing == nil {
		mu.Lock()
		defer mu.Unlock()
		if sing == nil{
			fmt.Println("create obj ")
			sing = &singleton{}
		}
	}
	return sing
}

func (s *singleton) show(){
	fmt.Println("a show",s.name)
}
func main() {
	obj := getInstance()
	obj.show()

	obj1 := getInstance()
	obj1.show()
}