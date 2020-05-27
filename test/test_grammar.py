import unittest
from datetime import datetime

from alteryx_formulas.visitor import calculate


class AlteryxGrammarTests(unittest.TestCase):
    def test_addition(self):
        result = calculate('1+2')
        self.assertEqual(3, result)

    def test_concatenation(self):
        result = calculate("'A'+'B'")
        self.assertEqual('AB', result)

    def test_subtraction(self):
        result = calculate('4-3')
        self.assertEqual(1, result)

    def test_multiplication(self):
        result = calculate('4*6')
        self.assertEqual(24, result)

    def test_division(self):
        result = calculate('21/7')
        self.assertEqual(3, result)

    def test_number_greater_than(self):
        result = calculate('3>2')
        self.assertEqual(True, result)
        result = calculate('2>2')
        self.assertEqual(False, result)

    def test_string_greater_than(self):
        result = calculate('"B">"A"')
        self.assertEqual(True, result)
        result = calculate('"A">"A"')
        self.assertEqual(False, result)

    def test_date_greater_than(self):
        result = calculate("`2020-02-02`>`2020-02-01`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-01`>`2020-01-01`")
        self.assertEqual(False, result)

    def test_number_greater_equal(self):
        result = calculate('3>=2')
        self.assertEqual(True, result)
        result = calculate('2>=2')
        self.assertEqual(True, result)
        result = calculate('1>=2')
        self.assertEqual(False, result)

    def test_string_greater_equal(self):
        result = calculate("'C'>='B'")
        self.assertEqual(True, result)
        result = calculate("'B'>='B'")
        self.assertEqual(True, result)
        result = calculate("'A'>='B'")
        self.assertEqual(False, result)

    def test_date_greater_equal(self):
        result = calculate("`2020-01-03`>=`2020-01-02`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-02`>=`2020-01-02`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-01`>=`2020-01-02`")
        self.assertEqual(False, result)

    def test_number_less_than(self):
        result = calculate('2<3')
        self.assertEqual(True, result)
        result = calculate('3<3')
        self.assertEqual(False, result)

    def test_string_less_than(self):
        result = calculate("'B'<'C'")
        self.assertEqual(True, result)
        result = calculate("'C'<'C'")
        self.assertEqual(False, result)

    def test_date_less_than(self):
        result = calculate("`2020-01-02`<`2020-01-03`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-03`<`2020-01-03`")
        self.assertEqual(False, result)

    def test_number_less_equal(self):
        result = calculate('2<=2')
        self.assertEqual(True, result)
        result = calculate('1<=2')
        self.assertEqual(True, result)
        result = calculate('3<=2')
        self.assertEqual(False, result)

    def test_string_less_equal(self):
        result = calculate("'B'<='B'")
        self.assertEqual(True, result)
        result = calculate("'A'<='B'")
        self.assertEqual(True, result)
        result = calculate("'C'<='B'")
        self.assertEqual(False, result)

    def test_date_less_equal(self):
        result = calculate("`2020-01-02`<=`2020-01-02`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-01`<=`2020-01-02`")
        self.assertEqual(True, result)
        result = calculate("`2020-01-03`<=`2020-01-02`")
        self.assertEqual(False, result)

    def test_number_equal(self):
        result = calculate('2=2')
        self.assertEqual(True, result)
        result = calculate('1=2')
        self.assertEqual(False, result)
        result = calculate('3=2')
        self.assertEqual(False, result)

    def test_string_equal(self):
        result = calculate("'B'='B'")
        self.assertEqual(True, result)
        result = calculate("'A'='B'")
        self.assertEqual(False, result)
        result = calculate("'C'='B'")
        self.assertEqual(False, result)

    def test_date_equal(self):
        result = calculate('`2020-01-02`=`2020-01-02`')
        self.assertEqual(True, result)
        result = calculate('`2020-01-01`=`2020-01-02`')
        self.assertEqual(False, result)
        result = calculate('`2020-01-03`=`2020-01-02`')
        self.assertEqual(False, result)

    def test_number_not_equal(self):
        result = calculate('2!=2')
        self.assertEqual(False, result)
        result = calculate('1!=2')
        self.assertEqual(True, result)
        result = calculate('3!=2')
        self.assertEqual(True, result)

    def test_string_not_equal(self):
        result = calculate("'b'!='b'")
        self.assertEqual(False, result)
        result = calculate("'a'!='b'")
        self.assertEqual(True, result)
        result = calculate("'c'!='b'")
        self.assertEqual(True, result)

    def test_date_not_equal(self):
        result = calculate('`2020-01-02`!=`2020-01-02`')
        self.assertEqual(False, result)
        result = calculate('`2020-01-01`!=`2020-01-02`')
        self.assertEqual(True, result)
        result = calculate('`2020-01-03`!=`2020-01-02`')
        self.assertEqual(True, result)

    def test_number_in(self):
        result = calculate('2 in (1,2,3)')
        self.assertEqual(True, result)
        result = calculate('2 in (3,4,5)')
        self.assertEqual(False, result)

    def test_string_in(self):
        result = calculate("'b' in ('a','b','c')")
        self.assertEqual(True, result)
        result = calculate("'b' in ('c','d','e')")
        self.assertEqual(False, result)

    def test_date_in(self):
        result = calculate('`2020-01-02` in (`2020-01-01`,`2020-01-02`,`2020-01-03`)')
        self.assertEqual(True, result)
        result = calculate('`2020-01-02` in (`2020-01-03`,`2020-01-04`,`2020-01-05`)')
        self.assertEqual(False, result)

    def test_number_not_in(self):
        result = calculate('2 not in (1,2,3)')
        self.assertEqual(False, result)
        result = calculate('2 not in (3,4,5)')
        self.assertEqual(True, result)

    def test_string_not_in(self):
        result = calculate("'b' not in ('a','b','c')")
        self.assertEqual(False, result)
        result = calculate("'b' not in ('c','d','e')")
        self.assertEqual(True, result)

    def test_date_not_in(self):
        result = calculate('`2020-01-02` not in (`2020-01-01`,`2020-01-02`,`2020-01-03`)')
        self.assertEqual(False, result)
        result = calculate('`2020-01-02` not in (`2020-01-03`,`2020-01-04`,`2020-01-05`)')
        self.assertEqual(True, result)

    def test_and(self):
        result = calculate('true and true')
        self.assertTrue(result)
        result = calculate('false and false')
        self.assertFalse(result)
        result = calculate('true and false')
        self.assertFalse(result)

    def test_or(self):
        result = calculate('true or true')
        self.assertTrue(result)
        result = calculate('false or false')
        self.assertFalse(result)
        result = calculate('true or false')
        self.assertTrue(result)

    def test_then(self):
        result = calculate('if true then 1 else 2 endif')
        self.assertEqual(1, result)
        result = calculate("if true then 'a' else 'b' endif")
        self.assertEqual('a', result)
        result = calculate('if true then `2020-01-01` else `2020-01-02` endif')
        self.assertEqual(datetime(2020, 1, 1), result)

    def test_else(self):
        result = calculate('if false then 1 else 2 endif')
        self.assertEqual(2, result)
        result = calculate("if false then 'a' else 'b' endif")
        self.assertEqual('b', result)
        result = calculate('if false then `2020-01-01` else `2020-01-02` endif')
        self.assertEqual(datetime(2020, 1, 2), result)

    def test_elseifthen(self):
        result = calculate('if false then 1 elseif true then 3 else 2 endif')
        self.assertEqual(3, result)
        result = calculate("if false then 'A' elseif true then 'C' else 'B' endif")
        self.assertEqual('C', result)
        result = calculate('if false then `2020-01-01` elseif true then `2020-01-03` else `2020-01-02` endif')
        self.assertEqual(datetime(2020, 1, 3), result)

    def test_elseifelse(self):
        result = calculate('if false then 1 elseif false then 3 else 2 endif')
        self.assertEqual(2, result)
        result = calculate("if false then 'a' elseif false then 'c' else 'b' endif")
        self.assertEqual('b', result)
        result = calculate('if false then `2020-01-01` elseif false then `2020-01-03` else `2020-01-02` endif')
        self.assertEqual(datetime(2020, 1, 2), result)


if __name__ == '__main__':
    unittest.main()
