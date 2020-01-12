package gopack

import (
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	pack := NewPack()
	actual := pack.Container()
	expected := "demo"
	if !strings.HasSuffix(actual, expected) {
		t.Fatalf("Expected container expected %v but got %v", expected, actual)
	}
}
