import datetime

from antlr4 import *
from antlr4.error.ErrorListener import ErrorListener
from alteryx_formulas.generated_grammar.AlteryxFormulasLexer import AlteryxFormulasLexer
from alteryx_formulas.generated_grammar.AlteryxFormulasParser import AlteryxFormulasParser
from typing import Callable, Dict


class MissingFieldException(Exception):
    def __init__(self, missing_field: str):
        self.MissingField = missing_field

    def __str__(self):
        return "Missing field {field}".format(field=self.MissingField)


def calculate(expression: str, fields=None):
    if fields is None:
        fields = {}
    visitor = AlteryxFormulaVisitor(expression, fields)
    return visitor.calculate()


class AlteryxFormulasErrorListener(ErrorListener):

    def __init__(self):
        super(AlteryxFormulasErrorListener, self).__init__()

    def syntaxError(self, recognizer, offending_symbol, line, column, msg, e):
        raise Exception("{line}:{col} {msg}".format(line=line, col=column, msg=msg))


class AlteryxFormulaVisitor(ParseTreeVisitor):
    def __init__(self, expression: str, fields: Dict[str, Callable]):
        self.Expression = expression
        self.Fields = fields
        lexer = AlteryxFormulasLexer(InputStream(expression))
        lexer.addErrorListener(AlteryxFormulasErrorListener())
        stream = CommonTokenStream(lexer)
        parser = AlteryxFormulasParser(stream)
        parser.addErrorListener(AlteryxFormulasErrorListener())
        self._tree = parser.formula()

    def _left_right_check(self, ctx, operator):
        left = self.visit(ctx.left)
        right = self.visit(ctx.right)
        return operator(left, right)

    def calculate(self):
        return self.visit(self._tree)

    def visitFormula(self, ctx: AlteryxFormulasParser.FormulaContext):
        return self.visitChildren(ctx)

    def visitAdd(self, ctx: AlteryxFormulasParser.AddContext):
        return self._left_right_check(ctx, lambda l, r: l + r)

    def visitOr(self, ctx: AlteryxFormulasParser.OrContext):
        return self._left_right_check(ctx, lambda l, r: l or r)

    def visitStringIn(self, ctx: AlteryxFormulasParser.StringInContext):
        return self._visit_in(ctx.stringExpr)

    def visitNumberIn(self, ctx: AlteryxFormulasParser.StringInContext):
        return self._visit_in(ctx.numberExpr)

    def visitDateIn(self, ctx: AlteryxFormulasParser.StringInContext):
        return self._visit_in(ctx.dateExpr)

    def _visit_in(self, expr):
        value = self.visit(expr(0))
        compare_to = 1
        while compare_to < len(expr()):
            if value == self.visit(expr(compare_to)):
                return True
            compare_to += 1
        return False

    def visitSubtract(self, ctx: AlteryxFormulasParser.SubtractContext):
        return self._left_right_check(ctx, lambda l, r: l - r)

    def visitStringNotEqual(self, ctx: AlteryxFormulasParser.StringNotEqualContext):
        return self._left_right_check(ctx, lambda l, r: l != r)

    def visitNumberNotEqual(self, ctx: AlteryxFormulasParser.NumberNotEqualContext):
        return self._left_right_check(ctx, lambda l, r: l != r)

    def visitDateNotEqual(self, ctx: AlteryxFormulasParser.DateNotEqualContext):
        return self._left_right_check(ctx, lambda l, r: l != r)

    def visitInteger(self, ctx: AlteryxFormulasParser.IntegerContext):
        return int(ctx.getText())

    def visitBoolParenthesis(self, ctx: AlteryxFormulasParser.BoolParenthesisContext):
        return self.visit(ctx.boolExpr())

    def visitDateParenthesis(self, ctx: AlteryxFormulasParser.DateParenthesisContext):
        return self.visit(ctx.dateExpr())

    def visitNumberParenthesis(self, ctx: AlteryxFormulasParser.DateParenthesisContext):
        return self.visit(ctx.numberExpr())

    def visitStringParenthesis(self, ctx: AlteryxFormulasParser.NumberParenthesisContext):
        return self.visit(ctx.stringExpr())

    def visitStringEqual(self, ctx: AlteryxFormulasParser.StringEqualContext):
        return self._left_right_check(ctx, lambda l, r: l == r)

    def visitNumberEqual(self, ctx: AlteryxFormulasParser.NumberEqualContext):
        return self._left_right_check(ctx, lambda l, r: l == r)

    def visitDateEqual(self, ctx: AlteryxFormulasParser.DateEqualContext):
        return self._left_right_check(ctx, lambda l, r: l == r)

    def visitStringField(self, ctx: AlteryxFormulasParser.StringFieldContext):
        return self._visit_field(ctx.getText())

    def visitNumberField(self, ctx: AlteryxFormulasParser.NumberFieldContext):
        return self._visit_field(ctx.getText())

    def visitDateField(self, ctx: AlteryxFormulasParser.DateFieldContext):
        return self._visit_field(ctx.getText())

    def visitBoolField(self, ctx: AlteryxFormulasParser.BoolFieldContext):
        return self._visit_field(ctx.getText())

    def _visit_field(self, field_name):
        field_name = field_name[1:-1]
        value_getter = self.Fields.get(field_name)
        if value_getter is None:
            raise MissingFieldException(missing_field=field_name)
        return value_getter()

    def visitStringLiteral(self, ctx: AlteryxFormulasParser.StringLiteralContext):
        return ctx.getText()[1:-1]

    def visitAnd(self, ctx: AlteryxFormulasParser.AndContext):
        return self._left_right_check(ctx, lambda l, r: l and r)

    def visitNumberElseIf(self, ctx: AlteryxFormulasParser.NumberElseIfContext):
        return self._visit_else_if(ctx.boolExpr, ctx.numberExpr)

    def _visit_else_if(self, bool_expr, then_else_expr):
        condition = self.visit(bool_expr(0))
        if condition:
            return self.visit(then_else_expr(0))

        elseifs = len(bool_expr())-1
        elseif = 0
        while elseif < elseifs:
            start_index = (elseif + 1)
            condition = self.visit(bool_expr(start_index))
            if condition:
                return self.visit(then_else_expr(start_index))
            elseif += 1

        else_expr = len(then_else_expr())-1
        return self.visit(then_else_expr(else_expr))

    def visitStringLessThan(self, ctx: AlteryxFormulasParser.StringLessThanContext):
        return self._left_right_check(ctx, lambda l, r: l < r)

    def visitNumberLessThan(self, ctx: AlteryxFormulasParser.NumberLessThanContext):
        return self._left_right_check(ctx, lambda l, r: l < r)

    def visitDateLessThan(self, ctx: AlteryxFormulasParser.DateLessThanContext):
        return self._left_right_check(ctx, lambda l, r: l < r)

    def visitDateLiteral(self, ctx: AlteryxFormulasParser.DateLiteralContext):
        return datetime.datetime.strptime(ctx.getText()[1:-1], "%Y-%m-%d")

    def visitDatetimeLiteral(self, ctx: AlteryxFormulasParser.DatetimeLiteralContext):
        return datetime.datetime.strptime(ctx.getText()[1:-1], "%Y-%m-%d %H:%M:%S")

    def visitDivide(self, ctx: AlteryxFormulasParser.DivideContext):
        return self._left_right_check(ctx, lambda l, r: l / r)

    def visitStringGreaterEqual(self, ctx: AlteryxFormulasParser.StringGreaterEqualContext):
        return self._left_right_check(ctx, lambda l, r: l >= r)

    def visitNumberGreaterEqual(self, ctx: AlteryxFormulasParser.NumberGreaterEqualContext):
        return self._left_right_check(ctx, lambda l, r: l >= r)

    def visitDateGreaterEqual(self, ctx: AlteryxFormulasParser.DateGreaterEqualContext):
        return self._left_right_check(ctx, lambda l, r: l >= r)

    def visitStringNotIn(self, ctx: AlteryxFormulasParser.StringNotInContext):
        return self._visit_not_in(ctx.stringExpr)

    def visitNumberNotIn(self, ctx: AlteryxFormulasParser.NumberNotInContext):
        return self._visit_not_in(ctx.numberExpr)

    def visitDateNotIn(self, ctx: AlteryxFormulasParser.DateNotInContext):
        return self._visit_not_in(ctx.dateExpr)

    def _visit_not_in(self, expr):
        value = self.visit(expr(0))
        compare_to = 1
        while compare_to < len(expr()):
            if value == self.visit(expr(compare_to)):
                return False
            compare_to += 1
        return True

    def visitStringLessEqual(self, ctx: AlteryxFormulasParser.StringLessEqualContext):
        return self._left_right_check(ctx, lambda l, r: l <= r)

    def visitNumberLessEqual(self, ctx: AlteryxFormulasParser.NumberLessEqualContext):
        return self._left_right_check(ctx, lambda l, r: l <= r)

    def visitDateLessEqual(self, ctx: AlteryxFormulasParser.DateLessEqualContext):
        return self._left_right_check(ctx, lambda l, r: l <= r)

    def visitDecimal(self, ctx: AlteryxFormulasParser.DecimalContext):
        return float(ctx.getText())

    def visitMultiply(self, ctx: AlteryxFormulasParser.MultiplyContext):
        return self._left_right_check(ctx, lambda l, r: l * r)

    def visitNumberIf(self, ctx: AlteryxFormulasParser.NumberIfContext):
        return self._visit_if(ctx.boolExpr, ctx.numberExpr)

    def _visit_if(self, bool_expr, then_else_expr):
        condition = self.visit(bool_expr())
        if condition:
            return self.visit(then_else_expr(0))
        else:
            return self.visit(then_else_expr(1))

    def visitStringGreaterThan(self, ctx: AlteryxFormulasParser.StringGreaterThanContext):
        return self._left_right_check(ctx, lambda l, r: l > r)

    def visitNumberGreaterThan(self, ctx: AlteryxFormulasParser.NumberGreaterThanContext):
        return self._left_right_check(ctx, lambda l, r: l > r)

    def visitDateGreaterThan(self, ctx: AlteryxFormulasParser.DateGreaterThanContext):
        return self._left_right_check(ctx, lambda l, r: l > r)

    def visitPow(self, ctx: AlteryxFormulasParser.PowContext):
        return pow(self.visit(ctx.expr(0)), self.visit(ctx.expr(1)))

    def visitNumberMin(self, ctx: AlteryxFormulasParser.NumberMinContext):
        min_value = self.visit(ctx.expr(0))
        index = 1
        while index < len(ctx.expr()):
            compare_to = self.visit(ctx.expr(index))
            if compare_to < min_value:
                min_value = compare_to
            index += 1
        return min_value

    def visitNumberMax(self, ctx: AlteryxFormulasParser.NumberMaxContext):
        max_value = self.visit(ctx.expr(0))
        index = 1
        while index < len(ctx.expr()):
            compare_to = self.visit(ctx.expr(index))
            if compare_to > max_value:
                max_value = compare_to
            index += 1
        return max_value

    def visitBoolLiteral(self, ctx: AlteryxFormulasParser.BoolLiteralContext):
        return ctx.getText().lower() == 'true'
