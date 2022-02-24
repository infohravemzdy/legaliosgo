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
	AnnualsBasisCut(incomeList []ParticyHealthTarget, annuityBasis int32) ParticyHealthResultTriple
}

type propsHealthBase struct {
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

func (p propsHealthBase) MinMonthlyBasis() int32 {
	return p.minMonthlyBasis
}

func (p propsHealthBase) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p propsHealthBase) LimMonthlyState() int32 {
	return p.limMonthlyState
}

func (p propsHealthBase) LimMonthlyDis50() int32 {
	return p.limMonthlyDis50
}

func (p propsHealthBase) FactorCompound() Decimal {
	return p.factorCompound
}

func (p propsHealthBase) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p propsHealthBase) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p propsHealthBase) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func (p propsHealthBase) ValueEquals(otherHealth IPropsHealth) bool {
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

func (p propsHealthBase) HasParticy(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p propsHealthBase) hasTermExemptionParticy(_term types.WorkHealthTerms) bool {
	return false
}
func (p propsHealthBase) hasIncomeBasedEmploymentParticy(_term types.WorkHealthTerms) bool {
	return _term == types.HEALTH_TERM_AGREEM_WORK
}
func (p propsHealthBase) hasIncomeBasedAgreementsParticy(_term types.WorkHealthTerms) bool {
	return _term == types.HEALTH_TERM_AGREEM_TASK
}
func (p propsHealthBase) hasIncomeCumulatedParticy(_term types.WorkHealthTerms) bool {
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

func (p propsHealthBase) decInsuranceRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}

func (p propsHealthBase) intInsuranceRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}

func (p propsHealthBase) HasParticyWithAdapters(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32,
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

func (p propsHealthBase) RoundedCompoundPaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(basisResult))

	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorCompound))
}

func (p propsHealthBase) RoundedEmployeePaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))
	return p.intInsuranceRoundUp(types.MultiplyAndDivide(NewFromInt32(basisResult), factorCompound, p.FactorEmployee()))
}

func (p propsHealthBase) RoundedAugmentEmployeePaym(basisGenerals int32, basisAugment int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))

	return p.intInsuranceRoundUp(
		types.Multiply(NewFromInt32(basisAugment), factorCompound).Add(
			types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))
}

func (p propsHealthBase) RoundedAugmentEmployerPaym(basisGenerals int32, baseEmployee int32, baseEmployer int32) int32 {
	factorCompound := types.Divide(p.factorCompound, NewFromInt32(100))

	compoundBasis := baseEmployer + baseEmployee + basisGenerals

	compoundPayment := p.intInsuranceRoundUp(types.Multiply(NewFromInt32(compoundBasis), factorCompound))
	employeePayment := p.intInsuranceRoundUp(types.Multiply(NewFromInt32(baseEmployee), factorCompound).Add(
		types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))

	return max32(0, compoundPayment - employeePayment)
}

func (p propsHealthBase) RoundedEmployerPaym(basisResult int32) int32 {
	compoundPayment := p.RoundedCompoundPaym(basisResult)
	employeePayment := p.RoundedEmployeePaym(basisResult)

	return max32(0, compoundPayment - employeePayment)
}

func (p propsHealthBase) AnnualsBasisCut(incomeList []ParticyHealthTarget, annuityBasis int32) ParticyHealthResultTriple {
	var annualyMaxim int32 = p.MaxAnnualsBasis()

	annualsBasis := max32(0, annualyMaxim - annuityBasis)

	var resultList = ParticyHealthResultTriple {annualyMaxim, annualsBasis, []ParticyHealthResult{} }

	for _, x := range incomeList {
		var cutAnnualsBasis int32 = 0
		var rawAnnualsBasis int32 = x.targetsBase
		var remAnnualsBasis int32 = resultList.remBase

		if x.particyCode != 0 {
			cutAnnualsBasis = rawAnnualsBasis
			if resultList.maxBase > 0 {
				ovrAnnualsBasis := max32(0, rawAnnualsBasis - resultList.remBase)
				cutAnnualsBasis = rawAnnualsBasis - ovrAnnualsBasis
			}
			remAnnualsBasis = max32(0, resultList.remBase - cutAnnualsBasis)
		}

		r := ParticyHealthResult{
			contractCode: x.contractCode,
			subjectType:  x.subjectType,
			interestCode: x.interestCode,
			subjectTerm:  x.subjectTerm,
			particyCode:  x.particyCode,
			targetsBase:  x.targetsBase,
			resultsBase:  max32(0, cutAnnualsBasis),
		}
		resultList = ParticyHealthResultTriple {resultList.maxBase, remAnnualsBasis, append(resultList.resList, r) }
	}

	return resultList
}

