package test

import (
	"fmt"
	"errors"
)

// Hello returns a string of a greeting to someone
func Hiya(testName string) (string, error) {
	if testName == "" {
		return "", errors.New("no name provided")
	} 
	
	testMessage := fmt.Sprintf("Hiya, %v. These are your new begginings!", testName)
	return testMessage, nil
}