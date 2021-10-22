package period2013

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSalary2013 struct {
	providers.ProviderBase
}

func NewProviderSalary2013() providers.IProviderSalary {
	return providerSalary2013{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2013)},
	}
}

func (b providerSalary2013) GetProps(period types.IPeriod) props.IPropsSalary {
	return props.NewPropsSalary(b.Version,
		b.WorkingShiftWeek(period),
		b.WorkingShiftTime(period),
		b.MinMonthlyWage(period),
		b.MinHourlyWage(period))
}


func (b providerSalary2013) WorkingShiftWeek(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_WEEK
}

func (b providerSalary2013) WorkingShiftTime(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_TIME
}

func (b providerSalary2013) MinMonthlyWage(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 8) {
		return VAR08_SALARY_MIN_MONTHLY_WAGE
	}
	return SALARY_MIN_MONTHLY_WAGE
}

func (b providerSalary2013) MinHourlyWage(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2013, 8) {
		return VAR08_SALARY_MIN_HOURLY_WAGE
	}
	return SALARY_MIN_HOURLY_WAGE
}
