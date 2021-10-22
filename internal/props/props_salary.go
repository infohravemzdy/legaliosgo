package props

import (
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type IPropsSalary interface {
	IProps
	WorkingShiftWeek() int32
	WorkingShiftTime() int32
	MinMonthlyWage() int32
	MinHourlyWage() int32
}

type PropsSalary struct {
	propsBase
	workingShiftWeek int32
	workingShiftTime int32
	minMonthlyWage int32
	minHourlyWage int32
}

func (p PropsSalary) WorkingShiftWeek() int32 {
	return p.workingShiftWeek
}

func (p PropsSalary) WorkingShiftTime() int32 {
	return p.workingShiftTime
}

func (p PropsSalary) MinMonthlyWage() int32 {
	return p.minMonthlyWage
}

func (p PropsSalary) MinHourlyWage() int32 {
	return p.minHourlyWage
}

func NewPropsSalary(versionId types.IVersionId,
	workingShiftWeek int32,
	workingShiftTime int32,
	minMonthlyWage int32,
	minHourlyWage int32) IPropsSalary {
	return PropsSalary{
		propsBase:        propsBase{ Version: versionId },
		workingShiftWeek: workingShiftWeek,
		workingShiftTime: workingShiftTime,
		minMonthlyWage:   minMonthlyWage,
		minHourlyWage:    minHourlyWage,
	}
}

func EmptyPropsSalary() IPropsSalary {
	return PropsSalary{
		propsBase:        propsBase{ Version: types.GetVersionId(types.VERSION_ZERO) },
		workingShiftWeek: 0,
		workingShiftTime: 0,
		minMonthlyWage:   0,
		minHourlyWage:    0,
	}
}

