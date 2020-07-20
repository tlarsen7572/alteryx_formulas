package alteryx_formulas

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

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

func (calc *calculator) contains() {
	value := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	caseInsensitive := calc.popValue().(float64)
	if caseInsensitive != 0 {
		value = strings.ToLower(value)
		lookFor = strings.ToLower(lookFor)
	}
	contains := strings.Contains(value, lookFor)
	calc.pushValue(contains)
}

func (calc *calculator) endsWith() {
	value := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	caseInsensitive := calc.popValue().(float64)
	startAt := len(value) - len(lookFor)
	if startAt < 0 {
		calc.pushValue(false)
		return
	}
	if caseInsensitive != 0 {
		value = strings.ToLower(value)
		lookFor = strings.ToLower(lookFor)
	}

	contains := value[startAt:] == lookFor
	calc.pushValue(contains)
}

func (calc *calculator) concatenate() {
	concatenate := func(left interface{}, right interface{}) interface{} {
		return left.(string) + right.(string)
	}
	calc.ifNonNullLeftRight(concatenate)
}

func (calc *calculator) countWords() {
	text := calc.popValue().(string)
	regex, err := regexp.Compile(`[^\s]+(?:\s+|$)`)
	if err != nil {
		calc.pushValue(nil)
		calc.errs = append(calc.errs, err)
		return
	}
	matches := regex.FindAllString(text, -1)
	calc.pushValue(float64(len(matches)))
}

func (calc *calculator) findString() {
	text := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	calc.pushValue(strings.Index(text, lookFor))
}
