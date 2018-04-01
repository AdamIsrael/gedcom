package parser

import (
	// "github.com/adamisrael/gedcom/parser"
	"strings"
	"testing"
)

type example struct {
	input string
	level int
	tag   string
	value string
	xref  string
}

var examples = []example{
	{string("1 SEX F\n"), 1, `SEX`, `F`, ""},
	{string(" 1 SEX F\n"), 1, `SEX`, `F`, ""},
	{string("  \r\n\t 1 SEX F\n"), 1, `SEX`, `F`, ""},
	{string("  \r\n\t 1     SEX      F\n"), 1, `SEX`, `F`, ""},
	{string("1 SEX F\r"), 1, `SEX`, `F`, ""},
	{string("1 SEX F \r"), 1, `SEX`, `F `, ""},
	{string("0 HEAD\r"), 0, `HEAD`, ``, ""},
	{string("0 @OTHER@ SUBM\n"), 0, `SUBM`, ``, "OTHER"},
	{string("2 DATE 13 März 1823\n"), 2, `DATE`, ``, "13 März 1823"},
}

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {

	for _, ex := range examples {
		s := NewScanner(strings.NewReader(ex.input))
		// s := &Scanner{}
		offset, err := s.Scan(ex.input)

		if err != nil {
			t.Fatalf(`Scan for "%s" returned error "%v", expected no error`, ex.input, err)
		}

		if offset == 0 {
			t.Fatalf(`Scan for "%s" did not find tag, expected it to find`, ex.input)
		}

		if s.level != ex.level {
			t.Errorf(`Scan for "%s" returned level %d, expected %d`, ex.input, s.level, ex.level)
		}
		//
		// if string(s.tag) != ex.tag {
		// 	t.Errorf(`nextTag for "%s" returned tag "%s", expected "%s"`, ex.input, s.tag, ex.tag)
		// }
		//
		// if string(s.value) != ex.value {
		// 	t.Errorf(`nextTag for "%s" returned value "%s", expected "%s"`, ex.input, s.value, ex.value)
		// }
		//
		// if string(s.xref) != ex.xref {
		// 	t.Errorf(`nextTag for "%s" returned xref "%s", expected "%s"`, ex.input, s.xref, ex.xref)
		// }
		s.Reset()

	}

	// var tests = []struct {
	// 	s   string
	// 	tok parser.Token
	// 	lit string
	// }{
	// 	// Special tokens (EOF, ILLEGAL, WS)
	// 	{s: ``, tok: parser.EOF},
	// 	{s: `#`, tok: parser.ILLEGAL, lit: `#`},
	// 	{s: ` `, tok: parser.WS, lit: " "},
	// 	{s: "\t", tok: parser.WS, lit: "\t"},
	// 	{s: "\n", tok: parser.WS, lit: "\n"},
	//
	// 	// Misc characters
	// 	{s: `*`, tok: parser.ASTERISK, lit: "*"},
	//
	// 	// Identifiers
	// 	{s: `foo`, tok: parser.IDENT, lit: `foo`},
	// 	{s: `Zx12_3U_-`, tok: parser.IDENT, lit: `Zx12_3U_`},
	//
	// 	// Keywords
	// 	{s: `FROM`, tok: parser.FROM, lit: "FROM"},
	// 	{s: `SELECT`, tok: parser.SELECT, lit: "SELECT"},
	// }
	//
	// for i, tt := range tests {
	// 	s := parser.NewScanner(strings.NewReader(tt.s))
	// 	tok, lit := s.Scan()
	// 	if tt.tok != tok {
	// 		t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
	// 	} else if tt.lit != lit {
	// 		t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
	// 	}
	// }
}
