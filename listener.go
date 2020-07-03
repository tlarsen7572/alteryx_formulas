package alteryx_formulas

import (
	"github.com/tlarsen7572/alteryx_formulas/parser"
	"strconv"
)

type listener struct {
	functions []func()
	numbers   []int
	parser.BaseAlteryxFormulasListener
}

func (l *listener) calculateInt() int {
	for funcIndex := len(l.functions) - 1; funcIndex >= 0; funcIndex-- {
		l.functions[funcIndex]()
	}
	return l.numbers[0]
}

func (l *listener) popNumbers() int {
	item := len(l.numbers) - 1
	returnVal := l.numbers[item]
	l.numbers = l.numbers[:item]
	return returnVal
}

func (l *listener) add() {
	value1 := l.popNumbers()
	value2 := l.popNumbers()
	calc := value1 + value2
	l.numbers = append(l.numbers, calc)
}

func (l *listener) EnterAdd(c *parser.AddContext) {
	l.functions = append(l.functions, l.add)
}

func (l *listener) EnterInteger(c *parser.IntegerContext) {
	value, _ := strconv.Atoi(c.GetText())
	f := func() {
		l.numbers = append(l.numbers, value)
	}
	l.functions = append(l.functions, f)
}
