package parser

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

import (
	"fmt"
	"io"
	"strings"

	"github.com/adamisrael/gedcom/types"
)

// Parser represents a parser.
type Parser struct {
	s    *Scanner
	refs map[string]interface{}
	// r io.Reader
	// buf struct {
	// 	// tok Token  // last read token
	// 	lit string // last read literal
	// 	n   int    // buffer size (max=1)
	// }
	parsers []parser
	// Individual []types.Individual
}

type parser func(level int, tag string, value string, xref string) error

func (P *Parser) pushParser(p parser) {
	P.parsers = append(P.parsers, p)
}

func (p *Parser) popParser(level int, tag string, value string, xref string) error {
	n := len(p.parsers) - 1
	if n < 1 {
		panic("MASSIVE ERROR") // TODO
	}
	p.parsers = p.parsers[0:n]

	return p.parsers[len(p.parsers)-1](level, tag, value, xref)
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*types.Gedcom, error) {
	g := &types.Gedcom{
		Header: types.Header{ID: "test"},
	}

	p.refs = make(map[string]interface{})
	p.parsers = []parser{makeRootParser(p, g)}
	p.Scan(g)

	return g, nil
}

func (p *Parser) Scan(g *types.Gedcom) {
	s := &Scanner{}

	// buf := make([]byte, 512)
	// buf := ""
	pos := 0
	for {
		line, err := p.s.r.ReadString('\n')
		if err != nil {
			// TODO
		}
		// fmt.Print(line)
		s.Reset()
		offset, err := s.Scan(line)
		pos += offset
		if err != nil {
			if err != io.EOF {
				println(err.Error())
				return
			}
			break
		}

		p.parsers[len(p.parsers)-1](s.level, string(s.tag), string(s.value), string(s.xref))

		// switch s.level {
		// case 0:
		// 	switch s.tag {
		// 	case "HEAD":
		// 		print("asfd\n")
		//
		// 	}
		// }
		// fmt.Printf("%d %s %s\n", s.level, s.tag, s.value)
		// d.parsers[len(d.parsers)-1](s.level, string(s.tag), string(s.value), string(s.xref))

	}
}

func makeRootParser(p *Parser, g *types.Gedcom) parser {
	return func(level int, tag string, value string, xref string) error {
		// println(level, tag, value, xref)
		if level == 0 {
			switch tag {
			case "HEAD":
				// println(level, tag, value, xref)
				obj := p.head(xref)
				// println("obj: %x", obj)
				fmt.Printf("%#v\n", obj)
			case "INDI":
				obj := p.individual(xref)
				g.Individual = append(g.Individual, obj)
				p.pushParser(makeIndividualParser(p, obj, level))
			// fmt.Printf("%#v\n", obj)
			case "SUBM":
				g.Submitter = append(g.Submitter, &types.Submitter{})
			case "FAM":
				obj := p.family(xref)
				g.Family = append(g.Family, obj)
				p.pushParser(makeFamilyParser(p, obj, level))
			case "SOUR":
				obj := p.source(xref)
				g.Source = append(g.Source, obj)
				//d.pushParser(makeSourceParser(d, s, level))
			}
		}
		return nil
	}
}

func (p *Parser) head(xref string) *types.Header {
	// if xref == "" {
	// 	println("xref not set")
	// 	return &types.Header{}
	// }
	ref, _ := p.refs[xref].(*types.Header)
	// if !found {
	// 	fmt.Printf("%s not found", xref)
	// }
	return ref
}

func makeHeaderParser(p *Parser, h *types.Header, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "CHAR":
			h.CharacterSet.Name = value
		}
		return nil
	}
}

