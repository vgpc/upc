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
	ndc             string
	family          int
	value           int
}

var tests = map[string]breakdown{
	"045496830434": { // EarthBound for SNES
		numberSystem:    0,
		checkDigit:      4,
		isGlobalProduct: true,
		manufacturer:    "045496",
		product:         83043,
	},
	"298765432109": {
		numberSystem: 2,
		checkDigit:   9,
		isLocal:      true,
	},
	"412345678903": {
		numberSystem: 4,
		checkDigit:   3,
		isLocal:      true,
	},
	"363824057361": { // Mucinex D
		numberSystem: 3,
		checkDigit:   1,
		isDrug:       true,
		ndc:          "6382405736",
	},
	"512345678900": {
		numberSystem: 5,
		checkDigit:   0,
		isCoupon:     true,
		manufacturer: "12345",
		family:       678,
		value:        90,
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
		b.isDrug = u.IsDrug()
		b.isLocal = u.IsLocal()
		b.isCoupon = u.IsCoupon()
		if b.isGlobalProduct {
			b.manufacturer = u.Manufacturer()
			b.product = u.Product()
		}
		if b.isDrug {
			b.ndc = u.Ndc()
		}
		if b.isCoupon {
			b.manufacturer = u.Manufacturer()
			b.family = u.Family()
			b.value = u.Value()
		}
	} else {
		b.err = err.Error()
	}

	return u, b
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("045496830434")
	}
}
