package forth

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

type token int

const (
	invalid token = iota
	eof
	ws

	word
	term

	plus
	minus
	times
	divide

	dup
	drop
	swap
	over
)

type scanner struct {
	r *bufio.Reader
}

func newScanner(r io.Reader) *scanner {
	return &scanner{r: bufio.NewReader(r)}
}

func (s *scanner) scan() (t token, l string) {
	ch := s.read()

	// if whitespace consume all following whitespace
	// if letter consume as word
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isDigit(ch) {
		s.unread()
		return s.scanTerm()
	} else if isWordLetter(ch) {
		s.unread()
		return s.scanWord()
	}

	if ch == endOfFile {
		return eof, ""
	}

	return invalid, string(ch)
}

func (s *scanner) scanWhitespace() (t token, l string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer
	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return ws, buf.String()
}

func (s *scanner) scanTerm() (t token, l string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent digit into the buffer
	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if !isDigit(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return term, buf.String()
}

func (s *scanner) scanWord() (t token, l string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent digit into the buffer
	for {
		if ch := s.read(); ch == endOfFile {
			break
		} else if !isWordLetter(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	// return specific words
	switch strings.ToUpper(buf.String()) {
	case "+":
		return plus, buf.String()
	case "-":
		return minus, buf.String()
	case "*":
		return times, buf.String()
	case "/":
		return divide, buf.String()
	case "DUP":
		return dup, buf.String()
	case "DROP":
		return drop, buf.String()
	case "SWAP":
		return swap, buf.String()
	case "OVER":
		return over, buf.String()
	}

	return word, buf.String()
}

func (s *scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return endOfFile
	}
	return ch
}

func (s *scanner) unread() {
	_ = s.r.UnreadRune()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isWordLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

var endOfFile = rune(0)

type parser struct {
	s   *scanner
	buf struct {
		tok token
		lit string
		n   int
	}
}

func newParser(r io.Reader) *parser {
	return &parser{s: newScanner(r)}
}

func (p *parser) scan() (tok token, lit string) {
	// if we have a token on buffer, then return it
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// otherwise read next token from scanner
	tok, lit = p.s.scan()

	// save it to buffer in case we unscan later
	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *parser) unscan() { p.buf.n = 1 }

func (p *parser) scanIgnoreWhitespace() (tok token, lit string) {
	tok, lit = p.scan()
	if tok == ws {
		tok, lit = p.scan()
	}
	return
}

func (p *parser) Parse() (forthstack, error) {
	stack := make(forthstack, 0)
	for {
		tok, lit := p.scanIgnoreWhitespace()
		if tok == eof {
			break
		}

		stack = append(stack, newNode(tok, lit))
	}

	return stack, nil
}

type node struct {
	tok token
	lit string
}

func newNode(tok token, lit string) *node {
	return &node{tok: tok, lit: lit}
}

type forthstack []*node

// Forth evaluates the given Forth inputs and returns the result.
func Forth(inputs []string) ([]int, error) {
	parser := newParser(strings.NewReader(inputs[0]))
	stack, _ := parser.Parse()
	return evaluate(stack)
}

func evaluate(input forthstack) ([]int, error) {
	stack := make([]int, 0)
	for _, n := range input {
		switch n.tok {
		case term:
			i, _ := strconv.Atoi(n.lit)
			stack = append(stack, i)
		case plus:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			p2 := stack[len(stack)-1]
			p1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			r := p1 + p2
			stack = append(stack, r)
		case minus:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			p2 := stack[len(stack)-1]
			p1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			r := p1 - p2
			stack = append(stack, r)
		case times:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			p2 := stack[len(stack)-1]
			p1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			r := p1 * p2
			stack = append(stack, r)
		case divide:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			p2 := stack[len(stack)-1]
			if p2 == 0 {
				return stack, errors.New("divide by zero")
			}
			p1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			r := p1 / p2
			stack = append(stack, r)
		case dup:
			if len(stack) < 1 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			stack = append(stack, stack[len(stack)-1])
		case drop:
			if len(stack) < 1 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			stack = stack[0 : len(stack)-1]
		case swap:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			p1 := stack[len(stack)-1]
			stack[len(stack)-1] = stack[len(stack)-2]
			stack[len(stack)-2] = p1
		case over:
			if len(stack) < 2 {
				return stack, errors.New("not enough numbers in stack to perform operation")
			}
			stack = append(stack, stack[len(stack)-2])
		}
	}

	return stack, nil
}
