// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type BaseAlteryxFormulasListener struct{}

var _ AlteryxFormulasListener = &BaseAlteryxFormulasListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlteryxFormulasListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlteryxFormulasListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlteryxFormulasListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlteryxFormulasListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMaxFunc is called when production maxFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterMaxFunc(ctx *MaxFuncContext) {}

// ExitMaxFunc is called when production maxFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitMaxFunc(ctx *MaxFuncContext) {}

// EnterCosFunc is called when production cosFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCosFunc(ctx *CosFuncContext) {}

// ExitCosFunc is called when production cosFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCosFunc(ctx *CosFuncContext) {}

// EnterReplaceFunc is called when production replaceFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterReplaceFunc(ctx *ReplaceFuncContext) {}

// ExitReplaceFunc is called when production replaceFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitReplaceFunc(ctx *ReplaceFuncContext) {}

// EnterRandIntFunc is called when production randIntFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRandIntFunc(ctx *RandIntFuncContext) {}

// ExitRandIntFunc is called when production randIntFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRandIntFunc(ctx *RandIntFuncContext) {}

// EnterCharToIntFunc is called when production charToIntFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCharToIntFunc(ctx *CharToIntFuncContext) {}

// ExitCharToIntFunc is called when production charToIntFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCharToIntFunc(ctx *CharToIntFuncContext) {}

// EnterTrimRightFunc is called when production trimRightFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterTrimRightFunc(ctx *TrimRightFuncContext) {}

// ExitTrimRightFunc is called when production trimRightFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitTrimRightFunc(ctx *TrimRightFuncContext) {}

// EnterLeftFunc is called when production leftFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterLeftFunc(ctx *LeftFuncContext) {}

// ExitLeftFunc is called when production leftFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitLeftFunc(ctx *LeftFuncContext) {}

// EnterTanFunc is called when production tanFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterTanFunc(ctx *TanFuncContext) {}

// ExitTanFunc is called when production tanFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitTanFunc(ctx *TanFuncContext) {}

// EnterToDateFunc is called when production toDateFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterToDateFunc(ctx *ToDateFuncContext) {}

// ExitToDateFunc is called when production toDateFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitToDateFunc(ctx *ToDateFuncContext) {}

// EnterRegexReplaceFunc is called when production regexReplaceFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRegexReplaceFunc(ctx *RegexReplaceFuncContext) {}

// ExitRegexReplaceFunc is called when production regexReplaceFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRegexReplaceFunc(ctx *RegexReplaceFuncContext) {}

// EnterFloorFunc is called when production floorFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterFloorFunc(ctx *FloorFuncContext) {}

// ExitFloorFunc is called when production floorFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitFloorFunc(ctx *FloorFuncContext) {}

// EnterNotIn is called when production notIn is entered.
func (s *BaseAlteryxFormulasListener) EnterNotIn(ctx *NotInContext) {}

// ExitNotIn is called when production notIn is exited.
func (s *BaseAlteryxFormulasListener) ExitNotIn(ctx *NotInContext) {}

// EnterTanhFunc is called when production tanhFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterTanhFunc(ctx *TanhFuncContext) {}

// ExitTanhFunc is called when production tanhFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitTanhFunc(ctx *TanhFuncContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseAlteryxFormulasListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseAlteryxFormulasListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterAsinFunc is called when production asinFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterAsinFunc(ctx *AsinFuncContext) {}

// ExitAsinFunc is called when production asinFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitAsinFunc(ctx *AsinFuncContext) {}

// EnterAdd is called when production add is entered.
func (s *BaseAlteryxFormulasListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseAlteryxFormulasListener) ExitAdd(ctx *AddContext) {}

// EnterCoshFunc is called when production coshFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCoshFunc(ctx *CoshFuncContext) {}

// ExitCoshFunc is called when production coshFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCoshFunc(ctx *CoshFuncContext) {}

// EnterIsNullFunc is called when production isNullFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterIsNullFunc(ctx *IsNullFuncContext) {}

// ExitIsNullFunc is called when production isNullFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitIsNullFunc(ctx *IsNullFuncContext) {}

