package alteryx_formulas

type NullableValue interface {
	Value() interface{}
}

type nullableNum struct {
	value  float64
	isNull bool
}

func (v nullableNum) Value() interface{} {
	if v.isNull {
		return nil
	}
	return v.value
}
