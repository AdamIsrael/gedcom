package main

import (
	"flag"
	"fmt"
	// "github.com/adamisrael/gedcom"
	"gedcom"
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

	if *gedcomFile == nil {
		fmt.Println("Invalid gedcom")
		return
	}
	// Check if file exists

	// f, err := os.Open(*gedcomFile)
	// check(err)
	//
	// defer f.Close()

	g := gedcom.Gedcom(*gedcomFile)
	fmt.Println("Opened Gedcom")
	for _, i := range g.Individual {
		if i.Xref == "P1" {
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
}
