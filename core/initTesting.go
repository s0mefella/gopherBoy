package main

import (
	"fmt"
	"example.com/test"
)

func main(){
	message := test.Hiya("Steve")
	fmt.Println(message)
}