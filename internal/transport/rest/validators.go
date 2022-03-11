package rest

import (
	"strconv"
)

func validateValuesAndUnpack(values []string, result ...*int) error {
	for i, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		*result[i] = num
	}
	return nil
}
