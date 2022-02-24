package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
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
	FactorSolidary() Decimal
	FactorTaxRate2() Decimal
	MinAmountOfTaxBonus() int32
	MaxAmountOfTaxBonus() int32
	MarginIncomeOfTaxBonus() int32
	MarginIncomeOfRounding() int32
	MarginIncomeOfWithhold() int32
	MarginIncomeOfSolidary() int32
	MarginIncomeOfTaxRate2() int32
	MarginIncomeOfWthEmp() int32
	MarginIncomeOfWthAgr() int32

	ValueEquals(otherTaxing IPropsTaxing) bool
	HasWithholdIncome(termOpt types.WorkTaxingTerms, signOpt types.TaxDeclSignOption, noneOpt types.TaxNoneSignOption, incomeSum int32) bool
	BenefitAllowancePayer(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32
	BenefitAllowanceDisab(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclDisabOption) int32
	BenefitAllowanceStudy(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32
	BenefitAllowanceChild(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption, benefitOrds int32, disabelOpts int32) int32
	BonusChildRaw(income int32, benefit int32, rebated int32) int32
	BonusChildFix(income int32, benefit int32, rebated int32) int32
	TaxableIncomeSupers(incomeResult int32, healthResult int32, socialResult int32) int32
	TaxableIncomeBasis(incomeResult int32) int32
	RoundedRawBaseAdvances(incomeResult int32) int32
	RoundedBaseAdvances(incomeResult int32, healthResult int32, socialResult int32) int32
	RoundedBaseSolidary(incomeResult int32) int32
	RoundedAdvancesPaym(supersResult int32, basisResult int32) int32
	RoundedSolidaryPaym(basisResult int32) int32
	RoundedBaseWithhold(incomeResult int32) int32
	RoundedWithholdPaym(supersResult int32, basisResult int32) int32
}

type propsTaxingBase struct {
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
	factorSolidary Decimal
	factorTaxRate2 Decimal
	minAmountOfTaxBonus int32
	maxAmountOfTaxBonus int32
	marginIncomeOfTaxBonus int32
	marginIncomeOfRounding int32
	marginIncomeOfWithhold int32
	marginIncomeOfSolidary int32
	marginIncomeOfTaxRate2 int32
	marginIncomeOfWthEmp int32
	marginIncomeOfWthAgr int32
}

func (p propsTaxingBase) AllowancePayer() int32 {
	return p.allowancePayer
}

func (p propsTaxingBase) AllowanceDisab1st() int32 {
	return p.allowanceDisab1st
}

func (p propsTaxingBase) AllowanceDisab2nd() int32 {
	return p.allowanceDisab2nd
}

func (p propsTaxingBase) AllowanceDisab3rd() int32 {
	return p.allowanceDisab3rd
}

func (p propsTaxingBase) AllowanceStudy() int32 {
	return p.allowanceStudy
}

func (p propsTaxingBase) AllowanceChild1st() int32 {
	return p.allowanceChild1st
}

func (p propsTaxingBase) AllowanceChild2nd() int32 {
	return p.allowanceChild2nd
}

func (p propsTaxingBase) AllowanceChild3rd() int32 {
	return p.allowanceChild3rd
}

func (p propsTaxingBase) FactorAdvances() Decimal {
	return p.factorAdvances
}

func (p propsTaxingBase) FactorWithhold() Decimal {
	return p.factorWithhold
}

func (p propsTaxingBase) FactorSolidary() Decimal {
	return p.factorSolidary
}

func (p propsTaxingBase) FactorTaxRate2() Decimal {
	return p.factorTaxRate2
}

func (p propsTaxingBase) MinAmountOfTaxBonus() int32 {
	return p.minAmountOfTaxBonus
}

func (p propsTaxingBase) MaxAmountOfTaxBonus() int32 {
	return p.maxAmountOfTaxBonus
}

func (p propsTaxingBase) MarginIncomeOfTaxBonus() int32 {
	return p.marginIncomeOfTaxBonus
}

func (p propsTaxingBase) MarginIncomeOfRounding() int32 {
	return p.marginIncomeOfRounding
}

func (p propsTaxingBase) MarginIncomeOfWithhold() int32 {
	return p.marginIncomeOfWithhold
}

func (p propsTaxingBase) MarginIncomeOfSolidary() int32 {
	return p.marginIncomeOfSolidary
}

func (p propsTaxingBase) MarginIncomeOfTaxRate2() int32 {
	return p.marginIncomeOfTaxRate2
}

func (p propsTaxingBase) MarginIncomeOfWthEmp() int32 {
	return p.marginIncomeOfWthEmp
}

func (p propsTaxingBase) MarginIncomeOfWthAgr() int32 {
	return p.marginIncomeOfWthAgr
}

func (p propsTaxingBase) ValueEquals(otherTaxing IPropsTaxing) bool {
	if otherTaxing == nil {
		return false
	}
	return  p.allowancePayer == otherTaxing.AllowancePayer() &&
			p.allowanceDisab1st == otherTaxing.AllowanceDisab1st() &&
			p.allowanceDisab2nd == otherTaxing.AllowanceDisab2nd() &&
			p.allowanceDisab3rd == otherTaxing.AllowanceDisab3rd() &&
			p.allowanceStudy == otherTaxing.AllowanceStudy() &&
			p.allowanceChild1st == otherTaxing.AllowanceChild1st() &&
			p.allowanceChild2nd == otherTaxing.AllowanceChild2nd() &&
			p.allowanceChild3rd == otherTaxing.AllowanceChild3rd() &&
			p.factorAdvances.Equal(otherTaxing.FactorAdvances()) &&
			p.factorWithhold.Equal(otherTaxing.FactorWithhold()) &&
			p.factorSolidary.Equal(otherTaxing.FactorSolidary()) &&
			p.factorTaxRate2.Equal(otherTaxing.FactorTaxRate2()) &&
			p.minAmountOfTaxBonus == otherTaxing.MinAmountOfTaxBonus() &&
			p.maxAmountOfTaxBonus == otherTaxing.MaxAmountOfTaxBonus() &&
			p.marginIncomeOfTaxBonus == otherTaxing.MarginIncomeOfTaxBonus() &&
			p.marginIncomeOfRounding == otherTaxing.MarginIncomeOfRounding() &&
			p.marginIncomeOfWithhold == otherTaxing.MarginIncomeOfWithhold() &&
			p.marginIncomeOfSolidary == otherTaxing.MarginIncomeOfSolidary() &&
			p.marginIncomeOfTaxRate2 == otherTaxing.MarginIncomeOfTaxRate2() &&
			p.marginIncomeOfWthEmp == otherTaxing.MarginIncomeOfWthEmp() &&
			p.marginIncomeOfWthAgr == otherTaxing.MarginIncomeOfWthAgr()
}

func (p propsTaxingBase) intTaxRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}
func (p propsTaxingBase) intTaxRoundNearUp(valueDec Decimal, nearest int32) int32 {
	return types.NearRoundUp(valueDec, nearest)
}
func (p propsTaxingBase) intTaxRoundNearUp100(valueDec Decimal) int32 {
	return types.NearRoundUp(valueDec, 100)
}
func (p propsTaxingBase) intTaxRoundDown(valueDec Decimal) int32 {
	return types.RoundDown(valueDec)
}
func (p propsTaxingBase) intTaxRoundNearDown(valueDec Decimal, nearest int32) int32 {
	return types.NearRoundDown(valueDec, nearest)
}
func (p propsTaxingBase) intTaxRoundNearDown100(valueDec Decimal) int32 {
	return types.NearRoundDown(valueDec, 100)
}
func (p propsTaxingBase) decTaxRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}
func (p propsTaxingBase) decTaxRoundNearUp(valueDec Decimal, nearest int32) Decimal {
	return types.DecNearRoundUp(valueDec, nearest)
}
func (p propsTaxingBase) decTaxRoundNearUp100(valueDec Decimal) Decimal {
	return types.DecNearRoundUp(valueDec, 100)
}
func (p propsTaxingBase) decTaxRoundDown(valueDec Decimal) Decimal {
	return types.DecRoundDown(valueDec)
}
func (p propsTaxingBase) decTaxRoundNearDown(valueDec Decimal, nearest int32) Decimal {
	return types.DecNearRoundDown(valueDec, nearest)
}
func (p propsTaxingBase) decTaxRoundNearDown100(valueDec Decimal) Decimal {
	return types.DecNearRoundDown(valueDec, 100)
}

