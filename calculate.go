package alteryx_formulas

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
)

func Calculate(formula string) (NullableValue, error) {
	inputStream := antlr.NewInputStream(formula)
	lexer := parser.NewAlteryxFormulasLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewAlteryxFormulasParser(tokens)
	listener := &listener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Formula())
	result := listener.Calculate()
	return result, nil
}
