package period2010

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type providerHealth2010 struct {
	providers.ProviderBase
}

func NewProviderHealth2010() providers.IProviderHealth {
	return providerHealth2010{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2010)},
	}
}

func (b providerHealth2010) GetProps(period types.IPeriod) props.IPropsHealth {
	return props.NewPropsHealth2010(b.Version,
		b.MinMonthlyBasis(period),
		b.MaxAnnualsBasis(period),
		b.LimMonthlyState(period),
		b.LimMonthlyDis50(period),
		b.FactorCompound(period),
		b.FactorEmployee(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerHealth2010) MinMonthlyBasis(period types.IPeriod) int32 {
	return HEALTH_MIN_MONTHLY_BASIS
}

func (b providerHealth2010) MaxAnnualsBasis(period types.IPeriod) int32 {
	return HEALTH_MAX_ANNUALS_BASIS
}

func (b providerHealth2010) LimMonthlyState(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_STATE
}

func (b providerHealth2010) LimMonthlyDis50(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_DIS50
}

func (b providerHealth2010) FactorCompound(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_COMPOUND)
}

func (b providerHealth2010) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_EMPLOYEE)
}

func (b providerHealth2010) MarginIncomeEmp(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_EMP
}

func (b providerHealth2010) MarginIncomeAgr(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_AGR
}
