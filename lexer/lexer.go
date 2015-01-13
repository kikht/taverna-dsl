
package lexer

import (
	
	// "fmt"
	// "gitlab.ict.sbras.ru/taverna/dsl/util"
	
	"io/ioutil"
	"unicode/utf8"
	"gitlab.ict.sbras.ru/taverna/dsl/token"
)

const(
	NoState = -1
	NumStates = 93
	NumSymbols = 105
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: '_'
1: '"'
2: '"'
3: 'n'
4: 'a'
5: 'm'
6: 'e'
7: ':'
8: 't'
9: 'y'
10: 'p'
11: 'e'
12: 's'
13: 'a'
14: 'n'
15: 'i'
16: 't'
17: 'i'
18: 'z'
19: 'e'
20: 'd'
21: 'p'
22: 'a'
23: 'r'
24: 'a'
25: 'l'
26: 'l'
27: 'e'
28: 'l'
29: 'i'
30: 'n'
31: 't'
32: 'e'
33: 'r'
34: 'p'
35: 'r'
36: 'e'
37: 't'
38: 'e'
39: 'r'
40: 's'
41: 'c'
42: 'h'
43: 'e'
44: 'd'
45: 'u'
46: 'l'
47: 'e'
48: 'r'
49: 's'
50: 'c'
51: 'r'
52: 'i'
53: 'p'
54: 't'
55: 'o'
56: 'u'
57: 't'
58: 'p'
59: 'u'
60: 't'
61: 's'
62: ','
63: 'w'
64: 'a'
65: 'i'
66: 't'
67: 'i'
68: 'n'
69: 'p'
70: 'u'
71: 't'
72: 's'
73: '*'
74: '<'
75: '('
76: ')'
77: '='
78: 't'
79: 'r'
80: 'u'
81: 'e'
82: 'f'
83: 'a'
84: 'l'
85: 's'
86: 'e'
87: 't'
88: 'f'
89: '0'
90: '1'
91: '_'
92: '_'
93: '-'
94: '.'
95: '/'
96: ' '
97: '\t'
98: '\n'
99: '\r'
100: '1'-'9'
101: '0'-'9'
102: 'a'-'z'
103: 'A'-'Z'
104: .

*/
