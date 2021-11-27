package errors

import "fmt"

// Handles all error generically, followed by a panic
func HandleErrFatal(err error) {
	if err != nil {
		err = fmt.Errorf("generic error handled: %w", err)
		panic(err)
	}
}
