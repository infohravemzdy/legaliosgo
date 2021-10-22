package period2022

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type providerSalary2022 struct {
	providers.ProviderBase
}

func NewProviderSalary2022() providers.IProviderSalary {
	return providerSalary2022{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2022)},
	}
}

func (b providerSalary2022) GetProps(period types.IPeriod) props.IPropsSalary {
	return props.NewPropsSalary(b.Version,
		b.WorkingShiftWeek(period),
		b.WorkingShiftTime(period),
		b.MinMonthlyWage(period),
		b.MinHourlyWage(period))
}


func (b providerSalary2022) WorkingShiftWeek(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_WEEK
}

func (b providerSalary2022) WorkingShiftTime(period types.IPeriod) int32 {
	return SALARY_WORKING_SHIFT_TIME
}

func (b providerSalary2022) MinMonthlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_MONTHLY_WAGE
}

func (b providerSalary2022) MinHourlyWage(period types.IPeriod) int32 {
	return SALARY_MIN_HOURLY_WAGE
}
