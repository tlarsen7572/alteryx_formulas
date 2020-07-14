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

// EnterBoolLiteral is called when production boolLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production boolLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterMinFunc is called when production minFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterMinFunc(ctx *MinFuncContext) {}

// ExitMinFunc is called when production minFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitMinFunc(ctx *MinFuncContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseAlteryxFormulasListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseAlteryxFormulasListener) ExitAnd(ctx *AndContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseAlteryxFormulasListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseAlteryxFormulasListener) ExitLessThan(ctx *LessThanContext) {}

// EnterDivide is called when production divide is entered.
func (s *BaseAlteryxFormulasListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production divide is exited.
func (s *BaseAlteryxFormulasListener) ExitDivide(ctx *DivideContext) {}

// EnterNotIn is called when production notIn is entered.
func (s *BaseAlteryxFormulasListener) EnterNotIn(ctx *NotInContext) {}

// ExitNotIn is called when production notIn is exited.
func (s *BaseAlteryxFormulasListener) ExitNotIn(ctx *NotInContext) {}

// EnterExprIf is called when production exprIf is entered.
func (s *BaseAlteryxFormulasListener) EnterExprIf(ctx *ExprIfContext) {}

// ExitExprIf is called when production exprIf is exited.
func (s *BaseAlteryxFormulasListener) ExitExprIf(ctx *ExprIfContext) {}

// EnterMultiply is called when production multiply is entered.
func (s *BaseAlteryxFormulasListener) EnterMultiply(ctx *MultiplyContext) {}

// ExitMultiply is called when production multiply is exited.
func (s *BaseAlteryxFormulasListener) ExitMultiply(ctx *MultiplyContext) {}

// EnterExprField is called when production exprField is entered.
func (s *BaseAlteryxFormulasListener) EnterExprField(ctx *ExprFieldContext) {}

// ExitExprField is called when production exprField is exited.
func (s *BaseAlteryxFormulasListener) ExitExprField(ctx *ExprFieldContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseAlteryxFormulasListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseAlteryxFormulasListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterAdd is called when production add is entered.
func (s *BaseAlteryxFormulasListener) EnterAdd(ctx *AddContext) {}

// ExitAdd is called when production add is exited.
func (s *BaseAlteryxFormulasListener) ExitAdd(ctx *AddContext) {}

// EnterOr is called when production or is entered.
func (s *BaseAlteryxFormulasListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseAlteryxFormulasListener) ExitOr(ctx *OrContext) {}

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

// EnterPowFunc is called when production powFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterPowFunc(ctx *PowFuncContext) {}

// ExitPowFunc is called when production powFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitPowFunc(ctx *PowFuncContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseAlteryxFormulasListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseAlteryxFormulasListener) ExitEqual(ctx *EqualContext) {}

// EnterNullFunc is called when production nullFunc is entered.
func (s *BaseAlteryxFormulasListener) EnterNullFunc(ctx *NullFuncContext) {}

// ExitNullFunc is called when production nullFunc is exited.
func (s *BaseAlteryxFormulasListener) ExitNullFunc(ctx *NullFuncContext) {}

// EnterDatetimeLiteral is called when production datetimeLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterDatetimeLiteral(ctx *DatetimeLiteralContext) {}

// ExitDatetimeLiteral is called when production datetimeLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitDatetimeLiteral(ctx *DatetimeLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterGreaterEqual is called when production greaterEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterGreaterEqual(ctx *GreaterEqualContext) {}

// ExitGreaterEqual is called when production greaterEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitGreaterEqual(ctx *GreaterEqualContext) {}

// EnterLessEqual is called when production lessEqual is entered.
func (s *BaseAlteryxFormulasListener) EnterLessEqual(ctx *LessEqualContext) {}

// ExitLessEqual is called when production lessEqual is exited.
func (s *BaseAlteryxFormulasListener) ExitLessEqual(ctx *LessEqualContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseAlteryxFormulasListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseAlteryxFormulasListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}
