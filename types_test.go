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
