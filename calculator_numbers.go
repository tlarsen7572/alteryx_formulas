package alteryx_formulas

func number(value float64) nullableNum {
	return nullableNum{value: value}
}

func nullNumber() nullableNum {
	return nullableNum{isNull: true}
}

func (calc *calculator) addNumbers() {
	value1 := calc.popNumber()
	value2 := calc.popNumber()
	if value1.isNull || value2.isNull {
		calc.pushNumber(nullNumber())
		return
	}
	calc.pushNumber(number(value1.value + value2.value))
}
