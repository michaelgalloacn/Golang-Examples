package errorhandling

import "fmt"

func TestErrorHandling() {
	str, err := stringWithError(false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)
}
func stringWithError(causeError bool) (string, error) {
	if causeError {
		return "failure", fmt.Errorf("Do not ignore this error")
	}
	return "string", nil
}

func TestPanic() {
	// The compiler will not allow the denominator to be a const 0 so it has to be assigned to a
	// A var first to trick the compiler and allow this panic to happen at runtime
	var denominator = 0
	_ = 1 / denominator
}
