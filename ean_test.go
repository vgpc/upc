package upc
import "testing"

// a breakdown of a UPC into each of its possible attributes
type eanBreakdown struct {
	err             string
	checkDigit      int
	isJan           bool
}

var eanTests = map[string]eanBreakdown{
	"0045496401771": { // mario party 8 wii
		checkDigit: 1,
	},
	"5055856406112": { // fallout 4 ps4
		checkDigit: 2,
	},
	"5030938121923": { // fifa 19 xbox one
		checkDigit: 3,
		isJan:      false,
	},
	"4549673590600": { // mario kart ds
		checkDigit: 0,
		isJan:      true,
	},
	"0045496738730": { // ds browser
		checkDigit: 0,
		isJan:      false,
	},
}

func EanTestBreakdown(t *testing.T) {
	for s, expect := range eanTests {
		if e, got := getBreakdownEan(s); got != expect {
			t.Errorf("%s: wrong breakdown\n got: %#v\nwant: %#v\n", s, got, expect)
		} else if e.String() != s {
			t.Errorf("%s: wrong string: got %s", s, e)
		}
	}
}

func getBreakdownEan(s string) (Ean, eanBreakdown) {
	var b eanBreakdown
	e, err := ParseEan(s)
	if err == nil {
		b.checkDigit = e.CheckDigit()
		b.isJan =      e.IsJan()
	} else {
		b.err = err.Error()
	}

	return e, b
}

func EanTestWrong(t *testing.T) {
	short := []string{
		"",
		"$19.99",
		"J7D-00001",
	}
	for _, s := range short {
		if _, err := ParseEan(s); err != ErrEanTooShort {
			t.Errorf("%s: expected ErrEanTooShort got %q", s, err)
		}
	}

	long := []string{
		"012345678912345",
	}
	for _, s := range long {
		if _, err := ParseEan(s); err != ErrEanTooLong {
			t.Errorf("%s: expected ErrEanTooLong got %q", s, err)
		}
	}

	check := []string{
		"0123456789199",
	}
	for _, s := range check {
		if _, err := ParseEan(s); err != ErrEanInvalidCheckDigit {
			t.Errorf("%s: expected ErrEanInvalidCheckDigit got %q", s, err)
		}
	}

	digit := []string{
		"x123456789128",
		"012345678910x",
		"01234x6789128",
		"0123456x89128",
	}
	for _, s := range digit {
		if _, err := ParseEan(s); err == nil {
			t.Errorf("%s: expected an error got none", s)
		}
	}
}

func EanBenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseEan("045496830434")
	}
}
