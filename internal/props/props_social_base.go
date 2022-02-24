package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type IPropsSocial interface {
	IProps
	MaxAnnualsBasis() int32
	FactorEmployer() Decimal
	FactorEmployerHigher() Decimal
	FactorEmployee() Decimal
	FactorEmployeeGarant() Decimal
	FactorEmployeeReduce() Decimal
	MarginIncomeEmp() int32
	MarginIncomeAgr() int32

	ValueEquals(otherSocial IPropsSocial) bool
	HasParticy(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32) bool
	RoundedEmployeePaym(basisResult int32) int32
	RoundedEmployerPaym(basisResult int32) int32
	ResultOvercaps(baseSuma int32, overCaps int32) OvercapsResultPair
	AnnualsBasisCut(incomeList []ParticySocialTarget, annuityBasis int32) ParticySocialResultTriple
}

type propsSocialBase struct {
	propsBase
	maxAnnualsBasis int32
	factorEmployer Decimal
	factorEmployerHigher Decimal
	factorEmployee Decimal
	factorEmployeeGarant Decimal
	factorEmployeeReduce Decimal
	marginIncomeEmp int32
	marginIncomeAgr int32
}

func (p propsSocialBase) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p propsSocialBase) FactorEmployer() Decimal {
	return p.factorEmployer
}

func (p propsSocialBase) FactorEmployerHigher() Decimal {
	return p.factorEmployerHigher
}

func (p propsSocialBase) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p propsSocialBase) FactorEmployeeGarant() Decimal {
	return p.factorEmployeeGarant
}

func (p propsSocialBase) FactorEmployeeReduce() Decimal {
	return p.factorEmployeeReduce
}

func (p propsSocialBase) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p propsSocialBase) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func (p propsSocialBase) ValueEquals(otherSocial IPropsSocial) bool {
	if otherSocial == nil {
		return false
	}
	return  p.maxAnnualsBasis == otherSocial.MaxAnnualsBasis() &&
			p.factorEmployer.Equal(otherSocial.FactorEmployer()) &&
			p.factorEmployerHigher.Equal(otherSocial.FactorEmployerHigher()) &&
			p.factorEmployee.Equal(otherSocial.FactorEmployee()) &&
			p.factorEmployeeGarant.Equal(otherSocial.FactorEmployeeGarant()) &&
			p.factorEmployeeReduce.Equal(otherSocial.FactorEmployeeReduce()) &&
			p.marginIncomeEmp == otherSocial.MarginIncomeEmp() &&
			p.marginIncomeAgr == otherSocial.MarginIncomeAgr()
}

func (p propsSocialBase) HasParticy(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p propsSocialBase) hasTermExemptionParticy(_term types.WorkSocialTerms) bool {
	return false
}
func (p propsSocialBase) hasIncomeBasedEmploymentParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_SMALLS_EMPL

}
func (p propsSocialBase) hasIncomeBasedAgreementsParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_AGREEM_TASK
}
func (p propsSocialBase) hasIncomeCumulatedParticy(_term types.WorkSocialTerms) bool {
	var particy bool = false
	switch _term {
	case types.SOCIAL_TERM_EMPLOYMENTS: particy = false
	case types.SOCIAL_TERM_AGREEM_TASK: particy = true
	case types.SOCIAL_TERM_SMALLS_EMPL: particy = true
	case types.SOCIAL_TERM_SHORTS_MEET: particy = false
	case types.SOCIAL_TERM_SHORTS_DENY: particy = false
	case types.SOCIAL_TERM_BY_CONTRACT: particy = false
	default:
		particy = false
	}
	return particy
}

func (p propsSocialBase) decInsuranceRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}

func (p propsSocialBase) intInsuranceRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}

func (p propsSocialBase) HasParticyWithAdapters(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32,
	hasTermExemptionParticy func (types.WorkSocialTerms) bool,
	hasIncomeBasedEmploymentParticy func (types.WorkSocialTerms) bool,
	hasIncomeBasedAgreementsParticy func (types.WorkSocialTerms) bool,
	hasIncomeCumulatedParticy func (types.WorkSocialTerms) bool) bool {
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

func (p propsSocialBase) RoundedEmployeePaym(basisResult int32) int32 {
	factorEmployee := types.Divide(p.FactorEmployee(), NewFromInt32(100))
	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployee))
}

func (p propsSocialBase) RoundedEmployerPaym(basisResult int32) int32 {
	factorEmployer := types.Divide(p.FactorEmployer(), NewFromInt32(100))
	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployer))
}

func (p propsSocialBase) ResultOvercaps(baseSuma int32, overCaps int32) OvercapsResultPair {
	maxBaseEmployee := max32(0, baseSuma - overCaps)
	empBaseOvercaps := max32(0, baseSuma - maxBaseEmployee)
	valBaseOvercaps := max32(0, overCaps - empBaseOvercaps)
	return OvercapsResultPair {maxBaseEmployee, valBaseOvercaps }
}

func (p propsSocialBase) AnnualsBasisCut(incomeList []ParticySocialTarget, annuityBasis int32) ParticySocialResultTriple {
	var annualyMaxim int32 = p.MaxAnnualsBasis()

	annualsBasis := max32(0, annualyMaxim - annuityBasis)

	var resultList = ParticySocialResultTriple {annualyMaxim, annualsBasis, []ParticySocialResult{} }

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

		r := ParticySocialResult{
			contractCode: x.contractCode,
			subjectType:  x.subjectType,
			interestCode: x.interestCode,
			subjectTerm:  x.subjectTerm,
			particyCode:  x.particyCode,
			targetsBase:  x.targetsBase,
			resultsBase:  max32(0, cutAnnualsBasis),
		}
		resultList = ParticySocialResultTriple {resultList.maxBase, remAnnualsBasis, append(resultList.resList, r) }
	}

	return resultList
}

