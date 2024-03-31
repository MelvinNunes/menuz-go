package util

import (
	"math"
	"strings"
)

func StringToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func TransformToDecimalAmount(amount float64, decimals int) int64 {
	factor := math.Pow(10, float64(decimals))
	return int64(amount * factor)
}
