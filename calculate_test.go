package alteryx_formulas_test

import (
	"fmt"
	f "github.com/tlarsen7572/alteryx_formulas"
	"testing"
	"time"
)

func TestAddition(t *testing.T) {
	result, err := f.Calculate(`1.0+2+4`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 7.0 {
		t.Fatalf(`expected 7 but got %v`, result)
	}
}

func TestSubtraction(t *testing.T) {
	result, err := f.Calculate(`4-1.0-2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 1.0 {
		t.Fatalf(`expected 1 but got %v`, result)
	}
}

func TestMultiplication(t *testing.T) {
	result, err := f.Calculate(`10*4.0`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 40.0 {
		t.Fatalf(`expected 40 but got %v`, result)
	}
}

func TestNullNumber(t *testing.T) {
	result, err := f.Calculate(`1+NULL()`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestDivision(t *testing.T) {
	result, err := f.Calculate(`40.0/10`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 4.0 {
		t.Fatalf(`expected 4 but got %v`, result)
	}
}

func TestNegativeNumberAddition(t *testing.T) {
	result, err := f.Calculate(`-1.0+-3`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != -4.0 {
		t.Fatalf(`expected -4 but got %v`, result)
	}
}

func TestNumberEquals(t *testing.T) {
	result, err := f.Calculate(`1=1.0`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}

	result, err = f.Calculate(`1=2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != false {
		t.Fatalf(`expected false but got: %v`, result)
	}

	result, err = f.Calculate(`1=NULL()`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != false {
		t.Fatalf(`expected false but got: %v`, result)
	}

	result, err = f.Calculate(`NULL()=NULL()`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestNumberGreaterThan(t *testing.T) {
	result, err := f.Calculate(`1 > 2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, err = f.Calculate(`2 > 1`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, err = f.Calculate(`2 > NULL()`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, err = f.Calculate(`NULL() > 2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberGreaterEqual(t *testing.T) {
	result, err := f.Calculate(`1 >= 2.0`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
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
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`2 < 2`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberLessEqual(t *testing.T) {
	result, err := f.Calculate(`1 <= 2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`2 <= 2`, nil)
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`3 <= 2`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberNotEqual(t *testing.T) {
	result, err := f.Calculate(`1 != 2`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`2 != 2`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberIn(t *testing.T) {
	result, err := f.Calculate(`1 in (1,2,3,4)`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`1 in (2,3,4)`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestNumberNotIn(t *testing.T) {
	result, err := f.Calculate(`1 not in (2,3,4)`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}

	result, _ = f.Calculate(`1 not in (1,2,3,4)`, nil)
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestIf(t *testing.T) {
	result, err := f.Calculate(`if 1=1 THEN 2 else 3 ENDIF`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2 but got %v`, result)
	}
}

func TestElse(t *testing.T) {
	result, err := f.Calculate(`if 1=0 THEN 2 else 3 ENDIF`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestElseIf(t *testing.T) {
	result, err := f.Calculate(`if 1=0 THEN 2 elseif 1=1 THEN 4 ELSE 3 ENDIF`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 4.0 {
		t.Fatalf(`expected 4 but got %v`, result)
	}
}

func TestIntField(t *testing.T) {
	recordInfo := &mockSingleFieldRecord{
		value:     1,
		isNull:    false,
		fieldType: f.Int64Type,
	}
	result, err := f.Calculate(`[MyField]+2`, recordInfo)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
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
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
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
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != `hello world` {
		t.Fatalf(`expected 'hello world' but got %v`, result)
	}
}

func TestAddTwoNumberFields(t *testing.T) {
	recordInfo := &twoValues{fieldType: f.Int64Type}
	result, err := f.Calculate(`[Value1]+[Value2]`, recordInfo)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestNestedNumberEqualOperators(t *testing.T) {
	result, err := f.Calculate(`1+2=3+0`, nil)
	if len(err) > 0 {
		t.Fatalf(`expected no error but got %v`, err)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestAddDifferentTypes(t *testing.T) {
	result, err := f.Calculate(`1 + '2'`, nil)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(fmt.Sprintf(`%v`, err))
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestBadSyntax(t *testing.T) {
	_, err := f.Calculate(`1 ++ 2`, nil)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(fmt.Sprintf(`%v`, err))
}

func TestPow(t *testing.T) {
	result, errs := f.Calculate(`POW(2,3)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 8.0 {
		t.Fatalf(`expected 8 but got: %v`, result)
	}
}

func TestMinNumber(t *testing.T) {
	result, errs := f.Calculate(`min(10, 3,5,7,8)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestLotsOfNesting(t *testing.T) {
	result, errs := f.Calculate(`min(10,min(6,7,min(min(3,4),11),16,12),5,7,8)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestMaxNumber(t *testing.T) {
	result, errs := f.Calculate(`max(3,5,10,8)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 10.0 {
		t.Fatalf(`expected 10 but got %v`, result)
	}
}

func TestIif(t *testing.T) {
	result, errs := f.Calculate(`iif(1=1,2,3)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2 but got %v`, result)
	}

	result, _ = f.Calculate(`iif(1=2,4,6)`, nil)
	if result != 6.0 {
		t.Fatalf(`expected 6 but got %v`, result)
	}
}

func TestAbs(t *testing.T) {
	result, errs := f.Calculate(`abs(10)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 10.0 {
		t.Fatalf(`expected 10 but got %v`, result)
	}

	result, _ = f.Calculate(`abs(-20)`, nil)
	if result != 20.0 {
		t.Fatalf(`expected 20 but got %v`, result)
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
