package alteryx_formulas

import (
	"fmt"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
)

type secondPassListener struct {
	calc    *calculator
	symbols map[antlr.ParserRuleContext]int
	parser.BaseAlteryxFormulasListener
}

func (l *secondPassListener) getSymbol(c antlr.ParserRuleContext) (int, bool) {
	symbol, ok := l.symbols[c]
	return symbol, ok
}

type HasLeftRightTypes interface {
	GetLeft() parser.IExprContext
	GetRight() parser.IExprContext
}

func (l *secondPassListener) getLeftRightTypes(c HasLeftRightTypes) (int, int) {
	leftType, ok := l.getSymbol(c.GetLeft())
	if !ok {
		panic(`left symbol does not exist`)
	}
	rightType, ok := l.getSymbol(c.GetRight())
	if !ok {
		panic(`right symbol does not exist`)
	}
	return leftType, rightType
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

func (l *secondPassListener) EnterNullFunc(_ *parser.NullFuncContext) {
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
	leftType, rightType := l.getLeftRightTypes(c)

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
	c.GetParser().NotifyErrorListeners(`left and right arguments are not the same type`, c.GetStart(), InvalidType(``, c))
}

func (l *secondPassListener) EnterSubtract(c *parser.SubtractContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		panic(`invalid left type`)
	}
	if rightType != Number && rightType != Null {
		panic(`invalid right type`)
	}

	l.calc.pushFunction(l.calc.subtractNumbers)
}

func (l *secondPassListener) EnterMultiply(c *parser.MultiplyContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		panic(`invalid left type`)
	}
	if rightType != Number && rightType != Null {
		panic(`invalid right type`)
	}

	l.calc.pushFunction(l.calc.multiplyNumbers)
}

func (l *secondPassListener) EnterDivide(c *parser.DivideContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		panic(`invalid left type`)
	}
	if rightType != Number && rightType != Null {
		panic(`invalid right type`)
	}

	l.calc.pushFunction(l.calc.divideNumbers)
}

func (l *secondPassListener) EnterEqual(_ *parser.EqualContext) {
	l.calc.pushFunction(l.calc.equal)
}

func (l *secondPassListener) EnterNotEqual(_ *parser.NotEqualContext) {
	l.calc.pushFunction(l.calc.notEqual)
}

func (l *secondPassListener) EnterGreaterThan(c *parser.GreaterThanContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if (leftType == Number || leftType == Null) && (rightType == Number || rightType == Null) {
		l.calc.pushFunction(l.calc.numberGreaterThan)
		return
	}
	panic(`invalid left or right type`)
}

func (l *secondPassListener) EnterGreaterEqual(c *parser.GreaterEqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if (leftType == Number || leftType == Null) && (rightType == Number || rightType == Null) {
		l.calc.pushFunction(l.calc.numberGreaterEqual)
		return
	}
	panic(`invalid left or right type`)
}

func (l *secondPassListener) EnterLessThan(c *parser.LessThanContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if (leftType == Number || leftType == Null) && (rightType == Number || rightType == Null) {
		l.calc.pushFunction(l.calc.numberLessThan)
		return
	}
	panic(`invalid left or right type`)
}

func (l *secondPassListener) EnterLessEqual(c *parser.LessEqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if (leftType == Number || leftType == Null) && (rightType == Number || rightType == Null) {
		l.calc.pushFunction(l.calc.numberLessEqual)
		return
	}
	panic(`invalid left or right type`)
}

func (l *secondPassListener) EnterIn(c *parser.InContext) {
	inSymbol := l.checkTypeOfMultiExprs(c)

	if inSymbol == Number || inSymbol == Null {
		l.calc.pushFunction(l.calc.numberIn)
		exprs := len(c.AllExpr())
		l.calc.pushValueFunc(exprs)
		return
	}
	panic(`invalid type`)
}

func (l *secondPassListener) EnterNotIn(c *parser.NotInContext) {
	inSymbol := l.checkTypeOfMultiExprs(c)

	if inSymbol == Number || inSymbol == Null {
		l.calc.pushFunction(l.calc.numberNotIn)
		exprs := len(c.AllExpr())
		l.calc.pushValueFunc(exprs)
		return
	}
	panic(`invalid type`)
}

type HasAllExpr interface {
	AllExpr() []parser.IExprContext
}

func (l *secondPassListener) checkTypeOfMultiExprs(c HasAllExpr) int {
	exprs := c.AllExpr()
	exprCount := len(exprs)
	inSymbol := Null
	for i := 0; i < exprCount; i++ {
		argType, ok := l.getSymbol(exprs[i])
		if !ok {
			panic(`symbol not found for in argument`)
		}
		if argType == Null {
			continue
		}
		if inSymbol != Null && argType != inSymbol {
			panic(fmt.Sprintf(`invalid type at parameter %v`, i))
		}
		inSymbol = argType
	}
	return inSymbol
}

