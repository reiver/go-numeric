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
			fmt.Printf("Although %s is numeric, it is not in the range of an int64.\n", r)
		default:
			fmt.Println("Some other error.")
		}
	}

	fmt.Printf("The int64 number is: %d.\n", i64)
}
