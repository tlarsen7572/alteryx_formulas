package alteryx_formulas_test

import (
	"fmt"
	f "github.com/tlarsen7572/alteryx_formulas"
	"math"
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

func TestAcos(t *testing.T) {
	result, errs := f.Calculate(`acos(0.5)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expectedVal := math.Acos(0.5); result != expectedVal {
		t.Fatalf(`expected %v but got %v`, expectedVal, result)
	}
}

func TestAtan(t *testing.T) {
	result, errs := f.Calculate(`atan(0.8)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expectedVal := math.Atan(0.8); result != expectedVal {
		t.Fatalf(`expected %v but got %v`, expectedVal, result)
	}
}

func TestAtan2(t *testing.T) {
	result, errs := f.Calculate(`atan2(4, -3)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expectedVal := math.Atan2(4, -3); result != expectedVal {
		t.Fatalf(`expected %v but got %v`, expectedVal, result)
	}
}

func TestAverage(t *testing.T) {
	result, errs := f.Calculate(`average(20, 30, 55)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 35.0 {
		t.Fatalf(`expected 35 but got %v`, result)
	}
}

func TestAverageWithNull(t *testing.T) {
	result, errs := f.Calculate(`average(20, 30, 55, NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 26.25 {
		t.Fatalf(`expected 26.25 but got %v`, result)
	}
}

func TestCeil(t *testing.T) {
	result, errs := f.Calculate(`ceil(1.1)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2 but got %v`, result)
	}
}

func TestCos(t *testing.T) {
	result, errs := f.Calculate(`cos(0.5)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expectedVal := math.Cos(0.5); result != expectedVal {
		t.Fatalf(`expected 2.0 but got %v`, result)
	}
}

func TestDistance(t *testing.T) {
	result, errs := f.Calculate(`distance(42,-90,43,-80)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	value := result.(float64)
	if value < 514.5 || value > 515.5 {
		t.Fatalf(`expected between 515 and 516 but got %v`, result)
	}
	t.Logf(`%v`, value)
}

func TestExp(t *testing.T) {
	result, errs := f.Calculate(`exp(4)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expected := math.Exp(4); result != expected {
		t.Fatalf(`expected %v but got %v`, expected, result)
	}
}

func TestFloor(t *testing.T) {
	result, errs := f.Calculate(`fLoor(5.5)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 5.0 {
		t.Fatalf(`expected 5 but got %v`, result)
	}
}

func TestLog(t *testing.T) {
	result, errs := f.Calculate(`log(14)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expected := math.Log(14); result != expected {
		t.Fatalf(`expected %v but got %v`, expected, result)
	}
}

func TestLog10(t *testing.T) {
	result, errs := f.Calculate(`log10(14)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expected := math.Log10(14); result != expected {
		t.Fatalf(`expected %v but got %v`, expected, result)
	}
}

func TestMedianWithOddElements(t *testing.T) {
	result, errs := f.Calculate(`median(5,4,3,7,6)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 5.0 {
		t.Fatalf(`expected 5 but got %v`, result)
	}
}

func TestMedianWithEvenElements(t *testing.T) {
	result, errs := f.Calculate(`median(5,4,3,7)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 4.5 {
		t.Fatalf(`expected 4.5 but got %v`, result)
	}
}

func TestMod(t *testing.T) {
	result, errs := f.Calculate(`mod(6,4)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2 but got %v`, result)
	}
}

func TestModWithDecimal(t *testing.T) {
	result, errs := f.Calculate(`mod(6,0.4)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestDivideByZero(t *testing.T) {
	result, errs := f.Calculate(`2/0`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestPi(t *testing.T) {
	result, errs := f.Calculate(`pi()`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != math.Pi {
		t.Fatalf(`expected %v but got %v`, math.Pi, result)
	}
}

func TestRand(t *testing.T) {
	result, errs := f.Calculate(`rand()`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	value := result.(float64)
	if value < 0 || value > 1 {
		t.Fatalf(`expected a random value between 0 and 1 but got %v`, value)
	}
	t.Logf(`first value: %v`, result)
	result, _ = f.Calculate(`rand()`, nil)
	if result == value {
		t.Fatalf(`expected a second, different value but got the same`)
	}
	t.Logf(`second value: %v`, result)
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
