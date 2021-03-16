package gedcom

import (
	"bufio"
	"errors"
	"os"

	"github.com/adamisrael/gedcom/parser"
	"github.com/adamisrael/gedcom/types"
)

// How do I export types from types.* here as gedcom.Gedcom?

// OpenGedcom will open a filename, if it exists, and parse it as a GEDCOM
func OpenGedcom(filename string) (*types.Gedcom, error) {

	if _, err := os.Stat(filename); err == nil {
		f, err := os.Open(filename)
		check(err)

		defer f.Close()

		p := parser.NewParser(bufio.NewReader(f))
		g, err := p.Parse()
		check(err)

		return g, err
	}
	return nil, errors.New("invalid GEDCOM file")
}

// Gedcom is the main entrypoint
func Gedcom(gedcomFile string) *types.Gedcom {

	f, err := os.Open(gedcomFile)
	check(err)

	defer f.Close()

	p := parser.NewParser(bufio.NewReader(f))
	g, err := p.Parse()
	check(err)

	return g
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
