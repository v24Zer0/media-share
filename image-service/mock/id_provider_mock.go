package mock

type MockIDProvider struct{}

func (provider MockIDProvider) GenerateID() string {
	return ""
}

func (provider MockIDProvider) ValidateUUID(id string) bool {
	return true
}
