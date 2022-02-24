package period2013

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerHealth2013 struct {
	providers.ProviderBase
}

func NewProviderHealth2013() providers.IProviderHealth {
	return providerHealth2013{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2013)},
	}
}

func (b providerHealth2013) GetProps(period types.IPeriod) props.IPropsHealth {
	return props.NewPropsHealth2012(b.Version,
		b.MinMonthlyBasis(period),
		b.MaxAnnualsBasis(period),
		b.LimMonthlyState(period),
		b.LimMonthlyDis50(period),
		b.FactorCompound(period),
		b.FactorEmployee(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerHealth2013) MinMonthlyBasis(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 8) {
		return VAR08_HEALTH_MIN_MONTHLY_BASIS
	}
	return HEALTH_MIN_MONTHLY_BASIS
}

func (b providerHealth2013) MaxAnnualsBasis(period types.IPeriod) int32 {
	return HEALTH_MAX_ANNUALS_BASIS
}

func (b providerHealth2013) LimMonthlyState(period types.IPeriod) int32 {
	return HEALTH_LIM_MONTHLY_STATE
}

func (b providerHealth2013) LimMonthlyDis50(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 11) {
		return VAR11_HEALTH_LIM_MONTHLY_DIS50
	}
	return HEALTH_LIM_MONTHLY_DIS50
}

func (b providerHealth2013) FactorCompound(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_COMPOUND)
}

func (b providerHealth2013) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(HEALTH_FACTOR_EMPLOYEE)
}

func (b providerHealth2013) MarginIncomeEmp(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_EMP
}

func (b providerHealth2013) MarginIncomeAgr(period types.IPeriod) int32 {
	return HEALTH_MARGIN_INCOME_AGR
}
