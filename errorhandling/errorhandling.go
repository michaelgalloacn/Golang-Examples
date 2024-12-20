package errorhandling

import "fmt"

func StringWithError(causeError bool) (string, error) {
	if causeError {
		return "failure", fmt.Errorf("Do not ignore this error")
	}
	return "string", nil
}
