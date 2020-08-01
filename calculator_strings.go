package alteryx_formulas

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func (calc *calculator) charFromInt() {
	popped := calc.popValue()
	if popped == nil {
		calc.pushValue(nil)
		return
	}
	charCode := rune(popped.(float64))
	if !utf8.ValidRune(charCode) {
		calc.pushValue(nil)
		return
	}
	value := string(charCode)
	calc.pushValue(value)
}

func (calc *calculator) charToInt() {
	popped := calc.popValue()
	if popped == nil {
		calc.pushValue(nil)
		return
	}
	value := popped.(string)
	if len(value) == 0 {
		calc.pushValue(nil)
		return
	}
	runes := []rune(value)
	code := float64(runes[0])
	calc.pushValue(code)
}

func (calc *calculator) contains() {
	popped1 := calc.popValue()
	popped2 := calc.popValue()
	popped3 := calc.popValue()
	if popped1 == nil || popped2 == nil {
		calc.pushValue(false)
		return
	}
	value := popped1.(string)
	lookFor := popped2.(string)
	caseInsensitive := 1.0
	if popped3 != nil {
		caseInsensitive = popped3.(float64)
	}
	if caseInsensitive != 0 {
		value = strings.ToLower(value)
		lookFor = strings.ToLower(lookFor)
	}
	contains := strings.Contains(value, lookFor)
	calc.pushValue(contains)
}

func (calc *calculator) endsWith() {
	popped1 := calc.popValue()
	popped2 := calc.popValue()
	popped3 := calc.popValue()
	if popped1 == nil || popped2 == nil {
		calc.pushValue(false)
		return
	}
	value := popped1.(string)
	lookFor := popped2.(string)
	caseInsensitive := 1.0
	if popped3 != nil {
		caseInsensitive = popped3.(float64)
	}

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
	text := calc.popValue()
	if text == nil {
		calc.pushValue(0.0)
		return
	}
	matches := calc.wordExp.FindAllString(text.(string), -1)
	calc.pushValue(float64(len(matches)))
}

func (calc *calculator) findString() {
	text := calc.popValue()
	lookFor := calc.popValue()
	if text == nil && lookFor == nil {
		calc.pushValue(true)
		return
	}
	if text == nil {
		calc.pushValue(false)
		return
	}
	if lookFor == nil {
		calc.pushValue(true)
		return
	}
	calc.pushValue(strings.Index(text.(string), lookFor.(string)))
}

func (calc *calculator) getWord() {
	text := calc.popValue()
	poppedIndex := calc.popValue()
	if text == nil {
		calc.pushValue(nil)
		return
	}
	wordIndex := 0
	if poppedIndex != nil {
		wordIndex = int(poppedIndex.(float64))
	}
	match := calc.wordExp.FindAllStringSubmatch(text.(string), -1)
	if wordIndex < 0 || wordIndex >= len(match) {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(match[wordIndex][1])
}

func (calc *calculator) left() {
	text := calc.popValue()
	lengthPopped := calc.popValue()
	if text == nil {
		calc.pushValue(nil)
		return
	}
	if lengthPopped == nil {
		calc.pushValue(``)
		return
	}

	length := int(lengthPopped.(float64))
	runes := []rune(text.(string))
	if length < 0 || length >= len(runes) {
		calc.pushValue(text.(string))
		return
	}
	newText := string(runes[:length])
	calc.pushValue(newText)
}

func (calc *calculator) length() {
	text := calc.popValue()
	if text == nil {
		calc.pushValue(0.0)
		return
	}
	calc.pushValue(float64(len(text.(string))))
}

func (calc *calculator) lowercase() {
	text := calc.popValue()
	if text == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(strings.ToLower(text.(string)))
}

func (calc *calculator) padLeft() {
	calc.pad(`left`)
}

func (calc *calculator) padRight() {
	calc.pad(`right`)
}

func (calc *calculator) pad(side string) {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	expr3 := calc.popValue()

	if expr1 == nil || expr2 == nil {
		calc.pushValue(nil)
		return
	}
	if expr3 == nil {
		expr3 = ``
	}

	text := expr1.(string)
	length := int(expr2.(float64))
	padCharStr := expr3.(string)
	if len(padCharStr) == 0 {
		calc.pushValue(text)
		return
	}
	padRunes := []rune(padCharStr)
	var newText string
	if side == `left` {
		newText = strings.Repeat(string(padRunes[0]), length) + text
	} else {
		newText = text + strings.Repeat(string(padRunes[0]), length)
	}
	calc.pushValue(newText)
}

func (calc *calculator) regexCountMatches() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	expr3 := calc.popValue()

	if expr1 == nil || expr2 == nil {
		calc.pushValue(0.0)
		return
	}

	if expr3 == nil {
		expr3 = 0.0
	}

	text := expr1.(string)
	regex := expr2.(string)
	caseInsensitive := expr3.(float64)

	if caseInsensitive != 0 {
		regex = `(?i)` + regex
	}
	r, err := regexp.Compile(regex)
	if err != nil {
		calc.pushValue(0)
		calc.errs = append(calc.errs, err)
	}
	result := r.FindAllString(text, -1)
	calc.pushValue(float64(len(result)))
}

