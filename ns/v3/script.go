// Package ns exposes NoxScript API used in maps.
//
// This package is for documentation and type-checking only, all functions do nothing.
//
// Names of functions are synchronized with the latest NoxScript implementation:
// https://noxtools.github.io/noxscript/.
package ns

import (
	"strconv"

	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

// https://github.com/NoxTools/noxscript/blob/master/builtins.h

type Func = any

// RandomFloat generate a random float.
func RandomFloat(min, max float32) float32 {
	return ns4.RandomFloat(min, max)
}

// Random generate a random int.
func Random(min, max int) int {
	return ns4.Random(min, max)
}

// StopScript exit current script function.
func StopScript(value any) {
	ns4.StopScript(value)
}

// IntToString convert an integer to a string.
//
// This will convert an integer to a string, and add the to string the string
// table temporarily.
func IntToString(number int) string {
	return strconv.Itoa(number)
}

// FloatToString converts a float to a string.
//
// This will convert a float to a string, and add the to string the string
// table temporarily.
func FloatToString(number float32) string {
	return strconv.FormatFloat(float64(number), 'g', -1, 32)
}

// Distance calculate distance between two locations.
func Distance(x1, y1, x2, y2 float32) float32 {
	return float32(ns4.Ptf(x1, y1).Sub(ns4.Ptf(x2, y2)).Len())
}
