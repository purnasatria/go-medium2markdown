package md2

import (
	"fmt"
	"runtime"
	"strings"
)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(2) // Changed to 2 to skip newError function
	fullName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

func newError(format string, args ...interface{}) error {
	wrappedErr := fmt.Errorf(format, args...)
	return fmt.Errorf("%s - %w", getFunctionName(), wrappedErr)
}
