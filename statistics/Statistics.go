package statistics

import (
	"fmt"
	"math"
)

type Statistics struct {
	value int
}

func ConstructStatistics() Statistics {
	return Statistics{value: 0}
}

func (stats *Statistics) CalculatePercentageAsString(totalValue int) string {

	const PercentMultiplier = 100

	percentage := int(math.Round(float64(stats.value) / float64(totalValue) * PercentMultiplier))

	percentageAsText := fmt.Sprintf("%v", percentage) + "%"

	return percentageAsText
}

func (stats *Statistics) GetValue() int {
	return stats.value
}

func (stats *Statistics) IncreaseValue() {
	stats.value++
}
