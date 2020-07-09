package alteryx_formulas

func stringVal(value string) nullableString {
	return nullableString{value: value}
}

func nullString() nullableString {
	return nullableString{isNull: true}
}

func (calc *calculator) concatenate() {
	value1 := calc.popString()
	value2 := calc.popString()
	if value1.isNull || value2.isNull {
		calc.pushString(nullString())
		return
	}
	calc.pushString(stringVal(value1.value + value2.value))
}
