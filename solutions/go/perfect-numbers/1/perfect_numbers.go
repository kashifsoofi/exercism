package perfect

import (
	"errors"
)

// Classification Nicomachus' (60 - 120 CE) classification
type Classification int

const (
	// ClassificationPerfect Perfect number
	ClassificationPerfect Classification = iota
	// ClassificationAbundant abundant number
	ClassificationAbundant
	// ClassificationDeficient deficient number
	ClassificationDeficient
)

// ErrOnlyPositive error if number is not positive
var ErrOnlyPositive error = errors.New("input not a positive number")

// Classify classifies a number per Nicomachus' (60 - 120 CE) classification
func Classify(number int64) (Classification, error) {
	if number < 1 {
		return ClassificationDeficient, ErrOnlyPositive
	}

	var aliquotSum int64 = 0
	var i int64
	for i = 1; i <= number/2; i++ {
		if number%i == 0 {
			aliquotSum += i
		}
	}

	if aliquotSum > number {
		return ClassificationAbundant, nil
	} else if aliquotSum < number {
		return ClassificationDeficient, nil
	}
	return ClassificationPerfect, nil
}