func (p *Parser) individual(xref string) *types.Individual {
	if xref == "" {
		return &types.Individual{}
	}

	ref, found := p.refs[xref].(*types.Individual)
	if !found {
		rec := &types.Individual{Xref: xref}
		p.refs[rec.Xref] = rec
		return rec
	}
	return ref

}
func makeIndividualParser(p *Parser, i *types.Individual, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "FAMC":
			family := p.family(stripXref(value))
			f := &types.FamilyLink{Family: family}
			i.Parents = append(i.Parents, f)
			p.pushParser(makeFamilyLinkParser(p, f, level))

		case "FAMS":
			family := p.family(stripXref(value))
			f := &types.FamilyLink{Family: family}
			i.Family = append(i.Family, f)
			p.pushParser(makeFamilyLinkParser(p, f, level))

		case "BIRT", "CHR", "DEAT", "BURI", "CREM", "ADOP", "BAPM", "BARM", "BASM", "BLES", "CHRA", "CONF", "FCOM", "ORDN", "NATU", "EMIG", "IMMI", "CENS", "PROB", "WILL", "GRAD", "RETI", "EVEN":
			e := &types.Event{Tag: tag, Value: value}
			i.Event = append(i.Event, e)
			p.pushParser(makeEventParser(p, e, level))

		case "CAST", "DSCR", "EDUC", "IDNO", "NATI", "NCHI", "NMR", "OCCU", "PROP", "RELI", "RESI", "SSN", "TITL", "FACT":
			e := &types.Event{Tag: tag, Value: value}
			i.Attribute = append(i.Attribute, e)
			p.pushParser(makeEventParser(p, e, level))

		case "NAME":
			// The Gedcom stores the name as "First Middle /Last/". Store the
			// original, but parse out the given and surname, too.
			var given, surname, suffix string

			/*
			*	Given the following examples:
			*	a) "Adam Michael /Israel/"
			*	b) "Adam Michael"
			*	c) "/Israel/"
			*   d) "Adam Michael /Israel/ Sr"
			*
			*	a, c, and d will return a three element slice:
			*		0 holding the given name(s)
			*		1 holding the surname
			*		2 holding the suffix.
			*	b will return a single element slice with just the given name.
			*
			 */

			names := strings.Split(value, "/")
			given = strings.TrimSpace(names[0])
			if len(names) == 3 {
				surname = names[1]
				suffix = names[2]
			}

			n := &types.Name{
				Name:    value,
				Given:   given,
				Surname: surname,
				Suffix:  suffix}

			i.Name = append(i.Name, n)
			p.pushParser(makeNameParser(p, n, level))
		case "SEX":
			i.Sex = value
		}
		return nil
	}

}

func makeNameParser(p *Parser, n *types.Name, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {

		case "SOUR":
			c := &types.Citation{Source: p.source(stripXref(value))}
			n.Citation = append(n.Citation, c)
			p.pushParser(makeCitationParser(p, c, level))
		case "NOTE":
			r := &types.Note{Note: value}
			n.Note = append(n.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		}

		return nil
	}
}

func makeSourceParser(p *Parser, s *types.Source, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "TITL":
			s.Title = value
			p.pushParser(makeTextParser(p, &s.Title, level))

		case "NOTE":
			r := &types.Note{Note: value}
			s.Note = append(s.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		}

		return nil
	}
}

func makeCitationParser(p *Parser, c *types.Citation, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "PAGE":
			c.Page = value
		case "QUAY":
			c.Quay = value
		case "NOTE":
			r := &types.Note{Note: value}
			c.Note = append(c.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		case "DATA":
			p.pushParser(makeDataParser(p, &c.Data, level))
		}

		return nil
	}
}

func makeNoteParser(p *Parser, n *types.Note, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "CONT":
			n.Note = n.Note + "\n" + value
		case "CONC":
			n.Note = n.Note + value
		case "SOUR":
			c := &types.Citation{Source: p.source(stripXref(value))}
			n.Citation = append(n.Citation, c)
			p.pushParser(makeCitationParser(p, c, level))
		}
		return nil
	}
}

func makeTextParser(p *Parser, s *string, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "CONT":
			*s = *s + "\n" + value
		case "CONC":
			*s = *s + value
		}

		return nil
	}
}

