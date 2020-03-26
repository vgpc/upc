package ean // import "github.com/vgpc/upc"
import (
	"errors"
	"fmt"
)

// Ean represents a European Article Number.  To reduce memory
// consumption, it's stored as a 64-bit integer without the check
// digit.
type Ean int64

var ErrEanTooShort = errors.New("EAN is too short (must be 13 digits)")
var ErrEanTooLong = errors.New("EAN is too long (must be 13 digits)")
var ErrEanInvalidCheckDigit = errors.New("EAN has an invalid check digit")

// Parse parses a string into a Ean value.  The following errors can
// be returned in addition to integer parsing errors:
//
//     ErrEanTooShort
//     ErrEanTooLong
//     ErrEanInvalidCheckDigit
func Parse(s string) (Ean, error) {
	if len(s) < 13 {
		return 0, ErrEanTooShort
	}
	if len(s) > 13 {
		return 0, ErrEanTooLong
	}

	var n int64
	var check int
	for i, b := range []byte(s) {
		if b < 48 || b > 57 {
			return 0, fmt.Errorf("Invalid UPC digit: %c", b)
		}
		if i == 12 {
			check = int(b - 48)
		} else {
			n *= 10
			n += int64(b - 48)
		}
	}
	e := Ean(n)
	if e.CheckDigit() != check {
		return 0, ErrEanInvalidCheckDigit
	}

	return e, nil
}

// String returns the standard, 13-digit string representation of this
// EAN.
func (e Ean) String() string {
	return fmt.Sprintf("%012d%d", int64(e), e.CheckDigit())
}

// CheckDigit returns the check digit that should be used as the 13th
// digit of the EAN.
func (e Ean) CheckDigit() int {
	var sum, multiplier, n int64
	n = int64(e)
	multiplier = 3
	for n > 0 {
		sum = sum + multiplier*(n%10)
		n = n / 10
		multiplier = 4 - multiplier // alternate between 3 and 1
	}
	return int((10 - (sum % 10)) % 10)
}
