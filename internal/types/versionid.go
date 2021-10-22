package types

const (
	VERSION_ZERO int16 = 0
)

type IVersionId interface {
	GetValue() int16
}

type versionId struct {
	value int16
}

func (v versionId) GetValue() int16 {
	return v.value
}

func NewVersionId() IVersionId {
	return versionId{	value: VERSION_ZERO }
}

func GetVersionId(id int16) IVersionId {
	return versionId{	value: id }
}

