package period2013

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSocial2013 struct {
	providers.ProviderBase
}

func NewProviderSocial2013() providers.IProviderSocial {
	return providerSocial2013{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2013)},
	}
}

func (b providerSocial2013) GetProps(period types.IPeriod) props.IPropsSocial {
	return props.NewPropsSocial2012(b.Version,
		b.MaxAnnualsBasis(period),
		b.FactorEmployer(period),
		b.FactorEmployerHigher(period),
		b.FactorEmployee(period),
		b.FactorEmployeeGarant(period),
		b.FactorEmployeeReduce(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerSocial2013) MaxAnnualsBasis(period types.IPeriod) int32 {
	return SOCIAL_MAX_ANNUALS_BASIS
}

func (b providerSocial2013) FactorEmployer(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER)
}

func (b providerSocial2013) FactorEmployerHigher(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER_HIGHER)
}

func (b providerSocial2013) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE)
}

func (b providerSocial2013) FactorEmployeeGarant(period types.IPeriod) Decimal {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 2) {
		return NewFromFloat(VAR02_SOCIAL_FACTOR_EMPLOYEE_GARANT)
	}
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_GARANT)
}

func (b providerSocial2013) FactorEmployeeReduce(period types.IPeriod) Decimal {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 2) {
		return NewFromFloat(VAR02_SOCIAL_FACTOR_EMPLOYEE_REDUCE)
	}
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_REDUCE)
}

func (b providerSocial2013) MarginIncomeEmp(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_EMP
}

func (b providerSocial2013) MarginIncomeAgr(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_AGR
}
