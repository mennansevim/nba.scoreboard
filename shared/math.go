package shared

import "math"

func RoundToTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
