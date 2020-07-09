package alteryx_formulas

func boolVal(value bool) nullableBool {
	return nullableBool{value: value}
}

func nullBool() nullableBool {
	return nullableBool{isNull: true}
}