func (p propsTaxingBase) HasWithholdIncome(termOpt types.WorkTaxingTerms, signOpt types.TaxDeclSignOption, noneOpt types.TaxNoneSignOption, incomeSum int32) bool {
	//*****************************************************************************
	// Tax income for advance from Year 2014 to Year 2017
	//*****************************************************************************
	// - withhold tax (non-signed declaration) and income
	// if (period.Year >= 2018)
	// -- income from DPP is less than X CZK
	// -- income from low-income employment is less than X CZK
	// -- income from statutory employment and non-resident is always withhold tax

	var withholdIncome bool = false
	if signOpt != types.DECL_TAX_NO_SIGNED {
		return withholdIncome
	}
	if noneOpt != types.NOSIGN_TAX_WITHHOLD {
		return withholdIncome
	}
	if termOpt == types.TAXING_TERM_AGREEM_TASK {
		if p.MarginIncomeOfWthAgr() == 0 || incomeSum <= p.MarginIncomeOfWthAgr() {
			if incomeSum > 0 {
				withholdIncome = true
			}
		}
		return withholdIncome
	}
	if termOpt == types.TAXING_TERM_EMPLOYMENTS {
		if p.MarginIncomeOfWthEmp() == 0 || incomeSum <= p.MarginIncomeOfWthEmp() {
			if incomeSum > 0 {
				withholdIncome = true
			}
		}
		return withholdIncome
	}
	if termOpt == types.TAXING_TERM_STATUT_PART {
		if incomeSum > 0 {
			withholdIncome = true
		}
		return withholdIncome
	}
	return withholdIncome
}

