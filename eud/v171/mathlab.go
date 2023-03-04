package eud

import "math"

func ABS(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func FABS(v float32) float32 {
	return float32(math.Abs(float64(v)))
}

func MathDirToDegree(direction int) int {
	return (direction * 45) >> 5
}

func MathDegreeToDir(degree int) int {
	return degree * 32 / 45
}
