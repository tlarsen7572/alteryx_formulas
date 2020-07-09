package alteryx_formulas

type calculator struct {
	functions  []func()
	strings    []nullableString
	numbers    []nullableNum
	bools      []nullableBool
	getResult  func() NullableValue
	recordInfo RecordInfo
}

func (calc *calculator) Calculate() NullableValue {
	for funcIndex := len(calc.functions) - 1; funcIndex >= 0; funcIndex-- {
		calc.functions[funcIndex]()
	}
	return calc.getResult()
}

func (calc *calculator) returnNumber() NullableValue {
	return calc.numbers[0]
}

func (calc *calculator) returnString() NullableValue {
	return calc.strings[0]
}

func (calc *calculator) returnBool() NullableValue {
	return calc.bools[0]
}

func (calc *calculator) pushFunction(f func()) {
	calc.functions = append(calc.functions, f)
}

func (calc *calculator) popNumber() nullableNum {
	item := len(calc.numbers) - 1
	returnVal := calc.numbers[item]
	calc.numbers = calc.numbers[:item]
	return returnVal
}

func (calc *calculator) pushNumber(value nullableNum) {
	calc.numbers = append(calc.numbers, value)
}

func (calc *calculator) pushNumberFunc(value nullableNum) {
	f := func() {
		calc.pushNumber(value)
	}
	calc.pushFunction(f)
}

func (calc *calculator) popBool() nullableBool {
	item := len(calc.bools) - 1
	returnVal := calc.bools[item]
	calc.bools = calc.bools[:item]
	return returnVal
}

func (calc *calculator) pushBool(value nullableBool) {
	calc.bools = append(calc.bools, value)
}

func (calc *calculator) pushBoolFunc(value nullableBool) {
	f := func() {
		calc.pushBool(value)
	}
	calc.pushFunction(f)
}

func (calc *calculator) popString() nullableString {
	item := len(calc.strings) - 1
	returnVal := calc.strings[item]
	calc.strings = calc.strings[:item]
	return returnVal
}

func (calc *calculator) pushString(value nullableString) {
	calc.strings = append(calc.strings, value)
}

func (calc *calculator) pushStringFunc(value nullableString) {
	f := func() {
		calc.pushString(value)
	}
	calc.pushFunction(f)
}

func (calc *calculator) pushStringField(fieldName string) {
	f := func() {
		value, isNull, err := calc.recordInfo.GetCurrentString(fieldName)
		if err != nil || isNull {
			calc.pushString(nullString())
			return
		}
		calc.pushString(stringVal(value))
	}
	calc.pushFunction(f)
}