func (p propsTaxingBase) BenefitAllowancePayer(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		if benefitOpts == types.DECL_TAX_BENEF1 {
			benefitValue = p.AllowancePayer()
		}
	}
	return benefitValue
}

func (p propsTaxingBase) BenefitAllowanceDisab(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclDisabOption) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		switch benefitOpts {
		case types.DISB_TAX_DISAB1: benefitValue = p.AllowanceDisab1st()
		case types.DISB_TAX_DISAB2: benefitValue = p.AllowanceDisab2nd()
		case types.DISB_TAX_DISAB3: benefitValue = p.AllowanceDisab3rd()
		default: benefitValue = 0
		}
	}
	return benefitValue
}

func (p propsTaxingBase) BenefitAllowanceStudy(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		if benefitOpts == types.DECL_TAX_BENEF1 {
			benefitValue = p.AllowanceStudy()
		}
	}
	return benefitValue
}

func (p propsTaxingBase) BenefitAllowanceChild(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption, benefitOrds int32, disabelOpts int32) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		var benefitUnits int32 = 0
		switch benefitOpts {
		case 0: benefitUnits = p.AllowanceChild1st()
		case 1: benefitUnits = p.AllowanceChild2nd()
		case 2: benefitUnits = p.AllowanceChild3rd()
		default: benefitUnits = 0
		}
		if benefitOpts == types.DECL_TAX_BENEF1 {
			if disabelOpts == 1 {
				benefitValue = benefitUnits * 2
			} else {
				benefitValue = benefitUnits
			}
		}
	}
	return benefitValue
}

