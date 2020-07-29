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
