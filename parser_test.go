package main

import (
	"github.com/gobs/pretty"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/ast"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/lexer"
	"gitlab.ict.sbras.ru/taverna/taverna-dsl/parser"
	"testing"
)

func TestParser(t *testing.T) {
	p := parser.NewParser()
	l, err := lexer.NewLexerFile("test.processor")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	res, err := p.Parse(l)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}
	processor := res.(*ast.Processor)
	str := pretty.PrettyFormat(processor)
	t.Log(str)
}

