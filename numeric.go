// Package numeric provides helper functions to deal with numeric data; and accounts for all UNICODE numeric characters.
package numeric

const (
	DigitZero  = '\u0030' // Digit Zero  (0)
	DigitOne   = '\u0031' // Digit One   (1)
	DigitTwo   = '\u0032' // Digit Two   (2)
	DigitThree = '\u0033' // Digit Three (3)
	DigitFour  = '\u0034' // Digit Four  (4)
	DigitFive  = '\u0035' // Digit Five  (5)
	DigitSix   = '\u0036' // Digit Six   (6)
	DigitSeven = '\u0037' // Digit Seven (7)
	DigitEight = '\u0038' // Digit Eight (8)
	DigitNine  = '\u0039' // Digit Nine  (9)

	ArabicIndicDigitZero  = '\u0660' // Arabic-Indic Digit Zero  (٠)
	ArabicIndicDigitOne   = '\u0661' // Arabic-Indic Digit One   (١)
	ArabicIndicDigitTwo   = '\u0662' // Arabic-Indic Digit Two   (٢)
	ArabicIndicDigitThree = '\u0663' // Arabic-Indic Digit Three (٣)
	ArabicIndicDigitFour  = '\u0664' // Arabic-Indic Digit Four  (٤)
	ArabicIndicDigitFive  = '\u0665' // Arabic-Indic Digit Five  (٥)
	ArabicIndicDigitSix   = '\u0666' // Arabic-Indic Digit Six   (٦)
	ArabicIndicDigitSeven = '\u0667' // Arabic-Indic Digit Seven (٧)
	ArabicIndicDigitEight = '\u0668' // Arabic-Indic Digit Eight (٨)
	ArabicIndicDigitNine  = '\u0669' // Arabic-Indic Digit Nine  (٩)

	ExtendedArabicIndicDigitZero  = '\u06F0' // Extended Arabic-Indic Digit Zero  (۰)
	ExtendedArabicIndicDigitOne   = '\u06F1' // Extended Arabic-Indic Digit One   (۱)
	ExtendedArabicIndicDigitTwo   = '\u06F2' // Extended Arabic-Indic Digit Two   (۲)
	ExtendedArabicIndicDigitThree = '\u06F3' // Extended Arabic-Indic Digit Three (۳)
	ExtendedArabicIndicDigitFour  = '\u06F4' // Extended Arabic-Indic Digit Four  (۴)
	ExtendedArabicIndicDigitFive  = '\u06F5' // Extended Arabic-Indic Digit Five  (۵)
	ExtendedArabicIndicDigitSix   = '\u06F6' // Extended Arabic-Indic Digit Six   (۶)
	ExtendedArabicIndicDigitSeven = '\u06F7' // Extended Arabic-Indic Digit Seven (۷)
	ExtendedArabicIndicDigitEight = '\u06F8' // Extended Arabic-Indic Digit Eight (۸)
	ExtendedArabicIndicDigitNine  = '\u06F9' // Extended Arabic-Indic Digit Nine  (۹)

	SuperscriptZero  = '\u2070' // Superscript Zero  (⁰)
	SuperscriptOne   = '\u00B9' // Superscript One   (¹)
	SuperscriptTwo   = '\u00B2' // Superscript Two   (²)
	SuperscriptThree = '\u00B3' // Superscript Three (³)
	SuperscriptFour  = '\u2074' // Superscript Four  (⁴)
	SuperscriptFive  = '\u2075' // Superscript Five  (⁵)
	SuperscriptSix   = '\u2076' // Superscript Six   (⁶)
	SuperscriptSeven = '\u2077' // Superscript Seven (⁷)
	SuperscriptEight = '\u2078' // Superscript Eight (⁸)
	SuperscriptNine  = '\u2079' // Superscript Nine  (⁹)

	SubscriptZero  = '\u2080' // Subscript Zero  (₀)
	SubscriptOne   = '\u2081' // Subscript One   (₁)
	SubscriptTwo   = '\u2082' // Subscript Two   (₂)
	SubscriptThree = '\u2083' // Subscript Three (₃)
	SubscriptFour  = '\u2084' // Subscript Four  (₄)
	SubscriptFive  = '\u2085' // Subscript Five  (₅)
	SubscriptSix   = '\u2086' // Subscript Six   (₆)
	SubscriptSeven = '\u2087' // Subscript Seven (₇)
	SubscriptEight = '\u2088' // Subscript Eight (₈)
	SubscriptNine  = '\u2089' // Subscript Nine  (₉)

	VulgarFractionOneQuarter    = '\u00BC' // Vulgar Fraction One Quarter    (¼)
	VulgarFractionOneHalf       = '\u00BD' // Vulgar Fraction One Half       (½)
	VulgarFractionThreeQuarters = '\u00BE' // Vulgar Fraction Three Quarters (¾)

	VulgarFractionOneSeventh    = '\u2150' // Vulgar Fraction One Seventh    (⅐)

	VulgarFractionOneNinth      = '\u2151' // Vulgar Fraction One Ninth      (⅑)

	VulgarFractionOneTenth      = '\u2152' // Vulgar Fraction One Tenth      (⅒)

	VulgarFractionZeroThirds    = '\u2189' // Vulgar Fraction Zero Thirds    (↉)
	VulgarFractionOneThird      = '\u2153' // Vulgar Fraction One Third      (⅓)
	VulgarFractionTwoThirds     = '\u2154' // Vulgar Fraction Two Thirds     (⅔)

	VulgarFractionOneFifth      = '\u2155' // Vulgar Fraction One Fifth      (⅕)
	VulgarFractionTwoFifths     = '\u2156' // Vulgar Fraction Two Fifths     (⅖)
	VulgarFractionThreeFifths   = '\u2157' // Vulgar Fraction Three Fifths   (⅗)
	VulgarFractionFourFifths    = '\u2158' // Vulgar Fraction Four Fifths    (⅘)

	VulgarFractionOneSixth      = '\u2159' // Vulgar Fraction One Sixth      (⅙)
	VulgarFractionFiveSixths    = '\u215A' // Vulgar Fraction Five Sixths    (⅚)

	VulgarFractionOneEighth     = '\u215B' // Vulgar Fraction One Eighth     (⅛)
	VulgarFractionThreeEighths  = '\u215C' // Vulgar Fraction Three Eighths  (⅜)
	VulgarFractionFiveEighths   = '\u215D' // Vulgar Fraction Five Eighths   (⅝)
	VulgarFractionSevenEighths  = '\u215E' // Vulgar Fraction Seven Eighths  (⅞)

	FractionNumeratorOne        = '\u215F' // Fraction Numerator One         (⅟)

	RomanNumeralOne         = '\u2160' // Roman Numeral One          (Ⅰ)
	RomanNumeralTwo         = '\u2161' // Roman Numeral Two          (Ⅱ)
	RomanNumeralThree       = '\u2162' // Roman Numeral Three        (Ⅲ)
	RomanNumeralFour        = '\u2163' // Roman Numeral Four         (Ⅳ)
	RomanNumeralFive        = '\u2164' // Roman Numeral Five         (Ⅴ)
	RomanNumeralSix         = '\u2165' // Roman Numeral Six          (Ⅵ)
	RomanNumeralSeven       = '\u2166' // Roman Numeral Seven        (Ⅶ)
	RomanNumeralEight       = '\u2167' // Roman Numeral Eight        (Ⅷ)
	RomanNumeralNine        = '\u2168' // Roman Numeral Nine         (Ⅸ)
	RomanNumeralTen         = '\u2169' // Roman Numeral Ten          (Ⅹ)
	RomanNumeralEleven      = '\u216A' // Roman Numeral Eleven       (Ⅺ)
	RomanNumeralTwelve      = '\u216B' // Roman Numeral Twelve       (Ⅻ)
	RomanNumeralFifty       = '\u216C' // Roman Numeral Fifty        (Ⅼ)
	RomanNumeralOneHundred  = '\u216D' // Roman Numeral One Hundred  (Ⅽ)
	RomanNumeralFiveHundred = '\u216E' // Roman Numeral Five Hundred (Ⅾ)
	RomanNumeralOneThousand = '\u216F' // Roman Numeral One Thousand (Ⅿ)

	SmallRomanNumeralOne         = '\u2170' // Small Roman Numeral One          (ⅰ)
	SmallRomanNumeralTwo         = '\u2171' // Small Roman Numeral Two          (ⅱ) 
	SmallRomanNumeralThree       = '\u2172' // Small Roman Numeral Three        (ⅲ)
	SmallRomanNumeralFour        = '\u2173' // Small Roman Numeral Four         (ⅳ)
	SmallRomanNumeralFive        = '\u2174' // Small Roman Numeral Five         (ⅴ)
	SmallRomanNumeralSix         = '\u2175' // Small Roman Numeral Six          (ⅵ)
	SmallRomanNumeralSeven       = '\u2176' // Small Roman Numeral Seven        (ⅶ)
	SmallRomanNumeralEight       = '\u2177' // Small Roman Numeral Eight        (ⅷ)
	SmallRomanNumeralNine        = '\u2178' // Small Roman Numeral Nine         (ⅸ)
	SmallRomanNumeralTen         = '\u2179' // Small Roman Numeral Ten          (ⅹ)
	SmallRomanNumeralEleven      = '\u217A' // Small Roman Numeral Eleven       (ⅺ)
	SmallRomanNumeralTwelve      = '\u217B' // Small Roman Numeral Twelve       (ⅻ)
	SmallRomanNumeralFifty       = '\u217C' // Small Roman Numeral Fifty        (ⅼ)
	SmallRomanNumeralOneHundred  = '\u217D' // Small Roman Numeral One Hundred  (ⅽ)
	SmallRomanNumeralFiveHundred = '\u217E' // Small Roman Numeral Five Hundred (ⅾ)
	SmallRomanNumeralOneThousand = '\u217F' // Small Roman Numeral One Thousand (ⅿ)

	RomanNumeralOneThousandCD      = '\u2180' // Roman Numeral One Thousand C D     (ↀ)
	RomanNumeralFiveThousand       = '\u2181' // Roman Numeral Five Thousand        (ↁ)
	RomanNumeralTenThousand        = '\u2182' // Roman Numeral Ten Thousand         (ↂ)
	RomanNumeralReservedOneHundred = '\u2183' // Roman Numeral Reversed One Hundred (Ↄ)
	LatinSmallLetterReversedC      = '\u2184' // Latin Small Letter Reversed C      (ↄ)
	RomanNumeralSixLateForm        = '\u2185' // Roman Numeral Six Late Form        (ↅ)
	RomanNumeralFiftyEarlyForm     = '\u2186' // Roman Numeral Fifty Early Form     (ↆ)
	RomanNumeralFiftyThousand      = '\u2187' // Roman Numeral Fifty Thousand       (ↇ)
	RomanNumeralOneHundredThousand = '\u2188' // Roman Numeral One Hundred Thousand (ↈ)


	TurnedDigitTwo   = '\u218A' // Turned Digit Two   (Duodecimal 10)
	TurnedDigitThree = '\u218B' // Turned Digit Three (Duodecimal 11)	
)



