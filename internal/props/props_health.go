package props

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IPropsHealth interface {
	IProps
	MinMonthlyBasis() int32
	MaxAnnualsBasis() int32
	LimMonthlyState() int32
	LimMonthlyDis50() int32
	FactorCompound() Decimal
	FactorEmployee() Decimal
	MarginIncomeEmp() int32
	MarginIncomeAgr() int32
}

type PropsHealth struct {
	propsBase
	minMonthlyBasis int32
	maxAnnualsBasis int32
	limMonthlyState int32
	limMonthlyDis50 int32
	factorCompound Decimal
	factorEmployee Decimal
	marginIncomeEmp int32
	marginIncomeAgr int32
}

func (p PropsHealth) MinMonthlyBasis() int32 {
	return p.minMonthlyBasis
}

func (p PropsHealth) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p PropsHealth) LimMonthlyState() int32 {
	return p.limMonthlyState
}

func (p PropsHealth) LimMonthlyDis50() int32 {
	return p.limMonthlyDis50
}

func (p PropsHealth) FactorCompound() Decimal {
	return p.factorCompound
}

func (p PropsHealth) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p PropsHealth) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p PropsHealth) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func NewPropsHealth(versionId types.IVersionId,
	minMonthlyBasis int32,
	maxAnnualsBasis int32,
	limMonthlyState int32,
	limMonthlyDis50 int32,
	factorCompound Decimal,
	factorEmployee Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsHealth {
	return PropsHealth{
		propsBase:       propsBase{ Version: versionId },
		minMonthlyBasis: minMonthlyBasis,
		maxAnnualsBasis: maxAnnualsBasis,
		limMonthlyState: limMonthlyState,
		limMonthlyDis50: limMonthlyDis50,
		factorCompound:  factorCompound,
		factorEmployee:  factorEmployee,
		marginIncomeEmp: marginIncomeEmp,
		marginIncomeAgr: marginIncomeAgr,
	}
}

func EmptyPropsHealth() IPropsHealth {
	return PropsHealth{
		propsBase:       propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		minMonthlyBasis: 0,
		maxAnnualsBasis: 0,
		limMonthlyState: 0,
		limMonthlyDis50: 0,
		factorCompound:  NewFromFloat(0),
		factorEmployee:  NewFromFloat(0),
		marginIncomeEmp: 0,
		marginIncomeAgr: 0,
	}
}