// EnterIn is called when production in is entered.
func (s *BaseAlteryxFormulasListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production in is exited.
func (s *BaseAlteryxFormulasListener) ExitIn(ctx *InContext) {}

// EnterSubtract is called when production subtract is entered.
func (s *BaseAlteryxFormulasListener) EnterSubtract(ctx *SubtractContext) {}

// ExitSubtract is called when production subtract is exited.
func (s *BaseAlteryxFormulasListener) ExitSubtract(ctx *SubtractContext) {}

// EnterNotEqual is called when production notEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterNotEqual(ctx *NotEqualContext) {}

// ExitNotEqual is called when production notEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitNotEqual(ctx *NotEqualContext) {}

// EnterParenthesis is called when production parenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production parenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterHexToNumberFunc is called when production hexToNumberFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterHexToNumberFunc(ctx *HexToNumberFuncContext) {}

// ExitHexToNumberFunc is called when production hexToNumberFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitHexToNumberFunc(ctx *HexToNumberFuncContext) {}

// EnterMedianFunc is called when production medianFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterMedianFunc(ctx *MedianFuncContext) {}

// ExitMedianFunc is called when production medianFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitMedianFunc(ctx *MedianFuncContext) {}

// EnterNullFunc is called when production nullFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterNullFunc(ctx *NullFuncContext) {}

// ExitNullFunc is called when production nullFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitNullFunc(ctx *NullFuncContext) {}

// EnterIsEmptyFunc is called when production isEmptyFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterIsEmptyFunc(ctx *IsEmptyFuncContext) {}

// ExitIsEmptyFunc is called when production isEmptyFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitIsEmptyFunc(ctx *IsEmptyFuncContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterRegexMatchFunc is called when production regexMatchFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRegexMatchFunc(ctx *RegexMatchFuncContext) {}

// ExitRegexMatchFunc is called when production regexMatchFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRegexMatchFunc(ctx *RegexMatchFuncContext) {}

// EnterAbsFunc is called when production absFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterAbsFunc(ctx *AbsFuncContext) {}

// ExitAbsFunc is called when production absFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitAbsFunc(ctx *AbsFuncContext) {}

// EnterAtan2Func is called when production atan2Func is entered.
func (s *BaseAlteryxFormulasListener) EnterAtan2Func(ctx *Atan2FuncContext) {}

// ExitAtan2Func is called when production atan2Func is exited.
func (s *BaseAlteryxFormulasListener) ExitAtan2Func(ctx *Atan2FuncContext) {}

// EnterPadLeftFunc is called when production padLeftFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterPadLeftFunc(ctx *PadLeftFuncContext) {}

// ExitPadLeftFunc is called when production padLeftFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitPadLeftFunc(ctx *PadLeftFuncContext) {}

// EnterFindStringFunc is called when production findStringFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterFindStringFunc(ctx *FindStringFuncContext) {}

// ExitFindStringFunc is called when production findStringFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitFindStringFunc(ctx *FindStringFuncContext) {}

// EnterLog10Func is called when production log10Func is entered.
func (s *BaseAlteryxFormulasListener) EnterLog10Func(ctx *Log10FuncContext) {}

// ExitLog10Func is called when production log10Func is exited.
func (s *BaseAlteryxFormulasListener) ExitLog10Func(ctx *Log10FuncContext) {}

// EnterCeilFunc is called when production ceilFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCeilFunc(ctx *CeilFuncContext) {}

// ExitCeilFunc is called when production ceilFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCeilFunc(ctx *CeilFuncContext) {}

// EnterEndsWithFunc is called when production endsWithFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterEndsWithFunc(ctx *EndsWithFuncContext) {}

// ExitEndsWithFunc is called when production endsWithFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitEndsWithFunc(ctx *EndsWithFuncContext) {}

// EnterDistanceFunc is called when production distanceFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterDistanceFunc(ctx *DistanceFuncContext) {}

// ExitDistanceFunc is called when production distanceFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitDistanceFunc(ctx *DistanceFuncContext) {}

// EnterCharFromIntFunc is called when production charFromIntFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCharFromIntFunc(ctx *CharFromIntFuncContext) {}

// ExitCharFromIntFunc is called when production charFromIntFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCharFromIntFunc(ctx *CharFromIntFuncContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterAcosFunc is called when production acosFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterAcosFunc(ctx *AcosFuncContext) {}

// ExitAcosFunc is called when production acosFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitAcosFunc(ctx *AcosFuncContext) {}

// EnterSqrtFunc is called when production sqrtFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterSqrtFunc(ctx *SqrtFuncContext) {}

// ExitSqrtFunc is called when production sqrtFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitSqrtFunc(ctx *SqrtFuncContext) {}

// EnterSinhFunc is called when production sinhFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterSinhFunc(ctx *SinhFuncContext) {}

// ExitSinhFunc is called when production sinhFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitSinhFunc(ctx *SinhFuncContext) {}

// EnterRandFunc is called when production randFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRandFunc(ctx *RandFuncContext) {}

// ExitRandFunc is called when production randFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRandFunc(ctx *RandFuncContext) {}

// EnterMinFunc is called when production minFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterMinFunc(ctx *MinFuncContext) {}

// ExitMinFunc is called when production minFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitMinFunc(ctx *MinFuncContext) {}

// EnterPiFunc is called when production piFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterPiFunc(ctx *PiFuncContext) {}

// ExitPiFunc is called when production piFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitPiFunc(ctx *PiFuncContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseAlteryxFormulasListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseAlteryxFormulasListener) ExitAnd(ctx *AndContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseAlteryxFormulasListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseAlteryxFormulasListener) ExitLessThan(ctx *LessThanContext) {}

// EnterPadRightFunc is called when production padRightFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterPadRightFunc(ctx *PadRightFuncContext) {}

// ExitPadRightFunc is called when production padRightFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitPadRightFunc(ctx *PadRightFuncContext) {}

// EnterRegexCountMatchesFunc is called when production regexCountMatchesFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRegexCountMatchesFunc(ctx *RegexCountMatchesFuncContext) {}

// ExitRegexCountMatchesFunc is called when production regexCountMatchesFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRegexCountMatchesFunc(ctx *RegexCountMatchesFuncContext) {}

// EnterDivide is called when production divide is entered.
func (s *BaseAlteryxFormulasListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production divide is exited.
func (s *BaseAlteryxFormulasListener) ExitDivide(ctx *DivideContext) {}

// EnterExprIf is called when production exprIf is entered.
func (s *BaseAlteryxFormulasListener) EnterExprIf(ctx *ExprIfContext) {}

// ExitExprIf is called when production exprIf is exited.
func (s *BaseAlteryxFormulasListener) ExitExprIf(ctx *ExprIfContext) {}

// EnterContainsFunc is called when production containsFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterContainsFunc(ctx *ContainsFuncContext) {}

// ExitContainsFunc is called when production containsFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitContainsFunc(ctx *ContainsFuncContext) {}

// EnterLowercaseFunc is called when production lowercaseFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterLowercaseFunc(ctx *LowercaseFuncContext) {}

// ExitLowercaseFunc is called when production lowercaseFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitLowercaseFunc(ctx *LowercaseFuncContext) {}

// EnterTrimFunc is called when production trimFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterTrimFunc(ctx *TrimFuncContext) {}

// ExitTrimFunc is called when production trimFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitTrimFunc(ctx *TrimFuncContext) {}

// EnterMultiply is called when production multiply is entered.
func (s *BaseAlteryxFormulasListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production multiply is exited.
func (s *BaseAlteryxFormulasListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterAtanFunc is called when production atanFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterAtanFunc(ctx *AtanFuncContext) {}

// ExitAtanFunc is called when production atanFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitAtanFunc(ctx *AtanFuncContext) {}

// EnterRoundFunc is called when production roundFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRoundFunc(ctx *RoundFuncContext) {}

// ExitRoundFunc is called when production roundFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRoundFunc(ctx *RoundFuncContext) {}

// EnterSwitchFunc is called when production switchFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterSwitchFunc(ctx *SwitchFuncContext) {}

// ExitSwitchFunc is called when production switchFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitSwitchFunc(ctx *SwitchFuncContext) {}

// EnterExprField is called when production exprField is entered.
func (s *BaseAlteryxFormulasListener) EnterExprField(ctx *ExprFieldContext) {}

// ExitExprField is called when production exprField is exited.
func (s *BaseAlteryxFormulasListener) ExitExprField(ctx *ExprFieldContext) {}

// EnterGetWordFunc is called when production getWordFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterGetWordFunc(ctx *GetWordFuncContext) {}

// ExitGetWordFunc is called when production getWordFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitGetWordFunc(ctx *GetWordFuncContext) {}

// EnterLengthFunc is called when production lengthFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterLengthFunc(ctx *LengthFuncContext) {}

// ExitLengthFunc is called when production lengthFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitLengthFunc(ctx *LengthFuncContext) {}

// EnterOr is called when production or is entered.
func (s *BaseAlteryxFormulasListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseAlteryxFormulasListener) ExitOr(ctx *OrContext) {}

// EnterExpFunc is called when production expFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterExpFunc(ctx *ExpFuncContext) {}

// ExitExpFunc is called when production expFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitExpFunc(ctx *ExpFuncContext) {}

// EnterCountWordsFunc is called when production countWordsFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterCountWordsFunc(ctx *CountWordsFuncContext) {}

// ExitCountWordsFunc is called when production countWordsFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitCountWordsFunc(ctx *CountWordsFuncContext) {}

// EnterUppercaseFunc is called when production uppercaseFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterUppercaseFunc(ctx *UppercaseFuncContext) {}

// ExitUppercaseFunc is called when production uppercaseFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitUppercaseFunc(ctx *UppercaseFuncContext) {}

// EnterPowFunc is called when production powFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterPowFunc(ctx *PowFuncContext) {}

// ExitPowFunc is called when production powFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitPowFunc(ctx *PowFuncContext) {}

// EnterToDatetimeFunc is called when production toDatetimeFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterToDatetimeFunc(ctx *ToDatetimeFuncContext) {}

// ExitToDatetimeFunc is called when production toDatetimeFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitToDatetimeFunc(ctx *ToDatetimeFuncContext) {}

// EnterIifFunc is called when production iifFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterIifFunc(ctx *IifFuncContext) {}

// ExitIifFunc is called when production iifFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitIifFunc(ctx *IifFuncContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseAlteryxFormulasListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseAlteryxFormulasListener) ExitEqual(ctx *EqualContext) {}

// EnterRightFunc is called when production rightFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterRightFunc(ctx *RightFuncContext) {}

// ExitRightFunc is called when production rightFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitRightFunc(ctx *RightFuncContext) {}

// EnterModFunc is called when production modFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterModFunc(ctx *ModFuncContext) {}

// ExitModFunc is called when production modFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitModFunc(ctx *ModFuncContext) {}

// EnterSinFunc is called when production sinFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterSinFunc(ctx *SinFuncContext) {}

// ExitSinFunc is called when production sinFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitSinFunc(ctx *SinFuncContext) {}

// EnterTrimLeftFunc is called when production trimLeftFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterTrimLeftFunc(ctx *TrimLeftFuncContext) {}

// ExitTrimLeftFunc is called when production trimLeftFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitTrimLeftFunc(ctx *TrimLeftFuncContext) {}

// EnterAverageFunc is called when production averageFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterAverageFunc(ctx *AverageFuncContext) {}

// ExitAverageFunc is called when production averageFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitAverageFunc(ctx *AverageFuncContext) {}

// EnterSubstringFunc is called when production substringFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterSubstringFunc(ctx *SubstringFuncContext) {}

// ExitSubstringFunc is called when production substringFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitSubstringFunc(ctx *SubstringFuncContext) {}

// EnterGreaterEqual is called when production greaterEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterGreaterEqual(ctx *GreaterEqualContext) {}

// ExitGreaterEqual is called when production greaterEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitGreaterEqual(ctx *GreaterEqualContext) {}

// EnterLessEqual is called when production lessEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterLessEqual(ctx *LessEqualContext) {}

// ExitLessEqual is called when production lessEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitLessEqual(ctx *LessEqualContext) {}

// EnterLogFunc is called when production logFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterLogFunc(ctx *LogFuncContext) {}

// ExitLogFunc is called when production logFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitLogFunc(ctx *LogFuncContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}
