package alteryx_formulas

import "math"

func (calc *calculator) addNumbers() {
	add := func(left interface{}, right interface{}) interface{} {
		return left.(float64) + right.(float64)
	}
	calc.ifNonNullLeftRight(add)
}

func (calc *calculator) subtractNumbers() {
	subtract := func(left interface{}, right interface{}) interface{} {
		return left.(float64) - right.(float64)
	}
	calc.ifNonNullLeftRight(subtract)
}

func (calc *calculator) multiplyNumbers() {
	multiply := func(left interface{}, right interface{}) interface{} {
		return left.(float64) * right.(float64)
	}
	calc.ifNonNullLeftRight(multiply)
}

func (calc *calculator) divideNumbers() {
	divide := func(left interface{}, right interface{}) interface{} {
		return left.(float64) / right.(float64)
	}
	calc.ifNonNullLeftRight(divide)
}

func (calc *calculator) numberIn() {
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

func (calc *calculator) numberNotIn() {
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

func (calc *calculator) pow() {
	value := calc.popValue().(float64)
	power := calc.popValue().(float64)
	calc.pushValue(math.Pow(value, power))
}

func (calc *calculator) numberMin() {
	arguments := calc.popValue().(int)
	var min float64
	for i := 0; i < arguments; i++ {
		if i == 0 {
			min = calc.popValue().(float64)
			continue
		}
		nextValue := calc.popValue().(float64)
		if nextValue < min {
			min = nextValue
		}
	}
	calc.pushValue(min)
}

func (calc *calculator) numberMax() {
	arguments := calc.popValue().(int)
	var max float64
	for i := 0; i < arguments; i++ {
		if i == 0 {
			max = calc.popValue().(float64)
			continue
		}
		nextValue := calc.popValue().(float64)
		if nextValue > max {
			max = nextValue
		}
	}
	calc.pushValue(max)
}

func (calc *calculator) numberIif() {
	question := calc.popValue().(bool)
	thenValue := calc.popValue().(float64)
	elseValue := calc.popValue().(float64)
	if question {
		calc.pushValue(thenValue)
	} else {
		calc.pushValue(elseValue)
	}
}

func (calc *calculator) abs() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Abs(expr))
}
