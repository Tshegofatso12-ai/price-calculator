package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for index, str := range strings {
		floatPrice, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to floar")
		}
		floats[index] = floatPrice
	}
	return floats, nil
}
