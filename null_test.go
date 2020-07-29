package alteryx_formulas_test

import (
	"math"
	"testing"
)
import f "github.com/tlarsen7572/alteryx_formulas"

func TestIfNull(t *testing.T) {
	result, errs := f.Calculate(`IF NULL() THEN 1 ELSE 2 ENDIF`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2.0 but got %v`, result)
	}
}

func TestAbsNull(t *testing.T) {
	result, errs := f.Calculate(`abs(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAcosNull(t *testing.T) {
	result, errs := f.Calculate(`acos(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAsinNull(t *testing.T) {
	result, errs := f.Calculate(`asin(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAtanNull(t *testing.T) {
	result, errs := f.Calculate(`atan(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAtan2Null(t *testing.T) {
	result, errs := f.Calculate(`atan2(1, NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expected := math.Atan2(1, 0); result != expected {
		t.Fatalf(`expected %v but got %v`, expected, result)
	}

	result, errs = f.Calculate(`atan2(NULL(), 1)`, nil)
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}
