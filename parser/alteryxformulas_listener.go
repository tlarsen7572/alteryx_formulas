// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type AlteryxFormulasListener interface {
	antlr.ParseTreeListener

	// EnterLog10Func is called when entering the log10Func production.
	EnterLog10Func(c *Log10FuncContext)

	// EnterMaxFunc is called when entering the maxFunc production.
	EnterMaxFunc(c *MaxFuncContext)

	// EnterCeilFunc is called when entering the ceilFunc production.
	EnterCeilFunc(c *CeilFuncContext)

	// EnterCosFunc is called when entering the cosFunc production.
	EnterCosFunc(c *CosFuncContext)

	// EnterDistanceFunc is called when entering the distanceFunc production.
	EnterDistanceFunc(c *DistanceFuncContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterAcosFunc is called when entering the acosFunc production.
	EnterAcosFunc(c *AcosFuncContext)

	// EnterMinFunc is called when entering the minFunc production.
	EnterMinFunc(c *MinFuncContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterFloorFunc is called when entering the floorFunc production.
	EnterFloorFunc(c *FloorFuncContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterNotIn is called when entering the notIn production.
	EnterNotIn(c *NotInContext)

	// EnterExprIf is called when entering the exprIf production.
	EnterExprIf(c *ExprIfContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterAtanFunc is called when entering the atanFunc production.
	EnterAtanFunc(c *AtanFuncContext)

	// EnterExprField is called when entering the exprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterAsinFunc is called when entering the asinFunc production.
	EnterAsinFunc(c *AsinFuncContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterCoshFunc is called when entering the coshFunc production.
	EnterCoshFunc(c *CoshFuncContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterExpFunc is called when entering the expFunc production.
	EnterExpFunc(c *ExpFuncContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterNotEqual is called when entering the notEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterPowFunc is called when entering the powFunc production.
	EnterPowFunc(c *PowFuncContext)

	// EnterIifFunc is called when entering the iifFunc production.
	EnterIifFunc(c *IifFuncContext)

	// EnterMedianFunc is called when entering the medianFunc production.
	EnterMedianFunc(c *MedianFuncContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterNullFunc is called when entering the nullFunc production.
	EnterNullFunc(c *NullFuncContext)

	// EnterDatetimeLiteral is called when entering the datetimeLiteral production.
	EnterDatetimeLiteral(c *DatetimeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterAverageFunc is called when entering the averageFunc production.
	EnterAverageFunc(c *AverageFuncContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterGreaterEqual is called when entering the greaterEqual production.
	EnterGreaterEqual(c *GreaterEqualContext)

	// EnterAbsFunc is called when entering the absFunc production.
	EnterAbsFunc(c *AbsFuncContext)

	// EnterLessEqual is called when entering the lessEqual production.
	EnterLessEqual(c *LessEqualContext)

	// EnterAtan2Func is called when entering the atan2Func production.
	EnterAtan2Func(c *Atan2FuncContext)

	// EnterLogFunc is called when entering the logFunc production.
	EnterLogFunc(c *LogFuncContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// ExitLog10Func is called when exiting the log10Func production.
	ExitLog10Func(c *Log10FuncContext)

	// ExitMaxFunc is called when exiting the maxFunc production.
	ExitMaxFunc(c *MaxFuncContext)

	// ExitCeilFunc is called when exiting the ceilFunc production.
	ExitCeilFunc(c *CeilFuncContext)

	// ExitCosFunc is called when exiting the cosFunc production.
	ExitCosFunc(c *CosFuncContext)

	// ExitDistanceFunc is called when exiting the distanceFunc production.
	ExitDistanceFunc(c *DistanceFuncContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitAcosFunc is called when exiting the acosFunc production.
	ExitAcosFunc(c *AcosFuncContext)

	// ExitMinFunc is called when exiting the minFunc production.
	ExitMinFunc(c *MinFuncContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitFloorFunc is called when exiting the floorFunc production.
	ExitFloorFunc(c *FloorFuncContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitNotIn is called when exiting the notIn production.
	ExitNotIn(c *NotInContext)

	// ExitExprIf is called when exiting the exprIf production.
	ExitExprIf(c *ExprIfContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitAtanFunc is called when exiting the atanFunc production.
	ExitAtanFunc(c *AtanFuncContext)

	// ExitExprField is called when exiting the exprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitAsinFunc is called when exiting the asinFunc production.
	ExitAsinFunc(c *AsinFuncContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitCoshFunc is called when exiting the coshFunc production.
	ExitCoshFunc(c *CoshFuncContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitExpFunc is called when exiting the expFunc production.
	ExitExpFunc(c *ExpFuncContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitNotEqual is called when exiting the notEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitPowFunc is called when exiting the powFunc production.
	ExitPowFunc(c *PowFuncContext)

	// ExitIifFunc is called when exiting the iifFunc production.
	ExitIifFunc(c *IifFuncContext)

	// ExitMedianFunc is called when exiting the medianFunc production.
	ExitMedianFunc(c *MedianFuncContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitNullFunc is called when exiting the nullFunc production.
	ExitNullFunc(c *NullFuncContext)

	// ExitDatetimeLiteral is called when exiting the datetimeLiteral production.
	ExitDatetimeLiteral(c *DatetimeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitAverageFunc is called when exiting the averageFunc production.
	ExitAverageFunc(c *AverageFuncContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitGreaterEqual is called when exiting the greaterEqual production.
	ExitGreaterEqual(c *GreaterEqualContext)

	// ExitAbsFunc is called when exiting the absFunc production.
	ExitAbsFunc(c *AbsFuncContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitAtan2Func is called when exiting the atan2Func production.
	ExitAtan2Func(c *Atan2FuncContext)

	// ExitLogFunc is called when exiting the logFunc production.
	ExitLogFunc(c *LogFuncContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)
}
