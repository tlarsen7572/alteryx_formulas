package alteryx_formulas

import (
	"fmt"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
	"strings"
	"time"
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

func boolChecker(c ErrorNotifier, errorMsg string) symbolChecker {
	return symbolChecker{
		c:            c,
		expectedType: Bool,
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

func (l *secondPassListener) checkBool(c ErrorNotifier, errorMsg string) {
	l.checkSymbol(boolChecker(c, errorMsg))
}

func (l *secondPassListener) checkDynamic(expectedType int, c ErrorNotifier, errorMsg string) {
	l.checkSymbol(symbolChecker{
		c:            c,
		expectedType: expectedType,
		allowNull:    true,
		errorMsg:     errorMsg,
	})
}

func (l *secondPassListener) checkSymbol(checker symbolChecker) {
	symbol, ok := l.getSymbol(checker.c)

	if !ok {
		panic(`could not find a symbol`)
	}
	if symbol != checker.expectedType && !(checker.allowNull && symbol == Null) {
		NotifyError(checker.c, checker.errorMsg)
	}

}

type HasLeftRightTypes interface {
	GetLeft() parser.IExprContext
	GetRight() parser.IExprContext
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
	NotifyError(c, `left and right arguments of add operation are not the same type`)
}

func (l *secondPassListener) EnterSubtract(c *parser.SubtractContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		NotifyError(c, `left argument of subtract operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		NotifyError(c, `right argument of subtract operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.subtractNumbers)
}

func (l *secondPassListener) EnterMultiply(c *parser.MultiplyContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		NotifyError(c, `left argument of multiply operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		NotifyError(c, `right argument of multiply operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.multiplyNumbers)
}

func (l *secondPassListener) EnterDivide(c *parser.DivideContext) {
	leftType, rightType := l.getLeftRightTypes(c)

	if leftType != Number && leftType != Null {
		NotifyError(c, `left argument of divide operation is not a number`)
		return
	}
	if rightType != Number && rightType != Null {
		NotifyError(c, `right argument of divide operation is not a number`)
		return
	}

	l.calc.pushFunction(l.calc.divideNumbers)
}

func (l *secondPassListener) EnterEqual(c *parser.EqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)
	if leftType != rightType && leftType != Null && rightType != Null {
		NotifyError(c, `left and right arguments of equal operation are not the same`)
		return
	}
	l.calc.pushFunction(l.calc.equal)
}

func (l *secondPassListener) EnterNotEqual(c *parser.NotEqualContext) {
	leftType, rightType := l.getLeftRightTypes(c)
	if leftType != rightType && leftType != Null && rightType != Null {
		NotifyError(c, `left and right arguments of equal operation are not the same`)
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
	NotifyError(c, `left and right arguments of greater than operation are not the same`)
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
	NotifyError(c, `left and right arguments of greater than or equal operation are not the same`)
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
	NotifyError(c, `left and right arguments of less than operation are not the same`)
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
	NotifyError(c, `left and right arguments of less than or equal operation are not the same`)
}

func (l *secondPassListener) EnterAnd(c *parser.AndContext) {
	leftSymbol, rightSymbol := l.getLeftRightTypes(c)
	if leftSymbol != Bool && leftSymbol != Null {
		NotifyError(c, `left argument of AND is not a boolean`)
	}
	if rightSymbol != Bool && rightSymbol != Null {
		NotifyError(c, `right argument of AND is not a boolean`)
	}
	l.calc.pushFunction(l.calc.and)
}

func (l *secondPassListener) EnterOr(c *parser.OrContext) {
	leftSymbol, rightSymbol := l.getLeftRightTypes(c)
	if leftSymbol != Bool && leftSymbol != Null {
		NotifyError(c, `left argument of OR is not a boolean`)
	}
	if rightSymbol != Bool && rightSymbol != Null {
		NotifyError(c, `right argument of OR is not a boolean`)
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
				NotifyError(c, `if/elseif statement is not a boolean`)
			}
		} else {
			symbol, ok := l.getSymbol(exprs[i])
			if !ok {
				panic(`then/else expr missing a symbol`)
			}
			if symbol != cSymbol && symbol != Null && cSymbol != Null {
				NotifyError(c, `THEN and ELSE statements are not the same type`)
			}
		}
	}

	l.calc.pushFunction(l.calc.exprIf)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterPowFunc(c *parser.PowFuncContext) {
	l.checkNumber(c.Expr(0), `pow value parameter is not a number`)
	l.checkNumber(c.Expr(1), `pow power parameter is not a number`)
	l.calc.pushFunction(l.calc.pow)
}

func (l *secondPassListener) EnterMinFunc(c *parser.MinFuncContext) {
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`min does not have a type`)
	}

	exprCount := len(c.AllExpr())
	for i := 0; i < exprCount; i++ {
		l.checkDynamic(dataType, c.Expr(i), fmt.Sprintf(`min parameter #%v is not the expected data type`, i+1))
	}
	switch dataType {
	case Null:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(min interface{}, nextValue interface{}) bool {
			return false
		}))
	case Number:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(min interface{}, nextValue interface{}) bool {
			return nextValue.(float64) < min.(float64)
		}))
	case String:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(min interface{}, nextValue interface{}) bool {
			return nextValue.(string) < min.(string)
		}))
	case Date:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(min interface{}, nextValue interface{}) bool {
			return nextValue.(time.Time).Before(min.(time.Time))
		}))
	default:
		NotifyError(c, `min function returns an invalid data type`)
	}
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterMaxFunc(c *parser.MaxFuncContext) {
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`max does not have a type`)
	}

	exprCount := len(c.AllExpr())
	for i := 0; i < exprCount; i++ {
		l.checkDynamic(dataType, c.Expr(i), fmt.Sprintf(`max parameter #%v is not the expected data type`, i+1))
	}
	switch dataType {
	case Null:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(max interface{}, nextValue interface{}) bool {
			return false
		}))
	case Number:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(max interface{}, nextValue interface{}) bool {
			return nextValue.(float64) > max.(float64)
		}))
	case String:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(max interface{}, nextValue interface{}) bool {
			return nextValue.(string) > max.(string)
		}))
	case Date:
		l.calc.pushFunction(l.calc.useNextIfTrue(func(max interface{}, nextValue interface{}) bool {
			return nextValue.(time.Time).After(max.(time.Time))
		}))
	default:
		NotifyError(c, `max function returns an invalid data type`)
	}
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterIifFunc(c *parser.IifFuncContext) {
	l.checkBool(c.Expr(0), `iif bool parameter is not a boolean`)
	dataType, ok := l.getSymbol(c)
	if !ok {
		panic(`iif does not have a type`)
	}

	l.checkDynamic(dataType, c.Expr(1), `iif function then value is not the expected data type`)
	l.checkDynamic(dataType, c.Expr(2), `iif function else value is not the expected data type`)
	l.calc.pushFunction(l.calc.iif)
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

func (l *secondPassListener) EnterCoshFunc(c *parser.CoshFuncContext) {
	l.checkNumber(c.Expr(), `cosh parameter is not a number`)
	l.calc.pushFunction(l.calc.cosh)
}

func (l *secondPassListener) EnterDistanceFunc(c *parser.DistanceFuncContext) {
	l.checkNumber(c.Expr(0), `distance from_lat parameter is not a number`)
	l.checkNumber(c.Expr(1), `distance from_lon parameter is not a number`)
	l.checkNumber(c.Expr(2), `distance to_lat parameter is not a number`)
	l.checkNumber(c.Expr(3), `distance to_lon parameter is not a number`)
	l.calc.geo = ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Mile, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	l.calc.pushFunction(l.calc.distance)
}

func (l *secondPassListener) EnterExpFunc(c *parser.ExpFuncContext) {
	l.checkNumber(c.Expr(), `exp parameter is not a number`)
	l.calc.pushFunction(l.calc.exp)
}

func (l *secondPassListener) EnterFloorFunc(c *parser.FloorFuncContext) {
	l.checkNumber(c.Expr(), `floor parameter is not a number`)
	l.calc.pushFunction(l.calc.floor)
}

func (l *secondPassListener) EnterLogFunc(c *parser.LogFuncContext) {
	l.checkNumber(c.Expr(), `log parameter is not a number`)
	l.calc.pushFunction(l.calc.log)
}

func (l *secondPassListener) EnterLog10Func(c *parser.Log10FuncContext) {
	l.checkNumber(c.Expr(), `log10 parameter is not a number`)
	l.calc.pushFunction(l.calc.log10)
}

func (l *secondPassListener) EnterMedianFunc(c *parser.MedianFuncContext) {
	exprCount := len(c.AllExpr())
	for i := 0; i < exprCount; i++ {
		l.checkNumber(c.Expr(i), fmt.Sprintf(`median parameter #%v is not a number`, i+1))
	}
	l.calc.pushFunction(l.calc.median)
	l.calc.pushValueFunc(exprCount)
}

func (l *secondPassListener) EnterModFunc(c *parser.ModFuncContext) {
	l.checkNumber(c.Expr(0), `mod dividend parameter is not a number`)
	l.checkNumber(c.Expr(1), `mod divisor is not a number`)
	l.calc.pushFunction(l.calc.mod)
}

func (l *secondPassListener) EnterPiFunc(_ *parser.PiFuncContext) {
	l.calc.pushFunction(l.calc.pi)
}

func (l *secondPassListener) EnterRandFunc(_ *parser.RandFuncContext) {
	l.calc.pushFunction(l.calc.rand)
}

func (l *secondPassListener) EnterRandIntFunc(c *parser.RandIntFuncContext) {
	l.checkNumber(c.Expr(), `randInt parameter is not a number`)
	l.calc.pushFunction(l.calc.randInt)
}

func (l *secondPassListener) EnterRoundFunc(c *parser.RoundFuncContext) {
	l.checkNumber(c.Expr(0), `round value parameter is not a number`)
	l.checkNumber(c.Expr(1), `round mult parameter is not a number`)
	l.calc.pushFunction(l.calc.round)
}

func (l *secondPassListener) EnterSinFunc(c *parser.SinFuncContext) {
	l.checkNumber(c.Expr(), `sin parameter is not a number`)
	l.calc.pushFunction(l.calc.sin)
}

func (l *secondPassListener) EnterSinhFunc(c *parser.SinhFuncContext) {
	l.checkNumber(c.Expr(), `sinh parameter is not a number`)
	l.calc.pushFunction(l.calc.sinh)
}

func (l *secondPassListener) EnterSqrtFunc(c *parser.SqrtFuncContext) {
	l.checkNumber(c.Expr(), `sqrt parameter is not a number`)
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

func (l *secondPassListener) EnterHexToNumberFunc(c *parser.HexToNumberFuncContext) {
	l.checkString(c.Expr(), `hexToNumber parameter is not a string`)
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

func (l *secondPassListener) EnterCountWordsFunc(c *parser.CountWordsFuncContext) {
	l.checkString(c.Expr(), `countWords parameter is not a string`)
	l.calc.pushFunction(l.calc.countWords)
}

func (l *secondPassListener) EnterEndsWithFunc(c *parser.EndsWithFuncContext) {
	l.checkString(c.Expr(0), `endsWith string parameter is not a string`)
	l.checkString(c.Expr(1), `endsWith target parameter is not a string`)
	if len(c.AllExpr()) == 3 {
		l.checkNumber(c.Expr(2), `endsWith caseInsensitive parameter is not a number`)
	}
	l.calc.pushFunction(l.calc.endsWith)
}

func (l *secondPassListener) ExitEndsWithFunc(c *parser.EndsWithFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterFindStringFunc(c *parser.FindStringFuncContext) {
	l.checkString(c.Expr(0), `findString string is not a string`)
	l.checkString(c.Expr(1), `findString target is not a string`)
	l.calc.pushFunction(l.calc.findString)
}

func (l *secondPassListener) EnterGetWordFunc(c *parser.GetWordFuncContext) {
	l.checkString(c.Expr(0), `getWord string parameter is not a string`)
	l.checkNumber(c.Expr(1), `getWord index parameter is not a number`)
	l.calc.pushFunction(l.calc.getWord)
}

func (l *secondPassListener) EnterLeftFunc(c *parser.LeftFuncContext) {
	l.checkString(c.Expr(0), `left string parameter is not a string`)
	l.checkNumber(c.Expr(1), `left len parameter is not a number`)
	l.calc.pushFunction(l.calc.left)
}

func (l *secondPassListener) EnterLengthFunc(c *parser.LengthFuncContext) {
	l.checkString(c.Expr(), `length parameter is not a string`)
	l.calc.pushFunction(l.calc.length)
}

func (l *secondPassListener) EnterLowercaseFunc(c *parser.LowercaseFuncContext) {
	l.checkString(c.Expr(), `lowercase parameter is not a string`)
	l.calc.pushFunction(l.calc.lowercase)
}

func (l *secondPassListener) EnterPadLeftFunc(c *parser.PadLeftFuncContext) {
	l.checkString(c.Expr(0), `padLeft string parameter is not a string`)
	l.checkNumber(c.Expr(1), `padLeft len parameter is not a number`)
	l.checkString(c.Expr(2), `padLeft char parameter is not a string`)
	l.calc.pushFunction(l.calc.padLeft)
}

func (l *secondPassListener) EnterPadRightFunc(c *parser.PadRightFuncContext) {
	l.checkString(c.Expr(0), `padRight string parameter is not a string`)
	l.checkNumber(c.Expr(1), `padRight len parameter is not a number`)
	l.checkString(c.Expr(2), `padRight char parameter is not a string`)
	l.calc.pushFunction(l.calc.padRight)
}

func (l *secondPassListener) EnterRegexCountMatchesFunc(c *parser.RegexCountMatchesFuncContext) {
	l.checkString(c.Expr(0), `regex_CountMatches string parameter is not a string`)
	l.checkString(c.Expr(1), `regex_CountMatches pattern parameter is not a string`)
	if len(c.AllExpr()) == 3 {
		l.checkNumber(c.Expr(2), `regex_CountMatches caseInsensitive parameter is not a number`)
	}
	l.calc.pushFunction(l.calc.regexCountMatches)
}

func (l *secondPassListener) ExitRegexCountMatchesFunc(c *parser.RegexCountMatchesFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterRegexMatchFunc(c *parser.RegexMatchFuncContext) {
	l.checkString(c.Expr(0), `regex_Match string parameter is not a string`)
	l.checkString(c.Expr(1), `regex_Match pattern parameter is not a string`)
	if len(c.AllExpr()) == 3 {
		l.checkNumber(c.Expr(2), `regex_Match caseInsensitive parameter is not a number`)
	}
	l.calc.pushFunction(l.calc.regexMatch)
}

func (l *secondPassListener) ExitRegexMatchFunc(c *parser.RegexMatchFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 2 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterRegexReplaceFunc(c *parser.RegexReplaceFuncContext) {
	l.checkString(c.Expr(0), `regex_Replace string parameter is not a string`)
	l.checkString(c.Expr(1), `regex_Replace pattern parameter is not a string`)
	l.checkString(c.Expr(2), `regex_Replace replace parameter is not a string`)
	if len(c.AllExpr()) == 4 {
		l.checkNumber(c.Expr(3), `regex_Replace caseInsensitive parameter is not a number`)
	}
	l.calc.pushFunction(l.calc.regexReplace)
}

func (l *secondPassListener) ExitRegexReplaceFunc(c *parser.RegexReplaceFuncContext) {
	exprs := len(c.AllExpr())
	if exprs == 3 {
		l.calc.pushValueFunc(1.0)
	}
}

func (l *secondPassListener) EnterReplaceFunc(c *parser.ReplaceFuncContext) {
	l.checkString(c.Expr(0), `replace string parameter is not a string`)
	l.checkString(c.Expr(1), `replace target parameter is not a string`)
	l.checkString(c.Expr(2), `replace replacement parameter is not a string`)
	l.calc.pushFunction(l.calc.replace)
}

func (l *secondPassListener) EnterRightFunc(c *parser.RightFuncContext) {
	l.checkString(c.Expr(0), `right string parameter is not a string`)
	l.checkNumber(c.Expr(1), `right len parameter is not a number`)
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
