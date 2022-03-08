package types

import (
	"fmt"
	"time"
)

const (
	TERM_BEG_FINISHED int16 = 32
	TERM_END_FINISHED int16 = 0
	WEEKMON_SUNDAY int16 = 7
	TIME_MULTIPLY_SIXTY int16 = 60
	WEEKDAYS_COUNT int16 = 7
)

// Max returns the largest of x or y.
func max(x, y int16) int16 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smallest of x or y.
func min(x, y int16) int16 {
	if x > y {
		return y
	}
	return x
}

// Max returns the largest of x or y.
func max32(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smallest of x or y.
func min32(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}

type zipInt32 struct {
	a, b int32
}

func zip(a, b []int32) ([]zipInt32, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]zipInt32, len(a), len(a))

	for i, e := range a {
		r[i] = zipInt32{e, b[i]}
	}

	return r, nil
}

func NewDate(year int16, month int16, day int16) time.Time {
	return time.Date(int(year), time.Month(int(month)), int(day), 0, 0, 0, 0, time.UTC)
}

func EmptyMonthSchedule() []int32 {
	return []int32 {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func TotalWeeksHours(template []int32) int32 {
	var result int32 = 0
	for idx, x := range template {
		if idx < 7 {
			result = result + x
		}
	}
	return result
}
func TotalMonthHours(template []int32) int32 {
	var result int32 = 0
	for idx, x := range template {
		if idx < 31 {
			result = result + x
		}
	}
	return result
}
func DaysInMonth(period IPeriod) int16 {
	date := NewDate(period.Year(), period.Month(), 1)

	return int16(date.AddDate(0, 1, -1).Day())
}
func DateOfMonth(period IPeriod, dayOrdinal int16) time.Time {
	var periodDay int16 = min(max(1, dayOrdinal), DaysInMonth(period))

	return NewDate(period.Year(), period.Month(), periodDay)
}
func DayOfWeekMonToSun(periodDateCwd time.Weekday) int16 {
	// DayOfWeek Sunday = 0,
	// Monday = 1, Tuesday = 2, Wednesday = 3, Thursday = 4, Friday = 5, Saturday = 6,
	var dayOfWeek int16 = 0;
	switch periodDateCwd {
	case time.Sunday:
		dayOfWeek = WEEKMON_SUNDAY
	case time.Monday:
		dayOfWeek = 1
	case time.Tuesday:
		dayOfWeek = 2
	case time.Wednesday:
		dayOfWeek = 3
	case time.Thursday:
		dayOfWeek = 4
	case time.Friday:
		dayOfWeek = 5
	case time.Saturday:
		dayOfWeek = 6
	}
	return dayOfWeek
}
func DayOfWeekFromOrdinal(dayOrdinal int16, periodBeginCwd int16) int16 {
	// dayOrdinal 1..31
	// periodBeginCwd 1..7
	// dayOfWeek 1..7

	dayOfWeek := (((dayOrdinal - 1) + (periodBeginCwd - 1)) % 7) + 1

	return dayOfWeek
}

func WeekDayOfMonth(period IPeriod, dayOrdinal int16) int16 {
	periodDate := DateOfMonth(period, dayOrdinal)

	periodDateCwd := periodDate.Weekday()

	return DayOfWeekMonToSun(periodDateCwd)
}
func DateFromInPeriod(period IPeriod, dateFrom *time.Time) int16 {
	var dayTermFrom = TERM_BEG_FINISHED

	periodDateBeg := NewDate(period.Year(), period.Month(), 1)

	if dateFrom != nil	{
		dayTermFrom = int16(dateFrom.Day())
	}

	if dateFrom == nil || dateFrom.Before(periodDateBeg) {
		dayTermFrom = 1
	}
	return dayTermFrom
}

func DateStopInPeriod(period IPeriod, dateEnds *time.Time) int16 {
	var dayTermEnd = TERM_END_FINISHED

	daysPeriod := DaysInMonth(period)

	periodDateEnd := NewDate(period.Year(), period.Month(), daysPeriod)

	if dateEnds != nil {
		dayTermEnd = int16(dateEnds.Day())
	}

	if dateEnds == nil || dateEnds.After(periodDateEnd) {
		dayTermEnd = daysPeriod
	}
	return dayTermEnd
}
func TimesheetWeekSchedule(period IPeriod, secondsWeekly int32, workdaysWeekly int16) []int32 {
	secondsDaily := secondsWeekly / int32(min(workdaysWeekly, WEEKDAYS_COUNT))

	secRemainder := secondsWeekly - (secondsDaily * int32(workdaysWeekly))

	weekSchedule := []int32{
		WeekDaySeconds(1, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(2, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(3, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(4, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(5, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(6, workdaysWeekly, secondsDaily, secRemainder),
		WeekDaySeconds(7, workdaysWeekly, secondsDaily, secRemainder),
	}

	return weekSchedule
}
func WeekDaySeconds(dayOrdinal int16, daysOfWork int16, secondsDaily int32, secRemainder int32) int32 {
	if dayOrdinal < daysOfWork {
		return secondsDaily
	} else {
		if dayOrdinal == daysOfWork {
			return secondsDaily + secRemainder
		}
	}
	return 0
}
func TimesheetFullSchedule(period IPeriod, weekSchedule []int32) []int32 {
	periodDaysCount := DaysInMonth(period)

	periodBeginCwd := WeekDayOfMonth(period, 1)

	var monthSchedule []int32
	for x := int16(1); x <= periodDaysCount; x++ {
		monthSchedule = append(monthSchedule, SecondsFromWeekSchedule(weekSchedule, x, periodBeginCwd))
	}
	return monthSchedule
}
func TimesheetWorkSchedule(monthSchedule []int32, dayTermFrom int16, dayTermStop int16) []int32 {
	var timeSheet []int32
	for idx, x := range monthSchedule {
		timeSheet = append(timeSheet, HoursFromCalendar(dayTermFrom, dayTermStop, int16(idx), x))
	}
	return timeSheet
}
func TimesheetWorkContract(monthContract []int32, monthPosition []int32, dayTermFrom int16, dayTermStop int16) []int32 {
	idxFrom := int(dayTermFrom - 1)
	idxStop := int(dayTermStop - 1)
	var zipedMonth, _ = zip(monthContract, monthPosition)
	var result []int32
	for idx, z := range zipedMonth {
		var res int32 = 0
		if idx >= idxFrom && idx <= idxStop {
			res = z.a + z.b
		}
		result = append(result, res)
	}
	return result
}
func SecondsFromPeriodWeekSchedule(period IPeriod, weekSchedule []int32, dayOrdinal int16) int32 {
	periodBeginCwd := WeekDayOfMonth(period, 1)

	return SecondsFromWeekSchedule(weekSchedule, dayOrdinal, periodBeginCwd)
}
func SecondsFromWeekSchedule(weekSchedule []int32, dayOrdinal int16, periodBeginCwd int16) int32 {
	dayOfWeek := DayOfWeekFromOrdinal(dayOrdinal, periodBeginCwd)

	indexWeek := int(dayOfWeek - 1)

	if indexWeek < 0 || indexWeek >= len(weekSchedule) {
		return 0
	}
	return weekSchedule[indexWeek]
}

func SecondsFromScheduleSeq(timeTable []int32, dayOrdinal int16, dayFromOrd int16, dayEndsOrd int16) int32 {
	if dayOrdinal < dayFromOrd || dayOrdinal > dayEndsOrd {
		return 0
	}

	indexTable := int(dayOrdinal - dayFromOrd)

	if indexTable < 0 || indexTable >= len(timeTable) {
		return 0
	}

	return timeTable[indexTable]
}
func ScheduleBaseSubtract(template []int32, subtract []int32, dayFrom int16, dayStop int16) []int32 {
	idxFrom := int(dayFrom - 1)
	idxStop := int(dayStop - 1)
	var zipedMonth, _ = zip(template, subtract)
	var result []int32
	for idx, z := range zipedMonth {
		var res int32 = 0
		if idx >= idxFrom && idx <= idxStop {
			res = max32(0, z.a - z.b)
		}
		result = append(result, res)
	}
	return result
}
func PlusHoursFromCalendar(dayTermFrom int16, dayTermStop int16, dayIndex int16, partSeconds int32, workSeconds int32) int32 {
	dayOrdinal := dayIndex + 1

	var plusSeconds = workSeconds

	if dayTermFrom > dayOrdinal {
		plusSeconds = 0
	}
	if dayTermStop < dayOrdinal {
		plusSeconds = 0
	}
	return plusSeconds + partSeconds
}
func HoursFromCalendar(dayTermFrom int16, dayTermStop int16, dayIndex int16, workSeconds int32) int32 {
	dayOrdinal := dayIndex + 1

	var workingDay = workSeconds

	if dayTermFrom > dayOrdinal {
		workingDay = 0
	}
	if dayTermStop < dayOrdinal {
		workingDay = 0
	}
	return workingDay
}
