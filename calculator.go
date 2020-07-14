package alteryx_formulas

import "github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"

type calculator struct {
	functions  []func()
	values     []interface{}
	recordInfo RecordInfo
	geo        ellipsoid.Ellipsoid
}

func (calc *calculator) Calculate() interface{} {
	for funcIndex := len(calc.functions) - 1; funcIndex >= 0; funcIndex-- {
		calc.functions[funcIndex]()
	}
	return calc.values[0]
}

func (calc *calculator) pushFunction(f func()) {
	calc.functions = append(calc.functions, f)
}

func (calc *calculator) popValue() interface{} {
	item := len(calc.values) - 1
	returnVal := calc.values[item]
	calc.values = calc.values[:item]
	return returnVal
}

func (calc *calculator) pushValue(value interface{}) {
	calc.values = append(calc.values, value)
}

func (calc *calculator) pushValueFunc(value interface{}) {
	f := func() {
		calc.pushValue(value)
	}
	calc.pushFunction(f)
}

func (calc *calculator) pushStringField(fieldName string) {
	f := func() {
		value, isNull, err := calc.recordInfo.GetCurrentString(fieldName)
		if err != nil || isNull {
			calc.pushValue(nil)
			return
		}
		calc.pushValue(value)
	}
	calc.pushFunction(f)
}

func (calc *calculator) pushNumberField(fieldName string) {
	fieldType, _ := calc.recordInfo.GetFieldTypeByName(fieldName)
	var f func()
	switch fieldType {
	case ByteType, Int16Type, Int32Type, Int64Type:
		f = func() {
			value, isNull, err := calc.recordInfo.GetCurrentInt(fieldName)
			if err != nil || isNull {
				calc.pushValue(nil)
				return
			}
			calc.pushValue(float64(value))
		}
	case FixedDecimalType, FloatType, DoubleType:
		f = func() {
			value, isNull, err := calc.recordInfo.GetCurrentFloat(fieldName)
			if err != nil || isNull {
				calc.pushValue(nil)
				return
			}
			calc.pushValue(value)
		}
	default:
		panic(`invalid field type`)
	}
	calc.pushFunction(f)
}

type leftRightOperator func(left interface{}, right interface{}) interface{}

func (calc *calculator) ifNonNullLeftRight(f leftRightOperator) {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(nil)
		return
	}
	calc.pushValue(f(value1, value2))
}

func (calc *calculator) exprIf() {
	exprCount := calc.popValue().(int)
	exprs := make([]interface{}, exprCount)
	for i := 0; i < exprCount; i++ {
		exprs[i] = calc.popValue()
	}
	ifCount := (exprCount - 1) / 2
	for i := 0; i < ifCount; i++ {
		if exprs[i*2] == true {
			valueIndex := (i * 2) + 1
			calc.pushValue(exprs[valueIndex])
			return
		}
	}
	calc.pushValue(exprs[exprCount-1])
}
