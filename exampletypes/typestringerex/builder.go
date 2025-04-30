package typestringerex

import (
	"errors"

)


// Builder is a generic function that creates an instance of a type based on the provided Type and value.
// It uses type assertion to convert the created instance to the desired type T.
// If the type assertion fails, it returns an error.
func Builder[T any](Type string, value string) (T, error) {
	var zero T // Declare a zero value of type T to return in case of errors

	if Type == "toml" {
		data := NewToML(value)
		if convertedData, ok := any(data).(T); ok {
			return convertedData, nil
		}
		return zero, errors.New("type assertion failed")
	}

	if Type == "binaryblob" {
		data, err := NewBinaryBlobFromString(value)
		if err != nil {
			return zero, err
		}
		if convertedData, ok := any(data).(T); ok {
			return convertedData, nil
		}
		return zero, errors.New("type assertion failed")
	}

	return zero, errors.New("unknown type")
}
