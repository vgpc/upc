package upc // import "github.com/vgpc/upc"
import (
	"errors"
	"fmt"
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

	var n int64
	var check int
	for i, b := range []byte(s) {
		if b < 48 || b > 57 {
			return 0, fmt.Errorf("Invalid UPC digit: %c", b)
		}
		if i == 11 {
			check = int(b - 48)
		} else {
			n *= 10
			n += int64(b - 48)
		}
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
// or 9.  These UPCs are intended for global use with products (in
// contrast to local use or coupon use, etc.)
func (u Upc) IsGlobalProduct() bool {
	switch u.NumberSystem() {
	case 0, 1, 6, 7, 8, 9:
		return true
	default:
		return false
	}
}

// IsLocal returns true if the number system is 2 or 4.  These UPCs
// are intended for local/warehouse use.
func (u Upc) IsLocal() bool {
	switch u.NumberSystem() {
	case 2, 4:
		return true
	default:
		return false
	}
}

// IsDrug returns true if the number system is 3.  These UPCs are
// intended for labeling drugs by their National Drug Council number.
// See Ndc method.
func (u Upc) IsDrug() bool {
	return u.NumberSystem() == 3
}

// IsCoupon returns true if the number system is 5.  These UPCs are
// intended for labeling coupons.  See Family and Value methods.
func (u Upc) IsCoupon() bool {
	return u.NumberSystem() == 5
}

// Manufacturer returns the 6-digit manufacturer code assigned by a
// GS1 organization.  The return value is a string because leading
// zeros are used for lookups in the standard GS1 databases. In the
// case of coupons, the manufacturer is only 5 digits long.
func (u Upc) Manufacturer() string {
	if u.NumberSystem() == 5 {
		return fmt.Sprintf("%05d", (u/100000)%100000)
	} else {
		return fmt.Sprintf("%06d", u/100000)
	}
}

// Product returns the product code assigned by a manufacturer.
func (u Upc) Product() int {
	return int(u % 100000)
}

// Ndc returns the National Drug Code associated with a UPC.  The
// value is only meaningful if IsDrug returns true for the UPC.
//
// The value is a string because leading zeros are meaningful in the
// FDA database of labeler codes.  No attempt is made to put the NDC
// code into standard format with dashes.
func (u Upc) Ndc() string {
	full := fmt.Sprintf("%011d", u)
	return full[1:]
}

// Family returns the coupon manufacturer's family code.  It
// designates the kind of products to which the coupon applies.
func (u Upc) Family() int {
	return int((u % 100000) / 100)
}

// Value returns the coupon value in pennies.  This is the amount
// saved by applying the coupon.  The value is always less than 100.
func (u Upc) Value() int {
	return int(u % 100)
}
