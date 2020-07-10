package alteryx_formulas

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
)

type secondPassListener struct {
	calc    *calculator
	symbols map[interface{}]int
	parser.BaseAlteryxFormulasListener
}

func (l *secondPassListener) getSymbol(c antlr.ParserRuleContext) (int, bool) {
	symbol, ok := l.symbols[c.GetStart()]
	return symbol, ok
}

func (l *secondPassListener) EnterExprField(c *parser.ExprFieldContext) {
	text := c.GetText()
	fieldName := text[1 : len(text)-1]
	fieldType, ok := l.getSymbol(c)
	if !ok {
		panic(`symbol not found`)
	}
	switch fieldType {
	case String:
		l.calc.pushStringField(fieldName)
	case Number:
		l.calc.pushNumberField(fieldName)
	default:
		panic("Invalid field type")
	}
}

func (l *secondPassListener) EnterNullFunc(c *parser.NullFuncContext) {
	l.calc.pushValueFunc(nil)
}

func (l *secondPassListener) EnterNumberLiteral(c *parser.NumberLiteralContext) {
	text := c.GetText()
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(fmt.Sprintf(`%v could not be parsed to a float64`, text))
	}
	l.calc.pushValueFunc(value)
}

func (l *secondPassListener) EnterStringLiteral(c *parser.StringLiteralContext) {
	text := c.GetText()
	value := text[1 : len(text)-1]
	l.calc.pushValueFunc(value)
}

func (l *secondPassListener) EnterAdd(c *parser.AddContext) {
	leftType, ok := l.getSymbol(c.GetLeft())
	if !ok {
		panic(`left symbol does not exist`)
	}
	rightType, ok := l.getSymbol(c.GetRight())
	if !ok {
		panic(`right symbol does not exist`)
	}

	if (leftType == Number && rightType == Number) || (leftType == Number && rightType == Null) || (leftType == Null && rightType == Number) {
		l.calc.pushFunction(l.calc.addNumbers)
		return
	}
	if (leftType == String && rightType == String) || (leftType == String && rightType == Null) || (leftType == Null && rightType == String) {
		l.calc.pushFunction(l.calc.concatenate)
		return
	}
	if leftType == Null && rightType == Null {
		l.calc.pushFunction(l.calc.addNumbers)
		return
	}
	panic(`invalid left and right type`)
}
