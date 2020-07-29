package alteryx_formulas

import (
	"fmt"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
	"strings"
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

type symbolChecker struct {
	c            ErrorNotifier
	expectedType int
	allowNull    bool
	errorMsg     string
}

func numberChecker(c ErrorNotifier, errorMsg string) symbolChecker {
	return symbolChecker{
		c:            c,
		expectedType: Number,
		allowNull:    true,
		errorMsg:     errorMsg,
	}
}

func stringChecker(c ErrorNotifier, errorMsg string) symbolChecker {
	return symbolChecker{
		c:            c,
		expectedType: String,
		allowNull:    true,
		errorMsg:     errorMsg,
	}
}

func (l *secondPassListener) checkNumber(c ErrorNotifier, errorMsg string) {
	l.checkSymbol(numberChecker(c, errorMsg))
}

func (l *secondPassListener) checkString(c ErrorNotifier, errorMsg string) {
	l.checkSymbol(stringChecker(c, errorMsg))
}

func (l *secondPassListener) checkSymbol(checker symbolChecker) {
	symbol, ok := l.getSymbol(checker.c)

	if !ok {
		panic(`could not find a symbol`)
	}
	if symbol != checker.expectedType && !(checker.allowNull && symbol == Null) {
		notifyTypeError(checker.c, checker.errorMsg)
	}

}

type HasLeftRightTypes interface {
	GetLeft() parser.IExprContext
	GetRight() parser.IExprContext
}

type ErrorNotifier interface {
	GetParser() antlr.Parser
	antlr.ParserRuleContext
}

func notifyTypeError(notifier ErrorNotifier, msg string) {
	notifier.GetParser().NotifyErrorListeners(msg, notifier.GetStart(), InvalidType(``, notifier))
}

func (l *secondPassListener) leftRightIsNullOr(formulaType int, c HasLeftRightTypes) bool {
	leftType, rightType := l.getLeftRightTypes(c)
	return (leftType == formulaType || leftType == Null) && (rightType == formulaType || rightType == Null)
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

func (l *secondPassListener) EnterBoolLiteral(c *parser.BoolLiteralContext) {
	value := strings.ToLower(c.GetText())
	if value == `true` {
		l.calc.pushValueFunc(true)
	} else {
		l.calc.pushValueFunc(false)
	}
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
	notifyTypeError(c, `left and right arguments of add operation are not the same type`)
}

func (l *secondPassListener) EnterSubtract(c *parser.SubtractContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		notifyTypeError(c, `left argument of subtract operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		notifyTypeError(c, `right argument of subtract operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.subtractNumbers)
}

func (l *secondPassListener) EnterMultiply(c *parser.MultiplyContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		notifyTypeError(c, `left argument of multiply operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		notifyTypeError(c, `right argument of multiply operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.multiplyNumbers)
}

func (l *secondPassListener) EnterDivide(c *parser.DivideContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		notifyTypeError(c, `left argument of divide operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		notifyTypeError(c, `right argument of divide operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.divideNumbers)
}

func (l *secondPassListener) EnterEqual(c *parser.EqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)
	if leftType != rightType && leftType != Null && rightType != Null {
		notifyTypeError(c, `left and right arguments of equal operation are not the same`)
		return
	}
	l.calc.pushFunction(l.calc.equal)
}

func (l *secondPassListener) EnterNotEqual(c *parser.NotEqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)
	if leftType != rightType && leftType != Null && rightType != Null {
		notifyTypeError(c, `left and right arguments of equal operation are not the same`)
		return
	}
	l.calc.pushFunction(l.calc.notEqual)
}

func (l *secondPassListener) EnterGreaterThan(c *parser.GreaterThanContext) {
	if l.leftRightIsNullOr(Number, c) {
		l.calc.pushFunction(l.calc.numberGreaterThan)
		return
	}
	if l.leftRightIsNullOr(String, c) {
		l.calc.pushFunction(l.calc.stringGreaterThan)
		return
	}
	if l.leftRightIsNullOr(Date, c) {
		l.calc.pushFunction(l.calc.dateGreaterThan)
		return
	}
	notifyTypeError(c, `left and right arguments of greater than operation are not the same`)
}

