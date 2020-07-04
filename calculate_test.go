package alteryx_formulas_test

import (
	f "github.com/tlarsen7572/alteryx_formulas"
	"testing"
)

func TestAddition(t *testing.T) {
	result, err := f.Calculate(`1+2+4`)
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if result.Value() != 7.0 {
		t.Fatalf(`expected 7 but got %v`, result.Value())
	}
}

func TestSubtraction(t *testing.T) {
	result, err := f.Calculate(`4-1-2`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result.Value() != 1.0 {
		t.Fatalf(`expected 1 but got %v`, result.Value())
	}
}

func TestMultiplication(t *testing.T) {
	result, err := f.Calculate(`10*4`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result.Value() != 40.0 {
		t.Fatalf(`expected 40 but got %v`, result.Value())
	}
}

func TestNullNumber(t *testing.T) {
	result, err := f.Calculate(`1+NULL()`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result.Value() != nil {
		t.Fatalf(`expected nil but got %v`, result.Value())
	}
}

func TestDivision(t *testing.T) {
	result, err := f.Calculate(`40/10`)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result.Value() != 4.0 {
		t.Fatalf(`expected 4 but got %v`, result.Value())
	}
}
