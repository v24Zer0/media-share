package mock

type MockIDProvider struct{}

func (provider MockIDProvider) GenerateID() string {
	return "mock-image-id-00"
}

func (provider MockIDProvider) ValidateUUID(id string) bool {
	return true
}
