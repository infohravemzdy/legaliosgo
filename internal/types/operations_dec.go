package types

import (. "github.com/shopspring/decimal")

func Multiply(op1 Decimal, op2 Decimal) Decimal {
	return op1.Mul(op2)
}

func Divide(op1 Decimal, div Decimal) Decimal {
	if div.IsZero() {
		return Zero
	}
	return op1.Div(div)
}

func MultiplyAndDivide(op1 Decimal, op2 Decimal, div Decimal) Decimal {
	if div.IsZero() {
		return Zero
	}
	multiRet := op1.Mul(op2)

	dividRet := multiRet.Div(div)

	return dividRet
}

func DecimalCast(number int32) Decimal {
	return NewFromInt32(number)
}

func MinIncMaxDecValue(valueToMinMax Decimal, accValueToMax Decimal, minLimitTo Decimal, maxLimitTo Decimal) Decimal {
	minBase := MinIncValue(valueToMinMax, minLimitTo)

	maxBase := MaxDecAccumValue(minBase, accValueToMax, maxLimitTo)

	return maxBase
}

func MaxDecAccumValue(valueToMax Decimal, accumToMax Decimal, maxLimitTo Decimal) Decimal {
	if maxLimitTo.GreaterThan(Zero) {
		valueToReduce := Max(valueToMax.Add(accumToMax), maxLimitTo)

		return Max(Zero, valueToReduce.Sub(accumToMax))
	}
	return valueToMax
}

func MaxDecAccumAbove(valueToMax Decimal, accumToMax Decimal, maxLimitTo Decimal) Decimal {
	if maxLimitTo.GreaterThan(Zero) {
		underToLimits := MaxDecAccumValue(valueToMax, accumToMax, maxLimitTo)

		return Max(Zero, valueToMax.Sub(underToLimits))
	}
	return Zero
}

func MinIncValue(valueToMin Decimal, minLimitTo Decimal) Decimal {
	if minLimitTo.GreaterThan(Zero) {
		if minLimitTo.GreaterThan(valueToMin) {
			return minLimitTo
		}
	}
	return valueToMin
}

func MaxDecValue(valueToMax Decimal, maxLimitTo Decimal) Decimal {
	if maxLimitTo.GreaterThan(Zero) {
		return Min(valueToMax, maxLimitTo)
	}
	return valueToMax
}

func SuppressNegative(suppress bool, valueDec Decimal) Decimal {
	if suppress && valueDec.LessThan(Zero) {
		return Zero
	}
	return valueDec
}
