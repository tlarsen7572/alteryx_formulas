// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type AlteryxFormulasListener interface {
	antlr.ParseTreeListener

	// EnterMaxFunc is called when entering the maxFunc production.
	EnterMaxFunc(c *MaxFuncContext)

	// EnterCosFunc is called when entering the cosFunc production.
	EnterCosFunc(c *CosFuncContext)

	// EnterRandIntFunc is called when entering the randIntFunc production.
	EnterRandIntFunc(c *RandIntFuncContext)

	// EnterCharToIntFunc is called when entering the charToIntFunc production.
	EnterCharToIntFunc(c *CharToIntFuncContext)

	// EnterTanFunc is called when entering the tanFunc production.
	EnterTanFunc(c *TanFuncContext)

	// EnterFloorFunc is called when entering the floorFunc production.
	EnterFloorFunc(c *FloorFuncContext)

	// EnterNotIn is called when entering the notIn production.
	EnterNotIn(c *NotInContext)

	// EnterTanhFunc is called when entering the tanhFunc production.
	EnterTanhFunc(c *TanhFuncContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterAsinFunc is called when entering the asinFunc production.
	EnterAsinFunc(c *AsinFuncContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterCoshFunc is called when entering the coshFunc production.
	EnterCoshFunc(c *CoshFuncContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterNotEqual is called when entering the notEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterHexToNumberFunc is called when entering the hexToNumberFunc production.
	EnterHexToNumberFunc(c *HexToNumberFuncContext)

	// EnterMedianFunc is called when entering the medianFunc production.
	EnterMedianFunc(c *MedianFuncContext)

	// EnterNullFunc is called when entering the nullFunc production.
	EnterNullFunc(c *NullFuncContext)

	// EnterDatetimeLiteral is called when entering the datetimeLiteral production.
	EnterDatetimeLiteral(c *DatetimeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterAbsFunc is called when entering the absFunc production.
	EnterAbsFunc(c *AbsFuncContext)

	// EnterAtan2Func is called when entering the atan2Func production.
	EnterAtan2Func(c *Atan2FuncContext)

	// EnterFindStringFunc is called when entering the findStringFunc production.
	EnterFindStringFunc(c *FindStringFuncContext)

	// EnterLog10Func is called when entering the log10Func production.
	EnterLog10Func(c *Log10FuncContext)

	// EnterCeilFunc is called when entering the ceilFunc production.
	EnterCeilFunc(c *CeilFuncContext)

	// EnterEndsWithFunc is called when entering the endsWithFunc production.
	EnterEndsWithFunc(c *EndsWithFuncContext)

	// EnterDistanceFunc is called when entering the distanceFunc production.
	EnterDistanceFunc(c *DistanceFuncContext)

	// EnterCharFromIntFunc is called when entering the charFromIntFunc production.
	EnterCharFromIntFunc(c *CharFromIntFuncContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterAcosFunc is called when entering the acosFunc production.
	EnterAcosFunc(c *AcosFuncContext)

	// EnterSqrtFunc is called when entering the sqrtFunc production.
	EnterSqrtFunc(c *SqrtFuncContext)

	// EnterSinhFunc is called when entering the sinhFunc production.
	EnterSinhFunc(c *SinhFuncContext)

	// EnterRandFunc is called when entering the randFunc production.
	EnterRandFunc(c *RandFuncContext)

	// EnterMinFunc is called when entering the minFunc production.
	EnterMinFunc(c *MinFuncContext)

	// EnterPiFunc is called when entering the piFunc production.
	EnterPiFunc(c *PiFuncContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterExprIf is called when entering the exprIf production.
	EnterExprIf(c *ExprIfContext)

	// EnterContainsFunc is called when entering the containsFunc production.
	EnterContainsFunc(c *ContainsFuncContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterAtanFunc is called when entering the atanFunc production.
	EnterAtanFunc(c *AtanFuncContext)

	// EnterRoundFunc is called when entering the roundFunc production.
	EnterRoundFunc(c *RoundFuncContext)

	// EnterSwitchFunc is called when entering the switchFunc production.
	EnterSwitchFunc(c *SwitchFuncContext)

	// EnterExprField is called when entering the exprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterExpFunc is called when entering the expFunc production.
	EnterExpFunc(c *ExpFuncContext)

	// EnterCountWordsFunc is called when entering the countWordsFunc production.
	EnterCountWordsFunc(c *CountWordsFuncContext)

	// EnterPowFunc is called when entering the powFunc production.
	EnterPowFunc(c *PowFuncContext)

	// EnterIifFunc is called when entering the iifFunc production.
	EnterIifFunc(c *IifFuncContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterModFunc is called when entering the modFunc production.
	EnterModFunc(c *ModFuncContext)

	// EnterSinFunc is called when entering the sinFunc production.
	EnterSinFunc(c *SinFuncContext)

	// EnterAverageFunc is called when entering the averageFunc production.
	EnterAverageFunc(c *AverageFuncContext)

	// EnterGreaterEqual is called when entering the greaterEqual production.
	EnterGreaterEqual(c *GreaterEqualContext)

	// EnterLessEqual is called when entering the lessEqual production.
	EnterLessEqual(c *LessEqualContext)

	// EnterLogFunc is called when entering the logFunc production.
	EnterLogFunc(c *LogFuncContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// ExitMaxFunc is called when exiting the maxFunc production.
	ExitMaxFunc(c *MaxFuncContext)

	// ExitCosFunc is called when exiting the cosFunc production.
	ExitCosFunc(c *CosFuncContext)

	// ExitRandIntFunc is called when exiting the randIntFunc production.
	ExitRandIntFunc(c *RandIntFuncContext)

	// ExitCharToIntFunc is called when exiting the charToIntFunc production.
	ExitCharToIntFunc(c *CharToIntFuncContext)

	// ExitTanFunc is called when exiting the tanFunc production.
	ExitTanFunc(c *TanFuncContext)

	// ExitFloorFunc is called when exiting the floorFunc production.
	ExitFloorFunc(c *FloorFuncContext)

	// ExitNotIn is called when exiting the notIn production.
	ExitNotIn(c *NotInContext)

	// ExitTanhFunc is called when exiting the tanhFunc production.
	ExitTanhFunc(c *TanhFuncContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitAsinFunc is called when exiting the asinFunc production.
	ExitAsinFunc(c *AsinFuncContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitCoshFunc is called when exiting the coshFunc production.
	ExitCoshFunc(c *CoshFuncContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitNotEqual is called when exiting the notEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitHexToNumberFunc is called when exiting the hexToNumberFunc production.
	ExitHexToNumberFunc(c *HexToNumberFuncContext)

	// ExitMedianFunc is called when exiting the medianFunc production.
	ExitMedianFunc(c *MedianFuncContext)

	// ExitNullFunc is called when exiting the nullFunc production.
	ExitNullFunc(c *NullFuncContext)

	// ExitDatetimeLiteral is called when exiting the datetimeLiteral production.
	ExitDatetimeLiteral(c *DatetimeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitAbsFunc is called when exiting the absFunc production.
	ExitAbsFunc(c *AbsFuncContext)

	// ExitAtan2Func is called when exiting the atan2Func production.
	ExitAtan2Func(c *Atan2FuncContext)

	// ExitFindStringFunc is called when exiting the findStringFunc production.
	ExitFindStringFunc(c *FindStringFuncContext)

	// ExitLog10Func is called when exiting the log10Func production.
	ExitLog10Func(c *Log10FuncContext)

	// ExitCeilFunc is called when exiting the ceilFunc production.
	ExitCeilFunc(c *CeilFuncContext)

	// ExitEndsWithFunc is called when exiting the endsWithFunc production.
	ExitEndsWithFunc(c *EndsWithFuncContext)

	// ExitDistanceFunc is called when exiting the distanceFunc production.
	ExitDistanceFunc(c *DistanceFuncContext)

	// ExitCharFromIntFunc is called when exiting the charFromIntFunc production.
	ExitCharFromIntFunc(c *CharFromIntFuncContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitAcosFunc is called when exiting the acosFunc production.
	ExitAcosFunc(c *AcosFuncContext)

	// ExitSqrtFunc is called when exiting the sqrtFunc production.
	ExitSqrtFunc(c *SqrtFuncContext)

	// ExitSinhFunc is called when exiting the sinhFunc production.
	ExitSinhFunc(c *SinhFuncContext)

	// ExitRandFunc is called when exiting the randFunc production.
	ExitRandFunc(c *RandFuncContext)

	// ExitMinFunc is called when exiting the minFunc production.
	ExitMinFunc(c *MinFuncContext)

	// ExitPiFunc is called when exiting the piFunc production.
	ExitPiFunc(c *PiFuncContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitExprIf is called when exiting the exprIf production.
	ExitExprIf(c *ExprIfContext)

	// ExitContainsFunc is called when exiting the containsFunc production.
	ExitContainsFunc(c *ContainsFuncContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitAtanFunc is called when exiting the atanFunc production.
	ExitAtanFunc(c *AtanFuncContext)

	// ExitRoundFunc is called when exiting the roundFunc production.
	ExitRoundFunc(c *RoundFuncContext)

	// ExitSwitchFunc is called when exiting the switchFunc production.
	ExitSwitchFunc(c *SwitchFuncContext)

	// ExitExprField is called when exiting the exprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitExpFunc is called when exiting the expFunc production.
	ExitExpFunc(c *ExpFuncContext)

	// ExitCountWordsFunc is called when exiting the countWordsFunc production.
	ExitCountWordsFunc(c *CountWordsFuncContext)

	// ExitPowFunc is called when exiting the powFunc production.
	ExitPowFunc(c *PowFuncContext)

	// ExitIifFunc is called when exiting the iifFunc production.
	ExitIifFunc(c *IifFuncContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitModFunc is called when exiting the modFunc production.
	ExitModFunc(c *ModFuncContext)

	// ExitSinFunc is called when exiting the sinFunc production.
	ExitSinFunc(c *SinFuncContext)

	// ExitAverageFunc is called when exiting the averageFunc production.
	ExitAverageFunc(c *AverageFuncContext)

	// ExitGreaterEqual is called when exiting the greaterEqual production.
	ExitGreaterEqual(c *GreaterEqualContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitLogFunc is called when exiting the logFunc production.
	ExitLogFunc(c *LogFuncContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)
}
