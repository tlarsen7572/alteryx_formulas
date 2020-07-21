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

func TestRandInt(t *testing.T) {
	result, errs := f.Calculate(`randint(10)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	value := result.(float64)
	if value < 0 || value > 10 {
		t.Fatalf(`expected a random value between 0 and 10 but got %v`, value)
	}
	if math.Mod(value, 1.0) > 0 {
		t.Fatalf(`expected an integer but got a decimal`)
	}
	t.Logf(`value: %v`, result)
}

func TestRandIntDistribution(t *testing.T) {
	results := map[float64]int{
		0:  0,
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
	}

	for i := 0; i < 1000; i++ {
		result, _ := f.Calculate(`randInt(10)`, nil)
		results[result.(float64)]++
	}
	for key, value := range results {
		if value == 0 {
			t.Fatalf(`expecting all integers from 0 to 10 to have a count of at least 1 but %v had 0`, key)
		}
	}
	t.Logf(`%v`, results)
}

func TestRound(t *testing.T) {
	result, errs := f.Calculate(`round(55.34, 2)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 56.0 {
		t.Fatalf(`expected 56 but got %v`, result)
	}
	result, _ = f.Calculate(`round(41.1, 10)`, nil)
	if result != 40.0 {
		t.Fatalf(`expected 40 but got %v`, result)
	}
	result, _ = f.Calculate(`round(1.227, 0.01)`, nil)
	if result != 1.23 {
		t.Fatalf(`expected 1.23 but got %v`, result)
	}
	result, _ = f.Calculate(`round(-1.227, 0.01)`, nil)
	if result != -1.23 {
		t.Fatalf(`expected -1.23 but got %v`, result)
	}
}

func TestRoundToMultipleOfZero(t *testing.T) {
	result, errs := f.Calculate(`round(11, 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got: %v`, result)
	}
}

