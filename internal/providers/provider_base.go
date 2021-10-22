package providers

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type VERSION int16

type ipropsProvider interface {
	GetVersion() types.IVersionId
	GetVersionValue() int16
}

type ProviderBase struct {
	Version types.IVersionId
}

func (b ProviderBase) GetVersion() types.IVersionId {
	return b.Version
}

func (b ProviderBase) GetVersionValue() int16 {
	return b.Version.GetValue()
}

func (b ProviderBase) IsPeriodGreaterOrEqualThan(period types.IPeriod, yearFrom int16, monthFrom int16) bool {
	return period.Year() > yearFrom || (period.Year() == yearFrom && period.Month() >= monthFrom)
}


