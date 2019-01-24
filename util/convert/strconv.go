package convert

import (
	"strconv"
)

// MustAtoi calls strconv.Atoi() and panics if the value cannot be converted
func MustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

// MustParseUint calls strconv.ParseUint() and panics if the value cannot be converted
func MustParseUint(s string, bitSize int) uint {
	result, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		panic(err)
	}
	return uint(result)
}

// MustParseUint8 calls strconv.ParseUint() and panics if the value cannot be converted
func MustParseUint8(s string) uint8 {
	return uint8(MustParseUint(s, 8))
}

// MustParseUint16 calls strconv.ParseUint() and panics if the value cannot be converted
func MustParseUint16(s string) uint16 {
	return uint16(MustParseUint(s, 16))
}

// MustParseUint32 calls strconv.ParseUint() and panics if the value cannot be converted
func MustParseUint32(s string) uint32 {
	return uint32(MustParseUint(s, 32))
}

// MustParseUint64 calls strconv.ParseUint() and panics if the value cannot be converted
func MustParseUint64(s string) uint64 {
	return uint64(MustParseUint(s, 64))
}
