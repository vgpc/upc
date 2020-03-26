
# Golang Library for UPC and EAN

A Go library for parsing, validating and analyzing UPCs and EAN-13/GTIN-13 codes.

* Validate the accuracy of UPC and EAN/GTIN codes.
* Parse UPC codes to determine the code type and other details (Manufacturer, Coupon Value, Product Code, etc)

# Code Support

* UPC (Universal Product Code)
* EAN/GTIN 13 (European Article Numbering or Global Trade Item Number)
* JAN (Japanese Article Numbering)

The library works with 12 digit UPC (Universal Product Code) codes and 13 digit EAN (European Article Number) codes.  EAN is also referred to as GTIN or Global Trade Item Number.  The numbers are the same.

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
