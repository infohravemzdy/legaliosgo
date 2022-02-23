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

func (p PropsTaxing) FactorSolidary() Decimal {
	return p.factorSolidary
}

func (p PropsTaxing) FactorTaxRate2() Decimal {
	return p.factorTaxRate2
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

func (p PropsTaxing) MarginIncomeOfSolidary() int32 {
	return p.marginIncomeOfSolidary
}

func (p PropsTaxing) MarginIncomeOfTaxRate2() int32 {
	return p.marginIncomeOfTaxRate2
}

func (p PropsTaxing) MarginIncomeOfWthEmp() int32 {
	return p.marginIncomeOfWthEmp
}

func (p PropsTaxing) MarginIncomeOfWthAgr() int32 {
	return p.marginIncomeOfWthAgr
}

func (p PropsTaxing) ValueEquals(otherTaxing IPropsTaxing) bool {
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

func (p PropsTaxing) intTaxRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}
func (p PropsTaxing) intTaxRoundNearUp(valueDec Decimal, nearest int32) int32 {
	return types.NearRoundUp(valueDec, nearest)
}
func (p PropsTaxing) intTaxRoundNearUp100(valueDec Decimal) int32 {
	return types.NearRoundUp(valueDec, 100)
}
func (p PropsTaxing) intTaxRoundDown(valueDec Decimal) int32 {
	return types.RoundDown(valueDec)
}
func (p PropsTaxing) intTaxRoundNearDown(valueDec Decimal, nearest int32) int32 {
	return types.NearRoundDown(valueDec, nearest)
}
func (p PropsTaxing) intTaxRoundNearDown100(valueDec Decimal) int32 {
	return types.NearRoundDown(valueDec, 100)
}
func (p PropsTaxing) decTaxRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}
func (p PropsTaxing) decTaxRoundNearUp(valueDec Decimal, nearest int32) Decimal {
	return types.DecNearRoundUp(valueDec, nearest)
}
func (p PropsTaxing) decTaxRoundNearUp100(valueDec Decimal) Decimal {
	return types.DecNearRoundUp(valueDec, 100)
}
func (p PropsTaxing) decTaxRoundDown(valueDec Decimal) Decimal {
	return types.DecRoundDown(valueDec)
}
func (p PropsTaxing) decTaxRoundNearDown(valueDec Decimal, nearest int32) Decimal {
	return types.DecNearRoundDown(valueDec, nearest)
}
func (p PropsTaxing) decTaxRoundNearDown100(valueDec Decimal) Decimal {
	return types.DecNearRoundDown(valueDec, 100)
}

func (p PropsTaxing) HasWithholdIncome(termOpt types.WorkTaxingTerms, signOpt types.TaxDeclSignOption, noneOpt types.TaxNoneSignOption, incomeSum int32) bool {
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

func (p PropsTaxing) BenefitAllowancePayer(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		if benefitOpts == types.DECL_TAX_BENEF1 {
			benefitValue = p.AllowancePayer()
		}
	}
	return benefitValue
}