func TestRoundToNegativeMultiple(t *testing.T) {
	result, errs := f.Calculate(`round(10.2, -1)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 10.0 {
		t.Fatalf(`expected 10 but got: %v`, result)
	}
}

func TestSin(t *testing.T) {
	result, errs := f.Calculate(`round(sin(0.523599),0.001)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0.5 {
		t.Fatalf(`expected 0.5 but got %v`, result)
	}
}

func TestSinh(t *testing.T) {
	result, errs := f.Calculate(`round(sinh(1),0.0001)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 1.1752 {
		t.Fatalf(`expected 1.1752 but got %v`, result)
	}
}

func TestSqrt(t *testing.T) {
	result, errs := f.Calculate(`sqrt(100)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 10.0 {
		t.Fatalf(`expected 10.0 but got %v`, result)
	}
}

func TestTan(t *testing.T) {
	result, errs := f.Calculate(`round(tan(0.785398),0.001)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 1.0 {
		t.Fatalf(`expected 1 but got %v`, result)
	}
}

func TestTanh(t *testing.T) {
	result, errs := f.Calculate(`round(tanh(0.785398),0.0001)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0.6558 {
		t.Fatalf(`expected 0.6558 but got %v`, result)
	}
}

func TestSwitchNumberDefault(t *testing.T) {
	result, errs := f.Calculate(`switch(1,10,2,20,3,30,4,40)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 10.0 {
		t.Fatalf(`expected 10 but got %v`, result)
	}
}

func TestSwitchNumberNotDefault(t *testing.T) {
	result, errs := f.Calculate(`switch(3,10,2,20,3,30,4,40)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 30.0 {
		t.Fatalf(`expected 30 but got %v`, result)
	}
}

func TestSwitchStringDefault(t *testing.T) {
	result, errs := f.Calculate(`switch('a','aa','b','bb','c','cc')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `aa` {
		t.Fatalf(`expected 'aa' but got %v`, result)
	}
}

func TestCharFromInt(t *testing.T) {
	result, errs := f.Calculate(`charFromInt(66)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `B` {
		t.Fatalf(`expected 'B' but got '%v'`, result)
	}

	result, _ = f.Calculate(`charFromInt(127944)`, nil)
	expected := "\U0001f3c8"
	if result != expected {
		t.Fatalf(`expected '%v' but got '%v'`, expected, result)
	}
	t.Logf(`result: '%v'`, result)
}

func TestCharFromIntInvalidCodes(t *testing.T) {
	result, errs := f.Calculate(`charFromInt(55300)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got '%v'`, result)
	}

	result, _ = f.Calculate(`charFromInt(-1)`, nil)
	if result != nil {
		t.Fatalf(`expected nil but got '%v'`, result)
	}
}

func TestCharToInt(t *testing.T) {
	result, errs := f.Calculate(`charToInt('B')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 66.0 {
		t.Fatalf(`expected 66 but got %v`, result)
	}

	result, _ = f.Calculate(`charToInt('🏈')`, nil)
	if result != 127944.0 {
		t.Fatalf(`expected 127944 but got %v`, result)
	}
}

func TestHexToNumber(t *testing.T) {
	result, errs := f.Calculate(`hexToNumber('dd')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 221.0 {
		t.Fatalf(`expected 221 but got %v`, result)
	}
}

func TestContainsTrueCaseInsensitive(t *testing.T) {
	result, errs := f.Calculate(`contains("ABC","a")`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestContainsFalseCaseInsensitive(t *testing.T) {
	result, errs := f.Calculate(`contains("ABC","d")`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestContainsTrueCaseSensitive(t *testing.T) {
	result, errs := f.Calculate(`contains("ABC","A", 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestContainsFalseCaseSensitive(t *testing.T) {
	result, errs := f.Calculate(`contains("ABC","a", 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected true but got: %v`, result)
	}
}

func TestCountWords(t *testing.T) {
	result, errs := f.Calculate(`countWords('  a b c  ')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 words but got %v`, result)
	}
}

func TestCountWordsWithNewLines(t *testing.T) {
	result, errs := f.Calculate(`countWords('a 
b c')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 3.0 {
		t.Fatalf(`expected 3 words but got %v`, result)
	}
}

func TestEndsWithTrueCaseInsentitive(t *testing.T) {
	result, errs := f.Calculate(`endsWith('123ABC','abc')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestEndsWithFalseCaseInsentitive(t *testing.T) {
	result, errs := f.Calculate(`endsWith('123ABC','abcd')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestEndsWithTrueCaseSentitive(t *testing.T) {
	result, errs := f.Calculate(`endsWith('123ABC','ABC', 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestEndsWithFalseCaseSentitive(t *testing.T) {
	result, errs := f.Calculate(`endsWith('123ABC','abc', 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestEndsWithTargetLenExceedsStringLen(t *testing.T) {
	result, errs := f.Calculate(`endsWith('123ABC','abcdefghijklmnop', 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}
}

func TestFindStringTrue(t *testing.T) {
	result, errs := f.Calculate(`findString('abcabc', 'ab')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0 {
		t.Fatalf(`expected 0 but got %v`, result)
	}

	result, _ = f.Calculate(`findString('123abc', 'ab')`, nil)
	if result != 3 {
		t.Fatalf(`expected 3 but got %v`, result)
	}
}

func TestFindStringFalse(t *testing.T) {
	result, errs := f.Calculate(`findString('abcabc', 'AB')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != -1 {
		t.Fatalf(`expected -1 but got %v`, result)
	}
}

func TestGetWord(t *testing.T) {
	result, errs := f.Calculate(`getWord(' ab cde fghi ', 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `ab` {
		t.Fatalf(`expected 'ab' but got '%v'`, result)
	}
}

func TestGetWordOutOfBounds(t *testing.T) {
	result, errs := f.Calculate(`getWord(' ab cde fghi ', 4)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got '%v'`, result)
	}
}

func TestGetWordNoMatch(t *testing.T) {
	result, errs := f.Calculate(`getWord('  ', 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got '%v'`, result)
	}
}

func TestLeft(t *testing.T) {
	result, errs := f.Calculate(`left('abcdef', 3)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `abc` {
		t.Fatalf(`expected abc but got %v`, result)
	}
}

func TestLeftOutOfBounds(t *testing.T) {
	result, errs := f.Calculate(`left('abcdef', -1)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `abcdef` {
		t.Fatalf(`expected abcdef but got %v`, result)
	}

	result, _ = f.Calculate(`left('abcdef', 20)`, nil)
	if result != `abcdef` {
		t.Fatalf(`expected abcdef but got %v`, result)
	}
}

func TestLen(t *testing.T) {
	result, errs := f.Calculate(`length('abcdef')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 6.0 {
		t.Fatalf(`expected 6 but got %v`, result)
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
