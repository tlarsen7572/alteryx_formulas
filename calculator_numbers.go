package alteryx_formulas

import (
	"math"
	"math/rand"
	"sort"
)

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
		if right.(float64) == 0 {
			return nil
		}
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

func (calc *calculator) acos() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Acos(expr))
}

func (calc *calculator) asin() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Asin(expr))
}

func (calc *calculator) atan() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Atan(expr))
}

func (calc *calculator) atan2() {
	expr1 := calc.popValue().(float64)
	expr2 := calc.popValue().(float64)
	calc.pushValue(math.Atan2(expr1, expr2))
}

func (calc *calculator) average() {
	exprCount := calc.popValue().(int)
	sum := 0.0
	for i := 0; i < exprCount; i++ {
		value := calc.popValue()
		if value == nil {
			continue
		}
		sum += value.(float64)
	}
	calc.pushValue(sum / float64(exprCount))
}

func (calc *calculator) ceil() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Ceil(expr))
}

func (calc *calculator) cos() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Cos(expr))
}

func (calc *calculator) cosh() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Cosh(expr))
}

func (calc *calculator) distance() {
	fromLat := calc.popValue().(float64)
	fromLon := calc.popValue().(float64)
	toLat := calc.popValue().(float64)
	toLon := calc.popValue().(float64)

	distance, _ := calc.geo.To(fromLat, fromLon, toLat, toLon)

	calc.pushValue(distance)
}

func (calc *calculator) exp() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Exp(expr))
}

func (calc *calculator) floor() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Floor(expr))
}

func (calc *calculator) log() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Log(expr))
}

func (calc *calculator) log10() {
	expr := calc.popValue().(float64)
	calc.pushValue(math.Log10(expr))
}

func (calc *calculator) median() {
	exprCount := calc.popValue().(int)
	exprs := make([]float64, exprCount)
	for i := 0; i < exprCount; i++ {
		exprs[i] = calc.popValue().(float64)
	}
	sort.Float64s(exprs)
	var value float64
	if exprCount%2 > 0 {
		valueIndex := int(math.Floor(float64(exprCount) / 2))
		value = exprs[valueIndex]
	} else {
		valueIndex := exprCount / 2
		value = (exprs[valueIndex] + exprs[valueIndex-1]) / 2
	}
	calc.pushValue(value)
}

func (calc *calculator) mod() {
	dividend := int(calc.popValue().(float64))
	divisor := int(calc.popValue().(float64))
	if divisor == 0 {
		calc.pushValue(nil)
		return
	}
	remainder := dividend % divisor
	calc.pushValue(float64(remainder))
}

func (calc *calculator) pi() {
	calc.pushValue(math.Pi)
}

func (calc *calculator) rand() {
	calc.pushValue(rand.Float64())
}

func (calc *calculator) randInt() {
	// adding 1 to the ceiling and using Floor() gives each integer the same probability.  If we use the ceiling
	// as-is and then round the final result, 0 and the ceiling both have half the probability of the other integers
	// and will, over time, have half the number of hits.
	ceiling := calc.popValue().(float64) + 1
	value := rand.Float64() * ceiling
	calc.pushValue(math.Floor(value))
}

func (calc *calculator) round() {
	value := calc.popValue().(float64)
	multiple := calc.popValue().(float64)
	if multiple == 0 {
		calc.pushValue(nil)
		return
	}
	floor, _ := math.Modf(value / multiple)
	floor = floor * multiple
	var ceiling float64
	if value >= 0 {
		ceiling = floor + multiple
	} else {
		ceiling = floor - multiple
	}
	if math.Abs(value-floor) < math.Abs(value-ceiling) {
		calc.pushValue(floor)
	} else {
		calc.pushValue(ceiling)
	}
}

func (calc *calculator) sin() {
	radians := calc.popValue().(float64)
	calc.pushValue(math.Sin(radians))
}

func (calc *calculator) sinh() {
	radians := calc.popValue().(float64)
	calc.pushValue(math.Sinh(radians))
}

func (calc *calculator) sqrt() {
	value := calc.popValue().(float64)
	calc.pushValue(math.Sqrt(value))
}

func (calc *calculator) tan() {
	radians := calc.popValue().(float64)
	calc.pushValue(math.Tan(radians))
}
