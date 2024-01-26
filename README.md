<p align="center">
  <img src="envsnatch.png" alt="EnvSnatch" />
</p>
<h1 align="center" style="padding-top: 36px; margin-top: -36px; margin-bottom: 36px">
 EnvSnatch
</h1>  

EnvSnatch is a Go package designed for efficient environment variable loading into Go structs. It supports a range of data types and handles both required and optional variables with ease.

## Features

- Loads environment variables directly into Go structs.
- Supports a variety of data types including strings, integers, unsigned integers, floats, and booleans.
- Treats all struct fields as required by default, unless explicitly marked as optional.
- Optional fields can be marked using the `optional:"true"` tag.

## Upcoming Features

We are constantly working to improve EnvSnatch and plan to add more features in the future. Any help or contributions to these features are highly welcome:

- **Additional Validations**: Alongside the current 'required' and 'optional' validations, we aim to introduce more validation types like 'size', 'min', 'max', etc., to provide more control over the data loaded from environment variables.
- **Support for Other File Types**: Currently, EnvSnatch supports `.env` files. We are looking to extend support to other file formats such as `.json`, `.yml`, etc., in future releases. This will allow more flexibility in how configuration data is structured and stored.

## Supported Data Types

EnvSnatch supports the following data types:

- `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`

## Optional Fields

All struct fields are considered required by default. Optional fields can be marked with `optional:"true"` to indicate that no error should be reported if the environment variable is not set.

Example:

```go
type Config struct {
// ...
OptionalField string `env:"OPTIONAL_FIELD" optional:"true"`
}
```

## Installation

```bash
go get github.com/mayckol/envsnatch
```

## Usage

Here's how to use EnvSnatch:

```go
package main

import (
	"fmt"
	"log"
	"github.com/mayckol/envsnatch"
)

type Config struct {
	Port int `env:"PORT"`
	// ... other fields
	OptionalField string `env:"OPTIONAL_FIELD" optional:"true"`
}

func main() {
	var cfg Config
	envReader, err := envsnatch.NewEnvSnatch()
	if err != nil {
		log.Fatal(err)
	}

	failedFields, err := envReader.Unmarshal(&cfg)
	// type UnmarshalingErr struct {
	//	Field  string
	//	Reason string
	//}
	
	// treat failedFields as a slice of UnmarshalingErr
	fmt.Println("Failed fields:", failedFields) // Failed fields is a slice of UnmarshalingErr
	if err != nil {
		// other errors
		log.Fatal(err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg)
}
```

## Contributing

Your contributions make the open-source community a great place to learn, inspire, and create. Any contributions you make to future features and enhancements are **greatly appreciated**.

- Fork the Project
- Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
- Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
- Push to the Branch (`git push origin feature/AmazingFeature`)
- Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
