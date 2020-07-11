package alteryx_formulas

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
)

const (
	Null   = 0
	Number = 1
	String = 2
	Date   = 3
	Bool   = 4
)

type firstPassListener struct {
	symbols    map[interface{}]int
	recordInfo RecordInfo
	parser.BaseAlteryxFormulasListener
}

func (l *firstPassListener) getSymbol(c antlr.ParserRuleContext) (int, bool) {
	symbol, ok := l.symbols[c.GetStart()]
	return symbol, ok
}

func (l *firstPassListener) setSymbol(c antlr.ParserRuleContext, value int) {
	l.symbols[c.GetStart()] = value
}

func (l *firstPassListener) EnterNumberLiteral(c *parser.NumberLiteralContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) EnterStringLiteral(c *parser.StringLiteralContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) EnterDateLiteral(c *parser.DateLiteralContext) {
	l.setSymbol(c, Date)
}

func (l *firstPassListener) EnterDatetimeLiteral(c *parser.DatetimeLiteralContext) {
	l.setSymbol(c, Date)
}

func (l *firstPassListener) EnterBoolLiteral(c *parser.BoolLiteralContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) EnterNullFunc(c *parser.NullFuncContext) {
	l.setSymbol(c, Null)
}

func (l *firstPassListener) ExitAdd(c *parser.AddContext) {
	leftType, ok := l.getSymbol(c.GetLeft())
	if !ok {
		panic(`left symbol does not have a type`)
	}
	rightType, ok := l.getSymbol(c.GetRight())
	if !ok {
		panic(`right symbol does not have a type`)
	}
	if leftType == Null {
		l.setSymbol(c, rightType)
	} else {
		l.setSymbol(c, leftType)
	}
}

func (l *firstPassListener) EnterExprField(c *parser.ExprFieldContext) {
	text := c.GetText()
	fieldName := text[1 : len(text)-1]
	fieldType, err := l.recordInfo.GetFieldTypeByName(fieldName)
	if err != nil {
		c.SetException(MissingField(fieldName, c))
		return
	}
	switch fieldType {
	case ByteType, Int16Type, Int32Type, Int64Type, FixedDecimalType, FloatType, DoubleType:
		l.setSymbol(c, Number)
	case BoolType:
		l.setSymbol(c, Bool)
	case DateType, DateTimeType:
		l.setSymbol(c, Date)
	case StringType, WStringType, V_StringType, V_WStringType:
		l.setSymbol(c, String)
	default:
		c.SetException(InvalidFieldType(fieldName, fieldType, c))
	}
}

func MissingField(missingField string, c antlr.ParserRuleContext) FormulasException {
	return FormulasException{
		Message:        fmt.Sprintf(`field '%v' does not exist`, missingField),
		InputStream:    c.GetStart().GetInputStream(),
		OffendingToken: c.GetStart(),
	}
}

func InvalidFieldType(field string, fieldType string, c antlr.ParserRuleContext) FormulasException {
	return FormulasException{
		Message:        fmt.Sprintf(`field '%v' has type '%v' which cannot be used in formulas`, field, fieldType),
		InputStream:    c.GetStart().GetInputStream(),
		OffendingToken: c.GetStart(),
	}
}

func InvalidType(message string, c antlr.ParserRuleContext) FormulasException {
	return FormulasException{
		Message:        message,
		InputStream:    c.GetStart().GetInputStream(),
		OffendingToken: c.GetStart(),
	}
}

type FormulasException struct {
	Message        string
	InputStream    antlr.IntStream
	OffendingToken antlr.Token
}

func (e FormulasException) GetOffendingToken() antlr.Token {
	return e.OffendingToken
}

func (e FormulasException) GetMessage() string {
	return e.Message
}

func (e FormulasException) GetInputStream() antlr.IntStream {
	return e.InputStream
}
