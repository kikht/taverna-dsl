
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		"name",
		":",
		"type",
		"sanitized",
		"parallel",
		"interpreter",
		"scheduler",
		"script",
		"outputs",
		",",
		"wait",
		"inputs",
		"*",
		"<",
		"(",
		")",
		"=",
		"id",
		"str_lit",
		"path",
		"true",
		"false",
		"t",
		"f",
		"0",
		"1",
		"pos_int",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"empty": 2,
		"name": 3,
		":": 4,
		"type": 5,
		"sanitized": 6,
		"parallel": 7,
		"interpreter": 8,
		"scheduler": 9,
		"script": 10,
		"outputs": 11,
		",": 12,
		"wait": 13,
		"inputs": 14,
		"*": 15,
		"<": 16,
		"(": 17,
		")": 18,
		"=": 19,
		"id": 20,
		"str_lit": 21,
		"path": 22,
		"true": 23,
		"false": 24,
		"t": 25,
		"f": 26,
		"0": 27,
		"1": 28,
		"pos_int": 29,
	},
}

