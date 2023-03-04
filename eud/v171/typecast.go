package eud

import "strconv"

// ToInt converts float to int.
//
// Deprecated: use native conversion int(x).
func ToInt(x float32) int {
	return int(x)
}

// ToFloat converts int to float.
//
// Deprecated: use native conversion float32(x).
func ToFloat(x int) float32 {
	return float32(x)
}

// SToInt converts string to int.
//
// Deprecated: use strconv.Atoi.
func SToInt(x string) int {
	v, _ := strconv.Atoi(x)
	return int(v)
}

// ToStr converts int to string.
//
// Deprecated: use strconv.Itoa.
func ToStr(x int) string {
	return strconv.Itoa(x)
}

// FloatToInt naively converts float to int.
//
// Deprecated: use native conversion int(x).
func FloatToInt(x float32) int {
	var result int32
	pos := x

	if pos < 0.0 {
		pos = -pos
	}
	pos = pos / 2147483648.0
	if pos < 2.0 {
		for i := 0; i < 32; i++ {
			if pos >= 1.0 {
				result++
				pos -= 1.0
			}
			if i != 31 {
				result = result << 1
			}
			pos *= 2.0
		}
	} else {
		result = 0x7fffffff
	}
	if x < 0.0 {
		return -int(result)
	} else {
		return int(result)
	}
}

// IntToFloat naively converts int to float.
//
// Deprecated: use native conversion float32(x).
func IntToFloat(x int) float32 {
	var i int32
	pos := int32(x)
	result := float32(0.0)
	if pos < 0 {
		pos = -pos
	}
	for i = 0; i < 32; i++ {
		if (uint32(pos) & 0x80000000) != 0 {
			result += 1.0
		}
		pos = pos << 1
		if i != 31 {
			result *= 2.0
		}
	}
	if x < 0 {
		return -result
	} else {
		return result
	}
}
