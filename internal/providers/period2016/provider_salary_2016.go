package period2016

import (
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSalary2016 struct {
	providers.ProviderBase
}

func NewProviderSalary2016() providers.IProviderSalary {
	return providerSalary2016{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2016)},
	}
}

func (b providerSalary2016) GetProps(period types.IPeriod) props.IPropsSalary {
	return props.NewPropsSalary(b.Version,
		b.WorkingShiftWeek(period),
		b.WorkingShiftTime(period),
		b.MinMonthlyWage(period),
		b.MinHourlyWage(period))
}


func (b providerSalary2016) WorkingShiftWeek(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_WEEK
}

func (b providerSalary2016) WorkingShiftTime(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_TIME
}

func (b providerSalary2016) MinMonthlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_MONTHLY_WAGE
}

func (b providerSalary2016) MinHourlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_HOURLY_WAGE
}
