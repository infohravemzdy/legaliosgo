package providers

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/props"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
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
	FactorSolitary(period types.IPeriod) Decimal
	MinAmountOfTaxBonus(period types.IPeriod) int32
	MaxAmountOfTaxBonus(period types.IPeriod) int32
	MarginIncomeOfTaxBonus(period types.IPeriod) int32
	MarginIncomeOfRounding(period types.IPeriod) int32
	MarginIncomeOfWithhold(period types.IPeriod) int32
	MarginIncomeOfSolitary(period types.IPeriod) int32
	MarginIncomeOfWthEmp(period types.IPeriod) int32
	MarginIncomeOfWthAgr(period types.IPeriod) int32
}

