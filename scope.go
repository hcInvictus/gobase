package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var x = 100

func main() {
	// fmt.Println(x)
	// foo()

	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

func foo(){
	fmt.Println(x)
}