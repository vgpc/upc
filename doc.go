/*

Package upc provides for parsing, validating and analyzing a UPC
(Universal Product Code).  For example,

	u, err := upc.Parse("045496830434")
	if err == upc.ErrInvalidCheckDigit {
		fmt.Println("There's a typo in the UPC")
	} else if err != nil {
		fmt.Printf("Something's wrong with the UPC: %s", err)
	} else {
		fmt.Printf("Number system: %d\n", u.NumberSystem())
		fmt.Printf("Check digit: %d\n", u.CheckDigit())
		if u.IsGlobalProduct() {
			fmt.Printf("Manufacturer code: %s\n", u.Manufacturer())
			fmt.Printf("Product code: %d\n", u.Product())
		} else if u.IsDrug() {
			fmt.Printf("Drug code: %d\n", u.Ndc())
		} else if u.IsLocal() {
			fmt.Println("UPC intended only for local use")
		} else if u.IsCoupon() {
			fmt.Printf("Manufacturer code: %s\n", u.Manufacturer())
			fmt.Printf("Family code: %d\n", u.Family())
			fmt.Printf("Coupon code: %d\n", u.Coupon())
		} else {
			panic("Preceeding categories are exhaustive")
		}
	}

*/
package upc // import "github.com/vgpc/upc"
