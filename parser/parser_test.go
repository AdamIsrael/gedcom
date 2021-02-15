package parser_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/adamisrael/gedcom/parser"
)

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
}
