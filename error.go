package numeric


var (
	errNotInRange = newNotInRangeComplainer()
	errNotNumeric = newNotNumericComplainer()
)


// NotInRangeComplainer represents the error situation where a rune that
// represents a numeric value is not in the range of the Go numeric
// type it is trying to be converted to.
//
// For example, although '¼' is numeric it cannot be converted to an
// int64 since int64 cannot represent fractional numbers.
type NotInRangeComplainer interface {
	error
	NotInRangeComplainer()
}

// NotNumericComplainer represents the error situation where a rune that
// does not represent a numeric value tries to be converted to a Go numeric
// type.
//
// For example, '&' cannot be converted to an int64 since it does not represent
// a numeric type.
type NotNumericComplainer interface {
	error
	NotNumericComplainer()
}


type internalNotInRangeComplainer struct{}
type internalNotNumericComplainer struct{}


func newNotInRangeComplainer() NotInRangeComplainer {
	complainer := internalNotInRangeComplainer{}

	return &complainer
}
func newNotNumericComplainer() NotNumericComplainer {
	complainer := internalNotNumericComplainer{}

	return &complainer
}


func (*internalNotInRangeComplainer) NotInRangeComplainer() {
	// Nothing here.
}
func (*internalNotNumericComplainer) NotNumericComplainer() {
	// Nothing here.
}


func (*internalNotInRangeComplainer) Error() string {
	return "Not In Range"
}
func (*internalNotNumericComplainer) Error() string {
	return "Not Numeric"
}

