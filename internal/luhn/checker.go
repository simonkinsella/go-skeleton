// Package luhn provides utilities for validating numeric strings using the
// Luhn checksum algorithm.
//
// The Luhn algorithm is commonly used to validate identification numbers such
// as credit card numbers, IMEI numbers, and other digit-based identifiers.
//
// See: https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

import (
	"errors"
	"log"
)

// Checker validates numeric strings using the Luhn checksum algorithm.
type Checker struct {
	transformerFuncs []transformerFunc
}

type transformerFunc func(int) int

// NewChecker returns a Checker instance with the standard Luhn digit transformations configured.
func NewChecker() *Checker {
	return &Checker{
		transformerFuncs: []transformerFunc{
			doubleSumNumeral,
			func(n int) int { return n },
		},
	}
}

// Check takes a string representation of a number and returns TRUE
// if it is a valid Luhn number as per
func (c *Checker) Check(str string) bool {
	if len(str) < 2 {
		return false
	}

	oddEvenCounter := len(str) // will be 0 for even length strings

	var total int
	for _, r := range str {
		d, err := digitFromRune(r)
		if err != nil {
			return false
		}
		total += c.transformerFuncs[oddEvenCounter%2](d)
		oddEvenCounter++
	}
	return total%10 == 0
}

func digitFromRune(r rune) (int, error) {
	d := int(r - '0')
	if d < 0 || d > 9 {
		return 0, errors.New("invalid number")
	}
	return d, nil
}

func doubleSumNumeral(i int) int {
	if i > 9 {
		log.Panic("numeral out of range")
	}
	i = 2 * i
	return i%10 + i/10
}
