// Code generated from C:/repositories/alteryx_formulas/grammar\AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

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

// EnterFormulaIsField is called when production formulaIsField is entered.
func (s *BaseAlteryxFormulasListener) EnterFormulaIsField(ctx *FormulaIsFieldContext) {}

// ExitFormulaIsField is called when production formulaIsField is exited.
func (s *BaseAlteryxFormulasListener) ExitFormulaIsField(ctx *FormulaIsFieldContext) {}

// EnterFormulaIsNumber is called when production formulaIsNumber is entered.
func (s *BaseAlteryxFormulasListener) EnterFormulaIsNumber(ctx *FormulaIsNumberContext) {}

// ExitFormulaIsNumber is called when production formulaIsNumber is exited.
func (s *BaseAlteryxFormulasListener) ExitFormulaIsNumber(ctx *FormulaIsNumberContext) {}

// EnterFormulaIsDate is called when production formulaIsDate is entered.
func (s *BaseAlteryxFormulasListener) EnterFormulaIsDate(ctx *FormulaIsDateContext) {}

// ExitFormulaIsDate is called when production formulaIsDate is exited.
func (s *BaseAlteryxFormulasListener) ExitFormulaIsDate(ctx *FormulaIsDateContext) {}

// EnterFormulaIsString is called when production formulaIsString is entered.
func (s *BaseAlteryxFormulasListener) EnterFormulaIsString(ctx *FormulaIsStringContext) {}

// ExitFormulaIsString is called when production formulaIsString is exited.
func (s *BaseAlteryxFormulasListener) ExitFormulaIsString(ctx *FormulaIsStringContext) {}

// EnterFormulaIsBool is called when production formulaIsBool is entered.
func (s *BaseAlteryxFormulasListener) EnterFormulaIsBool(ctx *FormulaIsBoolContext) {}

// ExitFormulaIsBool is called when production formulaIsBool is exited.
func (s *BaseAlteryxFormulasListener) ExitFormulaIsBool(ctx *FormulaIsBoolContext) {}

// EnterFieldParenthesis is called when production fieldParenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterFieldParenthesis(ctx *FieldParenthesisContext) {}

// ExitFieldParenthesis is called when production fieldParenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitFieldParenthesis(ctx *FieldParenthesisContext) {}

// EnterAnyField is called when production anyField is entered.
func (s *BaseAlteryxFormulasListener) EnterAnyField(ctx *AnyFieldContext) {}

// ExitAnyField is called when production anyField is exited.
func (s *BaseAlteryxFormulasListener) ExitAnyField(ctx *AnyFieldContext) {}

// EnterStringIf is called when production stringIf is entered.
func (s *BaseAlteryxFormulasListener) EnterStringIf(ctx *StringIfContext) {}

// ExitStringIf is called when production stringIf is exited.
func (s *BaseAlteryxFormulasListener) ExitStringIf(ctx *StringIfContext) {}

// EnterStringFunc is called when production stringFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterStringFunc(ctx *StringFuncContext) {}

// ExitStringFunc is called when production stringFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitStringFunc(ctx *StringFuncContext) {}

// EnterConcatenate is called when production concatenate is entered.
func (s *BaseAlteryxFormulasListener) EnterConcatenate(ctx *ConcatenateContext) {}

