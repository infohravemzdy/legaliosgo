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
	AnnualsBasisCut(particyList []IParticyResult, incomeList []IParticyResult, annuityBasis int32) ParticyResultTriple
}

type PropsSocial struct {
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

func (p PropsSocial) MaxAnnualsBasis() int32 {
	return p.maxAnnualsBasis
}

func (p PropsSocial) FactorEmployer() Decimal {
	return p.factorEmployer
}

func (p PropsSocial) FactorEmployerHigher() Decimal {
	return p.factorEmployerHigher
}

func (p PropsSocial) FactorEmployee() Decimal {
	return p.factorEmployee
}

func (p PropsSocial) FactorEmployeeGarant() Decimal {
	return p.factorEmployeeGarant
}

func (p PropsSocial) FactorEmployeeReduce() Decimal {
	return p.factorEmployeeReduce
}

func (p PropsSocial) MarginIncomeEmp() int32 {
	return p.marginIncomeEmp
}

func (p PropsSocial) MarginIncomeAgr() int32 {
	return p.marginIncomeAgr
}

func (p PropsSocial) ValueEquals(otherSocial IPropsSocial) bool {
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

func (p PropsSocial) HasParticy(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p PropsSocial) hasTermExemptionParticy(_term types.WorkSocialTerms) bool {
	return false
}
func (p PropsSocial) hasIncomeBasedEmploymentParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_SMALLS_EMPL

}
func (p PropsSocial) hasIncomeBasedAgreementsParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_AGREEM_TASK
}
func (p PropsSocial) hasIncomeCumulatedParticy(_term types.WorkSocialTerms) bool {
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

func (p PropsSocial) decInsuranceRoundUp(valueDec Decimal) Decimal {
	return types.DecRoundUp(valueDec)
}

func (p PropsSocial) intInsuranceRoundUp(valueDec Decimal) int32 {
	return types.RoundUp(valueDec)
}

func (p PropsSocial) HasParticyWithAdapters(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32,
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

func (p PropsSocial) RoundedEmployeePaym(basisResult int32) int32 {
	factorEmployee := types.Divide(p.FactorEmployee(), NewFromInt32(100))
	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployee))
}

func (p PropsSocial) RoundedEmployerPaym(basisResult int32) int32 {
	factorEmployer := types.Divide(p.FactorEmployer(), NewFromInt32(100))
	return p.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployer))
}

func (p PropsSocial) ResultOvercaps(baseSuma int32, overCaps int32) OvercapsResultPair {
	maxBaseEmployee := max32(0, baseSuma - overCaps)
	empBaseOvercaps := max32(0, baseSuma - maxBaseEmployee)
	valBaseOvercaps := max32(0, overCaps - empBaseOvercaps)
	return OvercapsResultPair {maxBaseEmployee, valBaseOvercaps }
}

func (p PropsSocial) AnnualsBasisCut(particyList []IParticyResult, incomeList []IParticyResult, annuityBasis int32) ParticyResultTriple {
	return MaximResultCut(particyList, incomeList, annuityBasis, p.MaxAnnualsBasis())
}

func NewPropsSocial(versionId types.IVersionId,
	maxAnnualsBasis int32,
	factorEmployer Decimal,
	factorEmployerHigher Decimal,
	factorEmployee Decimal,
	factorEmployeeGarant Decimal,
	factorEmployeeReduce Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsSocial {
	return PropsSocial{
		propsBase:            propsBase{ Version: versionId },
		maxAnnualsBasis:      maxAnnualsBasis,
		factorEmployer:       factorEmployer,
		factorEmployerHigher: factorEmployerHigher,
		factorEmployee:       factorEmployee,
		factorEmployeeGarant: factorEmployeeGarant,
		factorEmployeeReduce: factorEmployeeReduce,
		marginIncomeEmp:      marginIncomeEmp,
		marginIncomeAgr:      marginIncomeAgr,
	}
}

func EmptyPropsSocial() IPropsSocial {
	return PropsSocial{
		propsBase:            propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		maxAnnualsBasis:      0,
		factorEmployer:       NewFromFloat(0),
		factorEmployerHigher: NewFromFloat(0),
		factorEmployee:       NewFromFloat(0),
		factorEmployeeGarant: NewFromFloat(0),
		factorEmployeeReduce: NewFromFloat(0),
		marginIncomeEmp:      0,
		marginIncomeAgr:      0,
	}
}

