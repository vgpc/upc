# GOlang for UPC and EAN

A Go library for parsing, validating and analyzing UPCs and EAN-13 codes.

* Validate the accuracy of UPC and EAN codes.
* Parse UPC codes to determine the code type and other details (Manufacturer, Coupon Value, Product Code, etc)

The library works with 12 digit UPC (Universal Product Code) codes and 13 digit EAN (European Article Number) codes.

# Usage

```go
u, err := upc.Parse("045496830434")
if err == upc.ErrInvalidCheckDigit {
    fmt.Println("There's a typo in the UPC")
} else if err != nil {
    fmt.Printf("Something's wrong with the UPC: %s", err)
} else {
    fmt.Printf("Number system: %d\n", u.NumberSystem())
    fmt.Printf("Check digit: %d\n", u.CheckDigit())
    if u.IsGlobalProduct() {
        fmt.Printf("Manufacturer code: %d\n", u.Manufacturer())
        fmt.Printf("Product code: %d\n", u.Product())
    } else if u.IsDrug() {
        fmt.Printf("Drug code: %d\n", u.Ndc())
    } else if u.IsLocal() {
        fmt.Println("UPC intended only for local use")
    } else if u.IsCoupon() {
        fmt.Printf("Manufacturer code: %d\n", u.Manufacturer())
        fmt.Printf("Family code: %d\n", u.Family())
        fmt.Printf("Coupon value: $0.%02d\n", u.Value())
    } else {
        panic("Preceeding categories are exhaustive")
    }
}
```

# Installation

```
go get -u github.com/vgpc/upc
```
