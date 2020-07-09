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

type nullableBool struct {
	value  bool
	isNull bool
}

func (v nullableBool) Value() interface{} {
	if v.isNull {
		return nil
	}
	return v.value
}

type nullableString struct {
	value  string
	isNull bool
}

func (v nullableString) Value() interface{} {
	if v.isNull {
		return nil
	}
	return v.value
}
