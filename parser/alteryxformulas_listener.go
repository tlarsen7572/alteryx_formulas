// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlteryxFormulasListener is a complete listener for a parse tree produced by AlteryxFormulasParser.
type AlteryxFormulasListener interface {
	antlr.ParseTreeListener

	// EnterMaxFunc is called when entering the maxFunc production.
	EnterMaxFunc(c *MaxFuncContext)

	// EnterBoolLiteral is called when entering the boolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterMinFunc is called when entering the minFunc production.
	EnterMinFunc(c *MinFuncContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterDivide is called when entering the divide production.
	EnterDivide(c *DivideContext)

	// EnterNotIn is called when entering the notIn production.
	EnterNotIn(c *NotInContext)

	// EnterExprIf is called when entering the exprIf production.
	EnterExprIf(c *ExprIfContext)

	// EnterMultiply is called when entering the multiply production.
	EnterMultiply(c *MultiplyContext)

	// EnterExprField is called when entering the exprField production.
	EnterExprField(c *ExprFieldContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

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

	// EnterPowFunc is called when entering the powFunc production.
	EnterPowFunc(c *PowFuncContext)

	// EnterIifFunc is called when entering the iifFunc production.
	EnterIifFunc(c *IifFuncContext)

	// EnterEqual is called when entering the equal production.
	EnterEqual(c *EqualContext)

	// EnterNullFunc is called when entering the nullFunc production.
	EnterNullFunc(c *NullFuncContext)

	// EnterDatetimeLiteral is called when entering the datetimeLiteral production.
	EnterDatetimeLiteral(c *DatetimeLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterDateLiteral is called when entering the dateLiteral production.
	EnterDateLiteral(c *DateLiteralContext)

	// EnterGreaterEqual is called when entering the greaterEqual production.
	EnterGreaterEqual(c *GreaterEqualContext)

	// EnterLessEqual is called when entering the lessEqual production.
	EnterLessEqual(c *LessEqualContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// ExitMaxFunc is called when exiting the maxFunc production.
	ExitMaxFunc(c *MaxFuncContext)

	// ExitBoolLiteral is called when exiting the boolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitMinFunc is called when exiting the minFunc production.
	ExitMinFunc(c *MinFuncContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitDivide is called when exiting the divide production.
	ExitDivide(c *DivideContext)

	// ExitNotIn is called when exiting the notIn production.
	ExitNotIn(c *NotInContext)

	// ExitExprIf is called when exiting the exprIf production.
	ExitExprIf(c *ExprIfContext)

	// ExitMultiply is called when exiting the multiply production.
	ExitMultiply(c *MultiplyContext)

	// ExitExprField is called when exiting the exprField production.
	ExitExprField(c *ExprFieldContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

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

	// ExitPowFunc is called when exiting the powFunc production.
	ExitPowFunc(c *PowFuncContext)

	// ExitIifFunc is called when exiting the iifFunc production.
	ExitIifFunc(c *IifFuncContext)

	// ExitEqual is called when exiting the equal production.
	ExitEqual(c *EqualContext)

	// ExitNullFunc is called when exiting the nullFunc production.
	ExitNullFunc(c *NullFuncContext)

	// ExitDatetimeLiteral is called when exiting the datetimeLiteral production.
	ExitDatetimeLiteral(c *DatetimeLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitDateLiteral is called when exiting the dateLiteral production.
	ExitDateLiteral(c *DateLiteralContext)

	// ExitGreaterEqual is called when exiting the greaterEqual production.
	ExitGreaterEqual(c *GreaterEqualContext)

	// ExitLessEqual is called when exiting the lessEqual production.
	ExitLessEqual(c *LessEqualContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)
}
