package providers

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IProviderHealth interface {
	ipropsProvider
	GetProps(period types.IPeriod) props.IPropsHealth

	MinMonthlyBasis(period types.IPeriod) int32
	MaxAnnualsBasis(period types.IPeriod) int32
	LimMonthlyState(period types.IPeriod) int32
	LimMonthlyDis50(period types.IPeriod) int32
	FactorCompound(period types.IPeriod) Decimal
	FactorEmployee(period types.IPeriod) Decimal
	MarginIncomeEmp(period types.IPeriod) int32
	MarginIncomeAgr(period types.IPeriod) int32
}

