// Code generated from C:/repositories/alteryx_formulas/grammar\AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type AlteryxFormulasListener interface {
	antlr.ParseTreeListener

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterExprFunction is called when entering the exprFunction production.
	EnterExprFunction(c *ExprFunctionContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterIn is called when entering the in production.
	EnterIn(c *InContext)

	// EnterSubtract is called when entering the subtract production.
	EnterSubtract(c *SubtractContext)

	// EnterNotEqual is called when entering the notEqual production.
	EnterNotEqual(c *NotEqualContext)

	// EnterParenthesis is called when entering the parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterDatetimeLiteral is called when entering the datetimeLiteral production.
	EnterDatetimeLiteral(c *DatetimeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterExprElseIf is called when entering the exprElseIf production.
	EnterExprElseIf(c *ExprElseIfContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterGreaterEqual is called when entering the greaterEqual production.
	EnterGreaterEqual(c *GreaterEqualContext)

	// EnterNotIn is called when entering the notIn production.
	EnterNotIn(c *NotInContext)

	// EnterExprIf is called when entering the exprIf production.
	EnterExprIf(c *ExprIfContext)

	// EnterLessEqual is called when entering the lessEqual production.
	EnterLessEqual(c *LessEqualContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterExprField is called when entering the exprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterPowFunc is called when entering the powFunc production.
	EnterPowFunc(c *PowFuncContext)

	// EnterMinFunc is called when entering the minFunc production.
	EnterMinFunc(c *MinFuncContext)

	// EnterMaxFunc is called when entering the maxFunc production.
	EnterMaxFunc(c *MaxFuncContext)

	// EnterNullFunc is called when entering the nullFunc production.
	EnterNullFunc(c *NullFuncContext)

	// EnterStr is called when entering the str production.
	EnterStr(c *StrContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitExprFunction is called when exiting the exprFunction production.
	ExitExprFunction(c *ExprFunctionContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitIn is called when exiting the in production.
	ExitIn(c *InContext)

	// ExitSubtract is called when exiting the subtract production.
	ExitSubtract(c *SubtractContext)

	// ExitNotEqual is called when exiting the notEqual production.
	ExitNotEqual(c *NotEqualContext)

	// ExitParenthesis is called when exiting the parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitDatetimeLiteral is called when exiting the datetimeLiteral production.
	ExitDatetimeLiteral(c *DatetimeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitExprElseIf is called when exiting the exprElseIf production.
	ExitExprElseIf(c *ExprElseIfContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitGreaterEqual is called when exiting the greaterEqual production.
	ExitGreaterEqual(c *GreaterEqualContext)

	// ExitNotIn is called when exiting the notIn production.
	ExitNotIn(c *NotInContext)

	// ExitExprIf is called when exiting the exprIf production.
	ExitExprIf(c *ExprIfContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitExprField is called when exiting the exprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitPowFunc is called when exiting the powFunc production.
	ExitPowFunc(c *PowFuncContext)

	// ExitMinFunc is called when exiting the minFunc production.
	ExitMinFunc(c *MinFuncContext)

	// ExitMaxFunc is called when exiting the maxFunc production.
	ExitMaxFunc(c *MaxFuncContext)

	// ExitNullFunc is called when exiting the nullFunc production.
	ExitNullFunc(c *NullFuncContext)

	// ExitStr is called when exiting the str production.
	ExitStr(c *StrContext)
}