func (l *secondPassListener) EnterGreaterEqual(c *parser.GreaterEqualContext) {
	if l.leftRightIsNullOr(Number, c) {
		l.calc.pushFunction(l.calc.numberGreaterEqual)
		return
	}
	if l.leftRightIsNullOr(String, c) {
		l.calc.pushFunction(l.calc.stringGreaterEqual)
		return
	}
	if l.leftRightIsNullOr(Date, c) {
		l.calc.pushFunction(l.calc.dateGreaterEqual)
		return
	}
	notifyTypeError(c, `left and right arguments of greater than or equal operation are not the same`)
}

func (l *secondPassListener) EnterLessThan(c *parser.LessThanContext) {
	if l.leftRightIsNullOr(Number, c) {
		l.calc.pushFunction(l.calc.numberLessThan)
		return
	}
	if l.leftRightIsNullOr(String, c) {
		l.calc.pushFunction(l.calc.stringLessThan)
		return
	}
	if l.leftRightIsNullOr(Date, c) {
		l.calc.pushFunction(l.calc.dateLessThan)
		return
	}
	notifyTypeError(c, `left and right arguments of less than operation are not the same`)
}

func (l *secondPassListener) EnterLessEqual(c *parser.LessEqualContext) {
	if l.leftRightIsNullOr(Number, c) {
		l.calc.pushFunction(l.calc.numberLessEqual)
		return
	}
	if l.leftRightIsNullOr(String, c) {
		l.calc.pushFunction(l.calc.stringLessEqual)
		return
	}
	if l.leftRightIsNullOr(Date, c) {
		l.calc.pushFunction(l.calc.dateLessEqual)
		return
	}
	notifyTypeError(c, `left and right arguments of less than or equal operation are not the same`)
}

func (l *secondPassListener) EnterAnd(c *parser.AndContext) {
	leftSymbol, rightSymbol := l.getLeftRightTypes(c)
	if leftSymbol != Bool && leftSymbol != Null {
		notifyTypeError(c, `left argument of AND is not a boolean`)
	}
	if rightSymbol != Bool && rightSymbol != Null {
		notifyTypeError(c, `right argument of AND is not a boolean`)
	}
	l.calc.pushFunction(l.calc.and)
}

func (l *secondPassListener) EnterOr(c *parser.OrContext) {
	leftSymbol, rightSymbol := l.getLeftRightTypes(c)
	if leftSymbol != Bool && leftSymbol != Null {
		notifyTypeError(c, `left argument of OR is not a boolean`)
	}
	if rightSymbol != Bool && rightSymbol != Null {
		notifyTypeError(c, `right argument of OR is not a boolean`)
	}
	l.calc.pushFunction(l.calc.or)
}

func (l *secondPassListener) EnterIn(c *parser.InContext) {
	l.calc.pushFunction(l.calc.in)
	exprs := len(c.AllExpr())
	l.calc.pushValueFunc(exprs)
}

func (l *secondPassListener) EnterNotIn(c *parser.NotInContext) {
	l.calc.pushFunction(l.calc.notIn)
	exprs := len(c.AllExpr())
	l.calc.pushValueFunc(exprs)
}

