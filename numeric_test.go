package numeric


import (
	"testing"
)


func TestIsNumericTrue(t *testing.T) {
	tests := []struct{
		Rune rune
	}{
		{
			Rune: '0', // Digit Zero
		},
		{
			Rune: '1', // Digit One
		},
		{
			Rune: '2', // Digit Two
		},
		{
			Rune: '3', // Digit Three
		},
		{
			Rune: '4', // Digit Four
		},
		{
			Rune: '5', // Digit Five
		},
		{
			Rune: '6', // Digit Six
		},
		{
			Rune: '7', // Digit Seven
		},
		{
			Rune: '8', // Digit Eight
		},
		{
			Rune: '9', // Digit Nine
		},



		{
			Rune: '٠', // Arabic-Indic Digit Zero
		},
		{
			Rune: '١', // Arabic-Indic Digit One
		},
		{
			Rune: '٢', // Arabic-Indic Digit Two
		},
		{
			Rune: '۳', // Arabic-Indic Digit Three
		},
		{
			Rune: '٤', // Arabic-Indic Digit Four
		},
		{
			Rune: '٥', // Arabic-Indic Digit Five
		},
		{
			Rune: '٦', // Arabic-Indic Digit Six
		},
		{
			Rune: '٧', // Arabic-Indic Digit Seven
		},
		{
			Rune: '٨', // Arabic-Indic Digit Eight
		},
		{
			Rune: '٩', // Arabic-Indic Digit Nine
		},



		{
			Rune: '۰', // Extended Arabic-Indic Digit Zero
		},
		{
			Rune: '۱', // Extended Arabic-Indic Digit One
		},
		{
			Rune: '۲', // Extended Arabic-Indic Digit Two
		},
		{
			Rune: '۳', // Extended Arabic-Indic Digit Three
		},
		{
			Rune: '۴', // Extended Arabic-Indic Digit Four
		},
		{
			Rune: '۵', // Extended Arabic-Indic Digit Five
		},
		{
			Rune: '۶', // Extended Arabic-Indic Digit Six
		},
		{
			Rune: '۷', // Extended Arabic-Indic Digit Seven
		},
		{
			Rune: '۸', // Extended Arabic-Indic Digit Eight
		},
		{
			Rune: '۹', // Extended Arabic-Indic Digit Nine
		},



		{
			Rune: '⁰', // Superscript Zero
		},
		{
			Rune: '¹', // Superscript One
		},
		{
			Rune: '²', // Superscript Two
		},
		{
			Rune: '³', // Superscript Three
		},
		{
			Rune: '⁴', // Superscript Four
		},
		{
			Rune: '⁵', // Superscript Five
		},
		{
			Rune: '⁶', // Superscript Siz
		},
		{
			Rune: '⁷', // Superscript Seven
		},
		{
			Rune: '⁸', // Superscript Eight
		},
		{
			Rune: '⁹', // Superscript Nine
		},



		{
			Rune: '₀', // Subscript Zero
		},
		{
			Rune: '₁', // Subscript One
		},
		{
			Rune: '₂', // Subscript Two
		},
		{
			Rune: '₃', // Subscript Three
		},
		{
			Rune: '₄', // Subscript Four
		},
		{
			Rune: '₅', // Subscript Five
		},
		{
			Rune: '₆', // Subscript Six
		},
		{
			Rune: '₇', // Subscript Seven
		},
		{
			Rune: '₈', // Subscript Eight
		},
		{
			Rune: '₉', // Subscript Nine
		},



		{
			Rune: '¼', // Vulgar Fraction One Quarter
		},
		{
			Rune: '½', // Vulgar Fraction One Half
		},
		{
			Rune: '¾', // Vulgar Fraction Three Quarters
		},



		{
			Rune: '⅐', // Vulgar Fraction One Seventh
		},



		{
			Rune: '⅑', // Vulgar Fraction One Ninth
		},



		{
			Rune: '⅒', // Vulgar Fraction One Tenth
		},



		{
			Rune: '↉', // Vulgar Fraction Zero Thirds
		},
		{
			Rune: '⅓', // Vulgar Fraction One Thirds
		},
		{
			Rune: '⅔', // Vulgar Fraction Two Thirds
		},



		{
			Rune: '⅕', // Vulgar Fraction One Fifth
		},
		{
			Rune: '⅖', // Vulgar Fraction Two Fifth
		},
		{
			Rune: '⅗', // Vulgar Fraction Three Fifth
		},
		{
			Rune: '⅘', // Vulgar Fraction Four Fifth
		},



		{
			Rune: '⅙', // Vulgar Fraction One Sixth
		},
		{
			Rune: '⅚', // Vulgar Fraction Five Sixth
		},



		{
			Rune: '⅛', // Vulgar Fraction One Eighth
		},
		{
			Rune: '⅜', // Vulgar Fraction Three Eighth
		},
		{
			Rune: '⅝', // Vulgar Fraction Five Eighth
		},
		{
			Rune: '⅞', // Vulgar Fraction Seven Eighth
		},



		{
			Rune: '⅟', // Fraction Numerator One
		},



		{
			Rune: 'Ⅰ', // Roman Numeral One
		},
		{
			Rune: 'Ⅱ', // Roman Numeral Two
		},
		{
			Rune: 'Ⅲ', // Roman Numeral Three
		},
		{
			Rune: 'Ⅳ', // Roman Numeral Four
		},
		{
			Rune: 'Ⅴ', // Roman Numeral Five
		},
		{
			Rune: 'Ⅵ', // Roman Numeral Six
		},
		{
			Rune: 'Ⅶ', // Roman Numeral Seven
		},
		{
			Rune: 'Ⅷ', // Roman Numeral Eight
		},
		{
			Rune: 'Ⅸ', // Roman Numeral Nine
		},
		{
			Rune: 'Ⅹ', // Roman Numeral Ten
		},
		{
			Rune: 'Ⅺ', // Roman Numeral Eleven
		},
		{
			Rune: 'Ⅻ', // Roman Numeral Twelve
		},
		{
			Rune: 'Ⅼ', // Roman Numeral Fifty
		},
		{
			Rune: 'Ⅽ', // Roman Numeral One Hundred
		},
		{
			Rune: 'Ⅾ', // Roman Numeral Five Hundred
		},
		{
			Rune: 'Ⅿ', // Roman Numeral One Thousand
		},



		{
			Rune: 'ⅰ', // Small Roman Numeral One
		},
		{
			Rune: 'ⅱ', // Small Roman Numeral Two
		},
		{
			Rune: 'ⅲ', // Small Roman Numeral Three
		},
		{
			Rune: 'ⅳ', // Small Roman Numeral Four
		},
		{
			Rune: 'ⅴ', // Small Roman Numeral Five
		},
		{
			Rune: 'ⅵ', // Small Roman Numeral Six
		},
		{
			Rune: 'ⅶ', // Small Roman Numeral Seven
		},
		{
			Rune: 'ⅷ', // Small Roman Numeral Eight
		},
		{
			Rune: 'ⅸ', // Roman Numeral Nine
		},
		{
			Rune: 'ⅹ', // Small Roman Numeral Ten
		},
		{
			Rune: 'ⅺ', // Small Roman Numeral Eleven
		},
		{
			Rune: 'ⅻ', // Small Roman Numeral Twelve
		},
		{
			Rune: 'ⅼ', // Small Roman Numeral Fifty
		},
		{
			Rune: 'ⅽ', // Small Roman Numeral One Hundred
		},
		{
			Rune: 'ⅾ', // Small Roman Numeral Five Hundred
		},
		{
			Rune: 'ⅿ', // Small Roman Numeral One Thousand
		},



		{
			Rune: 'ↀ', // Roman Numeral One Thousand C D
		},
		{
			Rune: 'ↁ', // Roman Numeral Five Thousand
		},
		{
			Rune: 'ↂ', // Roman Numeral Ten Thousand
		},
		{
			Rune: 'Ↄ', // Roman Numeral Reversed One Hundred
		},
		{
			Rune: 'ↄ', // Latin Small Letter Reversed C
		},
		{
			Rune: 'ↅ', // Roman Numeral Six Late Form
		},
		{
			Rune: 'ↆ', // Roman Numeral Fifty Early Form
		},

		{
			Rune: 'ↇ', // Roman Numeral Fifty Thousand
		},
		{
			Rune: 'ↈ', // Roman Numeral One Hundred Thousand
		},



		{
			Rune: '↉', // Vulgar Fraction Zero Thirds
		},



		{
			Rune: '↊', // Turned Digit Two   (Duodecimal 10)
		},
		{
			Rune: '↋', // Turned Digit Three (Duodecimal 10)
		},



		{
			Rune: '〡', // Hangzhou Numeral One
		},
		{
			Rune: '〢', // Hangzhou Numeral Two
		},
		{
			Rune: '〣', // Hangzhou Numeral Three
		},
		{
			Rune: '〤', // Hangzhou Numeral Four
		},
		{
			Rune: '〥', // Hangzhou Numeral Five
		},
		{
			Rune: '〦', // Hangzhou Numeral Six
		},
		{
			Rune: '〧', // Hangzhou Numeral Seven
		},
		{
			Rune: '〨', // Hangzhou Numeral Eight
		},
		{
			Rune: '〩', // Hangzhou Numeral Nine
		},

		{
			Rune: '〸', // Hangzhou Numeral Ten
		},
		{
			Rune: '〹', // Hangzhou Numeral Twenty
		},

		{
			Rune: '〺', // Hangzhou Numeral Thirty
		},
	}


	for testNumber, test := range tests {
		if expected, actual := true, IsNumeric(test.Rune); expected != actual {
			t.Errorf("For test #%d, expected rune %q to be numeric (i.e., %t) but wasn't (i.e., %t).", testNumber, string(test.Rune), expected, actual)
			continue
		}
	}

}
