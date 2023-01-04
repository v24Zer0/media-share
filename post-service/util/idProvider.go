package util

import "github.com/segmentio/ksuid"

type IDProvider interface {
	GenerateID() string
}

type DefaultIDProvider struct {
}

func (u DefaultIDProvider) GenerateID() string {
	return ksuid.New().String()
}
