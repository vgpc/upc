package upc // import "github.com/vgpc/upc"
import (
	"errors"
	"fmt"
	"strconv"
)

// Upc represents a Universal Product Code.  To reduce memory
// consumption, it's stored as a 64-bit integer without the check
// digit.
type Upc int64

var ErrTooShort = errors.New("UPC is too short (must be 12 digits)")
var ErrTooLong = errors.New("UPC is too long (must be 12 digits)")
var ErrInvalidCheckDigit = errors.New("UPC has an invalid check digit")

// Parse parses a string into a Upc value.  The following errors can
// be returned in addition to integer parsing errors:
//
//     ErrTooShort
//     ErrTooLong
//     ErrInvalidCheckDigit
func Parse(s string) (Upc, error) {
	if len(s) < 12 {
		return 0, ErrTooShort
	}
	if len(s) > 12 {
		return 0, ErrTooLong
	}

	n, err := strconv.ParseInt(s[0:11], 10, 64)
	if err != nil {
		return 0, err
	}
	check, err := strconv.Atoi(s[11:])
	if err != nil {
		return 0, err
	}
	u := Upc(n)
	if u.CheckDigit() != check {
		return 0, ErrInvalidCheckDigit
	}

	return u, nil
}

// String returns the standard, 12-digit string representation of this
// UPC.
func (u Upc) String() string {
	return fmt.Sprintf("%011d%d", int64(u), u.CheckDigit())
}

// CheckDigit returns the check digit that should be used as the 12th
// digit of the UPC.
func (u Upc) CheckDigit() int {
	var sum, multiplier, n int64
	n = int64(u)
	multiplier = 3
	for n > 0 {
		sum = sum + multiplier*(n%10)
		n = n / 10
		multiplier = 4 - multiplier // alternate between 3 and 1
	}
	return int((10 - (sum % 10)) % 10)
}

// NumberSystem returns the first digit of the UPC, also known as the
// "number system".
func (u Upc) NumberSystem() int {
	return int(u / 10000000000)
}

// IsGlobalProduct returns true if the number system is 0, 1, 6, 7, 8
// or 9.  A UPC with one of these number systems is intended for
// global use with products (in contrast to local use or coupon use,
// etc.)
func (u Upc) IsGlobalProduct() bool {
	switch u.NumberSystem() {
	case 0, 1, 6, 7, 8, 9:
		return true
	default:
		return false
	}
}

// Manufacturer returns the 6-digit manufacturer code assigned by a
// GS1 organization.  The return value is a string because leading
// zeros are used for lookups in the standard GS1 databases.
func (u Upc) Manufacturer() string {
	return fmt.Sprintf("%06d", u/100000)
}

// Product returns the product code assigned by a manufacturer.
func (u Upc) Product() int {
	return int(u % 100000)
}
