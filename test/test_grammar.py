import unittest
from alteryx_formulas.visitor import calculate


class AlteryxGrammarTests(unittest.TestCase):
    def test_addition(self):
        result = calculate('1+2')
        self.assertEqual(3, result)


if __name__ == '__main__':
    unittest.main()
