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


    # Visit a parse tree produced by AlteryxFormulasParser#fieldParenthesis.
    def visitFieldParenthesis(self, ctx:AlteryxFormulasParser.FieldParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#anyField.
    def visitAnyField(self, ctx:AlteryxFormulasParser.AnyFieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringIf.
    def visitStringIf(self, ctx:AlteryxFormulasParser.StringIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringFunc.
    def visitStringFunc(self, ctx:AlteryxFormulasParser.StringFuncContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#concatenate.
    def visitConcatenate(self, ctx:AlteryxFormulasParser.ConcatenateContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringLiteral.
    def visitStringLiteral(self, ctx:AlteryxFormulasParser.StringLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringParenthesis.
    def visitStringParenthesis(self, ctx:AlteryxFormulasParser.StringParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringElseIf.
    def visitStringElseIf(self, ctx:AlteryxFormulasParser.StringElseIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringField.
    def visitStringField(self, ctx:AlteryxFormulasParser.StringFieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringMin.
    def visitStringMin(self, ctx:AlteryxFormulasParser.StringMinContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringMax.
    def visitStringMax(self, ctx:AlteryxFormulasParser.StringMaxContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberParenthesis.
    def visitNumberParenthesis(self, ctx:AlteryxFormulasParser.NumberParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#add.
    def visitAdd(self, ctx:AlteryxFormulasParser.AddContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberFunc.
    def visitNumberFunc(self, ctx:AlteryxFormulasParser.NumberFuncContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberField.
    def visitNumberField(self, ctx:AlteryxFormulasParser.NumberFieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#subtract.
    def visitSubtract(self, ctx:AlteryxFormulasParser.SubtractContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberIf.
    def visitNumberIf(self, ctx:AlteryxFormulasParser.NumberIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberElseIf.
    def visitNumberElseIf(self, ctx:AlteryxFormulasParser.NumberElseIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#divide.
    def visitDivide(self, ctx:AlteryxFormulasParser.DivideContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#integer.
    def visitInteger(self, ctx:AlteryxFormulasParser.IntegerContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#decimal.
    def visitDecimal(self, ctx:AlteryxFormulasParser.DecimalContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#multiply.
    def visitMultiply(self, ctx:AlteryxFormulasParser.MultiplyContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#pow.
    def visitPow(self, ctx:AlteryxFormulasParser.PowContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberMin.
    def visitNumberMin(self, ctx:AlteryxFormulasParser.NumberMinContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberMax.
    def visitNumberMax(self, ctx:AlteryxFormulasParser.NumberMaxContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateParenthesis.
    def visitDateParenthesis(self, ctx:AlteryxFormulasParser.DateParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateIf.
    def visitDateIf(self, ctx:AlteryxFormulasParser.DateIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateElseIf.
    def visitDateElseIf(self, ctx:AlteryxFormulasParser.DateElseIfContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateFunc.
    def visitDateFunc(self, ctx:AlteryxFormulasParser.DateFuncContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#datetimeLiteral.
    def visitDatetimeLiteral(self, ctx:AlteryxFormulasParser.DatetimeLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateLiteral.
    def visitDateLiteral(self, ctx:AlteryxFormulasParser.DateLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateField.
    def visitDateField(self, ctx:AlteryxFormulasParser.DateFieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateMin.
    def visitDateMin(self, ctx:AlteryxFormulasParser.DateMinContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateMax.
    def visitDateMax(self, ctx:AlteryxFormulasParser.DateMaxContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberIn.
    def visitNumberIn(self, ctx:AlteryxFormulasParser.NumberInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringGreaterThan.
    def visitStringGreaterThan(self, ctx:AlteryxFormulasParser.StringGreaterThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateIn.
    def visitDateIn(self, ctx:AlteryxFormulasParser.DateInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#boolParenthesis.
    def visitBoolParenthesis(self, ctx:AlteryxFormulasParser.BoolParenthesisContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberNotIn.
    def visitNumberNotIn(self, ctx:AlteryxFormulasParser.NumberNotInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#boolLiteral.
    def visitBoolLiteral(self, ctx:AlteryxFormulasParser.BoolLiteralContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringNotIn.
    def visitStringNotIn(self, ctx:AlteryxFormulasParser.StringNotInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#boolField.
    def visitBoolField(self, ctx:AlteryxFormulasParser.BoolFieldContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateGreaterEqual.
    def visitDateGreaterEqual(self, ctx:AlteryxFormulasParser.DateGreaterEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateLessThan.
    def visitDateLessThan(self, ctx:AlteryxFormulasParser.DateLessThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#and.
    def visitAnd(self, ctx:AlteryxFormulasParser.AndContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateNotEqual.
    def visitDateNotEqual(self, ctx:AlteryxFormulasParser.DateNotEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateEqual.
    def visitDateEqual(self, ctx:AlteryxFormulasParser.DateEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringLessEqual.
    def visitStringLessEqual(self, ctx:AlteryxFormulasParser.StringLessEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberEqual.
    def visitNumberEqual(self, ctx:AlteryxFormulasParser.NumberEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringLessThan.
    def visitStringLessThan(self, ctx:AlteryxFormulasParser.StringLessThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberLessEqual.
    def visitNumberLessEqual(self, ctx:AlteryxFormulasParser.NumberLessEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#or.
    def visitOr(self, ctx:AlteryxFormulasParser.OrContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberGreaterThan.
    def visitNumberGreaterThan(self, ctx:AlteryxFormulasParser.NumberGreaterThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateLessEqual.
    def visitDateLessEqual(self, ctx:AlteryxFormulasParser.DateLessEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringEqual.
    def visitStringEqual(self, ctx:AlteryxFormulasParser.StringEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringIn.
    def visitStringIn(self, ctx:AlteryxFormulasParser.StringInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringGreaterEqual.
    def visitStringGreaterEqual(self, ctx:AlteryxFormulasParser.StringGreaterEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateGreaterThan.
    def visitDateGreaterThan(self, ctx:AlteryxFormulasParser.DateGreaterThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#stringNotEqual.
    def visitStringNotEqual(self, ctx:AlteryxFormulasParser.StringNotEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberGreaterEqual.
    def visitNumberGreaterEqual(self, ctx:AlteryxFormulasParser.NumberGreaterEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#dateNotIn.
    def visitDateNotIn(self, ctx:AlteryxFormulasParser.DateNotInContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberLessThan.
    def visitNumberLessThan(self, ctx:AlteryxFormulasParser.NumberLessThanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#numberNotEqual.
    def visitNumberNotEqual(self, ctx:AlteryxFormulasParser.NumberNotEqualContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by AlteryxFormulasParser#string.
    def visitString(self, ctx:AlteryxFormulasParser.StringContext):
        return self.visitChildren(ctx)



del AlteryxFormulasParser