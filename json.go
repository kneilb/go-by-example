// https://gobyexample.com/json
//
// Go offers built-in support for JSON encoding & decoding.
// We can use built-in and custom data types.
//
// More docs here: https://go.dev/blog/json
// And here: https://pkg.go.dev/encoding/json

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Two structs, to show encoding & decoding custom types.
type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields are encoded / decoded in JSON.
// Fields must start with Capital Letters to be exported.
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Encode basic types to JSON
	bolB, _ := json.Marshal(true)
	// fmt.Println(bolB) // This outputs [116 114 117 101] - but why?!
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	// fmt.Println(intB) // This outputs [49] - but why?!
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	// fmt.Println(fltB) // This outputs [50 46 51 52] - but why?!
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	// fmt.Println(strB) // This outputs [34 103 111 112 104 101 114 34] - but why?!
	fmt.Println(string(strB))

	// Slices encode to JSON arrays
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	// fmt.Println(slcD)
	// fmt.Println(slcB)
	fmt.Println(string(slcB))

	// Maps encode to JSON objects
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Automatic encoding of structs to JSON.
	// Only includes exported fields, and uses those names as the JSON keys by default.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "banana"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Use tags on struct field declarations to change the encoded JSON key names.
	// See the struct definition above for examples.
	res2D := &response2{
		Page:   2,
		Fruits: []string{"bread", "cheese", "sausage"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON data into Go values!
	// Example of a generic data structure.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// This holds a map of strings to arbitrary types.
	var dat map[string]interface{}

	// Do the decoding, and check for errors (badly).
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// To use the values in the map, we need to convert them to the right type.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accessing nested data requires a series of conversions.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// We can also decode JSON into custom data types.
	// This adds additional type safety, and removes the need for type assertions ("casts") when accessing the data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the previous examples, we used bytes or strings.
	// We can also stream JSON encodings directly to os.Writers like os.Stdout or HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "herrings": 7}
	enc.Encode(d)
}
