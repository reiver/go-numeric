package numeric


var (
	runeToInt64LookupTable = map[rune]int64{
		DigitZero  : 0,
		DigitOne   : 1,
		DigitTwo   : 2,
		DigitThree : 3,
		DigitFour  : 4,
		DigitFive  : 5,
		DigitSix   : 6,
		DigitSeven : 7,
		DigitEight : 8,
		DigitNine  : 9,

		ArabicIndicDigitZero  : 0,
		ArabicIndicDigitOne   : 1,
		ArabicIndicDigitTwo   : 2,
		ArabicIndicDigitThree : 3,
		ArabicIndicDigitFour  : 4,
		ArabicIndicDigitFive  : 5,
		ArabicIndicDigitSix   : 6,
		ArabicIndicDigitSeven : 7,
		ArabicIndicDigitEight : 8,
		ArabicIndicDigitNine  : 9,

		ExtendedArabicIndicDigitZero  : 0,
		ExtendedArabicIndicDigitOne   : 1,
		ExtendedArabicIndicDigitTwo   : 2,
		ExtendedArabicIndicDigitThree : 3,
		ExtendedArabicIndicDigitFour  : 4,
		ExtendedArabicIndicDigitFive  : 5,
		ExtendedArabicIndicDigitSix   : 6,
		ExtendedArabicIndicDigitSeven : 7,
		ExtendedArabicIndicDigitEight : 8,
		ExtendedArabicIndicDigitNine  : 9,

		SuperscriptZero  : 0,
		SuperscriptOne   : 1,
		SuperscriptTwo   : 2,
		SuperscriptThree : 3,
		SuperscriptFour  : 4,
		SuperscriptFive  : 5,
		SuperscriptSix   : 6,
		SuperscriptSeven : 7,
		SuperscriptEight : 8,
		SuperscriptNine  : 9,

		SubscriptZero  : 0,
		SubscriptOne   : 1,
		SubscriptTwo   : 2,
		SubscriptThree : 3,
		SubscriptFour  : 4,
		SubscriptFive  : 5,
		SubscriptSix   : 6,
		SubscriptSeven : 7,
		SubscriptEight : 8,
		SubscriptNine  : 9,

		FractionNumeratorOne : 0,

		RomanNumeralOne         : 1,
		RomanNumeralTwo         : 2,
		RomanNumeralThree       : 3,
		RomanNumeralFour        : 4,
		RomanNumeralFive        : 5,
		RomanNumeralSix         : 6,
		RomanNumeralSeven       : 7,
		RomanNumeralEight       : 8,
		RomanNumeralNine        : 9,
		RomanNumeralTen         : 10,
		RomanNumeralEleven      : 11,
		RomanNumeralTwelve      : 12,
		RomanNumeralFifty       : 50,
		RomanNumeralOneHundred  : 100,
		RomanNumeralFiveHundred : 500,
		RomanNumeralOneThousand : 1000,

		SmallRomanNumeralOne         : 1,
		SmallRomanNumeralTwo         : 2,
		SmallRomanNumeralThree       : 3,
		SmallRomanNumeralFour        : 4,
		SmallRomanNumeralFive        : 5,
		SmallRomanNumeralSix         : 6,
		SmallRomanNumeralSeven       : 7,
		SmallRomanNumeralEight       : 8,
		SmallRomanNumeralNine        : 9,
		SmallRomanNumeralTen         : 10,
		SmallRomanNumeralEleven      : 11,
		SmallRomanNumeralTwelve      : 12,
		SmallRomanNumeralFifty       : 50,
		SmallRomanNumeralOneHundred  : 100,
		SmallRomanNumeralFiveHundred : 500,
		SmallRomanNumeralOneThousand : 1000,

		RomanNumeralOneThousandCD      : 1000,
		RomanNumeralFiveThousand       : 5000,
		RomanNumeralTenThousand        : 10000,
		RomanNumeralReservedOneHundred : 100,
		LatinSmallLetterReversedC      : 100,
		RomanNumeralSixLateForm        : 6,
		RomanNumeralFiftyEarlyForm     : 50,
		RomanNumeralFiftyThousand      : 50000,
		RomanNumeralOneHundredThousand : 100000,

		VulgarFractionZeroThirds : 0,

		TurnedDigitTwo   : 10,
		TurnedDigitThree : 11,

		HangzhouNumeralOne    : 1,
		HangzhouNumeralTwo    : 2,
		HangzhouNumeralThree  : 3,
		HangzhouNumeralFour   : 4,
		HangzhouNumeralFive   : 5,
		HangzhouNumeralSix    : 6,
		HangzhouNumeralSeven  : 7,
		HangzhouNumeralEight  : 8,
		HangzhouNumeralNine   : 9,

		HangzhouNumeralTen    : 10,
		HangzhouNumeralTwenty : 20,

		HangzhouNumeralThirty : 30,
	}
)


// Int64 tries to convert the numeric value that a rune represents to an
// int64.
//
// Note, of course, not all runes represent numeric values. In those cases
// Int64 returns an error of type NotNumericComplainer.
//
// Also, if the rune is numeric but cannot be represented as an int64,
// then Int64 returns an error of type NotInRangeComplainer.
func Int64(r rune) (int64, error) {
	if x, ok := runeToInt64LookupTable[r]; !ok {
		if !IsNumeric(r) {
			return 0, errNotNumeric
		} else {
			return 0, errNotInRange
		}
	} else {
		return x, nil
	}
}

