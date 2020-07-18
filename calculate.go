package alteryx_formulas

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"time"
)

type Calculator interface {
	Calculate() interface{}
}

type RecordInfo interface {
	GetCurrentBool(fieldName string) (bool, bool, error)
	GetCurrentInt(fieldName string) (int, bool, error)
	GetCurrentFloat(fieldName string) (float64, bool, error)
	GetCurrentString(fieldName string) (string, bool, error)
	GetCurrentDate(fieldName string) (time.Time, bool, error)
	GetFieldTypeByName(fieldName string) (string, error)
}

func Calculate(formula string, info RecordInfo) (interface{}, []error) {
	inputStream := antlr.NewInputStream(formula)
	lexer := parser.NewAlteryxFormulasLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewAlteryxFormulasParser(tokens)
	p.RemoveErrorListeners()
	errors := &errorListener{}
	p.AddErrorListener(errors)
	tree := p.Expr()
	walker := antlr.ParseTreeWalker{}
	firstListener := &firstPassListener{recordInfo: info, symbols: make(map[antlr.ParserRuleContext]int)}
	walker.Walk(firstListener, tree)

	if len(errors.errs) > 0 {
		return nil, errors.errs
	}

	secondListener := &secondPassListener{symbols: firstListener.symbols, calc: &calculator{recordInfo: info}}
	antlr.ParseTreeWalkerDefault.Walk(secondListener, tree)

	if len(errors.errs) > 0 {
		return nil, errors.errs
	}

	calc := secondListener.calc
	result, errs := calc.Calculate()
	return result, errs
}

type errorListener struct {
	errs []error
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Errorf(`line %v:%v: %v`, line, column, msg))
}

func (l *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

func (l *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	panic("implement me")
}

func (l *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	panic("implement me")
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
