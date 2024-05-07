package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(lines []string) ([]float64, error) {
	var floatslist []float64
	for _, stringVal := range lines {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}

		floatslist = append(floatslist, floatVal)
	}
	return floatslist, nil
}
