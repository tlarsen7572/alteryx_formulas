package alteryx_formulas

import (
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
)

type listener struct {
	functions []func()
	numbers   []nullableNum
	parser.BaseAlteryxFormulasListener
}

func number(value float64) nullableNum {
	return nullableNum{value: value}
}

func nullNumber() nullableNum {
	return nullableNum{isNull: true}
}

func (l *listener) calculateNumber() nullableNum {
	for funcIndex := len(l.functions) - 1; funcIndex >= 0; funcIndex-- {
		l.functions[funcIndex]()
	}
	return l.numbers[0]
}

func (l *listener) popNumbers() nullableNum {
	item := len(l.numbers) - 1
	returnVal := l.numbers[item]
	l.numbers = l.numbers[:item]
	return returnVal
}

func (l *listener) add() {
	value1 := l.popNumbers()
	value2 := l.popNumbers()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value+value2.value))
}

func (l *listener) subtract() {
	value1 := l.popNumbers()
	value2 := l.popNumbers()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value-value2.value))
}

func (l *listener) multiply() {
	value1 := l.popNumbers()
	value2 := l.popNumbers()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value*value2.value))
}

func (l *listener) EnterAdd(_ *parser.AddContext) {
	l.functions = append(l.functions, l.add)
}

func (l *listener) EnterSubtract(_ *parser.SubtractContext) {
	l.functions = append(l.functions, l.subtract)
}

func (l *listener) EnterMultiply(_ *parser.MultiplyContext) {
	l.functions = append(l.functions, l.multiply)
}

func (l *listener) EnterInteger(c *parser.IntegerContext) {
	value, _ := strconv.ParseFloat(c.GetText(), 64)
	f := func() {
		l.numbers = append(l.numbers, number(value))
	}
	l.functions = append(l.functions, f)
}
