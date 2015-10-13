# go-numeric

A library that provides helper functions to deal with runes that represent numeric values, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-numeric

[![GoDoc](https://godoc.org/github.com/reiver/go-numeric?status.svg)](https://godoc.org/github.com/reiver/go-numeric)


## Example

Here is an example using the numeric.Int64() func:

```
var r rune = '2'
//var r rune = '۲'
//var r rune = 'Ⅱ'
//var r rune = 'ⅱ'


i64, err := numeric.Int64(test.Rune)

if nil != err {
	switch err.(type) {
	case numeric.NotNumericComplainer:
		fmt.Println("The rune does not represent a numeric value.")
	case numeric.NotInRangeComplainer:
		fmt.Println("Although the rune represents a numeric value, it cannot be represents as an int64.")
	default:
		fmt.Println("This should never happen!")
	}
}

fmt.Printf("The numeric value that rune represents is: %d", i64)
```
Here is an example using the numeric.Frac64() func:

```
var r rune = '½'
//var r rune = '⅗'
//var r rune = 'Ⅵ'
//var r rune = 'ⅺ'


n64, d64, err := numeric.Frac64(test.Rune)

if nil != err {
	switch err.(type) {
	case numeric.NotNumericComplainer:
		fmt.Println("The rune does not represent a numeric value.")
	default:
		fmt.Println("This should never happen!")
	}
}

fmt.Printf("The numeric value fraction that rune represents has a numerator %d and denominator", n64, d64)
```
