// https://gobyexample.com/xml
//
// Go has built-in support for XML and similar formats in the encoding.xml package.

package main

import (
	"encoding/xml"
	"fmt"
)

// The Plant struct can be mapped to XML.
// Like the JSON example, field tags tell the encoder & decoder what to do.
// The XMLName field name dictates the name of the top-level XML element representing the struct.
// id,attr means that the Id field is an XML attribute rather than a nested element.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// Emit XML representing our plant, using MarshallIndent to produce human-readable output.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Explicitly add an XML header
	fmt.Println(xml.Header + string(out))

	// Parse a stream of bytes with XML into a data structure.
	// If the XML is malformed or can't be mapped onto a Plant, we'll get a nice error.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// The parent>child>plant directive tells the encoder to nest all Plants under <parent><child>
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
