package alteryx_formulas

import (
	"github.com/tlarsen7572/alteryx_formulas/parser"
)

type listener struct {
	functions []func()
	numbers   []nullableNum
	bools     []nullableBool
	getResult func() NullableValue
	parser.BaseAlteryxFormulasListener
}

func (l *listener) Calculate() NullableValue {
	for funcIndex := len(l.functions) - 1; funcIndex >= 0; funcIndex-- {
		l.functions[funcIndex]()
	}
	return l.getResult()
}

func (l *listener) pushFunction(f func()) {
	l.functions = append(l.functions, f)
}

func (l *listener) popNumber() nullableNum {
	item := len(l.numbers) - 1
	returnVal := l.numbers[item]
	l.numbers = l.numbers[:item]
	return returnVal
}

func (l *listener) pushNumber(number nullableNum) {
	l.numbers = append(l.numbers, number)
}

func (l *listener) popBool() nullableBool {
	item := len(l.bools) - 1
	returnVal := l.bools[item]
	l.bools = l.bools[:item]
	return returnVal
}

func (l *listener) pushBool(value nullableBool) {
	l.bools = append(l.bools, value)
}

func (l *listener) EnterFormulaIsNumber(_ *parser.FormulaIsNumberContext) {
	f := func() NullableValue {
		return l.numbers[0]
	}
	l.getResult = f
}

func (l *listener) EnterFormulaIsBool(_ *parser.FormulaIsBoolContext) {
	f := func() NullableValue {
		return l.bools[0]
	}
	l.getResult = f
}
