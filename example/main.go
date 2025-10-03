package main

import (
	"fmt"
	"github.com/unsafe0x0/trx"
)

func main() {
	err := trx.Csv2Json("input.csv", "output.json")
	if err != nil {
		fmt.Printf("Error converting CSV to JSON: %v\n", err)
		return
	}
	fmt.Println("CSV to JSON conversion completed successfully.")
	err = trx.Json2Csv("input.json", "output.csv")
	if err != nil {
		fmt.Printf("Error converting JSON back to CSV: %v\n", err)
		return
	}
	fmt.Println("JSON to CSV conversion completed successfully.")
}
