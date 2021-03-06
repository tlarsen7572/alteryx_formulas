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

	// EnterReplaceFunc is called when entering the replaceFunc production.
	EnterReplaceFunc(c *ReplaceFuncContext)

	// EnterRandIntFunc is called when entering the randIntFunc production.
	EnterRandIntFunc(c *RandIntFuncContext)

	// EnterCharToIntFunc is called when entering the charToIntFunc production.
	EnterCharToIntFunc(c *CharToIntFuncContext)

	// EnterTrimRightFunc is called when entering the trimRightFunc production.
	EnterTrimRightFunc(c *TrimRightFuncContext)

	// EnterLeftFunc is called when entering the leftFunc production.
	EnterLeftFunc(c *LeftFuncContext)

	// EnterTanFunc is called when entering the tanFunc production.
	EnterTanFunc(c *TanFuncContext)

	// EnterToDateFunc is called when entering the toDateFunc production.
	EnterToDateFunc(c *ToDateFuncContext)

	// EnterRegexReplaceFunc is called when entering the regexReplaceFunc production.
	EnterRegexReplaceFunc(c *RegexReplaceFuncContext)

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

	// EnterIsNullFunc is called when entering the isNullFunc production.
	EnterIsNullFunc(c *IsNullFuncContext)

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

	// EnterIsEmptyFunc is called when entering the isEmptyFunc production.
	EnterIsEmptyFunc(c *IsEmptyFuncContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterRegexMatchFunc is called when entering the regexMatchFunc production.
	EnterRegexMatchFunc(c *RegexMatchFuncContext)

	// EnterAbsFunc is called when entering the absFunc production.
	EnterAbsFunc(c *AbsFuncContext)

	// EnterAtan2Func is called when entering the atan2Func production.
	EnterAtan2Func(c *Atan2FuncContext)

	// EnterPadLeftFunc is called when entering the padLeftFunc production.
	EnterPadLeftFunc(c *PadLeftFuncContext)

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

	// EnterPadRightFunc is called when entering the padRightFunc production.
	EnterPadRightFunc(c *PadRightFuncContext)

	// EnterRegexCountMatchesFunc is called when entering the regexCountMatchesFunc production.
	EnterRegexCountMatchesFunc(c *RegexCountMatchesFuncContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterExprIf is called when entering the exprIf production.
	EnterExprIf(c *ExprIfContext)

	// EnterContainsFunc is called when entering the containsFunc production.
	EnterContainsFunc(c *ContainsFuncContext)

	// EnterLowercaseFunc is called when entering the lowercaseFunc production.
	EnterLowercaseFunc(c *LowercaseFuncContext)

	// EnterTrimFunc is called when entering the trimFunc production.
	EnterTrimFunc(c *TrimFuncContext)

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

	// EnterGetWordFunc is called when entering the getWordFunc production.
	EnterGetWordFunc(c *GetWordFuncContext)

	// EnterLengthFunc is called when entering the lengthFunc production.
	EnterLengthFunc(c *LengthFuncContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterExpFunc is called when entering the expFunc production.
	EnterExpFunc(c *ExpFuncContext)

	// EnterCountWordsFunc is called when entering the countWordsFunc production.
	EnterCountWordsFunc(c *CountWordsFuncContext)

	// EnterUppercaseFunc is called when entering the uppercaseFunc production.
	EnterUppercaseFunc(c *UppercaseFuncContext)

	// EnterPowFunc is called when entering the powFunc production.
	EnterPowFunc(c *PowFuncContext)

	// EnterToDatetimeFunc is called when entering the toDatetimeFunc production.
	EnterToDatetimeFunc(c *ToDatetimeFuncContext)

	// EnterIifFunc is called when entering the iifFunc production.
	EnterIifFunc(c *IifFuncContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterRightFunc is called when entering the rightFunc production.
	EnterRightFunc(c *RightFuncContext)

	// EnterModFunc is called when entering the modFunc production.
	EnterModFunc(c *ModFuncContext)

	// EnterSinFunc is called when entering the sinFunc production.
	EnterSinFunc(c *SinFuncContext)

	// EnterTrimLeftFunc is called when entering the trimLeftFunc production.
	EnterTrimLeftFunc(c *TrimLeftFuncContext)

	// EnterAverageFunc is called when entering the averageFunc production.
	EnterAverageFunc(c *AverageFuncContext)

	// EnterSubstringFunc is called when entering the substringFunc production.
	EnterSubstringFunc(c *SubstringFuncContext)

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

	// ExitReplaceFunc is called when exiting the replaceFunc production.
	ExitReplaceFunc(c *ReplaceFuncContext)

	// ExitRandIntFunc is called when exiting the randIntFunc production.
	ExitRandIntFunc(c *RandIntFuncContext)

	// ExitCharToIntFunc is called when exiting the charToIntFunc production.
	ExitCharToIntFunc(c *CharToIntFuncContext)

	// ExitTrimRightFunc is called when exiting the trimRightFunc production.
	ExitTrimRightFunc(c *TrimRightFuncContext)

	// ExitLeftFunc is called when exiting the leftFunc production.
	ExitLeftFunc(c *LeftFuncContext)

	// ExitTanFunc is called when exiting the tanFunc production.
	ExitTanFunc(c *TanFuncContext)

	// ExitToDateFunc is called when exiting the toDateFunc production.
	ExitToDateFunc(c *ToDateFuncContext)

	// ExitRegexReplaceFunc is called when exiting the regexReplaceFunc production.
	ExitRegexReplaceFunc(c *RegexReplaceFuncContext)

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

	// ExitIsNullFunc is called when exiting the isNullFunc production.
	ExitIsNullFunc(c *IsNullFuncContext)

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

	// ExitIsEmptyFunc is called when exiting the isEmptyFunc production.
	ExitIsEmptyFunc(c *IsEmptyFuncContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitRegexMatchFunc is called when exiting the regexMatchFunc production.
	ExitRegexMatchFunc(c *RegexMatchFuncContext)

	// ExitAbsFunc is called when exiting the absFunc production.
	ExitAbsFunc(c *AbsFuncContext)

	// ExitAtan2Func is called when exiting the atan2Func production.
	ExitAtan2Func(c *Atan2FuncContext)

	// ExitPadLeftFunc is called when exiting the padLeftFunc production.
	ExitPadLeftFunc(c *PadLeftFuncContext)

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

	// ExitPadRightFunc is called when exiting the padRightFunc production.
	ExitPadRightFunc(c *PadRightFuncContext)

	// ExitRegexCountMatchesFunc is called when exiting the regexCountMatchesFunc production.
	ExitRegexCountMatchesFunc(c *RegexCountMatchesFuncContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitExprIf is called when exiting the exprIf production.
	ExitExprIf(c *ExprIfContext)

	// ExitContainsFunc is called when exiting the containsFunc production.
	ExitContainsFunc(c *ContainsFuncContext)

	// ExitLowercaseFunc is called when exiting the lowercaseFunc production.
	ExitLowercaseFunc(c *LowercaseFuncContext)

	// ExitTrimFunc is called when exiting the trimFunc production.
	ExitTrimFunc(c *TrimFuncContext)

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

	// ExitGetWordFunc is called when exiting the getWordFunc production.
	ExitGetWordFunc(c *GetWordFuncContext)

	// ExitLengthFunc is called when exiting the lengthFunc production.
	ExitLengthFunc(c *LengthFuncContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitExpFunc is called when exiting the expFunc production.
	ExitExpFunc(c *ExpFuncContext)

	// ExitCountWordsFunc is called when exiting the countWordsFunc production.
	ExitCountWordsFunc(c *CountWordsFuncContext)

	// ExitUppercaseFunc is called when exiting the uppercaseFunc production.
	ExitUppercaseFunc(c *UppercaseFuncContext)

	// ExitPowFunc is called when exiting the powFunc production.
	ExitPowFunc(c *PowFuncContext)

	// ExitToDatetimeFunc is called when exiting the toDatetimeFunc production.
	ExitToDatetimeFunc(c *ToDatetimeFuncContext)

	// ExitIifFunc is called when exiting the iifFunc production.
	ExitIifFunc(c *IifFuncContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitRightFunc is called when exiting the rightFunc production.
	ExitRightFunc(c *RightFuncContext)

	// ExitModFunc is called when exiting the modFunc production.
	ExitModFunc(c *ModFuncContext)

	// ExitSinFunc is called when exiting the sinFunc production.
	ExitSinFunc(c *SinFuncContext)

	// ExitTrimLeftFunc is called when exiting the trimLeftFunc production.
	ExitTrimLeftFunc(c *TrimLeftFuncContext)

	// ExitAverageFunc is called when exiting the averageFunc production.
	ExitAverageFunc(c *AverageFuncContext)

	// ExitSubstringFunc is called when exiting the substringFunc production.
	ExitSubstringFunc(c *SubstringFuncContext)

	// ExitGreaterEqual is called when exiting the greaterEqual production.
	ExitGreaterEqual(c *GreaterEqualContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitLogFunc is called when exiting the logFunc production.
	ExitLogFunc(c *LogFuncContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)
}
