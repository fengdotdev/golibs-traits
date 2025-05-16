# golibs-traits

A Go library providing reusable traits (interfaces and patterns) to enhance modularity and code reusability in Go projects.

## Overview

`golibs-traits` is a collection of Go interfaces and patterns designed to simplify common programming tasks. It promotes modularity by offering reusable "traits" that can be implemented across different types, enabling flexible and maintainable codebases. This library is ideal for developers looking to standardize behaviors like serialization, validation, or iteration in their Go applications.

## Installation

To use `golibs-traits` in your Go project, run the following command:

```bash
go get github.com/fengdotdev/golibs-traits
```

Ensure you have Go 1.18 or later installed, as this library may leverage generics or other modern Go features.

## Usage

The library provides a set of interfaces (traits) that you can implement in your types. Below is a general guide to using `golibs-traits`:

1. **Import the library**:
   ```go
   import "github.com/fengdotdev/golibs-traits"
   ```

2. **Implement a trait**:
   Choose a trait (e.g., `Serializable`, `Validatable`) and implement its methods in your type.

3. **Use the trait**:
   Pass your type to functions or structs that expect the trait interface.

### Example

Suppose `golibs-traits` provides a `Serializable` trait for types that can be converted to JSON. Here's how you might use it:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fengdotdev/golibs-traits"
)

type User struct {
	Name string
	Age  int
}

// Implement the Serializable trait
func (u User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func main() {
	user := User{Name: "Alice", Age: 30}
	var serializable golibs_traits.Serializable = user

	data, err := serializable.ToJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON:", string(data))
}
```

## Available Traits

Below is a list of traits provided by `golibs-traits` (update this based on the actual interfaces in the repo):

- **`Serializable`**: For types that can be serialized (e.g., to JSON, YAML).
- **`Validatable`**: For types that support self-validation.
- **`Iterator`**: For types that can be iterated over.


## API Documentation

Detailed API documentation is coming soon. For now, check the source code comments for method signatures and usage details.

## Contributing

We welcome contributions! To contribute to `golibs-traits`:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -m "Add your feature"`).
4. Push to your branch (`git push origin feature/your-feature`).
5. Open a Pull Request.

Please ensure your code follows the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and includes tests.

## Testing

To run the tests for `golibs-traits`, use:

```bash
go test ./...
```

Ensure all tests pass before submitting a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, open an issue on the [GitHub repository](https://github.com/fengdotdev/golibs-traits) or reach out to the maintainer at [feng@vitruvio.cl].
