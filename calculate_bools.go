package alteryx_formulas

import "github.com/tlarsen7572/alteryx_formulas/parser"

func boolVal(value bool) nullableBool {
	return nullableBool{value: value}
}

func nullBool() nullableBool {
	return nullableBool{isNull: true}
}

func (l *listener) EnterNumberEqual(_ *parser.NumberEqualContext) {
	l.pushFunction(l.numberEqual)
}

func (l *listener) EnterNumberGreaterThan(_ *parser.NumberGreaterThanContext) {
	l.pushFunction(l.numberGreaterThan)
}

func (l *listener) numberEqual() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull && value2.isNull {
		l.pushBool(boolVal(true))
	}
	if value1.isNull || value2.isNull {
		l.pushBool(boolVal(false))
	}
	l.pushBool(boolVal(value1.value == value2.value))
}

func (l *listener) numberGreaterThan() {
	value1 := l.popNumber()
	value2 := l.popNumber()
	if value1.isNull || value2.isNull {
		l.pushBool(boolVal(false))
	}
	l.pushBool(boolVal(value1.value > value2.value))
}
