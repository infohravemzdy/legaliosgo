package period2014

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type providerHealth2014 struct {
	providers.ProviderBase
}

func NewProviderHealth2014() providers.IProviderHealth {
	return providerHealth2014{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2014)},
	}
}

func (b providerHealth2014) GetProps(period types.IPeriod) props.IPropsHealth {
	return props.NewPropsHealth2014(b.Version,
		b.MinMonthlyBasis(period),
		b.MaxAnnualsBasis(period),
		b.LimMonthlyState(period),
		b.LimMonthlyDis50(period),
		b.FactorCompound(period),
		b.FactorEmployee(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerHealth2014) MinMonthlyBasis(period types.IPeriod) int32 {
	return HEALTH_MIN_MONTHLY_BASIS
}

func (b providerHealth2014) MaxAnnualsBasis(period types.IPeriod) int32 {
	return HEALTH_MAX_ANNUALS_BASIS
}

func (b providerHealth2014) LimMonthlyState(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_STATE
}

func (b providerHealth2014) LimMonthlyDis50(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2014, 7) {
		return VAR07_HEALTH_LIM_MONTHLY_DIS50
	}
	return HEALTH_LIM_MONTHLY_DIS50
}

func (b providerHealth2014) FactorCompound(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_COMPOUND)
}

func (b providerHealth2014) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_EMPLOYEE)
}

func (b providerHealth2014) MarginIncomeEmp(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_EMP
}

func (b providerHealth2014) MarginIncomeAgr(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_AGR
}
