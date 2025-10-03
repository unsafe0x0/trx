package trx

import (
	"github.com/unsafe0x0/trx/tools"
)

func Json2Csv(input string, output string) error {
	return tools.Json2Csv(input, output)
}

func Csv2Json(input string, output string) error {
	return tools.Csv2Json(input, output)
}
