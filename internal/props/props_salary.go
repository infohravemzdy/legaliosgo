package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type IPropsSalary interface {
	IProps
	WorkingShiftWeek() int32
	WorkingShiftTime() int32
	MinMonthlyWage() int32
	MinHourlyWage() int32
	
	ValueEquals(otherSalary IPropsSalary) bool
	CoeffWithPartAndFullHours(fullWorkHours Decimal, partWorkHours Decimal) Decimal
	PaymentWithMonthlyAndFullWeekAndFullAndWorkHours(amountMonthly Decimal, fullWeekHours int32, partWeekHours int32, fullWorkHours int32, partWorkHours int32) Decimal
	PaymentRoundUpWithMonthlyAndFullWeekAndFullAndWorkHours(amountMonthly Decimal, fullWeekHours int32, partWeekHours int32, fullWorkHours int32, partWorkHours int32) Decimal
	PaymentWithMonthlyAndCoeffAndFullAndWorkHours(amountMonthly Decimal, monthlyCoeff Decimal, fullWorkHours int32, partWorkHours int32) Decimal
	PaymentRoundUpWithMonthlyAndCoeffAndFullAndWorkHours(amountMonthly Decimal, monthlyCoeff Decimal, fullWorkHours int32, partWorkHours int32) Decimal
	PaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	PaymentRoundUpWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	RelativeAmountWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	RelativeTariffWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	RelativePaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	ReverzedAmountWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	ReverzedTariffWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	ReverzedPaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal
	PaymentWithTariffAndHours(tariffHourly Decimal, workingsHours Decimal) Decimal
	PaymentRoundUpWithTariffAndHours(tariffHourly Decimal, workingsHours Decimal) Decimal
	TariffWithPaymentAndHours(amountHourly Decimal, workingsHours Decimal) Decimal
	PaymentWithAmountFixed(amountFixed Decimal) Decimal
	PaymentRoundUpWithAmountFixed(amountFixed Decimal) Decimal
	
	HoursToHalfHoursUp(hoursValue Decimal) Decimal
	HoursToQuartHoursUp(hoursValue Decimal) Decimal
	HoursToHalfHoursDown(hoursValue Decimal) Decimal
	HoursToQuartHoursDown(hoursValue Decimal) Decimal
	HoursToHalfHoursNorm(hoursValue Decimal) Decimal
	HoursToQuartHoursNorm(hoursValue Decimal) Decimal
	MoneyToRoundDown(moneyValue Decimal) Decimal
	MoneyToRoundUp(moneyValue Decimal) Decimal
	MoneyToRoundNorm(moneyValue Decimal) Decimal
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

func (p PropsSalary) ValueEquals(otherSalary IPropsSalary) bool {
	if otherSalary == nil {
		return false
	}
	return  p.workingShiftWeek == otherSalary.WorkingShiftWeek() &&
			p.workingShiftTime == otherSalary.WorkingShiftTime() &&
			p.minMonthlyWage == otherSalary.MinMonthlyWage() &&
			p.minHourlyWage == otherSalary.MinHourlyWage()
}

func (p PropsSalary) CoeffWithPartAndFullHours(fullWorkHours Decimal, partWorkHours Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentWithMonthlyAndFullWeekAndFullAndWorkHours(amountMonthly Decimal, fullWeekHours int32, partWeekHours int32, fullWorkHours int32, partWorkHours int32) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentRoundUpWithMonthlyAndFullWeekAndFullAndWorkHours(amountMonthly Decimal, fullWeekHours int32, partWeekHours int32, fullWorkHours int32, partWorkHours int32) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentWithMonthlyAndCoeffAndFullAndWorkHours(amountMonthly Decimal, monthlyCoeff Decimal, fullWorkHours int32, partWorkHours int32) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentRoundUpWithMonthlyAndCoeffAndFullAndWorkHours(amountMonthly Decimal, monthlyCoeff Decimal, fullWorkHours int32, partWorkHours int32) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentRoundUpWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) RelativeAmountWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) RelativeTariffWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) RelativePaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) ReverzedAmountWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) ReverzedTariffWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) ReverzedPaymentWithMonthlyAndCoeffAndWorkCoeff(amountMonthly Decimal, monthlyCoeff Decimal, workingCoeff Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentWithTariffAndHours(tariffHourly Decimal, workingsHours Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentRoundUpWithTariffAndHours(tariffHourly Decimal, workingsHours Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) TariffWithPaymentAndHours(amountHourly Decimal, workingsHours Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentWithAmountFixed(amountFixed Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) PaymentRoundUpWithAmountFixed(amountFixed Decimal) Decimal {
	panic("implement me")
}

func (p PropsSalary) HoursToHalfHoursUp(hoursValue Decimal) Decimal {
	return types.DecRoundUp50(hoursValue)
}
func (p PropsSalary) HoursToQuartHoursUp(hoursValue Decimal) Decimal {
	return types.DecRoundUp25(hoursValue)
}
func (p PropsSalary) HoursToHalfHoursDown(hoursValue Decimal) Decimal {
	return types.DecRoundDown50(hoursValue)
}
func (p PropsSalary) HoursToQuartHoursDown(hoursValue Decimal) Decimal {
	return types.DecRoundDown25(hoursValue)
}
func (p PropsSalary) HoursToHalfHoursNorm(hoursValue Decimal) Decimal {
	return types.DecRoundNorm50(hoursValue)
}
func (p PropsSalary) HoursToQuartHoursNorm(hoursValue Decimal) Decimal {
	return types.DecRoundNorm25(hoursValue)
}
func (p PropsSalary) MoneyToRoundDown(moneyValue Decimal) Decimal {
	return types.DecRoundDown01(moneyValue)
}
func (p PropsSalary) MoneyToRoundUp(moneyValue Decimal) Decimal {
	return types.DecRoundUp01(moneyValue)
}
func (p PropsSalary) MoneyToRoundNorm(moneyValue Decimal) Decimal {
	return types.DecRoundNorm01(moneyValue)
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

