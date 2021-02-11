package parser_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/adamisrael/gedcom/parser"
)

// func init() {
// 	var err error
// 	data, err = ioutil.ReadFile("testdata/allged.ged")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseGedcom(t *testing.T) {

	f, err := os.Open("../testdata/simple.ged")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := parser.NewParser(bufio.NewReader(f))
	g, err := p.Parse()

	if err != nil {
		t.Fatalf("Result of decoding gedcom gave error, expected no error")
	}

	if g == nil {
		t.Fatalf("Result of decoding gedcom was nil, expected valid object")
	}
	// t.Fatalf("Testing")

	// for _, i := range g.Individual {
	// 	fmt.Println(i.Xref)

	// 	// for _, n := range i.Name {
	// 	// 	fmt.Printf("Name: %q\n", n.Name)
	// 	// }

	// 	// for _, event := range i.Event {
	// 	// 	fmt.Printf("Event: %#v\n", event)
	// 	// }

	// }

}
