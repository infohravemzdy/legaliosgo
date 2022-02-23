package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type IPropsHealth interface {
	IProps
	MinMonthlyBasis() int32
	MaxAnnualsBasis() int32
	LimMonthlyState() int32
	LimMonthlyDis50() int32
	FactorCompound() Decimal
	FactorEmployee() Decimal
	MarginIncomeEmp() int32
	MarginIncomeAgr() int32

	ValueEquals(otherHealth IPropsHealth) bool
	HasParticy(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32) bool
	RoundedCompoundPaym(basisResult int32) int32
	RoundedEmployeePaym(basisResult int32) int32
	RoundedAugmentEmployeePaym(basisGenerals int32, basisAugment int32) int32
	RoundedAugmentEmployerPaym(basisGenerals int32, baseEmployee int32, baseEmployer int32) int32
	RoundedEmployerPaym(basisResult int32) int32
	AnnualsBasisCut(particyList []IParticyResult, incomeList []IParticyResult, annuityBasis int32) ParticyResultTriple
}

type PropsHealth struct {
	propsBase
	minMonthlyBasis int32
	maxAnnualsBasis int32
	limMonthlyState int32
	limMonthlyDis50 int32
	factorCompound Decimal
	factorEmployee Decimal
	marginIncomeEmp int32
	marginIncomeAgr int32
}

func (p PropsHealth) MinMonthlyBasis() int32 {
	return p.minMonthlyBasis
}

func (p PropsHealth) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p PropsHealth) LimMonthlyState() int32 {
	return p.limMonthlyState
}

func (p PropsHealth) LimMonthlyDis50() int32 {
	return p.limMonthlyDis50
}

func (p PropsHealth) FactorCompound() Decimal {
	return p.factorCompound
}

func (p PropsHealth) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p PropsHealth) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p PropsHealth) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func (p PropsHealth) ValueEquals(otherHealth IPropsHealth) bool {
	if otherHealth == nil {
		return false
	}
	return  p.minMonthlyBasis == otherHealth.MinMonthlyBasis() &&
			p.maxAnnualsBasis == otherHealth.MaxAnnualsBasis() &&
			p.limMonthlyState == otherHealth.LimMonthlyState() &&
			p.limMonthlyDis50 == otherHealth.LimMonthlyDis50() &&
			p.factorCompound.Equal(otherHealth.FactorCompound()) &&
			p.factorEmployee.Equal(otherHealth.FactorEmployee()) &&
			p.marginIncomeEmp == otherHealth.MarginIncomeEmp() &&
			p.marginIncomeAgr == otherHealth.MarginIncomeAgr()
}

func (p PropsHealth) HasParticy(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p PropsHealth) hasTermExemptionParticy(_term types.WorkHealthTerms) bool {
	return false
}
func (p PropsHealth) hasIncomeBasedEmploymentParticy(_term types.WorkHealthTerms) bool {
	return _term == types.HEALTH_TERM_AGREEM_WORK
}
func (p PropsHealth) hasIncomeBasedAgreementsParticy(_term types.WorkHealthTerms) bool {
	return _term == types.HEALTH_TERM_AGREEM_TASK
}
func (p PropsHealth) hasIncomeCumulatedParticy(_term types.WorkHealthTerms) bool {
	var particy bool = false
	switch _term {
	case types.HEALTH_TERM_EMPLOYMENTS: particy = false
	case types.HEALTH_TERM_AGREEM_WORK: particy = true
	case types.HEALTH_TERM_AGREEM_TASK: particy = true
	case types.HEALTH_TERM_BY_CONTRACT: particy = false
	default:
		particy = false
	}
	return particy
}

func (p PropsHealth) decInsuranceRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}

func (p PropsHealth) intInsuranceRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}

