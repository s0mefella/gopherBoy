package main

import (
	"fmt"
	"log"

	"example.com/test"
)

func main(){
	log.SetPrefix("goModTest: ")
	log.SetFlags(0)

	message, err := test.Hiya("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}