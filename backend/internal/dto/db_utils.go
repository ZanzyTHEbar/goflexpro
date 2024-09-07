package dto

import (
	"fmt"

	"github.com/ZanzyTHEbar/goflexpro/pkgs/errsx"
)

// Validator for DTO objects
// RetrieveValue retrieves a Prisma ORM value from an accessor function and sets an error if the value is not available
func RetrieveValue[T any](getValue func() (T, bool), key string, defaultValue T, errs *errsx.ErrorMap) T {
	value, ok := getValue()
	if !ok {
		errs.Set(key, fmt.Sprintf("No %s available for %v", key, defaultValue))
		return defaultValue
	}

	return value
}
