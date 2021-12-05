package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestHandleErrFatal(t *testing.T) {
	t.Run("test it calls panic for a given error if it is not nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered from panic", r)
			}
		}()
		HandleErrFatal(errors.New("boop"))
		t.Errorf("HandleErrFatal did not call panic()")
	})
}
