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
	if fieldType == String {
		l.calc.pushStringField(fieldName)
	}
}

func (l *secondPassListener) EnterNumberLiteral(c *parser.NumberLiteralContext) {
	text := c.GetText()
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(fmt.Sprintf(`%v could not be parsed to a float64`, text))
	}
	l.calc.pushNumberFunc(number(value))
}

func (l *secondPassListener) EnterStringLiteral(c *parser.StringLiteralContext) {
	text := c.GetText()
	value := text[1 : len(text)-1]
	l.calc.pushStringFunc(stringVal(value))
}

func (l *secondPassListener) ExitFormula(c *parser.FormulaContext) {
	fieldType, ok := l.getSymbol(c)
	if !ok {
		panic(`no field type found for formula`)
	}
	switch fieldType {
	case Number:
		l.calc.getResult = l.calc.returnNumber
	case String:
		l.calc.getResult = l.calc.returnString
	case Bool:
		l.calc.getResult = l.calc.returnBool
	default:
		panic(`invalid formula field type`)
	}
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
	if leftType == Number && rightType == Number {
		l.calc.pushFunction(l.calc.addNumbers)
		return
	}
	if leftType == String && rightType == String {
		l.calc.pushFunction(l.calc.concatenate)
		return
	}
	panic(`invalid left and right type`)
}
