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

func (l *firstPassListener) EnterNumberLiteral(c *parser.NumberLiteralContext) {
	l.symbols[c.GetStart()] = Number
}

func (l *firstPassListener) EnterStringLiteral(c *parser.StringLiteralContext) {
	l.symbols[c.GetStart()] = String
}

func (l *firstPassListener) EnterDateLiteral(c *parser.DateLiteralContext) {
	l.symbols[c.GetStart()] = Date
}

func (l *firstPassListener) EnterDatetimeLiteral(c *parser.DatetimeLiteralContext) {
	l.symbols[c.GetStart()] = Date
}

func (l *firstPassListener) EnterBoolLiteral(c *parser.BoolLiteralContext) {
	l.symbols[c.GetStart()] = Bool
}

func (l *firstPassListener) EnterNullFunc(c *parser.NullFuncContext) {
	l.symbols[c.GetStart()] = Null
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
		l.symbols[c.GetStart()] = Number
	case BoolType:
		l.symbols[c.GetStart()] = Bool
	case DateType, DateTimeType:
		l.symbols[c.GetStart()] = Date
	case StringType, WStringType, V_StringType, V_WStringType:
		l.symbols[c.GetStart()] = String
	default:
		c.SetException(InvalidFieldType(fieldName, fieldType, c))
	}
}

func MissingField(missingField string, c antlr.ParserRuleContext) FieldException {
	return FieldException{
		Message:        fmt.Sprintf(`field '%v' does not exist`, missingField),
		InputStream:    c.GetStart().GetInputStream(),
		OffendingToken: c.GetStart(),
	}
}

func InvalidFieldType(field string, fieldType string, c antlr.ParserRuleContext) FieldException {
	return FieldException{
		Message:        fmt.Sprintf(`field '%v' has type '%v' which cannot be used in formulas`, field, fieldType),
		InputStream:    c.GetStart().GetInputStream(),
		OffendingToken: c.GetStart(),
	}
}

type FieldException struct {
	Message        string
	InputStream    antlr.IntStream
	OffendingToken antlr.Token
}

func (e FieldException) GetOffendingToken() antlr.Token {
	return e.OffendingToken
}

func (e FieldException) GetMessage() string {
	return e.Message
}

func (e FieldException) GetInputStream() antlr.IntStream {
	return e.InputStream
}
