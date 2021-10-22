package period2021

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSalary2021 struct {
	providers.ProviderBase
}

func NewProviderSalary2021() providers.IProviderSalary {
	return providerSalary2021{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2021)},
	}
}

func (b providerSalary2021) GetProps(period types.IPeriod) props.IPropsSalary {
	return props.NewPropsSalary(b.Version,
		b.WorkingShiftWeek(period),
		b.WorkingShiftTime(period),
		b.MinMonthlyWage(period),
		b.MinHourlyWage(period))
}


func (b providerSalary2021) WorkingShiftWeek(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_WEEK
}

func (b providerSalary2021) WorkingShiftTime(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_TIME
}

func (b providerSalary2021) MinMonthlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_MONTHLY_WAGE
}

func (b providerSalary2021) MinHourlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_HOURLY_WAGE
}
