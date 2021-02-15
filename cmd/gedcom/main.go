package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/adamisrael/gedcom"
	// "gedcom"
	"github.com/araddon/dateparse"
)

var dateFormats = []string{
	"12 Feb 2006",
	"03 February 2013",
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// verbose := flag.Bool("verbose", false, "verbosity")
	gedcomFile := flag.String("gedcom", "", "The path to the GEDCOM file to analyze.")
	format := flag.String("format", "", "The format to output in (text, json).")
	anniversary := flag.Bool("anniversary", false, "Print all individuals who have an anniverary on this date.")

	flag.Parse()

	if *gedcomFile == "" {
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

				// for _, n := range i.Event {
				// 	fmt.Printf("Event: %q\n", n.Date)
				// }
			}

			// fmt.Printf("%+v\n", i)
			// fmt.Printf("%#v\n", i)
		}

		if *anniversary {

			currentTime := time.Now()

			for _, event := range i.Event {
				if len(event.Date) > 0 {
					t, err := dateparse.ParseLocal(event.Date)
					if err == nil {
						year, month, day := dateDiff(t, currentTime)
						if month == 0 && day == 0 {
							switch event.Tag {
							case "BIRT":
								fmt.Printf("%s was born on this date %d years ago\n", i.Name[0].Name, year)
							}

						}
					}
				}
			}
		}
	}
}

func dateDiff(a, b time.Time) (year, month, day int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)

	// Normalize negative values
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
