package props

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type IPropsSocial interface {
	IProps
	MaxAnnualsBasis() int32
	FactorEmployer() Decimal
	FactorEmployerHigher() Decimal
	FactorEmployee() Decimal
	FactorEmployeeGarant() Decimal
	FactorEmployeeReduce() Decimal
	MarginIncomeEmp() int32
	MarginIncomeAgr() int32
}
type PropsSocial struct {
	propsBase
	maxAnnualsBasis int32
	factorEmployer Decimal
	factorEmployerHigher Decimal
	factorEmployee Decimal
	factorEmployeeGarant Decimal
	factorEmployeeReduce Decimal
	marginIncomeEmp int32
	marginIncomeAgr int32
}

func (p PropsSocial) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p PropsSocial) FactorEmployer() Decimal {
	return p.factorEmployer
}

func (p PropsSocial) FactorEmployerHigher() Decimal {
	return p.factorEmployerHigher
}

func (p PropsSocial) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p PropsSocial) FactorEmployeeGarant() Decimal {
	return p.factorEmployeeGarant
}

func (p PropsSocial) FactorEmployeeReduce() Decimal {
	return p.factorEmployeeReduce
}

func (p PropsSocial) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p PropsSocial) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func NewPropsSocial(versionId types.IVersionId,
	maxAnnualsBasis int32,
	factorEmployer Decimal,
	factorEmployerHigher Decimal,
	factorEmployee Decimal,
	factorEmployeeGarant Decimal,
	factorEmployeeReduce Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsSocial {
	return PropsSocial{
		propsBase:            propsBase{ Version: versionId },
		maxAnnualsBasis:      maxAnnualsBasis,
		factorEmployer:       factorEmployer,
		factorEmployerHigher: factorEmployerHigher,
		factorEmployee:       factorEmployee,
		factorEmployeeGarant: factorEmployeeGarant,
		factorEmployeeReduce: factorEmployeeReduce,
		marginIncomeEmp:      marginIncomeEmp,
		marginIncomeAgr:      marginIncomeAgr,
	}
}

func EmptyPropsSocial() IPropsSocial {
	return PropsSocial{
		propsBase:            propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		maxAnnualsBasis:      0,
		factorEmployer:       NewFromFloat(0),
		factorEmployerHigher: NewFromFloat(0),
		factorEmployee:       NewFromFloat(0),
		factorEmployeeGarant: NewFromFloat(0),
		factorEmployeeReduce: NewFromFloat(0),
		marginIncomeEmp:      0,
		marginIncomeAgr:      0,
	}
}

