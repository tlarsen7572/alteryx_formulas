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
	symbols    map[antlr.ParserRuleContext]int
	recordInfo RecordInfo
	parser.BaseAlteryxFormulasListener
}

func (l *firstPassListener) getSymbol(c antlr.ParserRuleContext) (int, bool) {
	symbol, ok := l.symbols[c]
	return symbol, ok
}

func (l *firstPassListener) setSymbol(c antlr.ParserRuleContext, value int) {
	l.symbols[c] = value
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

func (l *firstPassListener) ExitNullFunc(c *parser.NullFuncContext) {
	l.setSymbol(c, Null)
}

func (l *firstPassListener) ExitAdd(c *parser.AddContext) {
	left, right := l.checkAndReturnLeftRightTypes(c)
	l.setLeftRightSymbol(c, left, right)
}

func (l *firstPassListener) ExitSubtract(c *parser.SubtractContext) {
	left, right := l.checkAndReturnLeftRightTypes(c)
	l.setLeftRightSymbol(c, left, right)
}

func (l *firstPassListener) ExitMultiply(c *parser.MultiplyContext) {
	left, right := l.checkAndReturnLeftRightTypes(c)
	l.setLeftRightSymbol(c, left, right)
}

func (l *firstPassListener) ExitDivide(c *parser.DivideContext) {
	left, right := l.checkAndReturnLeftRightTypes(c)
	l.setLeftRightSymbol(c, left, right)
}

func (l *firstPassListener) checkAndReturnLeftRightTypes(c HasLeftRightTypes) (int, int) {
	leftType, ok := l.getSymbol(c.GetLeft())
	if !ok {
		panic(`left symbol does not have a type`)
	}
	rightType, ok := l.getSymbol(c.GetRight())
	if !ok {
		panic(`right symbol does not have a type`)
	}
	return leftType, rightType
}

func (l *firstPassListener) setLeftRightSymbol(c antlr.ParserRuleContext, leftType int, rightType int) {
	if leftType == Null {
		l.setSymbol(c, rightType)
	} else {
		l.setSymbol(c, leftType)
	}
}

func (l *firstPassListener) ExitEqual(c *parser.EqualContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitGreaterThan(c *parser.GreaterThanContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitGreaterEqual(c *parser.GreaterEqualContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitLessThan(c *parser.LessThanContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitLessEqual(c *parser.LessEqualContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitNotEqual(c *parser.NotEqualContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitExprIf(c *parser.ExprIfContext) {
	thenType, ok := l.getSymbol(c.Expr(1))
	if !ok {
		panic(`then symbol does not have a type`)
	}
	l.setSymbol(c, thenType)
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

func (l *firstPassListener) ExitIn(c *parser.InContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitNotIn(c *parser.NotInContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitMinFunc(c *parser.MinFuncContext) {
	exprs := c.AllExpr()
	exprCount := len(exprs)
	for i := 0; i < exprCount; i++ {
		argType, ok := l.getSymbol(exprs[i])
		if !ok {
			panic(`min argument does not have a type`)
		}
		if argType == Null {
			continue
		}
		l.setSymbol(c, argType)
		return
	}
	l.setSymbol(c, Null)
}

func (l *firstPassListener) ExitMaxFunc(c *parser.MaxFuncContext) {
	exprs := c.AllExpr()
	exprCount := len(exprs)
	for i := 0; i < exprCount; i++ {
		argType, ok := l.getSymbol(exprs[i])
		if !ok {
			panic(`min argument does not have a type`)
		}
		if argType == Null {
			continue
		}
		l.setSymbol(c, argType)
		return
	}
	l.setSymbol(c, Null)
}

func (l *firstPassListener) ExitIifFunc(c *parser.IifFuncContext) {
	thenType, ok := l.getSymbol(c.Expr(1))
	if !ok {
		panic(`then symbol does not have a type`)
	}
	l.setSymbol(c, thenType)
}

func (l *firstPassListener) ExitAbsFunc(c *parser.AbsFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitAcosFunc(c *parser.AcosFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitAsinFunc(c *parser.AsinFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitAtanFunc(c *parser.AtanFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitAtan2Func(c *parser.Atan2FuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitAverageFunc(c *parser.AverageFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitCeilFunc(c *parser.CeilFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitCosFunc(c *parser.CosFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitCoshFunc(c *parser.CoshFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitDistanceFunc(c *parser.DistanceFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitExpFunc(c *parser.ExpFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitFloorFunc(c *parser.FloorFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitLogFunc(c *parser.LogFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitLog10Func(c *parser.Log10FuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitMedianFunc(c *parser.MedianFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitModFunc(c *parser.ModFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitPiFunc(c *parser.PiFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitRandFunc(c *parser.RandFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitRandIntFunc(c *parser.RandIntFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitRoundFunc(c *parser.RoundFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitSinFunc(c *parser.SinFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitSinhFunc(c *parser.SinhFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitSqrtFunc(c *parser.SqrtFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitTanFunc(c *parser.TanFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitTanhFunc(c *parser.TanhFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitSwitchFunc(c *parser.SwitchFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitCharFromIntFunc(c *parser.CharFromIntFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitCharToIntFunc(c *parser.CharToIntFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitHexToNumberFunc(c *parser.HexToNumberFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitContainsFunc(c *parser.ContainsFuncContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitCountWordsFunc(c *parser.CountWordsFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitEndsWithFunc(c *parser.EndsWithFuncContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitFindStringFunc(c *parser.FindStringFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitGetWordFunc(c *parser.GetWordFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitLeftFunc(c *parser.LeftFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitLengthFunc(c *parser.LengthFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitLowercaseFunc(c *parser.LowercaseFuncContext) {
	l.setSymbol(c, String)
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
