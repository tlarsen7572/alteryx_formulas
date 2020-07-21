package alteryx_formulas

import (
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
	matches := calc.wordExp.FindAllString(text, -1)
	calc.pushValue(float64(len(matches)))
}

func (calc *calculator) findString() {
	text := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	calc.pushValue(strings.Index(text, lookFor))
}

func (calc *calculator) getWord() {
	text := calc.popValue().(string)
	wordIndex := int(calc.popValue().(float64))
	match := calc.wordExp.FindAllStringSubmatch(text, -1)
	if wordIndex < 0 || wordIndex >= len(match) {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(match[wordIndex][1])
}

func (calc *calculator) left() {
	text := calc.popValue().(string)
	length := int(calc.popValue().(float64))
	if length < 0 || length >= len(text) {
		calc.pushValue(text)
		return
	}
	calc.pushValue(text[:length])
}
