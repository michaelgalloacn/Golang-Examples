package main

import (
	"github.com/michaelgalloacn/Golang-Examples/errorhandling"
	"github.com/michaelgalloacn/Golang-Examples/exported"
	Server "github.com/michaelgalloacn/Golang-Examples/server"
)

const (
	testErrors   = false
	testPanics   = false
	testExported = false
	runServer    = false
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

	if runServer {
		Server.RunServer()
	}
}