func (p PropsTaxing) BenefitAllowanceDisab(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclDisabOption) int32 {
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

func (p PropsTaxing) BenefitAllowanceStudy(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	var benefitValue int32 = 0
	if signOpts == types.DECL_TAX_DO_SIGNED {
		if benefitOpts == types.DECL_TAX_BENEF1 {
			benefitValue = p.AllowanceStudy()
		}
	}
	return benefitValue
}

func (p PropsTaxing) BenefitAllowanceChild(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption, benefitOrds int32, disabelOpts int32) int32 {
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

func (p PropsTaxing) BonusChildRaw(income int32, benefit int32, rebated int32) int32 {
	var bonusForChild Decimal = NewFromInt32(min32(0, rebated - benefit)).Neg()

	if p.MarginIncomeOfTaxBonus() > 0 {
		if income < p.MarginIncomeOfTaxBonus() {
			bonusForChild = Zero
		}
	}
	return types.RoundToInt(bonusForChild)
}

func (p PropsTaxing) BonusChildFix(income int32, benefit int32, rebated int32) int32 {
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

func (p PropsTaxing) TaxableIncomeSupers(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.TaxableIncomeBasis(incomeResult + healthResult + socialResult)
}

func (p PropsTaxing) TaxableIncomeBasis(incomeResult int32) int32 {
	taxableSuper := max32(0, incomeResult)
	return taxableSuper
}

func (p PropsTaxing) RoundedRawBaseAdvances(incomeResult int32) int32 {
	amountForCalc := p.TaxableIncomeBasis(incomeResult)

	var advanceBase int32 = 0
	if amountForCalc <= p.MarginIncomeOfRounding() {
		advanceBase = p.intTaxRoundUp(NewFromInt32(amountForCalc))
		return advanceBase
	}
	advanceBase = p.intTaxRoundNearUp(NewFromInt32(amountForCalc), 100)
	return advanceBase
}

func (p PropsTaxing) RoundedBaseAdvances(incomeResult int32, healthResult int32, socialResult int32) int32 {
	advanceBase := p.RoundedRawBaseAdvances(incomeResult + healthResult + socialResult)

	return advanceBase
}

func (p PropsTaxing) RoundedBaseSolidary(incomeResult int32) int32 {
	var solidaryBase int32 = 0

	taxableIncome := max32(0, incomeResult)
	if p.MarginIncomeOfSolidary() != 0 {
		solidaryBase = max32(0, taxableIncome - p.MarginIncomeOfSolidary())
	}
	return solidaryBase
}

func (p PropsTaxing) RoundedAdvancesPaym(supersResult int32, basisResult int32) int32 {
	factorAdvances := types.Divide(p.FactorAdvances(), NewFromInt32(100))
	factorTaxRate2 := types.Divide(p.FactorTaxRate2(), NewFromInt32(100))

	var taxRate1Basis int32 = basisResult
	var taxRate2Basis int32 = 0
	if p.MarginIncomeOfTaxRate2() != 0 {
		taxRate1Basis = min32(basisResult, p.MarginIncomeOfTaxRate2())
		taxRate2Basis = max32(0, basisResult - p.MarginIncomeOfTaxRate2())
	}
	var taxRate1Taxing Decimal = Zero
	if basisResult <= p.MarginIncomeOfRounding() {
		taxRate1Taxing = types.Multiply(NewFromInt32(taxRate1Basis), factorAdvances)
	} else 	{
		taxRate1Taxing = types.Multiply(NewFromInt32(taxRate1Basis), factorAdvances)
	}
	var taxRate2Taxing Decimal = Zero
	if p.MarginIncomeOfTaxRate2() != 0 {
		taxRate2Taxing = types.Multiply(NewFromInt32(taxRate2Basis), factorTaxRate2)
	}
	return p.intTaxRoundUp(taxRate1Taxing.Add(taxRate2Taxing))
}

func (p PropsTaxing) RoundedSolidaryPaym(basisResult int32) int32 {
	factorSolidary := types.Divide(p.FactorSolidary(), NewFromInt32(100))

	var solidaryTaxing int32 = 0
	if p.MarginIncomeOfSolidary() != 0 {
		solidaryTaxing = p.intTaxRoundUp(types.Multiply(NewFromInt32(basisResult), factorSolidary))
	}

	return solidaryTaxing
}

func (p PropsTaxing) RoundedBaseWithhold(incomeResult int32) int32 {
	amountForCalc := max32(0, incomeResult)
	withholdBase := p.intTaxRoundDown(NewFromInt32(amountForCalc))

	return withholdBase
}

func (p PropsTaxing) RoundedWithholdPaym(supersResult int32, basisResult int32) int32 {
	factorWithhold := types.Divide(p.FactorWithhold(), NewFromInt32(100))

	var withholdTaxing int32 = max32(0, supersResult)
	if withholdTaxing > 0 {
		withholdTaxing = p.intTaxRoundDown(types.Multiply(NewFromInt32(supersResult), factorWithhold))
	}
	return withholdTaxing
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
	factorSolidary Decimal,
	factorTaxRate2 Decimal,
	minAmountOfTaxBonus int32,
	maxAmountOfTaxBonus int32,
	marginIncomeOfTaxBonus int32,
	marginIncomeOfRounding int32,
	marginIncomeOfWithhold int32,
	marginIncomeOfSolidary int32,
	marginIncomeOfTaxRate2 int32,
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
		factorSolidary:    factorSolidary,
		factorTaxRate2:    factorTaxRate2,
		minAmountOfTaxBonus: minAmountOfTaxBonus,
		maxAmountOfTaxBonus: maxAmountOfTaxBonus,
		marginIncomeOfTaxBonus: marginIncomeOfTaxBonus,
		marginIncomeOfRounding: marginIncomeOfRounding,
		marginIncomeOfWithhold: marginIncomeOfWithhold,
		marginIncomeOfSolidary: marginIncomeOfSolidary,
		marginIncomeOfTaxRate2: marginIncomeOfTaxRate2,
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
		factorSolidary:    NewFromFloat(0),
		factorTaxRate2:    NewFromFloat(0),
		minAmountOfTaxBonus: 0,
		maxAmountOfTaxBonus: 0,
		marginIncomeOfTaxBonus: 0,
		marginIncomeOfRounding: 0,
		marginIncomeOfWithhold: 0,
		marginIncomeOfSolidary: 0,
		marginIncomeOfTaxRate2: 0,
		marginIncomeOfWthEmp: 0,
		marginIncomeOfWthAgr: 0,
	}
}

