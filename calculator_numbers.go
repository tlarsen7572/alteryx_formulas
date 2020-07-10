package alteryx_formulas

func (calc *calculator) addNumbers() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(value1.(float64) + value2.(float64))
}
