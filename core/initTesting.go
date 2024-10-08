package main

import (
	"fmt"
	"log"

	"example.com/test"
)

func main(){
	log.SetPrefix("goModTest: ")
	log.SetFlags(0)

	testNames := []string{"Stevie", "Dave", "Darwin"}
	testMessages, err := test.whoWhis(testNames)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(testMessages)
}