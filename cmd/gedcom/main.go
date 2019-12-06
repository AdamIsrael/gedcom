package main

import (
	// "bufio"
	"flag"
	"fmt"
	// "github.com/adamisrael/gedcom/parser"
	"github.com/adamisrael/gedcom"
	// "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// verbose := flag.Bool("verbose", false, "verbosity")
	gedcomFile := flag.String("gedcom", "", "The path to the GEDCOM file to analyze.")
	format := flag.String("format", "", "The format to output in (text, json).")

	flag.Parse()

	// Check if file exists

	// f, err := os.Open(*gedcomFile)
	// check(err)
	//
	// defer f.Close()

	g := gedcom.Gedcom(*gedcomFile)
	fmt.Println("Opened Gedcom")
	for _, i := range g.Individual {
		if i.Xref == "P1" {
			// if *verbose {
			// 	fmt.Printf("%#v\n", i)

			// }

			if *format == "json" {

				fmt.Println("JSON selected")
				fmt.Println(i.JSON())
			} else {
				for _, n := range i.Name {
					fmt.Printf("Name: %q\n", n.Name)
				}

				for _, n := range i.Event {
					fmt.Printf("Event: %q\n", n.Date)
				}
			}

			fmt.Printf("%+v\n", i)
			fmt.Printf("%#v\n", i)
		}
	}

	// p := parser.NewParser(bufio.NewReader(f))
	// g, err := p.Parse()
	// check(err)
	//
	// // fmt.Printf("Gedcom: %#v\n", g)
	//
	// for _, i := range g.Individual {
	// 	if i.Xref == "P1" {
	// 		if *verbose {
	// 			fmt.Printf("%#v\n", i)
	//
	// 		}
	//
	// 		for _, n := range i.Name {
	// 			fmt.Printf("Name: %q\n", n.Name)
	// 		}
	// 	}
	// }

}
