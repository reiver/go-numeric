package numeric_test


import (
	"github.com/reiver/go-numeric"

	"fmt"
)


func ExampleNotInRangeComplainer() {

	r := '½'

	i64, err := numeric.Int64(r)
	if nil != err {
		switch err.(type) {
		case numeric.NotInRangeComplainer:
			fmt.Printf("Although %q is numeric, it is not in the range of an int64.\n", r)
			return
		default:
			fmt.Println("Some other error.")
			return
		}
	}

	fmt.Printf("The int64 number is: %d.\n", i64)

	// Output: Although '½' is numeric, it is not in the range of an int64.
}


func ExampleNotNumericComplainer() {

	r := '&'

	i64, err := numeric.Int64(r)
	if nil != err {
		switch err.(type) {
		case numeric.NotNumericComplainer:
			fmt.Printf("The rune %q is NOT numeric.\n", r)
			return
		default:
			fmt.Println("Some other error.")
			return
		}
	}

	fmt.Printf("The int64 number is: %d.\n", i64)

	// Output: The rune '&' is NOT numeric.
}
