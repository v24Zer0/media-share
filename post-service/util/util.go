package util

type Util interface {
	GenerateID() string
}

type UtilImpl struct {
}

func (u UtilImpl) GenerateID() string {
	return ""
}
