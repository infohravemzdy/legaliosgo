package props

import (
	"github.com/mzdyhrave/payrollgo-legalios/internal/types"
)

type IProps interface {
	GetVersion() types.IVersionId
	GetVersionValue() int16
}

type propsBase struct {
	Version types.IVersionId
}

func (p propsBase) GetVersion() types.IVersionId {
	return p.Version
}

func (p propsBase) GetVersionValue() int16 {
	return p.Version.GetValue()
}

