package main

import (
	"fmt"
	"github.com/unsafe0x0/trx"
)

func main() {
	err := converter.Csv2Json("testings/input.csv", "testings/output.json")
	if err != nil {
		fmt.Printf("Error converting CSV to JSON: %v\n", err)
		return
	}
	fmt.Println("CSV to JSON conversion completed successfully.")
}
