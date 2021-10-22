package legaliosgo

import (
	"github.com/mzdyhrave/legaliosgo/internal/factories"
	"github.com/mzdyhrave/legaliosgo/internal/types"
)

type IPeriod = types.IPeriod
type IVersionId = types.IVersionId

func NewVersionId() IVersionId {
	return types.NewVersionId()
}

func GetVersionId(id int16) IVersionId {
	return types.GetVersionId(id)
}

func PeriodZero() IPeriod {
	return types.PeriodZero()
}

func NewPeriodWithCode(code int32) IPeriod {
	return types.GetPeriodWithCode(code)
}

func NewPeriodWithYearMonth(year int16, month int16) IPeriod {
	return types.GetPeriodWithYearMonth(year, month)
}

type IServiceLegalios interface {
	GetBundle(period types.IPeriod)  (provider factories.IBundleProps, err error)
}

type legaliosService struct {
	builder factories.IBundleBuilder
}

func (s legaliosService) GetBundle(period types.IPeriod) (provider factories.IBundleProps, err error) {
	bundle, exists := s.builder.GetBundle(period)
	if exists == false {
		err = newBundleNoneError()
		return nil, err
	}
	if bundle == nil {
		err = newBundleNullError()
		return nil, err
	}
	if bundle.GetPeriodProps().GetCode() == 0 {
		err = newBundleEmptyError()
		return nil, err
	}
	if bundle.GetPeriodProps().GetCode() != period.GetCode() {
		err = newBundleInvalidError()
		return nil, err
	}
	return bundle, nil
}

func NewLegaliosService() IServiceLegalios {
	return legaliosService{ builder: factories.NewBundleBuilder() }
}

