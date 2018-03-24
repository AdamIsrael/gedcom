package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	gedcomFile := flag.String("gedcom", "", "The path to the GEDCOM file to analyze.")

	flag.Parse()

	f, err := os.Open(*gedcomFile)
	check(err)

	/**
	*	GEDCOM grammar rules for gedcom_line(s)
	*	Source: http://www.phpgedview.net/ged551-5.pdf

	- Long values can be broken into shorter GEDCOM lines by using a
	subordinate CONC or CONT tag. The CONC tag assumes that the accompanying
	subordinate value is concatenated to the previous line value without saving
	the carriage return prior to the line terminator. If a concatenated line is
	broken at a space, then the space must be carried over to the next line.
	The CONT assumes that the subordinate line value is concatenated to the
	previous line, after inserting a carriage return.

	- The beginning of a new logical record is designated by a line whose level number is 0 (zero).

	- Level numbers must be between 0 to 99 and must not contain leading zeroes, for example, level one must be 1, not 01.

	- Each new level number must be no higher than the previous line plus 1.

	- All GEDCOM lines have either a value or a pointer unless the line
	contains subordinate GEDCOM lines. The presence of a level number and a tag
	alone should not be used to assert data (i.e. 1 FLAG Y not just 1 FLAG to
	imply that the flag is set).

	- Logical GEDCOM record sizes should be constrained so that they will fit
	in a memory buffer of less than 32K. GEDCOM files with records sizes
	greater than 32K run the risk of not being able	to be loaded in some
	programs. Use of pointers to records, particularly NOTE records, should
	ensure that this limit will be sufficient.

	- Any length constraints are given in characters, not bytes. When wide
	characters (characters wider than 8 bits) are used, byte buffer lengths
	should be adjusted accordingly.

	- The cross-reference ID has a maximum of 22 characters, including the
	enclosing ‘at’ signs (@), and it must be unique within the GEDCOM
	transmission.

	- Pointers to records imply that the record pointed to does actually exists
	within the transmission. Future pointer structures may allow pointing to
	records within a public accessible database as an alternative.

	- The length of the GEDCOM TAG is a maximum of 31 characters, with the
	first 15 characters being unique.

	- The total length of a GEDCOM line, including level number,
	cross-reference number, tag, value, delimiters, and terminator, must not
	exceed 255 (wide) characters.

	- Leading white space (tabs, spaces, and extra line terminators) preceding
	a GEDCOM line should be ignored by the reading system. Systems generating
	GEDCOM should not place any white space in front of the GEDCOM line.

	*/

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		/**
		* gedcom_line syntax
		*
		* A GEDCOM line has the following syntax:
		* level + delim + [optional_xref_ID] + tag + [optional_line_value] + terminator
		*
		 */

		// Right out of the gate, INDI have an extra trailing space.
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "0") {

			// Record types:
			// HEAD
			if strings.HasSuffix(line, "HEAD") {
				// Parse the Header
			}

			// INDI
			if strings.HasSuffix(line, "INDI") {
				fmt.Println(line) // Println will add back the final '\n'
			}
			// SOUR
			// FAM
			// REPO
			// TRLR
			if strings.HasSuffix(line, "TRLR") {
				// Parse the Trailer
			}

			// CONC or CONT
			// fmt.Println(scanner.Text()) // Println will add back the final '\n'

		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	f.Close()
}
