package providers

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type IProviderTaxing interface {
	ipropsProvider
	GetProps(period types.IPeriod) props.IPropsTaxing

	AllowancePayer(period types.IPeriod) int32
	AllowanceDisab1st(period types.IPeriod) int32
	AllowanceDisab2nd(period types.IPeriod) int32
	AllowanceDisab3rd(period types.IPeriod) int32
	AllowanceStudy(period types.IPeriod) int32
	AllowanceChild1st(period types.IPeriod) int32
	AllowanceChild2nd(period types.IPeriod) int32
	AllowanceChild3rd(period types.IPeriod) int32
	FactorAdvances(period types.IPeriod) Decimal
	FactorWithhold(period types.IPeriod) Decimal
	FactorSolidary(period types.IPeriod) Decimal
	FactorTaxRate2(period types.IPeriod) Decimal
	MinAmountOfTaxBonus(period types.IPeriod) int32
	MaxAmountOfTaxBonus(period types.IPeriod) int32
	MarginIncomeOfTaxBonus(period types.IPeriod) int32
	MarginIncomeOfRounding(period types.IPeriod) int32
	MarginIncomeOfWithhold(period types.IPeriod) int32
	MarginIncomeOfSolidary(period types.IPeriod) int32
	MarginIncomeOfTaxRate2(period types.IPeriod) int32
	MarginIncomeOfWthEmp(period types.IPeriod) int32
	MarginIncomeOfWthAgr(period types.IPeriod) int32
}

