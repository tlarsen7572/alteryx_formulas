package alteryx_formulas

import (
	"math"
	"math/rand"
	"sort"
	"strconv"
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

func (calc *calculator) pow() {
	value := calc.popValue()
	power := calc.popValue()
	if value == nil || power == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Pow(value.(float64), power.(float64)))
}

func (calc *calculator) abs() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Abs(expr.(float64)))
}

func (calc *calculator) acos() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Acos(expr.(float64)))
}

func (calc *calculator) asin() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Asin(expr.(float64)))
}

func (calc *calculator) atan() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Atan(expr.(float64)))
}

func (calc *calculator) atan2() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	if expr1 == nil {
		calc.pushValue(nil)
		return
	}
	if expr2 == nil {
		expr2 = 0.0
	}
	calc.pushValue(math.Atan2(expr1.(float64), expr2.(float64)))
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
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Ceil(expr.(float64)))
}

func (calc *calculator) cos() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Cos(expr.(float64)))
}

func (calc *calculator) cosh() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Cosh(expr.(float64)))
}

func (calc *calculator) distance() {
	values := make([]float64, 4)
	for i := 0; i < 4; i++ {
		poppedValue := calc.popValue()
		if poppedValue == nil {
			values[i] = 0.0
		} else {
			values[i] = poppedValue.(float64)
		}
	}

	distance, _ := calc.geo.To(values[0], values[1], values[2], values[3])

	calc.pushValue(distance)
}

func (calc *calculator) exp() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Exp(expr.(float64)))
}

func (calc *calculator) floor() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Floor(expr.(float64)))
}

func (calc *calculator) log() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Log(expr.(float64)))
}

func (calc *calculator) log10() {
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Log10(expr.(float64)))
}

func (calc *calculator) median() {
	exprCount := calc.popValue().(int)
	exprs := make([]float64, exprCount)
	for i := 0; i < exprCount; i++ {
		popped := calc.popValue()
		if popped == nil {
			exprs[i] = 0.0
		} else {
			exprs[i] = popped.(float64)
		}
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
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	if expr1 == nil || expr2 == nil {
		calc.pushValue(nil)
		return
	}

	dividend := int(expr1.(float64))
	divisor := int(expr2.(float64))
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
	expr := calc.popValue()
	if expr == nil {
		calc.pushValue(0.0)
		return
	}
	ceiling := expr.(float64) + 1
	value := rand.Float64() * ceiling
	calc.pushValue(math.Floor(value))
}

func (calc *calculator) round() {
	expr1 := calc.popValue()
	expr2 := calc.popValue()
	if expr1 == nil || expr2 == nil {
		calc.pushValue(nil)
		return
	}

	value := expr1.(float64)
	multiple := expr2.(float64)
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
	radians := calc.popValue()
	if radians == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Sin(radians.(float64)))
}

func (calc *calculator) sinh() {
	radians := calc.popValue()
	if radians == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(math.Sinh(radians.(float64)))
}

func (calc *calculator) sqrt() {
	value := calc.popValue().(float64)
	calc.pushValue(math.Sqrt(value))
}

func (calc *calculator) tan() {
	radians := calc.popValue().(float64)
	calc.pushValue(math.Tan(radians))
}

func (calc *calculator) tanh() {
	radians := calc.popValue().(float64)
	calc.pushValue(math.Tanh(radians))
}

func (calc *calculator) hexToNumber() {
	hexCode := calc.popValue()
	if hexCode == nil {
		calc.pushValue(nil)
		return
	}
	result, err := strconv.ParseInt(hexCode.(string), 16, 64)
	if err != nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(float64(result))
}
