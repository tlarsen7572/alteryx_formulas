package alteryx_formulas

import "unicode/utf8"

func (calc *calculator) charFromInt() {
	charCode := rune(calc.popValue().(float64))
	if !utf8.ValidRune(charCode) {
		calc.pushValue(nil)
		return
	}
	value := string(charCode)
	calc.pushValue(value)
}

func (calc *calculator) charToInt() {
	value := calc.popValue().(string)
	if len(value) == 0 {
		calc.pushValue(nil)
		return
	}
	runes := []rune(value)
	code := float64(runes[0])
	calc.pushValue(code)
}

func (calc *calculator) concatenate() {
	concatenate := func(left interface{}, right interface{}) interface{} {
		return left.(string) + right.(string)
	}
	calc.ifNonNullLeftRight(concatenate)
}
