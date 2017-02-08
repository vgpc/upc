package upc // import "github.com/vgpc/upc"
import "testing"

// a breakdown of a UPC into each of its possible attributes
type breakdown struct {
	err             string
	numberSystem    int
	checkDigit      int
	isGlobalProduct bool
	isDrug          bool
	isLocal         bool
	isCoupon        bool
	manufacturer    string
	product         int
	ndc             int
	family          int
	coupon          int
}

var tests = map[string]breakdown{
	"045496830434": {
		numberSystem:    0,
		checkDigit:      4,
		isGlobalProduct: true,
		manufacturer:    "045496",
		product:         83043,
	},
}

func TestBreakdown(t *testing.T) {
	for s, expect := range tests {
		if u, got := getBreakdown(s); got != expect {
			t.Errorf("%s: wrong breakdown\n got: %#v\nwant: %#v\n", s, got, expect)
		} else if u.String() != s {
			t.Errorf("%s: wrong string: got %s", s, u)
		}
	}
}

func getBreakdown(s string) (Upc, breakdown) {
	var b breakdown
	u, err := Parse(s)
	if err == nil {
		b.numberSystem = u.NumberSystem()
		b.checkDigit = u.CheckDigit()
		b.isGlobalProduct = u.IsGlobalProduct()
		// b.isDrug = u.IsDrug()
		// b.isLocal = u.IsLocal()
		// b.isCoupon = u.IsCoupon()
		b.manufacturer = u.Manufacturer()
		b.product = u.Product()
		// b.ndc = u.Ndc()
		// b.family = u.Family()
		// b.coupon = u.Coupon()
	} else {
		b.err = err.Error()
	}

	return u, b
}
