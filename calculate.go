package alteryx_formulas

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"time"
)

type RecordInfo interface {
	GetCurrentBool(fieldName string) (bool, bool, error)
	GetCurrentInt(fieldName string) (int, bool, error)
	GetCurrentFloat(fieldName string) (float64, bool, error)
	GetCurrentString(fieldName string) (string, bool, error)
	GetCurrentDate(fieldName string) (time.Time, bool, error)
	GetFieldTypeFromName(fieldName string) (string, error)
}

func Calculate(formula string, info RecordInfo) (NullableValue, error) {
	inputStream := antlr.NewInputStream(formula)
	lexer := parser.NewAlteryxFormulasLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewAlteryxFormulasParser(tokens)
	listener := &listener{recordInfo: info}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Formula())
	result := listener.Calculate()
	return result, nil
}

const (
	ByteType         = `Byte`
	BoolType         = `Bool`
	Int16Type        = `Int16`
	Int32Type        = `Int32`
	Int64Type        = `Int64`
	FixedDecimalType = `FixedDecimal`
	FloatType        = `Float`
	DoubleType       = `Double`
	StringType       = `String`
	WStringType      = `WString`
	V_StringType     = `V_String`
	V_WStringType    = `V_WString`
	DateType         = `Date`
	DateTimeType     = `DateTime`
)