func (l *secondPassListener) EnterExprIf(c *parser.ExprIfContext) {
	exprs := c.AllExpr()
	exprCount := len(exprs)
	cSymbol, ok := l.getSymbol(c)
	if !ok {
		panic(`IF statement does not have a symbol`)
	}
	for i := 0; i < exprCount; i++ {
		if i < exprCount-1 && i%2 == 0 { // IF/ELSEIF, expect a boolean
			ifSymbol, ok := l.getSymbol(exprs[i])
			if !ok {
				panic(`if expr missing a symbol`)
			}
			if ifSymbol != Bool && ifSymbol != Null {
				notifyTypeError(c, `if/elseif statement is not a boolean`)
			}
		} else {
			symbol, ok := l.getSymbol(exprs[i])
			if !ok {
				panic(`then/else expr missing a symbol`)
			}
			if symbol != cSymbol && symbol != Null && cSymbol != Null {
				notifyTypeError(c, `THEN and ELSE statements are not the same type`)
			}
		}
	}

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

func (l *secondPassListener) EnterAbsFunc(c *parser.AbsFuncContext) {
	l.checkNumber(c.Expr(), `abs parameter is not a number`)
	l.calc.pushFunction(l.calc.abs)
}

func (l *secondPassListener) EnterAcosFunc(c *parser.AcosFuncContext) {
	l.checkNumber(c.Expr(), `acos parameter is not a number`)
	l.calc.pushFunction(l.calc.acos)
}

func (l *secondPassListener) EnterAsinFunc(c *parser.AsinFuncContext) {
	l.checkNumber(c.Expr(), `asin parameter is not a number`)
	l.calc.pushFunction(l.calc.asin)
}

func (l *secondPassListener) EnterAtanFunc(c *parser.AtanFuncContext) {
	l.checkNumber(c.Expr(), `atan parameter is not a number`)
	l.calc.pushFunction(l.calc.atan)
}

func (l *secondPassListener) EnterAtan2Func(c *parser.Atan2FuncContext) {
	l.checkNumber(c.Expr(0), `first atan parameter is not a number`)
	l.checkNumber(c.Expr(1), `second atan parameter is not a number`)
	l.calc.pushFunction(l.calc.atan2)
}

func (l *secondPassListener) EnterAverageFunc(c *parser.AverageFuncContext) {
	exprCount := 0
	for i, expr := range c.AllExpr() {
		l.checkNumber(expr, fmt.Sprintf(`average parameter %v is not a number`, i+1))
		exprCount++
	}
	l.calc.pushFunction(l.calc.average)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterCeilFunc(c *parser.CeilFuncContext) {
	l.checkNumber(c.Expr(), `ceil parameter is not a number`)
	l.calc.pushFunction(l.calc.ceil)
}

func (l *secondPassListener) EnterCosFunc(c *parser.CosFuncContext) {
	l.checkNumber(c.Expr(), `cos parameter is not a number`)
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

func (l *secondPassListener) EnterCharFromIntFunc(c *parser.CharFromIntFuncContext) {
	l.checkNumber(c.Expr(), `charFromInt parameter is not a number`)
	l.calc.pushFunction(l.calc.charFromInt)
}

func (l *secondPassListener) EnterCharToIntFunc(c *parser.CharToIntFuncContext) {
	l.checkString(c.Expr(), `charToInt parameter is not a string`)
	l.calc.pushFunction(l.calc.charToInt)
}

func (l *secondPassListener) EnterHexToNumberFunc(_ *parser.HexToNumberFuncContext) {
	l.calc.pushFunction(l.calc.hexToNumber)
}

func (l *secondPassListener) EnterContainsFunc(c *parser.ContainsFuncContext) {
	l.checkString(c.Expr(0), `contains string parameter is not a string`)
	l.checkString(c.Expr(1), `contains target parameter is not a string`)
	if len(c.AllExpr()) == 3 {
		l.checkNumber(c.Expr(2), `contains caseInsensitive parameter is not a number`)
	}
	l.calc.pushFunction(l.calc.contains)
}

func (l *secondPassListener) ExitContainsFunc(c *parser.ContainsFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterCountWordsFunc(_ *parser.CountWordsFuncContext) {
	l.calc.pushFunction(l.calc.countWords)
}

func (l *secondPassListener) EnterEndsWithFunc(_ *parser.EndsWithFuncContext) {
	l.calc.pushFunction(l.calc.endsWith)
}

func (l *secondPassListener) ExitEndsWithFunc(c *parser.EndsWithFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterFindStringFunc(_ *parser.FindStringFuncContext) {
	l.calc.pushFunction(l.calc.findString)
}

func (l *secondPassListener) EnterGetWordFunc(_ *parser.GetWordFuncContext) {
	l.calc.pushFunction(l.calc.getWord)
}

func (l *secondPassListener) EnterLeftFunc(_ *parser.LeftFuncContext) {
	l.calc.pushFunction(l.calc.left)
}

func (l *secondPassListener) EnterLengthFunc(_ *parser.LengthFuncContext) {
	l.calc.pushFunction(l.calc.length)
}

func (l *secondPassListener) EnterLowercaseFunc(_ *parser.LowercaseFuncContext) {
	l.calc.pushFunction(l.calc.lowercase)
}

func (l *secondPassListener) EnterPadLeftFunc(_ *parser.PadLeftFuncContext) {
	l.calc.pushFunction(l.calc.padLeft)
}

func (l *secondPassListener) EnterPadRightFunc(_ *parser.PadRightFuncContext) {
	l.calc.pushFunction(l.calc.padRight)
}

func (l *secondPassListener) EnterRegexCountMatchesFunc(_ *parser.RegexCountMatchesFuncContext) {
	l.calc.pushFunction(l.calc.regexCountMatches)
}

func (l *secondPassListener) ExitRegexCountMatchesFunc(c *parser.RegexCountMatchesFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterRegexMatchFunc(_ *parser.RegexMatchFuncContext) {
	l.calc.pushFunction(l.calc.regexMatch)
}

func (l *secondPassListener) ExitRegexMatchFunc(c *parser.RegexMatchFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterRegexReplaceFunc(_ *parser.RegexReplaceFuncContext) {
	l.calc.pushFunction(l.calc.regexReplace)
}

func (l *secondPassListener) ExitRegexReplaceFunc(c *parser.RegexReplaceFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 3 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterReplaceFunc(_ *parser.ReplaceFuncContext) {
	l.calc.pushFunction(l.calc.replace)
}

func (l *secondPassListener) EnterRightFunc(_ *parser.RightFuncContext) {
	l.calc.pushFunction(l.calc.right)
}

func (l *secondPassListener) EnterSubstringFunc(_ *parser.SubstringFuncContext) {
	l.calc.pushFunction(l.calc.substring)
}

func (l *secondPassListener) EnterTrimFunc(_ *parser.TrimFuncContext) {
	l.calc.pushFunction(l.calc.trim)
}

func (l *secondPassListener) ExitTrimFunc(c *parser.TrimFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 1 {
		l.calc.pushValueFunc(``)
	}
}

func (l *secondPassListener) EnterTrimLeftFunc(_ *parser.TrimLeftFuncContext) {
	l.calc.pushFunction(l.calc.trimLeft)
}

func (l *secondPassListener) ExitTrimLeftFunc(c *parser.TrimLeftFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 1 {
		l.calc.pushValueFunc(``)
	}
}

func (l *secondPassListener) EnterTrimRightFunc(_ *parser.TrimRightFuncContext) {
	l.calc.pushFunction(l.calc.trimRight)
}

func (l *secondPassListener) ExitTrimRightFunc(c *parser.TrimRightFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 1 {
		l.calc.pushValueFunc(``)
	}
}

func (l *secondPassListener) EnterUppercaseFunc(_ *parser.UppercaseFuncContext) {
	l.calc.pushFunction(l.calc.uppercase)
}

func (l *secondPassListener) EnterIsNullFunc(_ *parser.IsNullFuncContext) {
	l.calc.pushFunction(l.calc.isNull)
}

func (l *secondPassListener) EnterIsEmptyFunc(_ *parser.IsEmptyFuncContext) {
	l.calc.pushFunction(l.calc.isEmpty)
}

func (l *secondPassListener) EnterToDateFunc(_ *parser.ToDateFuncContext) {
	l.calc.pushFunction(l.calc.toDate)
}

func (l *secondPassListener) EnterToDatetimeFunc(_ *parser.ToDatetimeFuncContext) {
	l.calc.pushFunction(l.calc.toDatetime)
}
