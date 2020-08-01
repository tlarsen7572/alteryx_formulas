package alteryx_formulas

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"regexp"
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

func NewCalculator(formula string, info RecordInfo) (*calculator, []error) {
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

	wordExp, _ := regexp.Compile(`([^\s]+)(?:\s+|$)`)
	secondListener := &secondPassListener{
		symbols: firstListener.symbols,
		calc: &calculator{
			recordInfo: info,
			wordExp:    wordExp,
		},
	}
	antlr.ParseTreeWalkerDefault.Walk(secondListener, tree)
	return secondListener.calc, errors.errs
}

func Calculate(formula string, info RecordInfo) (interface{}, []error) {
	calc, errs := NewCalculator(formula, info)
	if len(errs) > 0 {
		return nil, errs
	}
	return calc.Calculate()
}

type errorListener struct {
	errs []error
}

func (l *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errs = append(l.errs, fmt.Errorf(`line %v:%v: %v`, line, column, msg))
}

func (l *errorListener) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ bool, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
	panic("implement me")
}

func (l *errorListener) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _, _ int, _ *antlr.BitSet, _ antlr.ATNConfigSet) {
	panic("implement me")
}

func (l *errorListener) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _, _, _ int, _ antlr.ATNConfigSet) {
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
