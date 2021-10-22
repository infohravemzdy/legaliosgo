package props

import (
	. "github.com/shopspring/decimal"
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type IPropsTaxing interface {
	IProps
	AllowancePayer() int32
	AllowanceDisab1st() int32
	AllowanceDisab2nd() int32
	AllowanceDisab3rd() int32
	AllowanceStudy() int32
	AllowanceChild1st() int32
	AllowanceChild2nd() int32
	AllowanceChild3rd() int32
	FactorAdvances() Decimal
	FactorWithhold() Decimal
	FactorSolitary() Decimal
	MinAmountOfTaxBonus() int32
	MaxAmountOfTaxBonus() int32
	MarginIncomeOfTaxBonus() int32
	MarginIncomeOfRounding() int32
	MarginIncomeOfWithhold() int32
	MarginIncomeOfSolitary() int32
	MarginIncomeOfWthEmp() int32
	MarginIncomeOfWthAgr() int32
}

type PropsTaxing struct {
	propsBase
	allowancePayer    int32
	allowanceDisab1st int32
	allowanceDisab2nd int32
	allowanceDisab3rd int32
	allowanceStudy int32
	allowanceChild1st int32
	allowanceChild2nd int32
	allowanceChild3rd int32
	factorAdvances Decimal
	factorWithhold Decimal
	factorSolitary Decimal
	minAmountOfTaxBonus int32
	maxAmountOfTaxBonus int32
	marginIncomeOfTaxBonus int32
	marginIncomeOfRounding int32
	marginIncomeOfWithhold int32
	marginIncomeOfSolitary int32
	marginIncomeOfWthEmp int32
	marginIncomeOfWthAgr int32
}

func (p PropsTaxing) AllowancePayer() int32 {
	return p.allowancePayer
}

func (p PropsTaxing) AllowanceDisab1st() int32 {
	return p.allowanceDisab1st
}

func (p PropsTaxing) AllowanceDisab2nd() int32 {
	return p.allowanceDisab2nd
}

func (p PropsTaxing) AllowanceDisab3rd() int32 {
	return p.allowanceDisab3rd
}

func (p PropsTaxing) AllowanceStudy() int32 {
	return p.allowanceStudy
}

func (p PropsTaxing) AllowanceChild1st() int32 {
	return p.allowanceChild1st
}

func (p PropsTaxing) AllowanceChild2nd() int32 {
	return p.allowanceChild2nd
}

func (p PropsTaxing) AllowanceChild3rd() int32 {
	return p.allowanceChild3rd
}

func (p PropsTaxing) FactorAdvances() Decimal {
	return p.factorAdvances
}

func (p PropsTaxing) FactorWithhold() Decimal {
	return p.factorWithhold
}

func (p PropsTaxing) FactorSolitary() Decimal {
	return p.factorSolitary
}

func (p PropsTaxing) MinAmountOfTaxBonus() int32 {
	return p.minAmountOfTaxBonus
}

func (p PropsTaxing) MaxAmountOfTaxBonus() int32 {
	return p.maxAmountOfTaxBonus
}

func (p PropsTaxing) MarginIncomeOfTaxBonus() int32 {
	return p.marginIncomeOfTaxBonus
}

func (p PropsTaxing) MarginIncomeOfRounding() int32 {
	return p.marginIncomeOfRounding
}

func (p PropsTaxing) MarginIncomeOfWithhold() int32 {
	return p.marginIncomeOfWithhold
}

func (p PropsTaxing) MarginIncomeOfSolitary() int32 {
	return p.marginIncomeOfSolitary
}

func (p PropsTaxing) MarginIncomeOfWthEmp() int32 {
	return p.marginIncomeOfWthEmp
}

func (p PropsTaxing) MarginIncomeOfWthAgr() int32 {
	return p.marginIncomeOfWthAgr
}

func NewPropsTaxing(versionId types.IVersionId,
	allowancePayer int32,
	allowanceDisab1st int32,
	allowanceDisab2nd int32,
	allowanceDisab3rd int32,
	allowanceStudy int32,
	allowanceChild1st int32,
	allowanceChild2nd int32,
	allowanceChild3rd int32,
	factorAdvances Decimal,
	factorWithhold Decimal,
	factorSolitary Decimal,
	minAmountOfTaxBonus int32,
	maxAmountOfTaxBonus int32,
	marginIncomeOfTaxBonus int32,
	marginIncomeOfRounding int32,
	marginIncomeOfWithhold int32,
	marginIncomeOfSolitary int32,
	marginIncomeOfWthEmp int32,
	marginIncomeOfWthAgr int32) IPropsTaxing {
	return PropsTaxing{
		propsBase:         propsBase{ Version: versionId },
		allowancePayer:    allowancePayer,
		allowanceDisab1st: allowanceDisab1st,
		allowanceDisab2nd: allowanceDisab2nd,
		allowanceDisab3rd: allowanceDisab3rd,
		allowanceStudy:    allowanceStudy,
		allowanceChild1st: allowanceChild1st,
		allowanceChild2nd: allowanceChild2nd,
		allowanceChild3rd: allowanceChild3rd,
		factorAdvances:    factorAdvances,
		factorWithhold:    factorWithhold,
		factorSolitary:    factorSolitary,
		minAmountOfTaxBonus: minAmountOfTaxBonus,
		maxAmountOfTaxBonus: maxAmountOfTaxBonus,
		marginIncomeOfTaxBonus: marginIncomeOfTaxBonus,
		marginIncomeOfRounding: marginIncomeOfRounding,
		marginIncomeOfWithhold: marginIncomeOfWithhold,
		marginIncomeOfSolitary: marginIncomeOfSolitary,
		marginIncomeOfWthEmp: marginIncomeOfWthEmp,
		marginIncomeOfWthAgr: marginIncomeOfWthAgr,
	}
}

func EmptyPropsTaxing() IPropsTaxing {
	return PropsTaxing{
		propsBase:         propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		allowancePayer:    0,
		allowanceDisab1st: 0,
		allowanceDisab2nd: 0,
		allowanceDisab3rd: 0,
		allowanceStudy:    0,
		allowanceChild1st: 0,
		allowanceChild2nd: 0,
		allowanceChild3rd: 0,
		factorAdvances:    NewFromFloat(0),
		factorWithhold:    NewFromFloat(0),
		factorSolitary:    NewFromFloat(0),
		minAmountOfTaxBonus: 0,
		maxAmountOfTaxBonus: 0,
		marginIncomeOfTaxBonus: 0,
		marginIncomeOfRounding: 0,
		marginIncomeOfWithhold: 0,
		marginIncomeOfSolitary: 0,
		marginIncomeOfWthEmp: 0,
		marginIncomeOfWthAgr: 0,
	}
}

