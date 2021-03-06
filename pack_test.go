package pack

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	pack := NewPack()
	actual := pack.Container()
	expected := "demo"
	if !strings.HasSuffix(actual, expected) {
		t.Fatalf("Expected container expected %v but got %v", expected, actual)
		return
	}

	loaded, err := pack.Load()
	if err != nil {
		t.Fatalf("Unable to load actual %v", actual)
		return
	}

	if len(loaded) != 1 {
		t.Fatalf("Loaded %v files but expected %v", 1, len(loaded))
		return
	}

	html, htmlErr := pack.String("cmd/index.html")
	if htmlErr != nil {
		t.Fatalf("Unexpected error while reading contents %v", htmlErr)
		return
	}

	md5Expected := "b7a9a1f2ef4a3344f2e8f25df63d7ba1"
	md5Actual := md5.Sum([]byte(*html))
	if md5Actual != md5Actual {
		t.Fatalf("File md5 %v files but expected %v", md5Actual, md5Expected)
		return
	}

	pipe, pipeErr := pack.Pipe("cmd/index.html")
	if pipeErr != nil {
		t.Fatalf("Unexpected error while setting up pipe %v", pipeErr)
		return
	}

	data, readErr := ioutil.ReadAll(pipe)
	if readErr != nil {
		t.Fatalf("Unexpected error while setting up pipeing contents %v", readErr)
		return
	}

	if len(data) != 75 {
		t.Fatalf("Pipe len was %v but expected %v", len(data), 75)
		return
	}
}

func TestFile(t *testing.T) {
	pack := NewPack()

	_, err := pack.Load()
	if err != nil {
		t.Fatalf("Unable to load")
		return
	}

	filePath, err := pack.File("cmd/index.html")
	if err != nil {
		t.Fatal("Error while reading file from pack", err)
		return
	}
	defer os.Remove(*filePath)

	stat, err := os.Stat(*filePath)
	if err != nil {
		t.Fatal("Error while stat-ing file from pack", err)
		return
	}

	if stat.Size() != 75 {
		t.Fatalf("File sizes differ %v vs expected 75", stat.Size())
		return
	}
}
