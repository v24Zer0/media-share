package util

type IDProvider interface {
	GenerateID() string
}

type DefaultProvider struct {
}

func (provider DefaultProvider) GenerateID() string {
	return ""
}
