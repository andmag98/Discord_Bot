package tests

import (
	"project/src"
	"testing"
)

func TestRandomFox(t *testing.T) {
	_, err := src.RandomFox()
	if err != nil {
		t.Errorf("Test RandomFox() failed, received error: %v", err)
	}
}
