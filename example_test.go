package guid_test

import (
	"encoding/hex"
	"fmt"

	guid "github.com/bsm/go-guid"
)

func Example_New96() {
	// Create a new 12-byte globally-unique identifier
	id := guid.New96()
	fmt.Println(hex.EncodeToString(id.Bytes()))
}

func Example_New128() {
	// Create a new 16-byte globally-unique identifier
	id := guid.New128()
	fmt.Println(hex.EncodeToString(id.Bytes()))
}

func Example_NextPUID() {
	// Create a new 8-byte pseudo-unique identifier
	id := guid.NextPUID()
	fmt.Println(id)
}
