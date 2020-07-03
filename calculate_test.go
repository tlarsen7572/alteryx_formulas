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
	if result != 7 {
		t.Fatalf(`expected 7 but got %v`, result)
	}
}
