package period2014

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers"
	"github.com/mzdyhrave/payrollgo-legalios/internal/props"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type providerTaxing2014 struct {
	providers.ProviderBase
}

func NewProviderTaxing2014() providers.IProviderTaxing {
	return providerTaxing2014{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2014)},
	}
}

func (b providerTaxing2014) GetProps(period types.IPeriod) props.IPropsTaxing {
	return props.NewPropsTaxing(b.Version,
		b.AllowancePayer(period),
		b.AllowanceDisab1st(period),
		b.AllowanceDisab2nd(period),
		b.AllowanceDisab3rd(period),
		b.AllowanceStudy(period),
		b.AllowanceChild1st(period),
		b.AllowanceChild2nd(period),
		b.AllowanceChild3rd(period),
		b.FactorAdvances(period),
		b.FactorWithhold(period),
		b.FactorSolitary(period),
		b.MinAmountOfTaxBonus(period),
		b.MaxAmountOfTaxBonus(period),
		b.MarginIncomeOfTaxBonus(period),
		b.MarginIncomeOfRounding(period),
		b.MarginIncomeOfWithhold(period),
		b.MarginIncomeOfSolitary(period),
		b.MarginIncomeOfWthEmp(period),
		b.MarginIncomeOfWthAgr(period))
}

func (b providerTaxing2014) AllowancePayer(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_PAYER
}

func (b providerTaxing2014) AllowanceDisab1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_1ST
}

func (b providerTaxing2014) AllowanceDisab2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_2ND
}

func (b providerTaxing2014) AllowanceDisab3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_3RD
}

func (b providerTaxing2014) AllowanceStudy(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_STUDY
}

func (b providerTaxing2014) AllowanceChild1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_1ST
}

func (b providerTaxing2014) AllowanceChild2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_2ND
}

func (b providerTaxing2014) AllowanceChild3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_3RD
}

func (b providerTaxing2014) FactorAdvances(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_ADVANCES)
}

func (b providerTaxing2014) FactorWithhold(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_WITHHOLD)
}

func (b providerTaxing2014) FactorSolitary(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_SOLITARY)
}

func (b providerTaxing2014) MinAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MIN_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2014) MaxAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MAX_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2014) MarginIncomeOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXBONUS
}

func (b providerTaxing2014) MarginIncomeOfRounding(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_ROUNDING
}

func (b providerTaxing2014) MarginIncomeOfWithhold(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WITHHOLD
}

func (b providerTaxing2014) MarginIncomeOfSolitary(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_SOLITARY
}

func (b providerTaxing2014) MarginIncomeOfWthEmp(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_EMP
}

func (b providerTaxing2014) MarginIncomeOfWthAgr(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_AGR
}

