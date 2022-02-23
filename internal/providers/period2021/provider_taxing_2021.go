package period2021

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type providerTaxing2021 struct {
	providers.ProviderBase
}

func NewProviderTaxing2021() providers.IProviderTaxing {
	return providerTaxing2021{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2021)},
	}
}

func (b providerTaxing2021) GetProps(period types.IPeriod) props.IPropsTaxing {
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
		b.FactorSolidary(period),
		b.FactorTaxRate2(period),
		b.MinAmountOfTaxBonus(period),
		b.MaxAmountOfTaxBonus(period),
		b.MarginIncomeOfTaxBonus(period),
		b.MarginIncomeOfRounding(period),
		b.MarginIncomeOfWithhold(period),
		b.MarginIncomeOfSolidary(period),
		b.MarginIncomeOfTaxRate2(period),
		b.MarginIncomeOfWthEmp(period),
		b.MarginIncomeOfWthAgr(period))
}

func (b providerTaxing2021) AllowancePayer(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_PAYER
}

func (b providerTaxing2021) AllowanceDisab1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_1ST
}

func (b providerTaxing2021) AllowanceDisab2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_2ND
}

func (b providerTaxing2021) AllowanceDisab3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_3RD
}

func (b providerTaxing2021) AllowanceStudy(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_STUDY
}

func (b providerTaxing2021) AllowanceChild1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_1ST
}

func (b providerTaxing2021) AllowanceChild2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_2ND
}

func (b providerTaxing2021) AllowanceChild3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_3RD
}

func (b providerTaxing2021) FactorAdvances(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_ADVANCES)
}

func (b providerTaxing2021) FactorWithhold(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_WITHHOLD)
}

func (b providerTaxing2021) FactorSolidary(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_SOLIDARY)
}

func (b providerTaxing2021) FactorTaxRate2(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_TAXRATE2)
}

func (b providerTaxing2021) MinAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MIN_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2021) MaxAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MAX_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2021) MarginIncomeOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXBONUS
}

func (b providerTaxing2021) MarginIncomeOfRounding(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_ROUNDING
}

func (b providerTaxing2021) MarginIncomeOfWithhold(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WITHHOLD
}

func (b providerTaxing2021) MarginIncomeOfSolidary(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_SOLIDARY
}

func (b providerTaxing2021) MarginIncomeOfTaxRate2(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXRATE2
}

func (b providerTaxing2021) MarginIncomeOfWthEmp(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_EMP
}

func (b providerTaxing2021) MarginIncomeOfWthAgr(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_AGR
}

