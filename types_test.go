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
