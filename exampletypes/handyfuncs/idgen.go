package handyfuncs

import (
	"fmt"
	"math/rand"
	"time"
)

// rand + time
// This function generates a random ID using the current time and a random number.
// It is not suitable for production use as it is not cryptographically secure.
func IdGen() string {
	fmt.Println("only for testing purposes, do not use in production")
	// Generate a random ID using a combination of current time and random number
	// This is just an example, you can use any ID generation logic you prefer
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
}
