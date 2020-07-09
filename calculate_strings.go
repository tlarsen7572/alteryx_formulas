package alteryx_formulas

import "github.com/tlarsen7572/alteryx_formulas/parser"

func stringVal(value string) nullableString {
	return nullableString{value: value}
}

func nullString() nullableString {
	return nullableString{isNull: true}
}

func (l *listener) EnterConcatenate(_ *parser.ConcatenateContext) {
	l.pushFunction(l.concatenate)
}

func (l *listener) EnterStringLiteral(c *parser.StringLiteralContext) {
	text := c.GetText()
	value := text[1 : len(text)-1]
	f := func() {
		l.pushString(stringVal(value))
	}
	l.pushFunction(f)
}

func (l *listener) EnterStringField(c *parser.StringFieldContext) {
	text := c.GetText()
	fieldName := text[1 : len(text)-1]
	fieldType, err := l.recordInfo.GetFieldTypeByName(fieldName)
	if err != nil {
		return
	}

	var f func()
	switch fieldType {
	case StringType, WStringType, V_StringType, V_WStringType:
		f = func() {
			value, isNull, err := l.recordInfo.GetCurrentString(fieldName)
			if err != nil {
				return
			}
			if isNull {
				l.strings = append(l.strings, nullString())
				return
			}
			l.strings = append(l.strings, stringVal(value))
		}
	default:
		c.AddErrorNode(c.GetStart())
		return
	}
	l.pushFunction(f)
}

func (l *listener) concatenate() {
	value1 := l.popString()
	value2 := l.popString()
	if value1.isNull || value2.isNull {
		l.pushString(nullString())
		return
	}
	l.pushString(stringVal(value1.value + value2.value))
}
