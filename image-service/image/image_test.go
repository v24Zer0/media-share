package image_test

import (
	"os"
	"testing"
)

// Use sql-mock to test and suite from testify to mock and test repo
func TestMain(m *testing.M) {
	// Set working dir for tests
	os.Chdir("../")

	os.Exit(m.Run())
}
