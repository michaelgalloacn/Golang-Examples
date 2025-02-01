package main

import (
	"github.com/michaelgalloacn/Golang-Examples/errorhandling"
	"github.com/michaelgalloacn/Golang-Examples/exported"
)

const (
	testErrors   = false
	testPanics   = true
	testExported = true
)

func main() {

	if testExported {
		// privateFunction will fail to compile if uncommented
		// exported.privateFunction()
		exported.PublicFunction()
		exported.PublicWrapperForPrivatefunction()

	}

	if testErrors {
		errorhandling.TestErrorHandling()
	}

	if testPanics {
		errorhandling.TestPanic()
	}

}
