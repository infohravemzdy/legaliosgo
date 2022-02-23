package types

import . "github.com/shopspring/decimal"

var INT_ROUNDING_CONST = NewFromFloat(0.5)

func RoundToInt(valueDec Decimal) int32 {
	roundRet := valueDec.Abs().Add(INT_ROUNDING_CONST).Floor()

	if valueDec.LessThan(Zero) {
		return int32(roundRet.Neg().IntPart())
	}
	return int32(roundRet.IntPart())
}
func RoundUp(valueDec Decimal) int32 {
	roundRet := valueDec.Abs().Ceil()

	if valueDec.LessThan(Zero) {
		return int32(roundRet.Neg().IntPart())
	}
	return int32(roundRet.IntPart())
}

func RoundDown(valueDec Decimal) int32 {
	roundRet := valueDec.Abs().Floor()

	if valueDec.LessThan(Zero) {
		return int32(roundRet.Neg().IntPart())
	}
	return int32(roundRet.IntPart())
}
func RoundNorm(valueDec Decimal) int32 {
	roundRet := valueDec.Abs().Add(INT_ROUNDING_CONST).Floor()

	if valueDec.LessThan(Zero) {
		return int32(roundRet.Neg().IntPart())
	}
	return int32(roundRet.IntPart())
}

func NearRoundUp(valueDec Decimal, nearest int32) int32 {
	nearestBig := NewFromInt32(nearest)

	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundUp(dividRet), nearestBig)

	return RoundToInt(multiRet)
}
func NearRoundUp100(valueDec Decimal) int32 {
	nearestBig := NewFromInt32(100)

	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundUp(dividRet), nearestBig)

	return RoundToInt(multiRet)
}

func NearRoundDown(valueDec Decimal, nearest int32) int32 {
	nearestBig := NewFromInt32(nearest)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundDown(dividRet), nearestBig)

	return RoundToInt(multiRet)
}
func NearRoundDown100(valueDec Decimal) int32 {
	nearestBig := NewFromInt32(100)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundDown(dividRet), nearestBig)

	return RoundToInt(multiRet)
}
func RoundUp50(valueDec Decimal) int32 {
	divider := NewFromInt32(2)
	dividRet := Divide(DecRoundUp(Multiply(valueDec, divider)), divider)
	return RoundToInt(dividRet)
}
func RoundUp25(valueDec Decimal) int32 {
	divider := NewFromInt32(4)
	dividRet := Divide(DecRoundUp(Multiply(valueDec, divider)), divider)
	return RoundToInt(dividRet)
}

func DecRoundUp(valueDec Decimal) Decimal {
	roundRet := valueDec.Abs().Ceil()

	if valueDec.LessThan(Zero) {
		return roundRet.Neg()
	}
	return roundRet
}

func DecRoundDown(valueDec Decimal) Decimal {
	roundRet := valueDec.Abs().Floor()

	if valueDec.LessThan(Zero) {
		return roundRet.Neg()
	}
	return roundRet
}
func DecRoundNorm(valueDec Decimal) Decimal {
	roundRet := valueDec.Abs().Add(INT_ROUNDING_CONST).Floor()

	if valueDec.LessThan(Zero) {
		return roundRet.Neg()
	}
	return roundRet
}

func DecNearRoundUp(valueDec Decimal, nearest int32) Decimal {
	nearestBig := NewFromInt32(nearest)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundUp(dividRet), nearestBig)

	return multiRet
}
func DecNearRoundUp100(valueDec Decimal) Decimal {
	nearestBig := NewFromInt32(100)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundUp(dividRet), nearestBig)

	return multiRet
}

func DecNearRoundDown(valueDec Decimal, nearest int32) Decimal {
	nearestBig := NewFromInt32(nearest)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundDown(dividRet), nearestBig)

	return multiRet
}
func DecNearRoundDown100(valueDec Decimal) Decimal {
	nearestBig := NewFromInt32(100)
	dividRet := Divide(valueDec, nearestBig)

	multiRet := Multiply(DecRoundDown(dividRet), nearestBig)

	return multiRet
}
func DecRoundUp50(valueDec Decimal) Decimal {
	divider := NewFromInt32(2)
	return Divide(DecRoundUp(Multiply(valueDec, divider)), divider)
}
func DecRoundUp25(valueDec Decimal) Decimal {
	divider := NewFromInt32(4)
	return Divide(DecRoundUp(Multiply(valueDec, divider)), divider)
}
func DecRoundUp01(valueDec Decimal) Decimal {
	divider := NewFromInt32(100)
	return Divide(DecRoundUp(Multiply(valueDec, divider)), divider)
}
func DecRoundDown50(valueDec Decimal) Decimal {
	divider := NewFromInt32(2)
	return Divide(DecRoundDown(Multiply(valueDec, divider)), divider)
}
func DecRoundDown25(valueDec Decimal) Decimal {
	divider := NewFromInt32(4)
	return Divide(DecRoundDown(Multiply(valueDec, divider)), divider)
}
func DecRoundDown01(valueDec Decimal) Decimal {
	divider := NewFromInt32(100)
	return Divide(DecRoundDown(Multiply(valueDec, divider)), divider)
}
func DecRoundNorm50(valueDec Decimal) Decimal {
	divider := NewFromInt32(2)
	return Divide(DecRoundNorm(Multiply(valueDec, divider)), divider)
}
func DecRoundNorm25(valueDec Decimal) Decimal {
	divider := NewFromInt32(4)
	return Divide(DecRoundNorm(Multiply(valueDec, divider)), divider)
}
func DecRoundNorm01(valueDec Decimal) Decimal {
	divider := NewFromInt32(100)
	return Divide(DecRoundNorm(Multiply(valueDec, divider)), divider)
}
