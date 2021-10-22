package period2017

import (
	"github.com/mzdyhrave/payrollgo-legalios/internal/providers/period2016"
)

const (
	SOCIAL_VERSION_CODE int16 = 2017
	SOCIAL_MAX_ANNUALS_BASIS = 1355136
	SOCIAL_FACTOR_EMPLOYER float64 = period2016.SOCIAL_FACTOR_EMPLOYER
	SOCIAL_FACTOR_EMPLOYER_HIGHER float64 = period2016.SOCIAL_FACTOR_EMPLOYER_HIGHER
	SOCIAL_FACTOR_EMPLOYEE float64 = period2016.SOCIAL_FACTOR_EMPLOYEE
	SOCIAL_FACTOR_EMPLOYEE_REDUCE float64 = period2016.SOCIAL_FACTOR_EMPLOYEE_REDUCE
	SOCIAL_FACTOR_EMPLOYEE_GARANT float64 = period2016.SOCIAL_FACTOR_EMPLOYEE_GARANT
	SOCIAL_MARGIN_INCOME_EMP = period2016.SOCIAL_MARGIN_INCOME_EMP
	SOCIAL_MARGIN_INCOME_AGR = period2016.SOCIAL_MARGIN_INCOME_AGR
)
