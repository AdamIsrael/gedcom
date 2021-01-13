package gedcom

import (
	"bufio"
	// "bytes"
	"github.com/adamisrael/gedcom/parser"
	"github.com/adamisrael/gedcom/types"
	// "io/ioutil"
	"os"
)

// How do I export types from types.* here as gedcom.Gedcom?

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
