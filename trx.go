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

func Yaml2Json(input string, output string) error {
	return tools.Yaml2Json(input, output)
}

func Json2Yaml(input string, output string) error {
	return tools.Json2Yaml(input, output)
}

func Md2Html(input string, output string) error {
	return tools.Md2Html(input, output)
}

func Html2Md(input string, output string) error {
	return tools.Html2Md(input, output)
}
