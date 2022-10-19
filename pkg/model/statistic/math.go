package statistic

import (
	"fmt"
	"strconv"
)

func CalculatePercentage(total int, partial int) float64 {
	if total <= 0 {
		return 0.0
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(partial*100)/float64(total)), 64)
	return value
}
