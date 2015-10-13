package numeric


import (
	"testing"

	"fmt"
)


func TestInt64(t *testing.T) {
	tests := []struct{
		Rune     rune
		Expected int64
	}{
		{
			Rune:    '0', // Digit Zero
			Expected: 0,
		},
		{
			Rune:    '1', // Digit One
			Expected: 1,
		},
		{
			Rune:    '2', // Digit Two
			Expected: 2,
		},
		{
			Rune:    '3', // Digit Three
			Expected: 3,
		},
		{
			Rune:    '4', // Digit Four
			Expected: 4,
		},
		{
			Rune:    '5', // Digit Five
			Expected: 5,
		},
		{
			Rune:    '6', // Digit Six
			Expected: 6,
		},
		{
			Rune:    '7', // Digit Seven
			Expected: 7,
		},
		{
			Rune:    '8', // Digit Eight
			Expected: 8,
		},
		{
			Rune:    '9', // Digit Nine
			Expected: 9,
		},



		{
			Rune:    '٠', // Arabic-Indic Digit Zero
			Expected: 0,
		},
		{
			Rune:    '١', // Arabic-Indic Digit One
			Expected: 1,
		},
		{
			Rune:    '٢', // Arabic-Indic Digit Two
			Expected: 2,
		},
		{
			Rune:    '۳', // Arabic-Indic Digit Three
			Expected: 3,
		},
		{
			Rune:    '٤', // Arabic-Indic Digit Four
			Expected: 4,
		},
		{
			Rune:    '٥', // Arabic-Indic Digit Five
			Expected: 5,
		},
		{
			Rune:    '٦', // Arabic-Indic Digit Six
			Expected: 6,
		},
		{
			Rune:    '٧', // Arabic-Indic Digit Seven
			Expected: 7,
		},
		{
			Rune:    '٨', // Arabic-Indic Digit Eight
			Expected: 8,
		},
		{
			Rune:    '٩', // Arabic-Indic Digit Nine
			Expected: 9,
		},



		{
			Rune:    '۰', // Extended Arabic-Indic Digit Zero
			Expected: 0,
		},
		{
			Rune:    '۱', // Extended Arabic-Indic Digit One
			Expected: 1,
		},
		{
			Rune:    '۲', // Extended Arabic-Indic Digit Two
			Expected: 2,
		},
		{
			Rune:    '۳', // Extended Arabic-Indic Digit Three
			Expected: 3,
		},
		{
			Rune:    '۴', // Extended Arabic-Indic Digit Four
			Expected: 4,
		},
		{
			Rune:    '۵', // Extended Arabic-Indic Digit Five
			Expected: 5,
		},
		{
			Rune:    '۶', // Extended Arabic-Indic Digit Six
			Expected: 6,
		},
		{
			Rune:    '۷', // Extended Arabic-Indic Digit Seven
			Expected: 7,
		},
		{
			Rune:    '۸', // Extended Arabic-Indic Digit Eight
			Expected: 8,
		},
		{
			Rune:    '۹', // Extended Arabic-Indic Digit Nine
			Expected: 9,
		},



		{
			Rune:    '⁰', // Superscript Zero
			Expected: 0,
		},
		{
			Rune:    '¹', // Superscript One
			Expected: 1,
		},
		{
			Rune:    '²', // Superscript Two
			Expected: 2,
		},
		{
			Rune:    '³', // Superscript Three
			Expected: 3,
		},
		{
			Rune:    '⁴', // Superscript Four
			Expected: 4,
		},
		{
			Rune:    '⁵', // Superscript Five
			Expected: 5,
		},
		{
			Rune:    '⁶', // Superscript Six
			Expected: 6,
		},
		{
			Rune:    '⁷', // Superscript Seven
			Expected: 7,
		},
		{
			Rune:    '⁸', // Superscript Eight
			Expected: 8,
		},
		{
			Rune:    '⁹', // Superscript Nine
			Expected: 9,
		},



		{
			Rune:    '₀', // Subscript Zero
			Expected: 0,
		},
		{
			Rune:    '₁', // Subscript One
			Expected: 1,
		},
		{
			Rune:    '₂', // Subscript Two
			Expected: 2,
		},
		{
			Rune:    '₃', // Subscript Three
			Expected: 3,
		},
		{
			Rune:    '₄', // Subscript Four
			Expected: 4,
		},
		{
			Rune:    '₅', // Subscript Five
			Expected: 5,
		},
		{
			Rune:    '₆', // Subscript Six
			Expected: 6,
		},
		{
			Rune:    '₇', // Subscript Seven
			Expected: 7,
		},
		{
			Rune:    '₈', // Subscript Eight
			Expected: 8,
		},
		{
			Rune:    '₉', // Subscript Nine
			Expected: 9,
		},



		{
			Rune:    '⅟', // Fraction Numerator One
			Expected: 0,
		},



		{
			Rune:    'Ⅰ', // Roman Numeral One
			Expected: 1,
		},
		{
			Rune:    'Ⅱ', // Roman Numeral Two
			Expected: 2,
		},
		{
			Rune:    'Ⅲ', // Roman Numeral Three
			Expected: 3,
		},
		{
			Rune:    'Ⅳ', // Roman Numeral Four
			Expected: 4,
		},
		{
			Rune:    'Ⅴ', // Roman Numeral Five
			Expected: 5,
		},
		{
			Rune:    'Ⅵ', // Roman Numeral Six
			Expected: 6,
		},
		{
			Rune:    'Ⅶ', // Roman Numeral Seven
			Expected: 7,
		},
		{
			Rune:    'Ⅷ', // Roman Numeral Eight
			Expected: 8,
		},
		{
			Rune:    'Ⅸ', // Roman Numeral Nine
			Expected: 9,
		},
		{
			Rune:    'Ⅹ', // Roman Numeral Ten
			Expected: 10,
		},
		{
			Rune:    'Ⅺ', // Roman Numeral Eleven
			Expected: 11,
		},
		{
			Rune:    'Ⅻ', // Roman Numeral Twelve
			Expected: 12,
		},
		{
			Rune:    'Ⅼ', // Roman Numeral Fifty
			Expected: 50,
		},
		{
			Rune:    'Ⅽ', // Roman Numeral One Hundred
			Expected: 100,
		},
		{
			Rune:    'Ⅾ', // Roman Numeral Five Hundred
			Expected: 500,
		},
		{
			Rune:    'Ⅿ', // Roman Numeral One Thousand
			Expected: 1000,
		},



		{
			Rune:    'ⅰ', // Small Roman Numeral One
			Expected: 1,
		},
		{
			Rune:    'ⅱ', // Small Roman Numeral Two
			Expected: 2,
		},
		{
			Rune:    'ⅲ', // Small Roman Numeral Three
			Expected: 3,
		},
		{
			Rune:    'ⅳ', // Small Roman Numeral Four
			Expected: 4,
		},
		{
			Rune:    'ⅴ', // Small Roman Numeral Five
			Expected: 5,
		},
		{
			Rune:    'ⅵ', // Small Roman Numeral Six
			Expected: 6,
		},
		{
			Rune:    'ⅶ', // Small Roman Numeral Seven
			Expected: 7,
		},
		{
			Rune:    'ⅷ', // Small Roman Numeral Eight
			Expected: 8,
		},
		{
			Rune:    'ⅸ', // Roman Numeral Nine
			Expected: 9,
		},
		{
			Rune:    'ⅹ', // Small Roman Numeral Ten
			Expected: 10,
		},
		{
			Rune:    'ⅺ', // Small Roman Numeral Eleven
			Expected: 11,
		},
		{
			Rune:    'ⅻ', // Small Roman Numeral Twelve
			Expected: 12,
		},
		{
			Rune:    'ⅼ', // Small Roman Numeral Fifty
			Expected: 50,
		},
		{
			Rune:    'ⅽ', // Small Roman Numeral One Hundred
			Expected: 100,
		},
		{
			Rune:    'ⅾ', // Small Roman Numeral Five Hundred
			Expected: 500,
		},
		{
			Rune:    'ⅿ', // Small Roman Numeral One Thousand
			Expected: 1000,
		},



		{
			Rune:    'ↀ', // Roman Numeral One Thousand C D
			Expected: 1000,
		},
		{
			Rune:    'ↁ', // Roman Numeral Five Thousand
			Expected: 5000,
		},
		{
			Rune:    'ↂ', // Roman Numeral Ten Thousand
			Expected: 10000,
		},
		{
			Rune:    'Ↄ', // Roman Numeral Reversed One Hundred
			Expected: 100,
		},
		{
			Rune:    'ↄ', // Latin Small Letter Reversed C
			Expected: 100,
		},
		{
			Rune:    'ↅ', // Roman Numeral Six Late Form
			Expected: 6,
		},
		{
			Rune:    'ↆ', // Roman Numeral Fifty Early Form
			Expected: 50,
		},
		{
			Rune:    'ↇ', // Roman Numeral Fifty Thousand
			Expected: 50000,
		},
		{
			Rune:    'ↈ', // Roman Numeral One Hundred Thousand
			Expected: 100000,
		},



		{
			Rune:    '↉', // Vulgar Fraction Zero Thirds
			Expected: 0,
		},



		{
			Rune:    '↊', // Turned Digit Two   (Duodecimal 10)
			Expected: 10,
		},
		{
			Rune:    '↋', // Turned Digit Three (Duodecimal 10)
			Expected: 11,
		},
	}


	for testNumber, test := range tests {
		i64, err := Int64(test.Rune)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but received one: %v; for rune %q.", testNumber, err, string(test.Rune))
			continue
		}

		if expected, actual := test.Expected, i64; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for rune %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
	}
}



