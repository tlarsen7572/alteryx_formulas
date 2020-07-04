// Code generated from C:/repositories/alteryx_formulas/grammar\AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type AlteryxFormulasListener interface {
	antlr.ParseTreeListener

	// EnterFormulaIsField is called when entering the formulaIsField production.
	EnterFormulaIsField(c *FormulaIsFieldContext)

	// EnterFormulaIsNumber is called when entering the formulaIsNumber production.
	EnterFormulaIsNumber(c *FormulaIsNumberContext)

	// EnterFormulaIsDate is called when entering the formulaIsDate production.
	EnterFormulaIsDate(c *FormulaIsDateContext)

	// EnterFormulaIsString is called when entering the formulaIsString production.
	EnterFormulaIsString(c *FormulaIsStringContext)

	// EnterFormulaIsBool is called when entering the formulaIsBool production.
	EnterFormulaIsBool(c *FormulaIsBoolContext)

	// EnterFieldParenthesis is called when entering the fieldParenthesis production.
	EnterFieldParenthesis(c *FieldParenthesisContext)

	// EnterAnyField is called when entering the anyField production.
	EnterAnyField(c *AnyFieldContext)

	// EnterStringIf is called when entering the stringIf production.
	EnterStringIf(c *StringIfContext)

	// EnterStringFunc is called when entering the stringFunc production.
	EnterStringFunc(c *StringFuncContext)

	// EnterConcatenate is called when entering the concatenate production.
	EnterConcatenate(c *ConcatenateContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterStringParenthesis is called when entering the stringParenthesis production.
	EnterStringParenthesis(c *StringParenthesisContext)

	// EnterStringElseIf is called when entering the stringElseIf production.
	EnterStringElseIf(c *StringElseIfContext)

	// EnterStringField is called when entering the stringField production.
	EnterStringField(c *StringFieldContext)

	// EnterStringMin is called when entering the stringMin production.
	EnterStringMin(c *StringMinContext)

	// EnterStringMax is called when entering the stringMax production.
	EnterStringMax(c *StringMaxContext)

	// EnterNumberParenthesis is called when entering the numberParenthesis production.
	EnterNumberParenthesis(c *NumberParenthesisContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterNumberFunc is called when entering the numberFunc production.
	EnterNumberFunc(c *NumberFuncContext)

	// EnterNumberField is called when entering the numberField production.
	EnterNumberField(c *NumberFieldContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterNumberIf is called when entering the numberIf production.
	EnterNumberIf(c *NumberIfContext)

	// EnterNumberElseIf is called when entering the numberElseIf production.
	EnterNumberElseIf(c *NumberElseIfContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterDecimal is called when entering the decimal production.
	EnterDecimal(c *DecimalContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterPow is called when entering the pow production.
	EnterPow(c *PowContext)

	// EnterNumberMin is called when entering the numberMin production.
	EnterNumberMin(c *NumberMinContext)

	// EnterNumberMax is called when entering the numberMax production.
	EnterNumberMax(c *NumberMaxContext)

	// EnterNumberNull is called when entering the numberNull production.
	EnterNumberNull(c *NumberNullContext)

	// EnterDateParenthesis is called when entering the dateParenthesis production.
	EnterDateParenthesis(c *DateParenthesisContext)

	// EnterDateIf is called when entering the dateIf production.
	EnterDateIf(c *DateIfContext)

	// EnterDateElseIf is called when entering the dateElseIf production.
	EnterDateElseIf(c *DateElseIfContext)

	// EnterDateFunc is called when entering the dateFunc production.
	EnterDateFunc(c *DateFuncContext)

	// EnterDatetimeLiteral is called when entering the datetimeLiteral production.
	EnterDatetimeLiteral(c *DatetimeLiteralContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterDateField is called when entering the dateField production.
	EnterDateField(c *DateFieldContext)

	// EnterDateMin is called when entering the dateMin production.
	EnterDateMin(c *DateMinContext)

	// EnterDateMax is called when entering the dateMax production.
	EnterDateMax(c *DateMaxContext)

	// EnterNumberIn is called when entering the numberIn production.
	EnterNumberIn(c *NumberInContext)

	// EnterStringGreaterThan is called when entering the stringGreaterThan production.
	EnterStringGreaterThan(c *StringGreaterThanContext)

	// EnterDateIn is called when entering the dateIn production.
	EnterDateIn(c *DateInContext)

	// EnterBoolParenthesis is called when entering the boolParenthesis production.
	EnterBoolParenthesis(c *BoolParenthesisContext)

	// EnterNumberNotIn is called when entering the numberNotIn production.
	EnterNumberNotIn(c *NumberNotInContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterStringNotIn is called when entering the stringNotIn production.
	EnterStringNotIn(c *StringNotInContext)

	// EnterBoolField is called when entering the boolField production.
	EnterBoolField(c *BoolFieldContext)

	// EnterDateGreaterEqual is called when entering the dateGreaterEqual production.
	EnterDateGreaterEqual(c *DateGreaterEqualContext)

	// EnterDateLessThan is called when entering the dateLessThan production.
	EnterDateLessThan(c *DateLessThanContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterDateNotEqual is called when entering the dateNotEqual production.
	EnterDateNotEqual(c *DateNotEqualContext)

	// EnterDateEqual is called when entering the dateEqual production.
	EnterDateEqual(c *DateEqualContext)

	// EnterStringLessEqual is called when entering the stringLessEqual production.
	EnterStringLessEqual(c *StringLessEqualContext)

	// EnterNumberEqual is called when entering the numberEqual production.
	EnterNumberEqual(c *NumberEqualContext)

	// EnterStringLessThan is called when entering the stringLessThan production.
	EnterStringLessThan(c *StringLessThanContext)

	// EnterNumberLessEqual is called when entering the numberLessEqual production.
	EnterNumberLessEqual(c *NumberLessEqualContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterNumberGreaterThan is called when entering the numberGreaterThan production.
	EnterNumberGreaterThan(c *NumberGreaterThanContext)

	// EnterDateLessEqual is called when entering the dateLessEqual production.
	EnterDateLessEqual(c *DateLessEqualContext)

	// EnterStringEqual is called when entering the stringEqual production.
	EnterStringEqual(c *StringEqualContext)

	// EnterStringIn is called when entering the stringIn production.
	EnterStringIn(c *StringInContext)

	// EnterStringGreaterEqual is called when entering the stringGreaterEqual production.
	EnterStringGreaterEqual(c *StringGreaterEqualContext)

	// EnterDateGreaterThan is called when entering the dateGreaterThan production.
	EnterDateGreaterThan(c *DateGreaterThanContext)

	// EnterStringNotEqual is called when entering the stringNotEqual production.
	EnterStringNotEqual(c *StringNotEqualContext)

	// EnterNumberGreaterEqual is called when entering the numberGreaterEqual production.
	EnterNumberGreaterEqual(c *NumberGreaterEqualContext)

	// EnterDateNotIn is called when entering the dateNotIn production.
	EnterDateNotIn(c *DateNotInContext)

	// EnterNumberLessThan is called when entering the numberLessThan production.
	EnterNumberLessThan(c *NumberLessThanContext)

	// EnterNumberNotEqual is called when entering the numberNotEqual production.
	EnterNumberNotEqual(c *NumberNotEqualContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// ExitFormulaIsField is called when exiting the formulaIsField production.
	ExitFormulaIsField(c *FormulaIsFieldContext)

	// ExitFormulaIsNumber is called when exiting the formulaIsNumber production.
	ExitFormulaIsNumber(c *FormulaIsNumberContext)

	// ExitFormulaIsDate is called when exiting the formulaIsDate production.
	ExitFormulaIsDate(c *FormulaIsDateContext)

	// ExitFormulaIsString is called when exiting the formulaIsString production.
	ExitFormulaIsString(c *FormulaIsStringContext)

	// ExitFormulaIsBool is called when exiting the formulaIsBool production.
	ExitFormulaIsBool(c *FormulaIsBoolContext)

	// ExitFieldParenthesis is called when exiting the fieldParenthesis production.
	ExitFieldParenthesis(c *FieldParenthesisContext)

	// ExitAnyField is called when exiting the anyField production.
	ExitAnyField(c *AnyFieldContext)

	// ExitStringIf is called when exiting the stringIf production.
	ExitStringIf(c *StringIfContext)

	// ExitStringFunc is called when exiting the stringFunc production.
	ExitStringFunc(c *StringFuncContext)

	// ExitConcatenate is called when exiting the concatenate production.
	ExitConcatenate(c *ConcatenateContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitStringParenthesis is called when exiting the stringParenthesis production.
	ExitStringParenthesis(c *StringParenthesisContext)

	// ExitStringElseIf is called when exiting the stringElseIf production.
	ExitStringElseIf(c *StringElseIfContext)

	// ExitStringField is called when exiting the stringField production.
	ExitStringField(c *StringFieldContext)

	// ExitStringMin is called when exiting the stringMin production.
	ExitStringMin(c *StringMinContext)

	// ExitStringMax is called when exiting the stringMax production.
	ExitStringMax(c *StringMaxContext)

	// ExitNumberParenthesis is called when exiting the numberParenthesis production.
	ExitNumberParenthesis(c *NumberParenthesisContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitNumberFunc is called when exiting the numberFunc production.
	ExitNumberFunc(c *NumberFuncContext)

	// ExitNumberField is called when exiting the numberField production.
	ExitNumberField(c *NumberFieldContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitNumberIf is called when exiting the numberIf production.
	ExitNumberIf(c *NumberIfContext)

	// ExitNumberElseIf is called when exiting the numberElseIf production.
	ExitNumberElseIf(c *NumberElseIfContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitDecimal is called when exiting the decimal production.
	ExitDecimal(c *DecimalContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitPow is called when exiting the pow production.
	ExitPow(c *PowContext)

	// ExitNumberMin is called when exiting the numberMin production.
	ExitNumberMin(c *NumberMinContext)

	// ExitNumberMax is called when exiting the numberMax production.
	ExitNumberMax(c *NumberMaxContext)

	// ExitNumberNull is called when exiting the numberNull production.
	ExitNumberNull(c *NumberNullContext)

	// ExitDateParenthesis is called when exiting the dateParenthesis production.
	ExitDateParenthesis(c *DateParenthesisContext)

	// ExitDateIf is called when exiting the dateIf production.
	ExitDateIf(c *DateIfContext)

	// ExitDateElseIf is called when exiting the dateElseIf production.
	ExitDateElseIf(c *DateElseIfContext)

	// ExitDateFunc is called when exiting the dateFunc production.
	ExitDateFunc(c *DateFuncContext)

	// ExitDatetimeLiteral is called when exiting the datetimeLiteral production.
	ExitDatetimeLiteral(c *DatetimeLiteralContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitDateField is called when exiting the dateField production.
	ExitDateField(c *DateFieldContext)

	// ExitDateMin is called when exiting the dateMin production.
	ExitDateMin(c *DateMinContext)

	// ExitDateMax is called when exiting the dateMax production.
	ExitDateMax(c *DateMaxContext)

	// ExitNumberIn is called when exiting the numberIn production.
	ExitNumberIn(c *NumberInContext)

	// ExitStringGreaterThan is called when exiting the stringGreaterThan production.
	ExitStringGreaterThan(c *StringGreaterThanContext)

	// ExitDateIn is called when exiting the dateIn production.
	ExitDateIn(c *DateInContext)

	// ExitBoolParenthesis is called when exiting the boolParenthesis production.
	ExitBoolParenthesis(c *BoolParenthesisContext)

	// ExitNumberNotIn is called when exiting the numberNotIn production.
	ExitNumberNotIn(c *NumberNotInContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitStringNotIn is called when exiting the stringNotIn production.
	ExitStringNotIn(c *StringNotInContext)

	// ExitBoolField is called when exiting the boolField production.
	ExitBoolField(c *BoolFieldContext)

	// ExitDateGreaterEqual is called when exiting the dateGreaterEqual production.
	ExitDateGreaterEqual(c *DateGreaterEqualContext)

	// ExitDateLessThan is called when exiting the dateLessThan production.
	ExitDateLessThan(c *DateLessThanContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitDateNotEqual is called when exiting the dateNotEqual production.
	ExitDateNotEqual(c *DateNotEqualContext)

	// ExitDateEqual is called when exiting the dateEqual production.
	ExitDateEqual(c *DateEqualContext)

	// ExitStringLessEqual is called when exiting the stringLessEqual production.
	ExitStringLessEqual(c *StringLessEqualContext)

	// ExitNumberEqual is called when exiting the numberEqual production.
	ExitNumberEqual(c *NumberEqualContext)

	// ExitStringLessThan is called when exiting the stringLessThan production.
	ExitStringLessThan(c *StringLessThanContext)

	// ExitNumberLessEqual is called when exiting the numberLessEqual production.
	ExitNumberLessEqual(c *NumberLessEqualContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitNumberGreaterThan is called when exiting the numberGreaterThan production.
	ExitNumberGreaterThan(c *NumberGreaterThanContext)

	// ExitDateLessEqual is called when exiting the dateLessEqual production.
	ExitDateLessEqual(c *DateLessEqualContext)

	// ExitStringEqual is called when exiting the stringEqual production.
	ExitStringEqual(c *StringEqualContext)

	// ExitStringIn is called when exiting the stringIn production.
	ExitStringIn(c *StringInContext)

	// ExitStringGreaterEqual is called when exiting the stringGreaterEqual production.
	ExitStringGreaterEqual(c *StringGreaterEqualContext)

	// ExitDateGreaterThan is called when exiting the dateGreaterThan production.
	ExitDateGreaterThan(c *DateGreaterThanContext)

	// ExitStringNotEqual is called when exiting the stringNotEqual production.
	ExitStringNotEqual(c *StringNotEqualContext)

	// ExitNumberGreaterEqual is called when exiting the numberGreaterEqual production.
	ExitNumberGreaterEqual(c *NumberGreaterEqualContext)

	// ExitDateNotIn is called when exiting the dateNotIn production.
	ExitDateNotIn(c *DateNotInContext)

	// ExitNumberLessThan is called when exiting the numberLessThan production.
	ExitNumberLessThan(c *NumberLessThanContext)

	// ExitNumberNotEqual is called when exiting the numberNotEqual production.
	ExitNumberNotEqual(c *NumberNotEqualContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)
}
