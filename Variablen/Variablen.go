package variablen

import "errors"

func Divison(a, b float64) (float64, error) {
	division := a / b
	if b == 0 {
		return 0, errors.New("Divison kann nicht durch null Teilen")
	}
	return division, nil
}

func Addition(a, b float64) float64 {
	addition := a + b
	return addition
}

func Multiplikation(a, b float64) float64 {
	multiplikation := a * b
	return multiplikation
}

func Subtraction(a, b float64) float64 {
	subtration := a - b
	return subtration
}
