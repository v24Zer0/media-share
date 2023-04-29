package util

import (
	"regexp"

	"github.com/google/uuid"
)

type IDProvider interface {
	GenerateID() string
	ValidateUUID(id string) bool
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

func (provider DefaultProvider) ValidateUUID(id string) bool {
	regex, err := regexp.Compile("")
	if err != nil {
		return false
	}

	return regex.MatchString(id)
}
