package parser

import (
	"bufio"
	// "bytes"
	"fmt"
	"io"
	"strconv"
	// "strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r          *bufio.Reader
	parseState int
	tokenStart int
	level      int
	tag        string
	value      string
	xref       string
}

func (s *Scanner) Reset() {
	s.parseState = STATE_BEGIN
	s.tokenStart = 0
	s.level = 0
	// s.xref = make([]byte, 0)
	// s.tag = make([]byte, 0)
	// s.value = make([]byte, 0)
	s.xref = ""
	s.tag = ""
	s.value = ""
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
// func (s *Scanner) Scan() (tok Token, lit string) {

// Scan returns the next tag
func (s *Scanner) Scan(data string) (offset int, err error) {
	for i, c := range data {
		switch s.parseState {
		case STATE_BEGIN:
			switch {
			case c >= '0' && c <= '9':
				s.tokenStart = i
				s.parseState = STATE_LEVEL
				// fmt.Printf("Found level %c\n", c)
			case isWhitespace(c):
				continue
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Found non-whitespace before level: %q", data)
				return
			}
		case STATE_LEVEL:
			switch {
			case c >= '0' && c <= '9':
				continue
			case c == ' ':
				parsedLevel, perr := strconv.ParseInt(string(data[s.tokenStart:i]), 10, 64)
				if perr != nil {
					err = perr
					return
				}
				s.level = int(parsedLevel)
				s.parseState = SEEK_TAG_OR_XREF
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Level contained non-numerics")
				return

			}
		case SEEK_TAG:
			switch {
			case isAlphaNumeric(c):
				s.tokenStart = i
				s.parseState = STATE_TAG
			case c == ' ':
				continue
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Tag \"%s\" contained non-alphanumeric", string(data[s.tokenStart:i]))
				return
			}
		case SEEK_TAG_OR_XREF:
			switch {
			case isAlphaNumeric(c):
				s.tokenStart = i
				s.parseState = STATE_TAG
			case c == '@':
				s.tokenStart = i
				s.parseState = STATE_XREF
			case c == ' ':
				continue
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Xref \"%s\" contained non-alphanumeric", string(data[s.tokenStart:i]))
				return
			}
		case STATE_TAG:
			switch {
			case isAlphaNumeric(c):
				continue
			case c == '\n' || c == '\r':
				s.tag = data[s.tokenStart:i]
				s.parseState = STATE_END
				offset = i
				return
			case c == ' ':
				s.tag = data[s.tokenStart:i]
				s.parseState = SEEK_VALUE
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Tag contained non-alphanumeric")
				return
			}

		case STATE_XREF:
			switch {
			case isAlphaNumeric(c) || c == '@':
				continue
			case c == ' ':
				s.xref = data[s.tokenStart+1 : i-1]
				s.parseState = SEEK_TAG
			default:
				s.parseState = STATE_ERROR
				err = fmt.Errorf("Xref contained non-alphanumeric \"%c\"", c)
				return
			}

		case SEEK_VALUE:
			switch {
			case c == '\n' || c == '\r':
				s.parseState = STATE_END
				offset = i
				return
			case c == ' ':
				continue
			default:
				s.tokenStart = i
				s.parseState = STATE_VALUE
			}

		case STATE_VALUE:
			switch {
			case c == '\n' || c == '\r':
				s.value = data[s.tokenStart:i]
				s.parseState = STATE_END
				offset = i
				return
			default:
				continue
			}
		}
	}

	return 0, io.EOF
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
// func (s *Scanner) read() rune {
// 	ch, _, err := s.r.ReadRune()
// 	if err != nil {
// 		return eof
// 	}
// 	return ch
// }

// unread places the previously read rune back on the reader.
// func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

// isLetter returns true if the rune is a letter.
func isAlpha(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' }

// isDigit returns true if the rune is a digit.
func isNumeric(ch rune) bool { return (ch >= '0' && ch <= '9') }

func isAlphaNumeric(ch rune) bool { return isAlpha(ch) || isNumeric(ch) }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
