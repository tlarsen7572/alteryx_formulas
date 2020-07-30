package alteryx_formulas_test

import (
	f "github.com/tlarsen7572/alteryx_formulas"
	"testing"
)

func TestMultiplyWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 * '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' * 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestMultiplyNestedArgumentTypesBubbleUp(t *testing.T) {
	result, errs := f.Calculate(`(1) * Length('ABC')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestDivideWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 / '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' / 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAddWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 + '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' + 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestSubtractWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 - '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' - 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestEqualWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 = '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' = 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestNotEqualWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1 != '2'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`'1' != 2`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestGreaterThanWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`2 > '1'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestGreaterEqualWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`2 >= '1'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestLessThanWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`2 < '1'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestLessEqualWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`2 <= '1'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAndWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1=1 AND 'B'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestOrWrongTypes(t *testing.T) {
	_, errs := f.Calculate(`1=1 OR 'B'`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestIfNotBoolType(t *testing.T) {
	_, errs := f.Calculate(`IF 'A' THEN 1 ELSE 2 ENDIF`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestThenAndElseNotSameType(t *testing.T) {
	_, errs := f.Calculate(`IF true THEN 1 ELSE '2' ENDIF`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestElseIfNotBoolType(t *testing.T) {
	_, errs := f.Calculate(`IF true THEN 1 ELSEIF 'B' THEN 3 ELSE 2 ENDIF`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestThenAndElseNotSameTypeInElseIf(t *testing.T) {
	_, errs := f.Calculate(`IF true THEN 1 ELSEIF true THEN 3 ELSE '2' ENDIF`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAbsWrongType(t *testing.T) {
	_, errs := f.Calculate(`abs('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAcosWrongType(t *testing.T) {
	_, errs := f.Calculate(`acos('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAsinWrongType(t *testing.T) {
	_, errs := f.Calculate(`asin('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAtanWrongType(t *testing.T) {
	_, errs := f.Calculate(`atan('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAtan2WrongType(t *testing.T) {
	_, errs := f.Calculate(`atan2('a', 1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`atan2(1, '1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestAverageWrongType(t *testing.T) {
	_, errs := f.Calculate(`average(1,2,'3')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestCeilWrongType(t *testing.T) {
	_, errs := f.Calculate(`ceil('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestCharFromIntWrongType(t *testing.T) {
	_, errs := f.Calculate(`charFromInt('a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestCharToIntWrongType(t *testing.T) {
	_, errs := f.Calculate(`charToInt(1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestContainsWrongType(t *testing.T) {
	_, errs := f.Calculate(`contains('123',1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`%v`, errs)

	_, errs = f.Calculate(`contains(123,'1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`%v`, errs)

	_, errs = f.Calculate(`contains('123','1', '1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`%v`, errs)
}

func TestCosWrongType(t *testing.T) {
	_, errs := f.Calculate(`cos('1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestCoshWrongType(t *testing.T) {
	_, errs := f.Calculate(`cosh('1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestCountWordsWrongType(t *testing.T) {
	_, errs := f.Calculate(`countwords(1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestDistanceWrongType(t *testing.T) {
	_, errs := f.Calculate(`DISTANCE('42', -90, 43, -80)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`DISTANCE(42, "-90", 43, -80)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`DISTANCE(42, -90, '43', -80)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`DISTANCE(42, -90, 43, '-80')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestEndsWithWrongType(t *testing.T) {
	_, errs := f.Calculate(`endswith('123',1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`endswith(123, '1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`endswith('123', '3', '1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestExpWrongType(t *testing.T) {
	_, errs := f.Calculate(`exp('1')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}

func TestFindStringWrongType(t *testing.T) {
	_, errs := f.Calculate(`findstring(1, 'a')`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)

	_, errs = f.Calculate(`findstring('a', 1)`, nil)
	if len(errs) == 0 {
		t.Fatalf(`expected errors but got none`)
	}
	t.Logf(`errs: %v`, errs)
}