// ExitConcatenate is called when production concatenate is exited.
func (s *BaseAlteryxFormulasListener) ExitConcatenate(ctx *ConcatenateContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterStringParenthesis is called when production stringParenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterStringParenthesis(ctx *StringParenthesisContext) {}

// ExitStringParenthesis is called when production stringParenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitStringParenthesis(ctx *StringParenthesisContext) {}

// EnterStringElseIf is called when production stringElseIf is entered.
func (s *BaseAlteryxFormulasListener) EnterStringElseIf(ctx *StringElseIfContext) {}

// ExitStringElseIf is called when production stringElseIf is exited.
func (s *BaseAlteryxFormulasListener) ExitStringElseIf(ctx *StringElseIfContext) {}

// EnterStringField is called when production stringField is entered.
func (s *BaseAlteryxFormulasListener) EnterStringField(ctx *StringFieldContext) {}

// ExitStringField is called when production stringField is exited.
func (s *BaseAlteryxFormulasListener) ExitStringField(ctx *StringFieldContext) {}

// EnterStringMin is called when production stringMin is entered.
func (s *BaseAlteryxFormulasListener) EnterStringMin(ctx *StringMinContext) {}

// ExitStringMin is called when production stringMin is exited.
func (s *BaseAlteryxFormulasListener) ExitStringMin(ctx *StringMinContext) {}

// EnterStringMax is called when production stringMax is entered.
func (s *BaseAlteryxFormulasListener) EnterStringMax(ctx *StringMaxContext) {}

// ExitStringMax is called when production stringMax is exited.
func (s *BaseAlteryxFormulasListener) ExitStringMax(ctx *StringMaxContext) {}

// EnterNumberParenthesis is called when production numberParenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberParenthesis(ctx *NumberParenthesisContext) {}

// ExitNumberParenthesis is called when production numberParenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberParenthesis(ctx *NumberParenthesisContext) {}

// EnterAdd is called when production add is entered.
func (s *BaseAlteryxFormulasListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseAlteryxFormulasListener) ExitAdd(ctx *AddContext) {}

// EnterNumberFunc is called when production numberFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberFunc(ctx *NumberFuncContext) {}

// ExitNumberFunc is called when production numberFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberFunc(ctx *NumberFuncContext) {}

// EnterNumberField is called when production numberField is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberField(ctx *NumberFieldContext) {}

// ExitNumberField is called when production numberField is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberField(ctx *NumberFieldContext) {}

// EnterSubtract is called when production subtract is entered.
func (s *BaseAlteryxFormulasListener) EnterSubtract(ctx *SubtractContext) {}

// ExitSubtract is called when production subtract is exited.
func (s *BaseAlteryxFormulasListener) ExitSubtract(ctx *SubtractContext) {}

// EnterNumberIf is called when production numberIf is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberIf(ctx *NumberIfContext) {}

// ExitNumberIf is called when production numberIf is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberIf(ctx *NumberIfContext) {}

// EnterNumberElseIf is called when production numberElseIf is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberElseIf(ctx *NumberElseIfContext) {}

// ExitNumberElseIf is called when production numberElseIf is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberElseIf(ctx *NumberElseIfContext) {}

// EnterDivide is called when production divide is entered.
func (s *BaseAlteryxFormulasListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production divide is exited.
func (s *BaseAlteryxFormulasListener) ExitDivide(ctx *DivideContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseAlteryxFormulasListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseAlteryxFormulasListener) ExitInteger(ctx *IntegerContext) {}

// EnterDecimal is called when production decimal is entered.
func (s *BaseAlteryxFormulasListener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production decimal is exited.
func (s *BaseAlteryxFormulasListener) ExitDecimal(ctx *DecimalContext) {}

// EnterMultiply is called when production multiply is entered.
func (s *BaseAlteryxFormulasListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production multiply is exited.
func (s *BaseAlteryxFormulasListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterPow is called when production pow is entered.
func (s *BaseAlteryxFormulasListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production pow is exited.
func (s *BaseAlteryxFormulasListener) ExitPow(ctx *PowContext) {}

// EnterNumberMin is called when production numberMin is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberMin(ctx *NumberMinContext) {}

// ExitNumberMin is called when production numberMin is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberMin(ctx *NumberMinContext) {}

// EnterNumberMax is called when production numberMax is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberMax(ctx *NumberMaxContext) {}

// ExitNumberMax is called when production numberMax is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberMax(ctx *NumberMaxContext) {}

// EnterNumberNull is called when production numberNull is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberNull(ctx *NumberNullContext) {}

// ExitNumberNull is called when production numberNull is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberNull(ctx *NumberNullContext) {}

// EnterDateParenthesis is called when production dateParenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterDateParenthesis(ctx *DateParenthesisContext) {}

// ExitDateParenthesis is called when production dateParenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitDateParenthesis(ctx *DateParenthesisContext) {}

// EnterDateIf is called when production dateIf is entered.
func (s *BaseAlteryxFormulasListener) EnterDateIf(ctx *DateIfContext) {}

// ExitDateIf is called when production dateIf is exited.
func (s *BaseAlteryxFormulasListener) ExitDateIf(ctx *DateIfContext) {}

// EnterDateElseIf is called when production dateElseIf is entered.
func (s *BaseAlteryxFormulasListener) EnterDateElseIf(ctx *DateElseIfContext) {}

// ExitDateElseIf is called when production dateElseIf is exited.
func (s *BaseAlteryxFormulasListener) ExitDateElseIf(ctx *DateElseIfContext) {}

// EnterDateFunc is called when production dateFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterDateFunc(ctx *DateFuncContext) {}

// ExitDateFunc is called when production dateFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitDateFunc(ctx *DateFuncContext) {}

// EnterDatetimeLiteral is called when production datetimeLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterDatetimeLiteral(ctx *DatetimeLiteralContext) {}

// ExitDatetimeLiteral is called when production datetimeLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitDatetimeLiteral(ctx *DatetimeLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterDateField is called when production dateField is entered.
func (s *BaseAlteryxFormulasListener) EnterDateField(ctx *DateFieldContext) {}

// ExitDateField is called when production dateField is exited.
func (s *BaseAlteryxFormulasListener) ExitDateField(ctx *DateFieldContext) {}

// EnterDateMin is called when production dateMin is entered.
func (s *BaseAlteryxFormulasListener) EnterDateMin(ctx *DateMinContext) {}

// ExitDateMin is called when production dateMin is exited.
func (s *BaseAlteryxFormulasListener) ExitDateMin(ctx *DateMinContext) {}

// EnterDateMax is called when production dateMax is entered.
func (s *BaseAlteryxFormulasListener) EnterDateMax(ctx *DateMaxContext) {}

// ExitDateMax is called when production dateMax is exited.
func (s *BaseAlteryxFormulasListener) ExitDateMax(ctx *DateMaxContext) {}

// EnterNumberIn is called when production numberIn is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberIn(ctx *NumberInContext) {}

// ExitNumberIn is called when production numberIn is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberIn(ctx *NumberInContext) {}

// EnterStringGreaterThan is called when production stringGreaterThan is entered.
func (s *BaseAlteryxFormulasListener) EnterStringGreaterThan(ctx *StringGreaterThanContext) {}

// ExitStringGreaterThan is called when production stringGreaterThan is exited.
func (s *BaseAlteryxFormulasListener) ExitStringGreaterThan(ctx *StringGreaterThanContext) {}

// EnterDateIn is called when production dateIn is entered.
func (s *BaseAlteryxFormulasListener) EnterDateIn(ctx *DateInContext) {}

// ExitDateIn is called when production dateIn is exited.
func (s *BaseAlteryxFormulasListener) ExitDateIn(ctx *DateInContext) {}

// EnterBoolParenthesis is called when production boolParenthesis is entered.
func (s *BaseAlteryxFormulasListener) EnterBoolParenthesis(ctx *BoolParenthesisContext) {}

// ExitBoolParenthesis is called when production boolParenthesis is exited.
func (s *BaseAlteryxFormulasListener) ExitBoolParenthesis(ctx *BoolParenthesisContext) {}

// EnterNumberNotIn is called when production numberNotIn is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberNotIn(ctx *NumberNotInContext) {}

// ExitNumberNotIn is called when production numberNotIn is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberNotIn(ctx *NumberNotInContext) {}

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterStringNotIn is called when production stringNotIn is entered.
func (s *BaseAlteryxFormulasListener) EnterStringNotIn(ctx *StringNotInContext) {}

// ExitStringNotIn is called when production stringNotIn is exited.
func (s *BaseAlteryxFormulasListener) ExitStringNotIn(ctx *StringNotInContext) {}

// EnterBoolField is called when production boolField is entered.
func (s *BaseAlteryxFormulasListener) EnterBoolField(ctx *BoolFieldContext) {}

// ExitBoolField is called when production boolField is exited.
func (s *BaseAlteryxFormulasListener) ExitBoolField(ctx *BoolFieldContext) {}

// EnterDateGreaterEqual is called when production dateGreaterEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterDateGreaterEqual(ctx *DateGreaterEqualContext) {}

// ExitDateGreaterEqual is called when production dateGreaterEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitDateGreaterEqual(ctx *DateGreaterEqualContext) {}

// EnterDateLessThan is called when production dateLessThan is entered.
func (s *BaseAlteryxFormulasListener) EnterDateLessThan(ctx *DateLessThanContext) {}

// ExitDateLessThan is called when production dateLessThan is exited.
func (s *BaseAlteryxFormulasListener) ExitDateLessThan(ctx *DateLessThanContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseAlteryxFormulasListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseAlteryxFormulasListener) ExitAnd(ctx *AndContext) {}

// EnterDateNotEqual is called when production dateNotEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterDateNotEqual(ctx *DateNotEqualContext) {}

// ExitDateNotEqual is called when production dateNotEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitDateNotEqual(ctx *DateNotEqualContext) {}

// EnterDateEqual is called when production dateEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterDateEqual(ctx *DateEqualContext) {}

// ExitDateEqual is called when production dateEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitDateEqual(ctx *DateEqualContext) {}

// EnterStringLessEqual is called when production stringLessEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterStringLessEqual(ctx *StringLessEqualContext) {}

// ExitStringLessEqual is called when production stringLessEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitStringLessEqual(ctx *StringLessEqualContext) {}

// EnterNumberEqual is called when production numberEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberEqual(ctx *NumberEqualContext) {}

// ExitNumberEqual is called when production numberEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberEqual(ctx *NumberEqualContext) {}

// EnterStringLessThan is called when production stringLessThan is entered.
func (s *BaseAlteryxFormulasListener) EnterStringLessThan(ctx *StringLessThanContext) {}

// ExitStringLessThan is called when production stringLessThan is exited.
func (s *BaseAlteryxFormulasListener) ExitStringLessThan(ctx *StringLessThanContext) {}

// EnterNumberLessEqual is called when production numberLessEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberLessEqual(ctx *NumberLessEqualContext) {}

// ExitNumberLessEqual is called when production numberLessEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberLessEqual(ctx *NumberLessEqualContext) {}

// EnterOr is called when production or is entered.
func (s *BaseAlteryxFormulasListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseAlteryxFormulasListener) ExitOr(ctx *OrContext) {}

// EnterNumberGreaterThan is called when production numberGreaterThan is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberGreaterThan(ctx *NumberGreaterThanContext) {}

// ExitNumberGreaterThan is called when production numberGreaterThan is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberGreaterThan(ctx *NumberGreaterThanContext) {}

// EnterDateLessEqual is called when production dateLessEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterDateLessEqual(ctx *DateLessEqualContext) {}

// ExitDateLessEqual is called when production dateLessEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitDateLessEqual(ctx *DateLessEqualContext) {}

// EnterStringEqual is called when production stringEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterStringEqual(ctx *StringEqualContext) {}

// ExitStringEqual is called when production stringEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitStringEqual(ctx *StringEqualContext) {}

// EnterStringIn is called when production stringIn is entered.
func (s *BaseAlteryxFormulasListener) EnterStringIn(ctx *StringInContext) {}

// ExitStringIn is called when production stringIn is exited.
func (s *BaseAlteryxFormulasListener) ExitStringIn(ctx *StringInContext) {}

// EnterStringGreaterEqual is called when production stringGreaterEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterStringGreaterEqual(ctx *StringGreaterEqualContext) {}

// ExitStringGreaterEqual is called when production stringGreaterEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitStringGreaterEqual(ctx *StringGreaterEqualContext) {}

// EnterDateGreaterThan is called when production dateGreaterThan is entered.
func (s *BaseAlteryxFormulasListener) EnterDateGreaterThan(ctx *DateGreaterThanContext) {}

// ExitDateGreaterThan is called when production dateGreaterThan is exited.
func (s *BaseAlteryxFormulasListener) ExitDateGreaterThan(ctx *DateGreaterThanContext) {}

// EnterStringNotEqual is called when production stringNotEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterStringNotEqual(ctx *StringNotEqualContext) {}

// ExitStringNotEqual is called when production stringNotEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitStringNotEqual(ctx *StringNotEqualContext) {}

// EnterNumberGreaterEqual is called when production numberGreaterEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberGreaterEqual(ctx *NumberGreaterEqualContext) {}

// ExitNumberGreaterEqual is called when production numberGreaterEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberGreaterEqual(ctx *NumberGreaterEqualContext) {}

// EnterDateNotIn is called when production dateNotIn is entered.
func (s *BaseAlteryxFormulasListener) EnterDateNotIn(ctx *DateNotInContext) {}

// ExitDateNotIn is called when production dateNotIn is exited.
func (s *BaseAlteryxFormulasListener) ExitDateNotIn(ctx *DateNotInContext) {}

// EnterNumberLessThan is called when production numberLessThan is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberLessThan(ctx *NumberLessThanContext) {}

// ExitNumberLessThan is called when production numberLessThan is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberLessThan(ctx *NumberLessThanContext) {}

// EnterNumberNotEqual is called when production numberNotEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberNotEqual(ctx *NumberNotEqualContext) {}

// ExitNumberNotEqual is called when production numberNotEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberNotEqual(ctx *NumberNotEqualContext) {}

// EnterStr is called when production str is entered.
func (s *BaseAlteryxFormulasListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseAlteryxFormulasListener) ExitStr(ctx *StrContext) {}
