package post

import (
	"testing"
)

func TestNewPostHandler(t *testing.T) {
	handler := NewPostHandler(nil)

	if handler.service != nil {
		t.Fatal("handler service not nil")
	}

}
