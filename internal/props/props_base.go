package props

import (
	"github.com/mzdyhrave/legaliosgo/internal/types"
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

