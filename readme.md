# TRX

![Go](https://img.shields.io/badge/Go-1.25-blue?logo=go&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-orange.svg)
![Version](https://img.shields.io/badge/version-0.0.0-white)

Minimal converters written in Go for various file formats.

## Project Structure

- `go.mod`: Go module file specifying dependencies.
- `go.sum`: Go module checksum file.
- `readme.md`: This readme file.
- `trx.go`: Main module containing the conversion functions.
- `tools/`: Module containing various file format converters.

## Supported Converters

- JSON to CSV
- CSV to JSON
- JSON to YAML
- YAML to JSON

### JSON to CSV

```go
err := trx.Json2Csv("input.json", "output.csv")
if err != nil {
    fmt.Printf("Error converting JSON to CSV: %v\n", err)
}
```

### CSV to JSON

```go
err := trx.Csv2Json("input.csv", "output.json")
if err != nil {
    fmt.Printf("Error converting CSV to JSON: %v\n", err)
}
```

### JSON to YAML

```go
err := trx.Json2Yaml("input.json", "output.yaml")
if err != nil {
    fmt.Printf("Error converting JSON to YAML: %v\n", err)
}
```

### YAML to JSON

```go
err := trx.Yaml2Json("input.yaml", "output.json")
if err != nil {
    fmt.Printf("Error converting YAML to JSON: %v\n", err)
}
```

## Installation

You can clone the repository and run the application using the following commands:

```bash
git clone https://github.com/unsafezero/trx.git
cd trx
go run main.go
```

This will start the application and you can use the various converters provided in the `trx` module.

## Licensing

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.
