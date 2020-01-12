package pack

import (
	"crypto/md5"
	"io/ioutil"
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

	loaded, err := pack.Load()
	if err != nil {
		t.Fatalf("Unable to load actual %v", actual)
	}

	if len(loaded) != 1 {
		t.Fatalf("Loaded %v files but expected %v", 1, len(loaded))
	}

	html, htmlErr := pack.String("cmd/index.html")
	if htmlErr != nil {
		t.Fatalf("Unexpected error while reading contents %v", htmlErr)
	}

	md5Expected := "b7a9a1f2ef4a3344f2e8f25df63d7ba1"
	md5Actual := md5.Sum([]byte(html))
	if md5Actual != md5Actual {
		t.Fatalf("File md5 %v files but expected %v", md5Actual, md5Expected)
	}

	pipe, pipeErr := pack.Pipe("cmd/index.html")
	if pipeErr != nil {
		t.Fatalf("Unexpected error while setting up pipe %v", pipeErr)
	}

	data, readErr := ioutil.ReadAll(pipe)
	if readErr != nil {
		t.Fatalf("Unexpected error while setting up pipeing contents %v", readErr)
	}

	if len(data) != 75 {
		t.Fatalf("Pipe len was %v but expected %v", len(data), 75)
	}
}
