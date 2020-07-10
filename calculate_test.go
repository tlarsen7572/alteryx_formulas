package alteryx_formulas_test

import (
	"fmt"
	f "github.com/tlarsen7572/alteryx_formulas"
	"testing"
	"time"
)

func TestAddition(t *testing.T) {
	result, err := f.Calculate(`1.0+2+4`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got %v`, err.Error())
	}
	if result != 7.0 {
		t.Fatalf(`expected 7 but got %v`, result)
	}
}

func TestSubtraction(t *testing.T) {
	result, err := f.Calculate(`4-1.0-2`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 1.0 {
		t.Fatalf(`expected 1 but got %v`, result)
	}
}

func TestMultiplication(t *testing.T) {
	result, err := f.Calculate(`10*4.0`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 40.0 {
		t.Fatalf(`expected 40 but got %v`, result)
	}
}

func TestNullNumber(t *testing.T) {
	result, err := f.Calculate(`1+NULL()`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestDivision(t *testing.T) {
	result, err := f.Calculate(`40.0/10`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 4.0 {
		t.Fatalf(`expected 4 but got %v`, result)
	}
}

func TestNegativeNumberAddition(t *testing.T) {
	result, err := f.Calculate(`-1.0+-3`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != -4.0 {
		t.Fatalf(`expected -4 but got %v`, result)
	}
}

func TestNumberEquals(t *testing.T) {
	result, err := f.Calculate(`1=1.0`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}

	result, err = f.Calculate(`1=2`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got: %v`, result)
	}

	result, err = f.Calculate(`1=NULL()`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got: %v`, result)
	}

	result, err = f.Calculate(`NULL()=NULL()`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestNumberGreaterThan(t *testing.T) {
	result, err := f.Calculate(`1 > 2`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, err = f.Calculate(`2 > 1`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, err = f.Calculate(`2 > NULL()`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, err = f.Calculate(`NULL() > 2`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberGreaterEqual(t *testing.T) {
	result, err := f.Calculate(`1 >= 2.0`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, _ = f.Calculate(`2 >= 2.0`, nil)
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`3 >= 2.0`, nil)
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestNumberLessThan(t *testing.T) {
	result, err := f.Calculate(`1 < 2`, nil)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`2 < 2`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestIntField(t *testing.T) {
	recordInfo := &mockSingleFieldRecord{
		value:     1,
		isNull:    false,
		fieldType: f.Int64Type,
	}
	result, err := f.Calculate(`[MyField]+2`, recordInfo)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestFloatField(t *testing.T) {
	recordInfo := &mockSingleFieldRecord{
		value:     1.0,
		isNull:    false,
		fieldType: f.DoubleType,
	}
	result, err := f.Calculate(`[MyField]+2`, recordInfo)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestConcatenateField(t *testing.T) {
	recordInfo := &mockSingleFieldRecord{
		value:     `hello`,
		isNull:    false,
		fieldType: f.StringType,
	}
	result, err := f.Calculate(`[MyField]+' world'`, recordInfo)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != `hello world` {
		t.Fatalf(`expected 'hello world' but got %v`, result)
	}
}

func TestAddTwoNumberFields(t *testing.T) {
	recordInfo := &twoValues{fieldType: f.Int64Type}
	result, err := f.Calculate(`[Value1]+[Value2]`, recordInfo)
	if err != nil {
		t.Fatalf(`expected no error but got: %v`, err.Error())
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

type mockSingleFieldRecord struct {
	value     interface{}
	isNull    bool
	fieldType string
}

func (r *mockSingleFieldRecord) GetCurrentBool(_ string) (bool, bool, error) {
	if r.fieldType != f.BoolType {
		return false, false, fmt.Errorf(`wrong data type`)
	}
	return r.value.(bool), r.isNull, nil
}

func (r *mockSingleFieldRecord) GetCurrentInt(_ string) (int, bool, error) {
	switch r.fieldType {
	case f.ByteType, f.Int16Type, f.Int32Type, f.Int64Type:
	default:
		return 0, false, fmt.Errorf(`wrong data type`)
	}
	return r.value.(int), r.isNull, nil
}

func (r *mockSingleFieldRecord) GetCurrentFloat(_ string) (float64, bool, error) {
	switch r.fieldType {
	case f.FixedDecimalType, f.FloatType, f.DoubleType:
	default:
		return 0, false, fmt.Errorf(`wrong data type`)
	}
	return r.value.(float64), r.isNull, nil
}

func (r *mockSingleFieldRecord) GetCurrentString(_ string) (string, bool, error) {
	switch r.fieldType {
	case f.StringType, f.V_WStringType, f.WStringType, f.V_StringType:
	default:
		return ``, false, fmt.Errorf(`wrong data type`)
	}
	return r.value.(string), r.isNull, nil
}

func (r *mockSingleFieldRecord) GetCurrentDate(_ string) (time.Time, bool, error) {
	switch r.fieldType {
	case f.DateTimeType, f.DateType:
	default:
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), false, fmt.Errorf(`wrong data type`)
	}
	return r.value.(time.Time), r.isNull, nil
}

func (r *mockSingleFieldRecord) GetFieldTypeByName(_ string) (string, error) {
	return r.fieldType, nil
}

type twoValues struct {
	fieldType string
}

func (r *twoValues) GetCurrentBool(_ string) (bool, bool, error) {
	panic(`not implemented`)
}

func (r *twoValues) GetCurrentInt(fieldName string) (int, bool, error) {
	switch fieldName {
	case `Value1`:
		return 1, false, nil
	case `Value2`:
		return 2, false, nil
	default:
		return 0, false, fmt.Errorf(`field is not valid`)
	}
}

func (r *twoValues) GetCurrentFloat(_ string) (float64, bool, error) {
	panic(`not implemented`)
}

func (r *twoValues) GetCurrentString(fieldName string) (string, bool, error) {
	switch fieldName {
	case `Value1`:
		return `hello`, false, nil
	case `Value2`:
		return ` world`, false, nil
	default:
		return ``, false, fmt.Errorf(`field is not valid`)
	}
}

func (r *twoValues) GetCurrentDate(_ string) (time.Time, bool, error) {
	panic(`not implemented`)
}

func (r *twoValues) GetFieldTypeByName(_ string) (string, error) {
	return r.fieldType, nil
}
