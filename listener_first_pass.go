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

func (l *firstPassListener) ExitAnd(c *parser.AndContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitOr(c *parser.OrContext) {
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
		NotifyError(c, fmt.Sprintf(`field %v does not exist`, fieldName))
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
		NotifyError(c, fmt.Sprintf(`field %v has invalid type %v`, fieldName, fieldType))
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
	exprs := c.AllExpr()
	for i := 0; i < len(exprs)/2; i++ {
		symbol, ok := l.getSymbol(exprs[i*2+1])
		if !ok {
			panic(`no symbol for switch parameter`)
		}
		if symbol == Null {
			continue
		}
		l.setSymbol(c, symbol)
		return
	}
	l.setSymbol(c, Null)
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

func (l *firstPassListener) ExitPadLeftFunc(c *parser.PadLeftFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitPadRightFunc(c *parser.PadRightFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitRegexCountMatchesFunc(c *parser.RegexCountMatchesFuncContext) {
	l.setSymbol(c, Number)
}

func (l *firstPassListener) ExitRegexMatchFunc(c *parser.RegexMatchFuncContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitRegexReplaceFunc(c *parser.RegexReplaceFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitReplaceFunc(c *parser.ReplaceFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitRightFunc(c *parser.RightFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitSubstringFunc(c *parser.SubstringFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitTrimFunc(c *parser.TrimFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitTrimRightFunc(c *parser.TrimRightFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitTrimLeftFunc(c *parser.TrimLeftFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitUppercaseFunc(c *parser.UppercaseFuncContext) {
	l.setSymbol(c, String)
}

func (l *firstPassListener) ExitIsNullFunc(c *parser.IsNullFuncContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitIsEmptyFunc(c *parser.IsEmptyFuncContext) {
	l.setSymbol(c, Bool)
}

func (l *firstPassListener) ExitParenthesis(c *parser.ParenthesisContext) {
	childSymbol, ok := l.getSymbol(c.Expr())
	if !ok {
		panic(`parenthesis expr did not have a symbol`)
	}
	l.setSymbol(c, childSymbol)
}

func (l *firstPassListener) ExitToDateFunc(c *parser.ToDateFuncContext) {
	l.setSymbol(c, Date)
}

func (l *firstPassListener) ExitToDatetimeFunc(c *parser.ToDatetimeFuncContext) {
	l.setSymbol(c, Date)
}
