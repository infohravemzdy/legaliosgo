package period2018

import (
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSalary2018 struct {
	providers.ProviderBase
}

func NewProviderSalary2018() providers.IProviderSalary {
	return providerSalary2018{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2018)},
	}
}

func (b providerSalary2018) GetProps(period types.IPeriod) props.IPropsSalary {
	return props.NewPropsSalary(b.Version,
		b.WorkingShiftWeek(period),
		b.WorkingShiftTime(period),
		b.MinMonthlyWage(period),
		b.MinHourlyWage(period))
}


func (b providerSalary2018) WorkingShiftWeek(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_WEEK
}

func (b providerSalary2018) WorkingShiftTime(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_TIME
}

func (b providerSalary2018) MinMonthlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_MONTHLY_WAGE
}

func (b providerSalary2018) MinHourlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_HOURLY_WAGE
}
