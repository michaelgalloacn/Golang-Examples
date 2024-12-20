package main

import "github.com/michaelgalloacn/Golang-Examples/errorhandling"

const (
	testErrors = false
	testPanics = true
)

func main() {
	if testErrors {
		errorhandling.TestErrorHandling()
	}

	if testPanics {
		errorhandling.TestPanic()
	}

}
