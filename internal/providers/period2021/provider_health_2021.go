package period2021

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerHealth2021 struct {
	providers.ProviderBase
}

func NewProviderHealth2021() providers.IProviderHealth {
	return providerHealth2021{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2021)},
	}
}

func (b providerHealth2021) GetProps(period types.IPeriod) props.IPropsHealth {
	return props.NewPropsHealth(b.Version,
		b.MinMonthlyBasis(period),
		b.MaxAnnualsBasis(period),
		b.LimMonthlyState(period),
		b.LimMonthlyDis50(period),
		b.FactorCompound(period),
		b.FactorEmployee(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerHealth2021) MinMonthlyBasis(period types.IPeriod) int32 {
	return HEALTH_MIN_MONTHLY_BASIS
}

func (b providerHealth2021) MaxAnnualsBasis(period types.IPeriod) int32 {
	return HEALTH_MAX_ANNUALS_BASIS
}

func (b providerHealth2021) LimMonthlyState(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_STATE
}

func (b providerHealth2021) LimMonthlyDis50(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_DIS50
}

func (b providerHealth2021) FactorCompound(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_COMPOUND)
}

func (b providerHealth2021) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_EMPLOYEE)
}

func (b providerHealth2021) MarginIncomeEmp(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_EMP
}

func (b providerHealth2021) MarginIncomeAgr(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_AGR
}
