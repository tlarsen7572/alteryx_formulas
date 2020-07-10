package alteryx_formulas

func (calc *calculator) concatenate() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(value1.(string) + value2.(string))
}