func (p PropsHealth) HasParticyWithAdapters(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32,
	hasTermExemptionParticy func (types.WorkHealthTerms) bool,
	hasIncomeBasedEmploymentParticy func (types.WorkHealthTerms) bool,
	hasIncomeBasedAgreementsParticy func (types.WorkHealthTerms) bool,
	hasIncomeCumulatedParticy func (types.WorkHealthTerms) bool) bool {
	var particySpec bool = true
	if hasTermExemptionParticy(term) {
		particySpec = false
	} else if hasIncomeBasedAgreementsParticy(term) && p.MarginIncomeAgr() > 0 {
		particySpec = false
		if hasIncomeCumulatedParticy(term) {
			if incomeTerm >= p.MarginIncomeAgr() {
				particySpec = true
			}
		} else {
			if incomeSpec >= p.MarginIncomeAgr() {
				particySpec = true
			}
		}
	} else if hasIncomeBasedEmploymentParticy(term) && p.MarginIncomeEmp() > 0 {
		particySpec = false
		if hasIncomeCumulatedParticy(term) {
			if incomeTerm >= p.MarginIncomeEmp() {
				particySpec = true
			}
		} else {
			if incomeSpec >= p.MarginIncomeEmp() {
				particySpec = true
			}
		}
	}
	return particySpec
}

func (p PropsHealth) RoundedCompoundPaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(basisResult))

	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorCompound))
}

func (p PropsHealth) RoundedEmployeePaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))
	return p.intInsuranceRoundUp(types.MultiplyAndDivide(NewFromInt32(basisResult), factorCompound, p.FactorEmployee()))
}

func (p PropsHealth) RoundedAugmentEmployeePaym(basisGenerals int32, basisAugment int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))

	return p.intInsuranceRoundUp(
		types.Multiply(NewFromInt32(basisAugment), factorCompound).Add(
			types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))
}

func (p PropsHealth) RoundedAugmentEmployerPaym(basisGenerals int32, baseEmployee int32, baseEmployer int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))

	compoundBasis := baseEmployer + baseEmployee + basisGenerals

	compoundPayment := p.intInsuranceRoundUp(types.Multiply(NewFromInt32(compoundBasis), factorCompound))
	employeePayment := p.intInsuranceRoundUp(types.Multiply(NewFromInt32(baseEmployee), factorCompound).Add(
		types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))

	return max32(0, compoundPayment - employeePayment)
}

func (p PropsHealth) RoundedEmployerPaym(basisResult int32) int32 {
	compoundPayment := p.RoundedCompoundPaym(basisResult)
	employeePayment := p.RoundedEmployeePaym(basisResult)

	return max32(0, compoundPayment - employeePayment)
}

func (p PropsHealth) AnnualsBasisCut(particyList []IParticyResult, incomeList []IParticyResult, annuityBasis int32) ParticyResultTriple {
	return MaximResultCut(particyList, incomeList, annuityBasis, p.MaxAnnualsBasis())
}

func NewPropsHealth(versionId types.IVersionId,
	minMonthlyBasis int32,
	maxAnnualsBasis int32,
	limMonthlyState int32,
	limMonthlyDis50 int32,
	factorCompound Decimal,
	factorEmployee Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsHealth {
	return PropsHealth{
		propsBase:       propsBase{ Version: versionId },
		minMonthlyBasis: minMonthlyBasis,
		maxAnnualsBasis: maxAnnualsBasis,
		limMonthlyState: limMonthlyState,
		limMonthlyDis50: limMonthlyDis50,
		factorCompound:  factorCompound,
		factorEmployee:  factorEmployee,
		marginIncomeEmp: marginIncomeEmp,
		marginIncomeAgr: marginIncomeAgr,
	}
}

func EmptyPropsHealth() IPropsHealth {
	return PropsHealth{
		propsBase:       propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		minMonthlyBasis: 0,
		maxAnnualsBasis: 0,
		limMonthlyState: 0,
		limMonthlyDis50: 0,
		factorCompound:  NewFromFloat(0),
		factorEmployee:  NewFromFloat(0),
		marginIncomeEmp: 0,
		marginIncomeAgr: 0,
	}
}

