package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type PropsSocial2012 struct {
	propsSocialBase
}

func (p PropsSocial2012) MaxAnnualsBasis() int32 {
	return p.propsSocialBase.MaxAnnualsBasis()
}

func (p PropsSocial2012) FactorEmployer() Decimal {
	return p.propsSocialBase.factorEmployer
}

func (p PropsSocial2012) FactorEmployerHigher() Decimal {
	return p.propsSocialBase.factorEmployerHigher
}

func (p PropsSocial2012) FactorEmployee() Decimal {
	return p.propsSocialBase.factorEmployee
}

func (p PropsSocial2012) FactorEmployeeGarant() Decimal {
	return p.propsSocialBase.factorEmployeeGarant
}

func (p PropsSocial2012) FactorEmployeeReduce() Decimal {
	return p.propsSocialBase.factorEmployeeReduce
}

func (p PropsSocial2012) MarginIncomeEmp() int32 {
	return p.propsSocialBase.marginIncomeEmp
}

func (p PropsSocial2012) MarginIncomeAgr() int32 {
	return p.propsSocialBase.marginIncomeAgr
}

func (p PropsSocial2012) ValueEquals(otherSocial IPropsSocial) bool {
	return  p.propsSocialBase.ValueEquals(otherSocial)
}

func (p PropsSocial2012) HasParticy(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.propsSocialBase.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p PropsSocial2012) hasTermExemptionParticy(_term types.WorkSocialTerms) bool {
	return false
}
func (p PropsSocial2012) hasIncomeBasedEmploymentParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_SMALLS_EMPL

}
func (p PropsSocial2012) hasIncomeBasedAgreementsParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_AGREEM_TASK
}
func (p PropsSocial2012) hasIncomeCumulatedParticy(_term types.WorkSocialTerms) bool {
	var particy bool = false
	switch _term {
	case types.SOCIAL_TERM_EMPLOYMENTS: particy = false
	case types.SOCIAL_TERM_AGREEM_TASK: particy = true
	case types.SOCIAL_TERM_SMALLS_EMPL: particy = false
	case types.SOCIAL_TERM_SHORTS_MEET: particy = false
	case types.SOCIAL_TERM_SHORTS_DENY: particy = false
	case types.SOCIAL_TERM_BY_CONTRACT: particy = false
	default:
		particy = false
	}
	return particy
}

func (p PropsSocial2012) RoundedEmployeePaym(basisResult int32) int32 {
	factorEmployee := types.Divide(p.FactorEmployee(), NewFromInt32(100))
	return p.propsSocialBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployee))
}

func (p PropsSocial2012) RoundedEmployerPaym(basisResult int32) int32 {
	factorEmployer := types.Divide(p.FactorEmployer(), NewFromInt32(100))
	return p.propsSocialBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployer))
}

func (p PropsSocial2012) ResultOvercaps(baseSuma int32, overCaps int32) OvercapsResultPair {
	maxBaseEmployee := max32(0, baseSuma - overCaps)
	empBaseOvercaps := max32(0, baseSuma - maxBaseEmployee)
	valBaseOvercaps := max32(0, overCaps - empBaseOvercaps)
	return OvercapsResultPair {maxBaseEmployee, valBaseOvercaps }
}

func (p PropsSocial2012) AnnualsBasisCut(incomeList []ParticySocialTarget, annuityBasis int32) ParticySocialResultTriple {
	return p.propsSocialBase.AnnualsBasisCut(incomeList, annuityBasis)
}

func NewPropsSocial2012(versionId types.IVersionId,
	maxAnnualsBasis int32,
	factorEmployer Decimal,
	factorEmployerHigher Decimal,
	factorEmployee Decimal,
	factorEmployeeGarant Decimal,
	factorEmployeeReduce Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsSocial {
	return PropsSocial2012{
		propsSocialBase: propsSocialBase {
			propsBase:            propsBase{ Version: versionId },
			maxAnnualsBasis:      maxAnnualsBasis,
			factorEmployer:       factorEmployer,
			factorEmployerHigher: factorEmployerHigher,
			factorEmployee:       factorEmployee,
			factorEmployeeGarant: factorEmployeeGarant,
			factorEmployeeReduce: factorEmployeeReduce,
			marginIncomeEmp:      marginIncomeEmp,
			marginIncomeAgr:      marginIncomeAgr,
		},
	}
}

func EmptyPropsSocial2012() IPropsSocial {
	return PropsSocial2012{
		propsSocialBase: propsSocialBase {
			propsBase:            propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
			maxAnnualsBasis:      0,
			factorEmployer:       NewFromFloat(0),
			factorEmployerHigher: NewFromFloat(0),
			factorEmployee:       NewFromFloat(0),
			factorEmployeeGarant: NewFromFloat(0),
			factorEmployeeReduce: NewFromFloat(0),
			marginIncomeEmp:      0,
			marginIncomeAgr:      0,
		},
	}
}

