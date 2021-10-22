package types

const (
	periodZeroCode int32 = 0
)
type IPeriod interface {
	GetCode() int32
	Year() int16
	Month() int16
	Equals(period IPeriod) bool
}

type period struct {
	code int32
}

func (p period) GetCode() int32 {
	return p.code
}

func (p period) Year() int16 {
	return int16(p.code / 100)
}

func (p period) Month() int16 {
	return int16(p.code % 100)
}

func (p period) Equals(period IPeriod) bool {
	return p.code == period.GetCode()
}

func PeriodZero() IPeriod {
	return period{	code: periodZeroCode }
}

func PeriodNew() IPeriod {
	return period{	code: periodZeroCode }
}

func GetPeriodWithCode(code int32) IPeriod {
	return period{	code: code }
}

func GetPeriodWithYearMonth(year int16, month int16) IPeriod {
	return period{	code: int32(year) * 100 + int32(month) }
}


