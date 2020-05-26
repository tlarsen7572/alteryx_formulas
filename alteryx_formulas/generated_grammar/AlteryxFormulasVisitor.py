# Generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8
from antlr4 import *
if __name__ is not None and "." in __name__:
    from .AlteryxFormulasParser import AlteryxFormulasParser
else:
    from AlteryxFormulasParser import AlteryxFormulasParser

# This class defines a complete generic visitor for a parse tree produced by AlteryxFormulasParser.

class AlteryxFormulasVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by AlteryxFormulasParser#formula.
    def visitFormula(self, ctx:AlteryxFormulasParser.FormulaContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#add.
    def visitAdd(self, ctx:AlteryxFormulasParser.AddContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#or.
    def visitOr(self, ctx:AlteryxFormulasParser.OrContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#in.
    def visitIn(self, ctx:AlteryxFormulasParser.InContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#subtract.
    def visitSubtract(self, ctx:AlteryxFormulasParser.SubtractContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#notEqual.
    def visitNotEqual(self, ctx:AlteryxFormulasParser.NotEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#integer.
    def visitInteger(self, ctx:AlteryxFormulasParser.IntegerContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#parenthesis.
    def visitParenthesis(self, ctx:AlteryxFormulasParser.ParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#equal.
    def visitEqual(self, ctx:AlteryxFormulasParser.EqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#datetimeLiteral.
    def visitDatetimeLiteral(self, ctx:AlteryxFormulasParser.DatetimeLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#func.
    def visitFunc(self, ctx:AlteryxFormulasParser.FuncContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#field.
    def visitField(self, ctx:AlteryxFormulasParser.FieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringLiteral.
    def visitStringLiteral(self, ctx:AlteryxFormulasParser.StringLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#and.
    def visitAnd(self, ctx:AlteryxFormulasParser.AndContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#elseIf.
    def visitElseIf(self, ctx:AlteryxFormulasParser.ElseIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#lessThan.
    def visitLessThan(self, ctx:AlteryxFormulasParser.LessThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateLiteral.
    def visitDateLiteral(self, ctx:AlteryxFormulasParser.DateLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#divide.
    def visitDivide(self, ctx:AlteryxFormulasParser.DivideContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#greaterEqual.
    def visitGreaterEqual(self, ctx:AlteryxFormulasParser.GreaterEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#notIn.
    def visitNotIn(self, ctx:AlteryxFormulasParser.NotInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#lessEqual.
    def visitLessEqual(self, ctx:AlteryxFormulasParser.LessEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#decimal.
    def visitDecimal(self, ctx:AlteryxFormulasParser.DecimalContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#multiply.
    def visitMultiply(self, ctx:AlteryxFormulasParser.MultiplyContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#if.
    def visitIf(self, ctx:AlteryxFormulasParser.IfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#greaterThan.
    def visitGreaterThan(self, ctx:AlteryxFormulasParser.GreaterThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#string.
    def visitString(self, ctx:AlteryxFormulasParser.StringContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#pow.
    def visitPow(self, ctx:AlteryxFormulasParser.PowContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#min.
    def visitMin(self, ctx:AlteryxFormulasParser.MinContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#max.
    def visitMax(self, ctx:AlteryxFormulasParser.MaxContext):
        return self.visitChildren(ctx)



del AlteryxFormulasParser