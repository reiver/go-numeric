package numeric


var (
	runeToFrac64LoopupTable = map[rune]struct{Numerator int64 ; Denominator int64}{
		VulgarFractionOneQuarter    : {Numerator: 1, Denominator: 4},
		VulgarFractionOneHalf       : {Numerator: 1, Denominator: 2},
		VulgarFractionThreeQuarters : {Numerator: 3, Denominator: 4},

		VulgarFractionOneSeventh    : {Numerator: 1, Denominator: 7},

		VulgarFractionOneNinth      : {Numerator: 1, Denominator: 9},

		VulgarFractionOneTenth      : {Numerator: 1, Denominator: 10},

		VulgarFractionOneThird      : {Numerator: 1, Denominator: 3},
		VulgarFractionTwoThirds     : {Numerator: 2, Denominator: 3},

		VulgarFractionOneFifth      : {Numerator: 1, Denominator: 5},
		VulgarFractionTwoFifths     : {Numerator: 2, Denominator: 5},
		VulgarFractionThreeFifths   : {Numerator: 3, Denominator: 5},
		VulgarFractionFourFifths    : {Numerator: 4, Denominator: 5},

		VulgarFractionOneSixth      : {Numerator: 1, Denominator: 6},
		VulgarFractionFiveSixths    : {Numerator: 5, Denominator: 6},

		VulgarFractionOneEighth     : {Numerator: 1, Denominator: 8},
		VulgarFractionThreeEighths  : {Numerator: 3, Denominator: 8},
		VulgarFractionFiveEighths   : {Numerator: 5, Denominator: 8},
		VulgarFractionSevenEighths  : {Numerator: 7, Denominator: 8},
	}
)


// Frac64 tries to convert the numeric value that a rune represents to an
// fraction made of an int64 numerator and int64 denominator.
//
// Note, of course, not all runes represent numeric values. In those cases
// Frac64 returns an error of type NotNumericComplainer.
func Frac64(r rune) (numerator int64, denominator int64, err error) {
	if x, ok := runeToInt64LookupTable[r]; ok {
		return x, 1, nil
	} else if x, ok := runeToFrac64LoopupTable[r]; ok {
		return x.Numerator, x.Denominator, nil
	} else {
		return 0, 0, errNotNumeric
	}
}
