package thefarm

import (
	"errors"
	"fmt"
)

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")

// See types.go for the types defined for this exercise.
var ErrNegativeFodder = errors.New("negative fodder")
var ErrDivisionByZero = errors.New("division by zero")

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	Cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	switch {
	case cows == 0:
		return 0, ErrDivisionByZero
	case cows < 0:
		return 0, &SillyNephewError{Cows: cows}
	}

	fodder, err := weightFodder.FodderAmount()
	isFodderPositive := fodder >= 0
	isErrScaleMalfunction := errors.Is(err, ErrScaleMalfunction)

	if err != nil && !isErrScaleMalfunction {
		return 0, err
	}

	if isErrScaleMalfunction && isFodderPositive {
		return ((fodder * 2) / float64(cows)), nil
	}

	if !isFodderPositive {
		return 0, ErrNegativeFodder
	}

	return (fodder / float64(cows)), nil
}
