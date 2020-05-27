import unittest
from alteryx_formulas.visitor import calculate


class AlteryxGrammarTests(unittest.TestCase):
    def test_addition(self):
        result = calculate('1+2')
        self.assertEqual(3, result)

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

    def test_number_greater_equal(self):
        result = calculate('3>=2')
        self.assertEqual(True, result)
        result = calculate('2>=2')
        self.assertEqual(True, result)
        result = calculate('1>=2')
        self.assertEqual(False, result)

    def test_number_less_than(self):
        result = calculate('2<3')
        self.assertEqual(True, result)
        result = calculate('3<3')
        self.assertEqual(False, result)

    def test_number_less_equal(self):
        result = calculate('2<=2')
        self.assertEqual(True, result)
        result = calculate('1<=2')
        self.assertEqual(True, result)
        result = calculate('3<=2')
        self.assertEqual(False, result)

    def test_number_equal(self):
        result = calculate('2=2')
        self.assertEqual(True, result)
        result = calculate('1=2')
        self.assertEqual(False, result)
        result = calculate('3=2')
        self.assertEqual(False, result)

    def test_number_not_equal(self):
        result = calculate('2!=2')
        self.assertEqual(False, result)
        result = calculate('1!=2')
        self.assertEqual(True, result)
        result = calculate('3!=2')
        self.assertEqual(True, result)

    def test_number_in(self):
        result = calculate('2 in (1,2,3)')
        self.assertEqual(True, result)
        result = calculate('2 in (3,4,5)')
        self.assertEqual(False, result)

    def test_number_not_in(self):
        result = calculate('2 not in (1,2,3)')
        self.assertEqual(False, result)
        result = calculate('2 not in (3,4,5)')
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

    def test_else(self):
        result = calculate('if false then 1 else 2 endif')
        self.assertEqual(2, result)

    def test_elseifthen(self):
        result = calculate('if false then 1 elseif true then 3 else 2 endif')
        self.assertEqual(3, result)

    def test_elseifelse(self):
        result = calculate('if false then 1 elseif false then 3 else 2 endif')
        self.assertEqual(2, result)


if __name__ == '__main__':
    unittest.main()
