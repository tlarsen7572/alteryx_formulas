package alteryx_formulas

func (calc *calculator) equal() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 == value2)
}

func (calc *calculator) notEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 != value2)
}

func (calc *calculator) numberGreaterThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) > value2.(float64))
}

func (calc *calculator) stringGreaterThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(string) > value2.(string))
}

func (calc *calculator) numberGreaterEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) >= value2.(float64))
}

func (calc *calculator) stringGreaterEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(string) >= value2.(string))
}

func (calc *calculator) numberLessThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) < value2.(float64))
}

func (calc *calculator) stringLessThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(string) < value2.(string))
}

func (calc *calculator) numberLessEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) <= value2.(float64))
}

func (calc *calculator) stringLessEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(string) <= value2.(string))
}

func (calc *calculator) isNull() {
	value := calc.popValue()
	calc.pushValue(value == nil)
}

func (calc *calculator) isEmpty() {
	value := calc.popValue()
	calc.pushValue(value == nil || value == ``)
}
