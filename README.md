# upc

A Go library for parsing, validating and analyzing UPCs.

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