// IsNumeric return true it the value the rune represents is numeric, else returns false.
//
// For example, it would return true for '2', '⅗' and 'Ⅶ'; but would return false for '&', '😀' and '?'.
func IsNumeric(r rune) bool {

	result := false



	switch (r) {
		case
			DigitZero,
			DigitOne,
			DigitTwo,
			DigitThree,
			DigitFour,
			DigitFive,
			DigitSix,
			DigitSeven,
			DigitEight,
			DigitNine,

			ArabicIndicDigitZero,
			ArabicIndicDigitOne,
			ArabicIndicDigitTwo,
			ArabicIndicDigitThree,
			ArabicIndicDigitFour,
			ArabicIndicDigitFive,
			ArabicIndicDigitSix,
			ArabicIndicDigitSeven,
			ArabicIndicDigitEight,
			ArabicIndicDigitNine,

			ExtendedArabicIndicDigitZero,
			ExtendedArabicIndicDigitOne,
			ExtendedArabicIndicDigitTwo,
			ExtendedArabicIndicDigitThree,
			ExtendedArabicIndicDigitFour,
			ExtendedArabicIndicDigitFive,
			ExtendedArabicIndicDigitSix,
			ExtendedArabicIndicDigitSeven,
			ExtendedArabicIndicDigitEight,
			ExtendedArabicIndicDigitNine,

			SuperscriptZero,
			SuperscriptOne,
			SuperscriptTwo,
			SuperscriptThree,
			SuperscriptFour,
			SuperscriptFive,
			SuperscriptSix,
			SuperscriptSeven,
			SuperscriptEight,
			SuperscriptNine,

			SubscriptZero,
			SubscriptOne,
			SubscriptTwo,
			SubscriptThree,
			SubscriptFour,
			SubscriptFive,
			SubscriptSix,
			SubscriptSeven,
			SubscriptEight,
			SubscriptNine,

			VulgarFractionOneQuarter,
			VulgarFractionOneHalf,
			VulgarFractionThreeQuarters,

			VulgarFractionOneSeventh,

			VulgarFractionOneNinth,

			VulgarFractionOneTenth,

			VulgarFractionZeroThirds,
			VulgarFractionOneThird,
			VulgarFractionTwoThirds,

			VulgarFractionOneFifth,
			VulgarFractionTwoFifths,
			VulgarFractionThreeFifths,
			VulgarFractionFourFifths,

			VulgarFractionOneSixth,
			VulgarFractionFiveSixths,

			VulgarFractionOneEighth,
			VulgarFractionThreeEighths,
			VulgarFractionFiveEighths,
			VulgarFractionSevenEighths,

			FractionNumeratorOne,

			RomanNumeralOne,
			RomanNumeralTwo,
			RomanNumeralThree,
			RomanNumeralFour,
			RomanNumeralFive,
			RomanNumeralSix,
			RomanNumeralSeven,
			RomanNumeralEight,
			RomanNumeralNine,
			RomanNumeralTen,
			RomanNumeralEleven,
			RomanNumeralTwelve,
			RomanNumeralFifty,
			RomanNumeralOneHundred,
			RomanNumeralFiveHundred,
			RomanNumeralOneThousand,

			SmallRomanNumeralOne,
			SmallRomanNumeralTwo,
			SmallRomanNumeralThree,
			SmallRomanNumeralFour,
			SmallRomanNumeralFive,
			SmallRomanNumeralSix,
			SmallRomanNumeralSeven,
			SmallRomanNumeralEight,
			SmallRomanNumeralNine,
			SmallRomanNumeralTen,
			SmallRomanNumeralEleven,
			SmallRomanNumeralTwelve,
			SmallRomanNumeralFifty,
			SmallRomanNumeralOneHundred,
			SmallRomanNumeralFiveHundred,
			SmallRomanNumeralOneThousand,


			RomanNumeralOneThousandCD,
			RomanNumeralFiveThousand,
			RomanNumeralTenThousand,
			RomanNumeralReservedOneHundred,

			LatinSmallLetterReversedC,
			RomanNumeralSixLateForm,
			RomanNumeralFiftyEarlyForm,
			RomanNumeralFiftyThousand,
			RomanNumeralOneHundredThousand,

			TurnedDigitTwo,
			TurnedDigitThree:

			result = true
		default:
			result = false
	}



	return result
}
