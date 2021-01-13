package parser_test

import (
	"bufio"
	"fmt"
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
	// h := types.Header{}
	// if g.Header == h {
	// 	t.Fatalf("Header was nil, expected types.Header")
	// }

	for _, i := range g.Individual {
		if i.Xref == "P1" {
			fmt.Printf("%#v\n", i)

			for _, n := range i.Name {
				fmt.Printf("Name: %q\n", n.Name)
			}
		}
	}

}
