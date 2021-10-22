package period2011

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers"
	"github.com/mzdyhrave/payrollgo-legalios/internal/props"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type providerTaxing2011 struct {
	providers.ProviderBase
}

func NewProviderTaxing2011() providers.IProviderTaxing {
	return providerTaxing2011{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2011)},
	}
}

func (b providerTaxing2011) GetProps(period types.IPeriod) props.IPropsTaxing {
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

func (b providerTaxing2011) AllowancePayer(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_PAYER
}

func (b providerTaxing2011) AllowanceDisab1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_1ST
}

func (b providerTaxing2011) AllowanceDisab2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_2ND
}

func (b providerTaxing2011) AllowanceDisab3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_3RD
}

func (b providerTaxing2011) AllowanceStudy(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_STUDY
}

func (b providerTaxing2011) AllowanceChild1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_1ST
}

func (b providerTaxing2011) AllowanceChild2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_2ND
}

func (b providerTaxing2011) AllowanceChild3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_3RD
}

func (b providerTaxing2011) FactorAdvances(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_ADVANCES)
}

func (b providerTaxing2011) FactorWithhold(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_WITHHOLD)
}

func (b providerTaxing2011) FactorSolitary(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_SOLITARY)
}

func (b providerTaxing2011) MinAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MIN_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2011) MaxAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MAX_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2011) MarginIncomeOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXBONUS
}

func (b providerTaxing2011) MarginIncomeOfRounding(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_ROUNDING
}

func (b providerTaxing2011) MarginIncomeOfWithhold(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WITHHOLD
}

func (b providerTaxing2011) MarginIncomeOfSolitary(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_SOLITARY
}

func (b providerTaxing2011) MarginIncomeOfWthEmp(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_EMP
}

func (b providerTaxing2011) MarginIncomeOfWthAgr(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_AGR
}

