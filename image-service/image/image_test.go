package image_test

import (
	"log"
	"os"
	"testing"
)

var mockImage []byte

// Use sql-mock to test and suite from testify to mock and test repo
func TestMain(m *testing.M) {
	// Set working dir for tests
	os.Chdir("../")

	mockImage, _ = os.ReadFile("mock/image/mock_image_00.jpg")

	os.Exit(m.Run())
}

func RestoreStateAfterCreate() {
	// Delete image created
	err := os.Remove("upload/images/mock-image-id-00.jpg")
	if err != nil {
		log.Println("Failed to delete mock image")
	}
}

func RestoreStateAfterDelete() {
	// Restore image deleted
}
