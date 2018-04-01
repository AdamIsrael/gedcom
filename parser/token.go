package parser

// Token represents a lexical token.
// type Token int

const (
	// Special tokens
	// ILLEGAL Token = iota
	//
	// EOF
	// WS
	//
	// // Literals
	// IDENT // main
	//
	// // Misc characters
	// ASTERISK // *
	// COMMA    // ,
	//
	// // Keywords
	// SELECT
	// FROM

	// GEDCOM
	STATE_BEGIN = iota
	STATE_LEVEL
	STATE_TAG
	STATE_XREF
	STATE_VALUE
	STATE_ERROR

	STATE_END

	SEEK_TAG_OR_XREF
	SEEK_TAG
	SEEK_VALUE
)
