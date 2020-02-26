package tests

import (
	"fmt"
	"project/src"
	"testing"
)

func TestDiag(t *testing.T) {
	_, err := src.Diag()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
