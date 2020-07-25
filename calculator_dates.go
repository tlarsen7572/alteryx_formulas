package alteryx_formulas

import (
	"fmt"
	"time"
)

func (calc *calculator) toDate() {
	expr := calc.popValue()
	switch value := expr.(type) {
	case string:
		dateValue, err := time.Parse(`2006-01-02`, value)
		if err != nil {
			calc.pushValue(nil)
			calc.errs = append(calc.errs, fmt.Errorf(`'%v' is not a valid date: expecting format yyyy-MM-dd`, value))
			return
		}
		calc.pushValue(dateValue)
	case float64:
		addHours := time.Duration(int(value) * 24)
		dateValue := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(addHours * time.Hour)
		calc.pushValue(dateValue)
	case time.Time:
		calc.pushValue(value.Truncate(24 * time.Hour))
	default:
		calc.pushValue(nil)
		calc.errs = append(calc.errs, fmt.Errorf(`%v has incorrect type: ToDate can only convert strings, numbers, or datetimes`, value))
	}
}

func (calc *calculator) toDatetime() {
	expr := calc.popValue()
	switch value := expr.(type) {
	case string:
		dateValue, err := time.Parse(`2006-01-02 15:04:05`, value)
		if err != nil {
			calc.pushValue(nil)
			calc.errs = append(calc.errs, fmt.Errorf(`'%v' is not a valid datetime: expecting format yyyy-MM-dd hh:mm:ss`, value))
			return
		}
		calc.pushValue(dateValue)
	case float64:
		const secondsPerDay = 86400
		addSeconds := time.Duration(value * secondsPerDay)
		dateValue := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC).Add(addSeconds * time.Second)
		calc.pushValue(dateValue)
	case time.Time:
		calc.pushValue(value.Truncate(time.Second))
	default:
		calc.pushValue(nil)
		calc.errs = append(calc.errs, fmt.Errorf(`%v has incorrect type: ToDatetime can only convert strings, numbers, or datetimes`, value))
	}
}
