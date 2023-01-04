package mock

type MockIDProvider struct {
}

func (p MockIDProvider) GenerateID() string {
	return "6"
}
