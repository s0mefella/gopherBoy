package test

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a string of a greeting to someone
func Hiya(testName string) (string, error) {
	if testName == "" {
		return "", errors.New("no name provided")
	} 
	
	testMessage := fmt.Sprintf(randomHello(), testName)
	return testMessage, nil
}

func whoWhis(testNames []string) (map[string]string, error) {
	testMessages := make(map[string]string)
	for _, testNames := range testNames {
		testMessage, err := Hiya(testNames)
		if err != nil {
			return nil, err
		}
		
		testMessages[testNames] = testMessage
	}
	return testMessages, nil
}

func randomHello() string {
	testFormats := []string{
		"Ayup, %v. Come tut factory, there's cheddar to be made lad",
		"Seems like your here now, %v. Fancy a pint?", 
		"Honest I dunno whats happening either, %v.",
	}
	return testFormats[rand.Intn(len(testFormats))]
}