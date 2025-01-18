package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64

	for _, str := range strings {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error converting %s to float", str))
		}

		floats = append(floats, f)
	}

	return floats, nil
}
