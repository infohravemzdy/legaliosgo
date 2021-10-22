package providers

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IProviderSalary interface {
	ipropsProvider
	GetProps(period types.IPeriod) props.IPropsSalary

	WorkingShiftWeek(period types.IPeriod) int32
	WorkingShiftTime(period types.IPeriod) int32
	MinMonthlyWage(period types.IPeriod) int32
	MinHourlyWage(period types.IPeriod) int32
}

