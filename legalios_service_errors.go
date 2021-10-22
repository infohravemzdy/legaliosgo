package legalios

import (
	"errors"
	"fmt"
)

type historyResultError struct {
	Err error
}

func (e historyResultError) Error() string {
	return fmt.Sprintf("error: %v", e.Err)
}

func newBundleNoneError() historyResultError {
	return historyResultError{errors.New("service hasn't returned bundle") }
}

func newBundleNullError() historyResultError {
	return historyResultError{errors.New("service returned null bundle") }
}

func newBundleEmptyError() historyResultError {
	return historyResultError{errors.New("service returned empty bundle") }
}

func newBundleInvalidError() historyResultError {
	return historyResultError{errors.New("service returned invalid bundle") }
}

