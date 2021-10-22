package period2012

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSocial2012 struct {
	providers.ProviderBase
}

func NewProviderSocial2012() providers.IProviderSocial {
	return providerSocial2012{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2012)},
	}
}

func (b providerSocial2012) GetProps(period types.IPeriod) props.IPropsSocial {
	return props.NewPropsSocial(b.Version,
		b.MaxAnnualsBasis(period),
		b.FactorEmployer(period),
		b.FactorEmployerHigher(period),
		b.FactorEmployee(period),
		b.FactorEmployeeGarant(period),
		b.FactorEmployeeReduce(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerSocial2012) MaxAnnualsBasis(period types.IPeriod) int32 {
	return SOCIAL_MAX_ANNUALS_BASIS
}

func (b providerSocial2012) FactorEmployer(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER)
}

func (b providerSocial2012) FactorEmployerHigher(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER_HIGHER)
}

func (b providerSocial2012) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE)
}

func (b providerSocial2012) FactorEmployeeGarant(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_GARANT)
}

func (b providerSocial2012) FactorEmployeeReduce(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_REDUCE)
}

func (b providerSocial2012) MarginIncomeEmp(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_EMP
}

func (b providerSocial2012) MarginIncomeAgr(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_AGR
}
