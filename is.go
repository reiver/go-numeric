package numeric


// IsHangzhouNumeral returns true if the rune is a Hangzhou Numeral,
// and returns false otherwise.
//
// Namely:
//
// Hangzhou Numeral One    (〡)
//
// Hangzhou Numeral Two    (〢)
//
// Hangzhou Numeral Three  (〣)
//
// Hangzhou Numeral Four   (〤)
//
// Hangzhou Numeral Five   (〥)
//
// Hangzhou Numeral Six    (〦)
//
// Hangzhou Numeral Seven  (〧)
//
// Hangzhou Numeral Eight  (〨)
//
// Hangzhou Numeral Nine   (〩)
//
// Hangzhou Numeral Ten    (〸)
//
// Hangzhou Numeral Twenty (〹)
//
// Hangzhou Numeral Thrity (〺)
func IsHangzhouNumeral(r rune) bool {

	switch r {
	case
		HangzhouNumeralOne,
		HangzhouNumeralTwo,
		HangzhouNumeralThree,
		HangzhouNumeralFour,
		HangzhouNumeralFive,
		HangzhouNumeralSix,
		HangzhouNumeralSeven,
		HangzhouNumeralEight,
		HangzhouNumeralNine,

		HangzhouNumeralTen,
		HangzhouNumeralTwenty,

		HangzhouNumeralThirty:

		return true
	default:
		return false
	}
}


// IsOldPersianNumber returns true if the rune is an Old Persian Number,
// and returns false otherwise.
//
// Namely:
//
// Old Persian Number One     (𐏑)
//
// Old Persian Number Two     (𐏒)
//
// Old Persian Number Ten     (𐏓)
//
// Old Persian Number Twenty  (𐏔)
//
// Old Persian Number Hundred (𐏕)
func IsOldPersianNumber(r rune) bool {

	switch r {
	case
		OldPersianNumberOne,
		OldPersianNumberTwo,
		OldPersianNumberTen,
		OldPersianNumberTwenty,
		OldPersianNumberHundred:

		return true
	default:
		return false
	}
}