func makeDataParser(p *Parser, r *types.Data, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "DATE":
			r.Date = value
		case "TEXT":
			r.Text = append(r.Text, value)
			p.pushParser(makeTextParser(p, &r.Text[len(r.Text)-1], level))
		}

		return nil
	}
}

func makePlaceParser(p *Parser, pl *types.Place, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {

		case "SOUR":
			c := &types.Citation{Source: p.source(stripXref(value))}
			pl.Citation = append(pl.Citation, c)
			p.pushParser(makeCitationParser(p, c, level))
		case "NOTE":
			r := &types.Note{Note: value}
			pl.Note = append(pl.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		}

		return nil
	}
}

func makeFamilyLinkParser(p *Parser, f *types.FamilyLink, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "PEDI":
			f.Type = value
		case "NOTE":
			r := &types.Note{Note: value}
			f.Note = append(f.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		}

		return nil
	}
}

func makeFamilyParser(p *Parser, f *types.Family, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "HUSB":
			f.Husband = p.individual(stripXref(value))
		case "WIFE":
			f.Wife = p.individual(stripXref(value))
		case "CHIL":
			f.Child = append(f.Child, p.individual(stripXref(value)))
		case "ANUL", "CENS", "DIV", "DIVF", "ENGA", "MARR", "MARB", "MARC", "MARL", "MARS", "EVEN":
			e := &types.Event{Tag: tag, Value: value}
			f.Event = append(f.Event, e)
			p.pushParser(makeEventParser(p, e, level))
		}
		return nil
	}
}

func makeAddressParser(p *Parser, a *types.Address, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "CONT":
			a.Full = a.Full + "\n" + value
		case "ADR1":
			a.Line1 = value
		case "ADR2":
			a.Line2 = value
		case "CITY":
			a.City = value
		case "STAE":
			a.State = value
		case "POST":
			a.PostalCode = value
		case "CTRY":
			a.Country = value
		case "PHON":
			a.Phone = append(a.Phone, value)
		case "EMAIL":
			a.Email = append(a.Email, value)
		case "FAX":
			a.Fax = append(a.Fax, value)
		case "WWW":
			a.WWW = append(a.WWW, value)

		}

		return nil
	}
}

func makeEventParser(p *Parser, e *types.Event, minLevel int) parser {
	return func(level int, tag string, value string, xref string) error {
		if level <= minLevel {
			return p.popParser(level, tag, value, xref)
		}
		switch tag {
		case "TYPE":
			e.Type = value
		case "DATE":
			e.Date = value
		case "PLAC":
			e.Place.Name = value
			p.pushParser(makePlaceParser(p, &e.Place, level))
		case "ADDR":
			e.Address.Full = value
			p.pushParser(makeAddressParser(p, &e.Address, level))
		case "SOUR":
			c := &types.Citation{Source: p.source(stripXref(value))}
			e.Citation = append(e.Citation, c)
			p.pushParser(makeCitationParser(p, c, level))
		case "NOTE":
			r := &types.Note{Note: value}
			e.Note = append(e.Note, r)
			p.pushParser(makeNoteParser(p, r, level))
		}

		return nil
	}
}

func (p *Parser) source(xref string) *types.Source {
	if xref == "" {
		return &types.Source{}
	}

	ref, found := p.refs[xref].(*types.Source)
	if !found {
		rec := &types.Source{Xref: xref}
		p.refs[rec.Xref] = rec
		return rec
	}
	return ref
}

func (p *Parser) family(xref string) *types.Family {
	if xref == "" {
		return &types.Family{}
	}

	ref, found := p.refs[xref].(*types.Family)
	if !found {
		rec := &types.Family{Xref: xref}
		p.refs[rec.Xref] = rec
		return rec
	}
	return ref
}

func stripXref(value string) string {
	return strings.Trim(value, "@")
}
