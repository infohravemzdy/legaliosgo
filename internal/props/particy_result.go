package props

type IParticyResult interface {
	ParticyCode() int32
	ResultBasis() int32
	ResultValue() int32
	SetResultValue(value int32) int32
}

type ParticyResultTriple struct {
	maxBase int32
	remBase int32
	resList []IParticyResult
}

type OvercapsResultPair struct {
	maxBase int32
	remBase int32
}

func MaximResultCut(particyList []IParticyResult, incomeList []IParticyResult, annuityBasis int32, annualyMaxim int32) ParticyResultTriple {
	annualsBasis := max32(0, annualyMaxim - annuityBasis)

	var resultList = ParticyResultTriple {annualyMaxim, annualsBasis, particyList }

	for _, x := range incomeList {
		var cutAnnualsBasis int32 = 0
		var rawAnnualsBasis int32 = x.ResultBasis()
		var remAnnualsBasis int32 = resultList.remBase

		if x.ParticyCode() != 0 {
			cutAnnualsBasis = rawAnnualsBasis
			if resultList.maxBase > 0 {
				ovrAnnualsBasis := max32(0, rawAnnualsBasis - resultList.remBase)
				cutAnnualsBasis = rawAnnualsBasis - ovrAnnualsBasis
			}
			remAnnualsBasis = max32(0, resultList.remBase - cutAnnualsBasis)
		}

		x.SetResultValue(max32(0, cutAnnualsBasis))
		resultList = ParticyResultTriple {resultList.maxBase, remAnnualsBasis, append(resultList.resList, x) }
	}

	return resultList
}