func (p propsTaxingBase) BonusChildRaw(income int32, benefit int32, rebated int32) int32 {
	var bonusForChild Decimal = NewFromInt32(min32(0, rebated - benefit)).Neg()

	if p.MarginIncomeOfTaxBonus() > 0 {
		if income < p.MarginIncomeOfTaxBonus() {
			bonusForChild = Zero
		}
	}
	return types.RoundToInt(bonusForChild)
}

func (p propsTaxingBase) BonusChildFix(income int32, benefit int32, rebated int32) int32 {
	var childBonus = p.BonusChildRaw(income, benefit, rebated)

	if p.MinAmountOfTaxBonus() > 0 {
		if childBonus < p.MinAmountOfTaxBonus() {
			return 0
		}
	}
	if p.MaxAmountOfTaxBonus() > 0	{
		if childBonus > p.MaxAmountOfTaxBonus() {
			return p.MaxAmountOfTaxBonus()
		}
	}
	return childBonus
}

func (p propsTaxingBase) TaxableIncomeSupers(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.TaxableIncomeBasis(incomeResult + healthResult + socialResult)
}

func (p propsTaxingBase) TaxableIncomeBasis(incomeResult int32) int32 {
	taxableSuper := max32(0, incomeResult)
	return taxableSuper
}

func (p propsTaxingBase) RoundedRawBaseAdvances(incomeResult int32) int32 {
	amountForCalc := p.TaxableIncomeBasis(incomeResult)

	var advanceBase int32 = 0
	if amountForCalc <= p.MarginIncomeOfRounding() {
		advanceBase = p.intTaxRoundUp(NewFromInt32(amountForCalc))
		return advanceBase
	}
	advanceBase = p.intTaxRoundNearUp(NewFromInt32(amountForCalc), 100)
	return advanceBase
}

func (p propsTaxingBase) RoundedBaseAdvances(incomeResult int32, healthResult int32, socialResult int32) int32 {
	advanceBase := p.RoundedRawBaseAdvances(incomeResult + healthResult + socialResult)

	return advanceBase
}

func (p propsTaxingBase) RoundedBaseSolidary(incomeResult int32) int32 {
	var solidaryBase int32 = 0

	taxableIncome := max32(0, incomeResult)
	if p.MarginIncomeOfSolidary() != 0 {
		solidaryBase = max32(0, taxableIncome - p.MarginIncomeOfSolidary())
	}
	return solidaryBase
}

func (p propsTaxingBase) RoundedAdvancesPaym(supersResult int32, basisResult int32) int32 {
	return 0
}

func (p propsTaxingBase) RoundedSolidaryPaym(basisResult int32) int32 {
	factorSolidary := types.Divide(p.FactorSolidary(), NewFromInt32(100))

	var solidaryTaxing int32 = 0
	if p.MarginIncomeOfSolidary() != 0 {
		solidaryTaxing = p.intTaxRoundUp(types.Multiply(NewFromInt32(basisResult), factorSolidary))
	}

	return solidaryTaxing
}

func (p propsTaxingBase) RoundedBaseWithhold(incomeResult int32) int32 {
	amountForCalc := max32(0, incomeResult)
	withholdBase := p.intTaxRoundDown(NewFromInt32(amountForCalc))

	return withholdBase
}

func (p propsTaxingBase) RoundedWithholdPaym(supersResult int32, basisResult int32) int32 {
	factorWithhold := types.Divide(p.FactorWithhold(), NewFromInt32(100))

	var withholdTaxing int32 = max32(0, supersResult)
	if withholdTaxing > 0 {
		withholdTaxing = p.intTaxRoundDown(types.Multiply(NewFromInt32(supersResult), factorWithhold))
	}
	return withholdTaxing
}

