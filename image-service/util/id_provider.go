package util

import "github.com/google/uuid"

type IDProvider interface {
	GenerateID() string
}

type DefaultProvider struct {
}

func (provider DefaultProvider) GenerateID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return id.String()
}
