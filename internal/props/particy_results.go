package props

import "github.com/mzdyhrave/legaliosgo/internal/types"

func cmpInt32(a, b int32) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

func cmpInt16(a, b int16) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}

type ParticyHealthTarget struct {
	contractCode int16
	subjectType  types.WorkTaxingTerms
	interestCode int16
	subjectTerm  types.WorkHealthTerms
	particyCode  int16
	targetsBase  int32
}

func (p *ParticyHealthTarget) AddTargetBase(targetsBase int32) int32 {
	p.targetsBase += targetsBase
	return p.targetsBase
}

func (p ParticyHealthTarget) IncomeScore() int32 {
	var resultType int32 = 0
	switch p.subjectType {
	case types.TAXING_TERM_EMPLOYMENTS:
		resultType = 900
	case types.TAXING_TERM_AGREEM_TASK:
		resultType = 100
	case types.TAXING_TERM_STATUT_PART:
		resultType = 500
	case types.TAXING_TERM_BY_CONTRACT:
		resultType = 0
	}
	var resultBase int32 = 0
	switch p.subjectTerm {
	case types.HEALTH_TERM_EMPLOYMENTS:
		resultBase = 9000
	case types.HEALTH_TERM_AGREEM_WORK:
		resultBase = 5000
	case types.HEALTH_TERM_AGREEM_TASK:
		resultBase = 4000
	case types.HEALTH_TERM_BY_CONTRACT:
		resultBase = 0
	}
	var interestRes int32 = 0
	if p.interestCode == 1 {
		interestRes = 10000
	}
	var particyRes int32 = 0
	if p.particyCode == 1 {
		particyRes = 100000
	}
	return resultType + resultBase + interestRes + particyRes
}
func (p ParticyHealthTarget) ResultComparator() func (x ParticyHealthTarget, y ParticyHealthTarget) int {
	return func (x ParticyHealthTarget, y ParticyHealthTarget) int {
		var xIncomeScore = x.IncomeScore()
		var yIncomeScore = y.IncomeScore()

		if cmpInt32(xIncomeScore, yIncomeScore) == 0 {
			return cmpInt16(x.contractCode, y.contractCode)
		}
		return cmpInt32(yIncomeScore, xIncomeScore)
	}
}

type ParticyHealthResult struct {
	contractCode int16
	subjectType  types.WorkTaxingTerms
	interestCode int16
	subjectTerm  types.WorkHealthTerms
	particyCode  int16
	targetsBase  int32
	resultsBase  int32
}

type ParticySocialTarget struct {
	contractCode int16
	subjectType  types.WorkTaxingTerms
	interestCode int16
	subjectTerm  types.WorkSocialTerms
	particyCode  int16
	targetsBase  int32
}

func (p *ParticySocialTarget) AddTargetBase(targetsBase int32) int32 {
	p.targetsBase += targetsBase
	return p.targetsBase
}

func (p ParticySocialTarget) IncomeScore() int32 {
	var resultType int32 = 0
	switch p.subjectType {
	case types.TAXING_TERM_EMPLOYMENTS:
		resultType = 900
	case types.TAXING_TERM_AGREEM_TASK:
		resultType = 100
	case types.TAXING_TERM_STATUT_PART:
		resultType = 500
	case types.TAXING_TERM_BY_CONTRACT:
		resultType = 0
	}
	var resultBase int32 = 0
	switch p.subjectTerm {
	case types.SOCIAL_TERM_EMPLOYMENTS:
		resultBase = 9000
	case types.SOCIAL_TERM_SMALLS_EMPL:
		resultBase = 5000
	case types.SOCIAL_TERM_SHORTS_MEET:
		resultBase = 4000
	case types.SOCIAL_TERM_SHORTS_DENY:
		resultBase = 0
	case types.SOCIAL_TERM_BY_CONTRACT:
		resultBase = 0
	case types.SOCIAL_TERM_AGREEM_TASK:
		resultBase = 0
	}
	var interestRes int32 = 0
	if p.interestCode == 1 {
		interestRes = 10000
	}
	var particyRes int32 = 0
	if p.particyCode == 1 {
		particyRes = 100000
	}
	return resultType + resultBase + interestRes + particyRes
}

func (p ParticySocialTarget) ResultComparator() func (x ParticySocialTarget, y ParticySocialTarget) int {
	return func (x ParticySocialTarget, y ParticySocialTarget) int {
		var xIncomeScore = x.IncomeScore()
		var yIncomeScore = y.IncomeScore()

		if cmpInt32(xIncomeScore, yIncomeScore) == 0 {
			return cmpInt16(x.contractCode, y.contractCode)
		}
		return cmpInt32(yIncomeScore, xIncomeScore)
	}
}

type ParticySocialResult struct {
	contractCode int16
	subjectType  types.WorkTaxingTerms
	interestCode int16
	subjectTerm  types.WorkSocialTerms
	particyCode  int16
	targetsBase  int32
	resultsBase  int32
}

type ParticyHealthResultTriple struct {
	maxBase int32
	remBase int32
	resList []ParticyHealthResult
}

type ParticySocialResultTriple struct {
	maxBase int32
	remBase int32
	resList []ParticySocialResult
}

type OvercapsResultPair struct {
	maxBase int32
	remBase int32
}
