# goxl

`goxl` is a Golang library for reading Excel files in the same way as CSV.

## Installation

```bash
go get github.com/IamFaizanKhalid/goxl
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/IamFaizanKhalid/goxl"
)

func main() {
	xl, err := goxl.Open("test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer xl.Close()

	sheet, err := xl.OpenDefaultSheet()
	if err != nil {
		log.Fatal(err)
	}
	defer sheet.Close()

	for sheet.Next() {
		row := sheet.Row()
		fmt.Println(row)
	}
}
```
