package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type PropsHealth2010 struct {
	propsHealthBase
}

func (p PropsHealth2010) MinMonthlyBasis() int32 {
	return p.propsHealthBase.MinMonthlyBasis()
}

func (p PropsHealth2010) MaxAnnualsBasis() int32 {
	return p.propsHealthBase.MaxAnnualsBasis()
}

func (p PropsHealth2010) LimMonthlyState() int32 {
	return p.propsHealthBase.LimMonthlyState()
}

func (p PropsHealth2010) LimMonthlyDis50() int32 {
	return p.propsHealthBase.LimMonthlyDis50()
}

func (p PropsHealth2010) FactorCompound() Decimal {
	return p.propsHealthBase.FactorCompound()
}

func (p PropsHealth2010) FactorEmployee() Decimal {
	return p.propsHealthBase.FactorEmployee()
}

func (p PropsHealth2010) MarginIncomeEmp() int32 {
	return p.propsHealthBase.MarginIncomeEmp()
}

func (p PropsHealth2010) MarginIncomeAgr() int32 {
	return p.propsHealthBase.MarginIncomeAgr()
}

func (p PropsHealth2010) ValueEquals(otherHealth IPropsHealth) bool {
	return  p.propsHealthBase.ValueEquals(otherHealth)
}

func (p PropsHealth2010) HasParticy(term types.WorkHealthTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.propsHealthBase.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p PropsHealth2010) hasTermExemptionParticy(_term types.WorkHealthTerms) bool {
	return false
}
func (p PropsHealth2010) hasIncomeBasedEmploymentParticy(_term types.WorkHealthTerms) bool {
	return _term == types.HEALTH_TERM_AGREEM_WORK
}
func (p PropsHealth2010) hasIncomeBasedAgreementsParticy(_term types.WorkHealthTerms) bool {
	return false
}
func (p PropsHealth2010) hasIncomeCumulatedParticy(_term types.WorkHealthTerms) bool {
	var particy bool = false
	switch _term {
	case types.HEALTH_TERM_EMPLOYMENTS: particy = false
	case types.HEALTH_TERM_AGREEM_WORK: particy = false
	case types.HEALTH_TERM_AGREEM_TASK: particy = false
	case types.HEALTH_TERM_BY_CONTRACT: particy = false
	default:
		particy = false
	}
	return particy
}

func (p PropsHealth2010) RoundedCompoundPaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(basisResult))

	return p.propsHealthBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorCompound))
}

func (p PropsHealth2010) RoundedEmployeePaym(basisResult int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(100))
	return p.propsHealthBase.intInsuranceRoundUp(types.MultiplyAndDivide(NewFromInt32(basisResult), factorCompound, p.FactorEmployee()))
}

func (p PropsHealth2010) RoundedAugmentEmployeePaym(basisGenerals int32, basisAugment int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(100))

	return p.propsHealthBase.intInsuranceRoundUp(
		types.Multiply(NewFromInt32(basisAugment), factorCompound).Add(
			types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))
}

func (p PropsHealth2010) RoundedAugmentEmployerPaym(basisGenerals int32, baseEmployee int32, baseEmployer int32) int32 {
	factorCompound := types.Divide(p.FactorCompound(), NewFromInt32(100))

	compoundBasis := baseEmployer + baseEmployee + basisGenerals

	compoundPayment := p.propsHealthBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(compoundBasis), factorCompound))
	employeePayment := p.propsHealthBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(baseEmployee), factorCompound).Add(
		types.MultiplyAndDivide(NewFromInt32(basisGenerals), factorCompound, p.FactorEmployee())))

	return max32(0, compoundPayment - employeePayment)
}

func (p PropsHealth2010) RoundedEmployerPaym(basisResult int32) int32 {
	compoundPayment := p.RoundedCompoundPaym(basisResult)
	employeePayment := p.RoundedEmployeePaym(basisResult)

	return max32(0, compoundPayment - employeePayment)
}

func (p PropsHealth2010) AnnualsBasisCut(incomeList []ParticyHealthTarget, annuityBasis int32) ParticyHealthResultTriple {
	return p.propsHealthBase.AnnualsBasisCut(incomeList, annuityBasis)
}

func NewPropsHealth2010(versionId types.IVersionId,
	minMonthlyBasis int32,
	maxAnnualsBasis int32,
	limMonthlyState int32,
	limMonthlyDis50 int32,
	factorCompound Decimal,
	factorEmployee Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsHealth {
	return PropsHealth2010{
		propsHealthBase: propsHealthBase {
			propsBase:       propsBase{ Version: versionId },
			minMonthlyBasis: minMonthlyBasis,
			maxAnnualsBasis: maxAnnualsBasis,
			limMonthlyState: limMonthlyState,
			limMonthlyDis50: limMonthlyDis50,
			factorCompound:  factorCompound,
			factorEmployee:  factorEmployee,
			marginIncomeEmp: marginIncomeEmp,
			marginIncomeAgr: marginIncomeAgr,
		},
	}
}

func EmptyPropsHealth2010() IPropsHealth {
	return PropsHealth2010{
		propsHealthBase: propsHealthBase {
			propsBase:       propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
			minMonthlyBasis: 0,
			maxAnnualsBasis: 0,
			limMonthlyState: 0,
			limMonthlyDis50: 0,
			factorCompound:  NewFromFloat(0),
			factorEmployee:  NewFromFloat(0),
			marginIncomeEmp: 0,
			marginIncomeAgr: 0,
		},
	}
}

