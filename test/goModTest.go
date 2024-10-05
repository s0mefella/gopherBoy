package test

import "fmt"

// Hello returns a string of a greeting to someone
func Hiya(testName string) string { 
	testMessage := fmt.Sprintf("Hiya, %v. These are your new begginings!", testName)
	return testMessage
}