func (l *secondPassListener) EnterExprIf(c *parser.ExprIfContext) {
	exprCount := len(c.AllExpr())
	l.calc.pushFunction(l.calc.exprIf)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterPowFunc(c *parser.PowFuncContext) {
	valueType, ok := l.getSymbol(c.Expr(0))
	if !ok {
		panic(`value symbol does not have a type`)
	}
	powerType, ok := l.getSymbol(c.Expr(1))
	if !ok {
		panic(`power symbol does not have a type`)
	}
	if (valueType == Number || valueType == Null) && (powerType == Number || powerType == Null) {
		l.calc.pushFunction(l.calc.pow)
		return
	}
	panic(`value or power are not numbers`)
}

func (l *secondPassListener) EnterMinFunc(c *parser.MinFuncContext) {
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`min does not have a type`)
	}

	exprCount := len(c.AllExpr())
	switch dataType {
	case Number:
		l.calc.pushFunction(l.calc.numberMin)
	default:
		panic(`invalid type`)
	}
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterMaxFunc(c *parser.MaxFuncContext) {
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`max does not have a type`)
	}

	exprCount := len(c.AllExpr())
	switch dataType {
	case Number:
		l.calc.pushFunction(l.calc.numberMax)
	default:
		panic(`invalid type`)
	}
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterIifFunc(c *parser.IifFuncContext) {
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`iif does not have a type`)
	}

	switch dataType {
	case Number:
		l.calc.pushFunction(l.calc.numberIif)
	default:
		panic(`invalid type`)
	}
}

func (l *secondPassListener) EnterAbsFunc(_ *parser.AbsFuncContext) {
	l.calc.pushFunction(l.calc.abs)
}

func (l *secondPassListener) EnterAcosFunc(_ *parser.AcosFuncContext) {
	l.calc.pushFunction(l.calc.acos)
}

func (l *secondPassListener) EnterAsinFunc(_ *parser.AsinFuncContext) {
	l.calc.pushFunction(l.calc.asin)
}

func (l *secondPassListener) EnterAtanFunc(_ *parser.AtanFuncContext) {
	l.calc.pushFunction(l.calc.atan)
}

func (l *secondPassListener) EnterAtan2Func(_ *parser.Atan2FuncContext) {
	l.calc.pushFunction(l.calc.atan2)
}

func (l *secondPassListener) EnterAverageFunc(c *parser.AverageFuncContext) {
	exprCount := len(c.AllExpr())
	l.calc.pushFunction(l.calc.average)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterCeilFunc(_ *parser.CeilFuncContext) {
	l.calc.pushFunction(l.calc.ceil)
}

func (l *secondPassListener) EnterCosFunc(_ *parser.CosFuncContext) {
	l.calc.pushFunction(l.calc.cos)
}

func (l *secondPassListener) EnterCoshFunc(_ *parser.CoshFuncContext) {
	l.calc.pushFunction(l.calc.cosh)
}

func (l *secondPassListener) EnterDistanceFunc(_ *parser.DistanceFuncContext) {
	l.calc.geo = ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Mile, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	l.calc.pushFunction(l.calc.distance)
}

func (l *secondPassListener) EnterExpFunc(_ *parser.ExpFuncContext) {
	l.calc.pushFunction(l.calc.exp)
}

func (l *secondPassListener) EnterFloorFunc(_ *parser.FloorFuncContext) {
	l.calc.pushFunction(l.calc.floor)
}

func (l *secondPassListener) EnterLogFunc(_ *parser.LogFuncContext) {
	l.calc.pushFunction(l.calc.log)
}

func (l *secondPassListener) EnterLog10Func(_ *parser.Log10FuncContext) {
	l.calc.pushFunction(l.calc.log10)
}

func (l *secondPassListener) EnterMedianFunc(c *parser.MedianFuncContext) {
	exprCount := len(c.AllExpr())
	l.calc.pushFunction(l.calc.median)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterModFunc(_ *parser.ModFuncContext) {
	l.calc.pushFunction(l.calc.mod)
}

func (l *secondPassListener) EnterPiFunc(_ *parser.PiFuncContext) {
	l.calc.pushFunction(l.calc.pi)
}

func (l *secondPassListener) EnterRandFunc(_ *parser.RandFuncContext) {
	l.calc.pushFunction(l.calc.rand)
}

func (l *secondPassListener) EnterRandIntFunc(_ *parser.RandIntFuncContext) {
	l.calc.pushFunction(l.calc.randInt)
}

func (l *secondPassListener) EnterRoundFunc(_ *parser.RoundFuncContext) {
	l.calc.pushFunction(l.calc.round)
}

func (l *secondPassListener) EnterSinFunc(_ *parser.SinFuncContext) {
	l.calc.pushFunction(l.calc.sin)
}

func (l *secondPassListener) EnterSinhFunc(_ *parser.SinhFuncContext) {
	l.calc.pushFunction(l.calc.sinh)
}

func (l *secondPassListener) EnterSqrtFunc(_ *parser.SqrtFuncContext) {
	l.calc.pushFunction(l.calc.sqrt)
}

func (l *secondPassListener) EnterTanFunc(_ *parser.TanFuncContext) {
	l.calc.pushFunction(l.calc.tan)
}

func (l *secondPassListener) EnterTanhFunc(_ *parser.TanhFuncContext) {
	l.calc.pushFunction(l.calc.tanh)
}

func (l *secondPassListener) EnterSwitchFunc(c *parser.SwitchFuncContext) {
	_, ok := l.getSymbol(c)
	if !ok {
		panic(`symbol does not exist for switch`)
	}
	l.calc.pushFunction(l.calc.switchFunc)
	exprCount := len(c.AllExpr())
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterCharFromIntFunc(_ *parser.CharFromIntFuncContext) {
	l.calc.pushFunction(l.calc.charFromInt)
}
