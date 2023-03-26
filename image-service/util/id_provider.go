package util

type IDProvider interface {
	GenerateID() string
	GenerateFileID() string
}

type DefaultProvider struct {
}

func (provider DefaultProvider) GenerateID() string {
	return ""
}

func (provider DefaultProvider) GenerateFileID() string {
	return ""
}
