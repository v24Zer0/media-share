package util_test

import (
	"testing"

	"github.com/v24Zer0/media-share/image-service/util"
)

func TestGenerateID(t *testing.T) {
	idProvider := util.DefaultProvider{}

	id := idProvider.GenerateID()
	if id == "" {
		t.Fatal("Returned empty string")
	}
}
