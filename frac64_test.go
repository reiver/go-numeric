package numeric


import (
	"testing"
)


func TestFrac64(t *testing.T) {
	tests := []struct{
		Rune                rune
		ExpectedNumerator   int64
		ExpectedDenominator int64
	}{
		{
			Rune:               '0', // Digit Zero
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '1', // Digit One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '2', // Digit Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '3', // Digit Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '4', // Digit Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '5', // Digit Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '6', // Digit Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '7', // Digit Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '8', // Digit Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '9', // Digit Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '٠', // Arabic-Indic Digit Zero
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '١', // Arabic-Indic Digit One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٢', // Arabic-Indic Digit Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۳', // Arabic-Indic Digit Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٤', // Arabic-Indic Digit Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٥', // Arabic-Indic Digit Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٦', // Arabic-Indic Digit Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٧', // Arabic-Indic Digit Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٨', // Arabic-Indic Digit Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '٩', // Arabic-Indic Digit Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '۰', // Extended Arabic-Indic Digit Zero
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۱', // Extended Arabic-Indic Digit One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۲', // Extended Arabic-Indic Digit Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۳', // Extended Arabic-Indic Digit Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۴', // Extended Arabic-Indic Digit Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۵', // Extended Arabic-Indic Digit Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۶', // Extended Arabic-Indic Digit Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۷', // Extended Arabic-Indic Digit Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۸', // Extended Arabic-Indic Digit Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '۹', // Extended Arabic-Indic Digit Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '⁰', // Superscript Zero
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '¹', // Superscript One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '²', // Superscript Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '³', // Superscript Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁴', // Superscript Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁵', // Superscript Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁶', // Superscript Siz
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁷', // Superscript Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁸', // Superscript Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⁹', // Superscript Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '₀', // Subscript Zero
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₁', // Subscript One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₂', // Subscript Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₃', // Subscript Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₄', // Subscript Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₅', // Subscript Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₆', // Subscript Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₇', // Subscript Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₈', // Subscript Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '₉', // Subscript Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '¼', // Vulgar Fraction One Quarter
			ExpectedNumerator:   1,
			ExpectedDenominator: 4,
		},
		{
			Rune:               '½', // Vulgar Fraction One Half
			ExpectedNumerator:   1,
			ExpectedDenominator: 2,
		},
		{
			Rune:               '¾', // Vulgar Fraction Three Quarters
			ExpectedNumerator:   3,
			ExpectedDenominator: 4,
		},



		{
			Rune:               '⅐', // Vulgar Fraction One Seventh
			ExpectedNumerator:   1,
			ExpectedDenominator: 7,
		},



		{
			Rune:               '⅑', // Vulgar Fraction One Ninth
			ExpectedNumerator:   1,
			ExpectedDenominator: 9,
		},



		{
			Rune:               '⅒', // Vulgar Fraction One Tenth
			ExpectedNumerator:   1,
			ExpectedDenominator: 10,
		},



		{
			Rune:               '↉', // Vulgar Fraction Zero Thirds
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '⅓', // Vulgar Fraction One Thirds
			ExpectedNumerator:   1,
			ExpectedDenominator: 3,
		},
		{
			Rune:               '⅔', // Vulgar Fraction Two Thirds
			ExpectedNumerator:   2,
			ExpectedDenominator: 3,
		},



		{
			Rune:               '⅕', // Vulgar Fraction One Fifth
			ExpectedNumerator:   1,
			ExpectedDenominator: 5,
		},
		{
			Rune:               '⅖', // Vulgar Fraction Two Fifth
			ExpectedNumerator:   2,
			ExpectedDenominator: 5,
		},
		{
			Rune:               '⅗', // Vulgar Fraction Three Fifth
			ExpectedNumerator:   3,
			ExpectedDenominator: 5,
		},
		{
			Rune:               '⅘', // Vulgar Fraction Four Fifth
			ExpectedNumerator:   4,
			ExpectedDenominator: 5,
		},



		{
			Rune:               '⅙', // Vulgar Fraction One Sixth
			ExpectedNumerator:   1,
			ExpectedDenominator: 6,
		},
		{
			Rune:               '⅚', // Vulgar Fraction Five Sixth
			ExpectedNumerator:   5,
			ExpectedDenominator: 6,
		},



		{
			Rune:               '⅛', // Vulgar Fraction One Eighth
			ExpectedNumerator:   1,
			ExpectedDenominator: 8,
		},
		{
			Rune:               '⅜', // Vulgar Fraction Three Eighth
			ExpectedNumerator:   3,
			ExpectedDenominator: 8,
		},
		{
			Rune:               '⅝', // Vulgar Fraction Five Eighth
			ExpectedNumerator:   5,
			ExpectedDenominator: 8,
		},
		{
			Rune:               '⅞', // Vulgar Fraction Seven Eighth
			ExpectedNumerator:   7,
			ExpectedDenominator: 8,
		},



		{
			Rune:               '⅟', // Fraction Numerator One
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},



		{
			Rune:               'Ⅰ', // Roman Numeral One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅱ', // Roman Numeral Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅲ', // Roman Numeral Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅳ', // Roman Numeral Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅴ', // Roman Numeral Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅵ', // Roman Numeral Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅶ', // Roman Numeral Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅷ', // Roman Numeral Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅸ', // Roman Numeral Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅹ', // Roman Numeral Ten
			ExpectedNumerator:   10,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅺ', // Roman Numeral Eleven
			ExpectedNumerator:   11,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅻ', // Roman Numeral Twelve
			ExpectedNumerator:   12,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅼ', // Roman Numeral Fifty
			ExpectedNumerator:   50,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅽ', // Roman Numeral One Hundred
			ExpectedNumerator:   100,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅾ', // Roman Numeral Five Hundred
			ExpectedNumerator:   500,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ⅿ', // Roman Numeral One Thousand
			ExpectedNumerator:   1000,
			ExpectedDenominator: 1,
		},



		{
			Rune:               'ⅰ', // Small Roman Numeral One
			ExpectedNumerator:   1,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅱ', // Small Roman Numeral Two
			ExpectedNumerator:   2,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅲ', // Small Roman Numeral Three
			ExpectedNumerator:   3,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅳ', // Small Roman Numeral Four
			ExpectedNumerator:   4,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅴ', // Small Roman Numeral Five
			ExpectedNumerator:   5,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅵ', // Small Roman Numeral Six
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅶ', // Small Roman Numeral Seven
			ExpectedNumerator:   7,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅷ', // Small Roman Numeral Eight
			ExpectedNumerator:   8,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅸ', // Roman Numeral Nine
			ExpectedNumerator:   9,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅹ', // Small Roman Numeral Ten
			ExpectedNumerator:   10,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅺ', // Small Roman Numeral Eleven
			ExpectedNumerator:   11,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅻ', // Small Roman Numeral Twelve
			ExpectedNumerator:   12,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅼ', // Small Roman Numeral Fifty
			ExpectedNumerator:   50,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅽ', // Small Roman Numeral One Hundred
			ExpectedNumerator:   100,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅾ', // Small Roman Numeral Five Hundred
			ExpectedNumerator:   500,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ⅿ', // Small Roman Numeral One Thousand
			ExpectedNumerator:   1000,
			ExpectedDenominator: 1,
		},



		{
			Rune:               'ↀ', // Roman Numeral One Thousand C D
			ExpectedNumerator:   1000,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↁ', // Roman Numeral Five Thousand
			ExpectedNumerator:   5000,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↂ', // Roman Numeral Ten Thousand
			ExpectedNumerator:   10000,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'Ↄ', // Roman Numeral Reversed One Hundred
			ExpectedNumerator:   100,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↄ', // Latin Small Letter Reversed C
			ExpectedNumerator:   100,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↅ', // Roman Numeral Six Late Form
			ExpectedNumerator:   6,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↆ', // Roman Numeral Fifty Early Form
			ExpectedNumerator:   50,
			ExpectedDenominator: 1,
		},

		{
			Rune:               'ↇ', // Roman Numeral Fifty Thousand
			ExpectedNumerator:   50000,
			ExpectedDenominator: 1,
		},
		{
			Rune:               'ↈ', // Roman Numeral One Hundred Thousand
			ExpectedNumerator:   100000,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '↉', // Vulgar Fraction Zero Thirds
			ExpectedNumerator:   0,
			ExpectedDenominator: 1,
		},



		{
			Rune:               '↊', // Turned Digit Two   (Duodecimal 10)
			ExpectedNumerator:   10,
			ExpectedDenominator: 1,
		},
		{
			Rune:               '↋', // Turned Digit Three (Duodecimal 10)
			ExpectedNumerator:   11,
			ExpectedDenominator: 1,
		},

	}


	for testNumber, test := range tests {
		n64, d64, err := Frac64(test.Rune)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but received one: %v; for rune %q.", testNumber, err, string(test.Rune))
			continue
		}

		if expected, actual := test.ExpectedNumerator, n64; expected != actual {
			t.Errorf("For test #%d, expected numerator to be %d, but actually got %d; for rune %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}

		if expected, actual := test.ExpectedDenominator, d64; expected != actual {
			t.Errorf("For test #%d, expected denominator to be %d, but actually got %d; for rune %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
	}

}
