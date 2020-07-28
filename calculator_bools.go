package alteryx_formulas

import "time"

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

func (calc *calculator) dateGreaterThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(time.Time).After(value2.(time.Time)))
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

func (calc *calculator) dateGreaterEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(time.Time).After(value2.(time.Time)) || value1 == value2)
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

func (calc *calculator) dateLessThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(time.Time).Before(value2.(time.Time)))
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

func (calc *calculator) dateLessEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(time.Time).Before(value2.(time.Time)) || value1 == value2)
}

func (calc *calculator) and() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 == true && value2 == true)
}

func (calc *calculator) or() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 == true || value2 == true)
}

func (calc *calculator) isNull() {
	value := calc.popValue()
	calc.pushValue(value == nil)
}

func (calc *calculator) isEmpty() {
	value := calc.popValue()
	calc.pushValue(value == nil || value == ``)
}

func (calc *calculator) in() {
	exprCount := calc.popValue().(int)
	exprs := make([]interface{}, exprCount)
	for i := 0; i < exprCount; i++ {
		exprs[i] = calc.popValue()
	}
	baseValue := exprs[0]
	for i := 1; i < exprCount; i++ {
		if baseValue == exprs[i] {
			calc.pushValue(true)
			return
		}
	}
	calc.pushValue(false)
}

func (calc *calculator) notIn() {
	exprCount := calc.popValue().(int)
	exprs := make([]interface{}, exprCount)
	for i := 0; i < exprCount; i++ {
		exprs[i] = calc.popValue()
	}
	baseValue := exprs[0]
	for i := 1; i < exprCount; i++ {
		if baseValue == exprs[i] {
			calc.pushValue(false)
			return
		}
	}
	calc.pushValue(true)
}
