// https://gobyexample.com/b64-encoding

package main

// Import the encoding/base64 package, aliased as b64.
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Test string
	data := "abc123!?$*&()'-=@~"

	// Encode with "standard" base64.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 2nd return value is an error, which you needn't check if you know the input is well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// Try again, with a URL-compatible format.
	// Trailing - instead of trailing +!
	// They both decode to the original string, though!
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