func TestInt64ErrNotInRange(t *testing.T) {
	tests := []struct{
		Rune rune
		ExpectedErr error
	}{
		{
			Rune: '¼',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '½',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '¾',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅐',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅑',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅒',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅓',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅔',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅕',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅖',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅗',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅘',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅙',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅚',
			ExpectedErr: errNotInRange,
		},

		{
			Rune: '⅛',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅜',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅝',
			ExpectedErr: errNotInRange,
		},
		{
			Rune: '⅞',
			ExpectedErr: errNotInRange,
		},
	}


	for testNumber, test := range tests {
		_, err := Int64(test.Rune)

		if expected, actual := fmt.Sprintf("%T", test.ExpectedErr), fmt.Sprintf("%T", err); expected != actual {
			t.Errorf("For test #%d, expected a %q error, but actually got a %q error; for rune %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
	}
}


func TestInt64ErrNotNumeric(t *testing.T) {
	tests := []struct{
		Rune rune
		ExpectedErr error
	}{
		{
			Rune: '\u0000',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0001',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0002',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0003',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0004',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0005',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0006',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0007',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0008',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0009',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000A',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000B',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000C',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000D',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000E',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u000F',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0010',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0011',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0012',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0013',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0014',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0015',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0016',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0017',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0018',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u0019',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001A',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001B',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001C',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001D',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001E',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\u001F',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: ' ',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '!',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '"',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '#',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '$',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '%',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '&',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '\'',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '(',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: ')',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '*',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '+',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: ',',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '-',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '.',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '/',
			ExpectedErr: errNotNumeric,
		},




		{
			Rune: '😀',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '😁',
			ExpectedErr: errNotNumeric,
		},
		{
			Rune: '😂',
			ExpectedErr: errNotNumeric,
		},

	}


	for testNumber, test := range tests {
		_, err := Int64(test.Rune)

		if expected, actual := fmt.Sprintf("%T", test.ExpectedErr), fmt.Sprintf("%T", err); expected != actual {
			t.Errorf("For test %d, expected a %q error, but actually got a %q error; for rune %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
	}
}
