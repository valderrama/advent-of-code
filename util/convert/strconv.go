package convert

import (
	"strconv"
)

// MustAtoi calls strconv.Atoi(string) and panics if the value cannot be converted
func MustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return result
}
