package upc
import (
	"errors"
	"fmt"
)

// Ean represents a European Article Number.  To reduce memory
// consumption, it's stored as a 64-bit integer without the check
// digit.
type Ean int64

var ErrEanTooShort = errors.New("EAN is too short (must be 12 digits)")
var ErrEanTooLong = errors.New("EAN is too long (must be 13 digits)")
var ErrEanInvalidCheckDigit = errors.New("EAN has an invalid check digit")

// Parse parses a string into a Ean value.  The following errors can
// be returned in addition to integer parsing errors:
//
//     ErrEanTooShort
//     ErrEanTooLong
//     ErrEanInvalidCheckDigit
func ParseEan(s string) (Ean, error) {
	if len(s) < 12 {
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
		if len(s) == 12 { // check if 12 digits
			if i == 11 {
				check = int(b - 48)
			} else {
				n *= 10
				n += int64(b - 48)
			}
		} else {
			if i == 12 { // check if 13 digits
				check = int(b - 48)
			} else {
				n *= 10
				n += int64(b - 48)
			}
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

// IsJan returns true if the number begins with 45 or 49
// JAN (Japanese Article Numbering) codes must begin with
// one of these two numbers
func (e Ean) IsJan() bool {
	firstTwo := int(e / 10000000000)
	switch firstTwo {
	case 45, 49:
		return true
	default:
		return false
	}
}