func (calc *calculator) regexMatch() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	expr3 := calc.popValue()
	if expr1 == nil || expr2 == nil {
		calc.pushValue(false)
		return
	}

	if expr3 == nil {
		expr3 = 0.0
	}

	text := expr1.(string)
	regex := expr2.(string)
	caseInsensitive := expr3.(float64)

	if caseInsensitive != 0 {
		regex = `(?i)` + regex
	}
	r, err := regexp.Compile(regex)
	if err != nil {
		calc.pushValue(false)
		calc.errs = append(calc.errs, err)
		return
	}
	matches := r.MatchString(text)
	if matches {
		calc.pushValue(true)
	} else {
		calc.pushValue(false)
	}
}

func (calc *calculator) regexReplace() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	expr3 := calc.popValue()
	expr4 := calc.popValue()
	if expr1 == nil || expr3 == nil {
		calc.pushValue(``)
		return
	}
	if expr2 == nil {
		expr2 = ``
	}
	if expr4 == nil {
		expr4 = 0.0
	}

	text := expr1.(string)
	regex := expr2.(string)
	replaceWith := expr3.(string)
	caseInsensitive := expr4.(float64)

	if caseInsensitive != 0 {
		regex = `(?i)` + regex
	}
	r, err := regexp.Compile(regex)
	if err != nil {
		calc.pushValue(nil)
		calc.errs = append(calc.errs, err)
		return
	}
	newText := r.ReplaceAllString(text, replaceWith)
	calc.pushValue(newText)
}

func (calc *calculator) replace() {
	text := calc.popValue()
	find := calc.popValue()
	replace := calc.popValue()
	if text == nil {
		calc.pushValue(``)
		return
	}
	if find == nil {
		calc.pushValue(text)
		return
	}
	if replace == nil {
		replace = ``
	}

	newText := strings.ReplaceAll(text.(string), find.(string), replace.(string))
	calc.pushValue(newText)
}

func (calc *calculator) right() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	if expr1 == nil {
		calc.pushValue(nil)
		return
	}
	if expr2 == nil {
		calc.pushValue(``)
		return
	}

	text := expr1.(string)
	length := int(expr2.(float64))
	runes := []rune(text)
	if length < 0 || length >= len(runes) {
		calc.pushValue(text)
		return
	}
	startAt := len(runes) - length
	newText := string(runes[startAt:])
	calc.pushValue(newText)
}

func (calc *calculator) substring() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	expr3 := calc.popValue()
	if expr1 == nil {
		calc.pushValue(nil)
		return
	}
	if expr3 == nil {
		calc.pushValue(``)
		return
	}
	if expr2 == nil {
		expr2 = 0.0
	}

	text := expr1.(string)
	startAt := int(expr2.(float64))
	length := int(expr3.(float64))

	if length < 0 {
		calc.pushValue(``)
		return
	}

	runes := []rune(text)
	runeLen := len(runes)

	startIndex := startAt
	endIndex := startAt + length

	if endIndex < 0 || startIndex > runeLen {
		calc.pushValue(``)
		return
	}

	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > runeLen {
		endIndex = runeLen
	}

	newText := string(runes[startIndex:endIndex])
	calc.pushValue(newText)
}

func (calc *calculator) trim() {
	expr1 := calc.popValue()
	trimChars := calc.popValue()
	if expr1 == nil {
		calc.pushValue(nil)
		return
	}
	text := expr1.(string)
	if trimChars == nil || trimChars == `` {
		calc.pushValue(strings.TrimSpace(text))
		return
	}
	calc.pushValue(strings.Trim(text, trimChars.(string)))
}

func (calc *calculator) trimLeft() {
	expr1 := calc.popValue()
	trimChars := calc.popValue()
	if expr1 == nil {
		calc.pushValue(nil)
		return
	}
	text := expr1.(string)
	if trimChars == nil || trimChars == `` {
		calc.pushValue(strings.TrimLeft(text, ` `))
		return
	}
	calc.pushValue(strings.TrimLeft(text, trimChars.(string)))
}

func (calc *calculator) trimRight() {
	text := calc.popValue().(string)
	trimChars := calc.popValue().(string)
	if trimChars == `` {
		calc.pushValue(strings.TrimRight(text, ` `))
		return
	}
	calc.pushValue(strings.TrimRight(text, trimChars))
}

func (calc *calculator) uppercase() {
	text := calc.popValue().(string)
	calc.pushValue(strings.ToUpper(text))
}
