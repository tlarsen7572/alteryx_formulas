package alteryx_formulas

import (
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
)

func number(value float64) nullableNum {
	return nullableNum{value: value}
}

func nullNumber() nullableNum {
	return nullableNum{isNull: true}
}

func (l *listener) EnterAdd(_ *parser.AddContext) {
	l.pushFunction(l.add)
}

func (l *listener) EnterSubtract(_ *parser.SubtractContext) {
	l.pushFunction(l.subtract)
}

func (l *listener) EnterMultiply(_ *parser.MultiplyContext) {
	l.pushFunction(l.multiply)
}

func (l *listener) EnterDivide(_ *parser.DivideContext) {
	l.pushFunction(l.divide)
}

func (l *listener) EnterInteger(c *parser.IntegerContext) {
	value, _ := strconv.ParseFloat(c.GetText(), 64)
	f := func() {
		l.pushNumber(number(value))
	}
	l.pushFunction(f)
}

func (l *listener) EnterNumberNull(_ *parser.NumberNullContext) {
	f := func() {
		l.pushNumber(nullNumber())
	}
	l.pushFunction(f)
}

func (l *listener) add() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value+value2.value))
}

func (l *listener) subtract() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value-value2.value))
}

func (l *listener) multiply() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value*value2.value))
}

func (l *listener) divide() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull || value2.isNull {
		l.numbers = append(l.numbers, nullNumber())
	}
	l.numbers = append(l.numbers, number(value1.value/value2.value))
}
