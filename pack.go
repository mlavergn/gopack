package pack

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	oslog "log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Version export
const Version = "1.2.0"

// logger stand-in
var log *oslog.Logger

// DEBUG toggle
var DEBUG = false

// Pack export
type Pack struct {
	ZipFiles map[string]*zip.File
}

// NewPack init
func NewPack() *Pack {
	if DEBUG {
		log = oslog.New(os.Stderr, "GoPack ", oslog.Ltime|oslog.Lshortfile)
	} else {
		log = oslog.New(ioutil.Discard, "", 0)
	}
	log.Println("NewPack")
	return &Pack{
		ZipFiles: map[string]*zip.File{},
	}
}

// Container export
func (id *Pack) Container() string {
	path, _ := os.Executable()
	if strings.HasSuffix(path, "gopack.test") {
		path, _ = os.Getwd()
		path = filepath.Join(path, "test", "demo")
	}
	return path
}

// Reader export
func (id *Pack) Reader() (*zip.Reader, error) {
	log.Println("Reader")
	executable := id.Container()
	file, _ := os.Open(executable)
	defer file.Close()

	// read the packed length from end of binary
	file.Seek(-10, 2)
	offsetBuffer := make([]byte, 10)
	readLen, readErr := file.Read(offsetBuffer)
	if readLen != 10 || readErr != nil {
		log.Println("Failed to read packed length")
		return nil, readErr
	}

	// convert packed length
	offsetString := string(offsetBuffer)
	// validate that we have a packed length
	const numerics = "1234567890"
	const alphas = "abcdefghijklmnopqrstuvwxyz"
	if strings.ContainsRune(offsetString, 0x00) || strings.ContainsAny(alphas, offsetString) || !strings.ContainsAny(numerics, offsetString[9:]) {
		// no packed content
		return nil, errors.New("No packed content found")
	}
	packLen, contentErr := strconv.Atoi(offsetString)
	if contentErr != nil {
		log.Println("Failed to convert packed length")
		return nil, contentErr
	}

	// read the packed data
	packOffset := int64((packLen + 10) * -1)
	file.Seek(packOffset, 2)
	packBuffer := make([]byte, packLen)
	readLen, readErr = file.Read(packBuffer)
	if readLen != packLen || readErr != nil {
		log.Println("Failed to read packed data")
		return nil, readErr
	}

	// create a packed data reader
	packReader := bytes.NewReader(packBuffer)

	// unzip the packed data
	zipReader, zipErr := zip.NewReader(packReader, int64(packLen))
	if zipErr != nil {
		log.Println("Failed to unzip packed data", zipErr)
		return nil, zipErr
	}

	return zipReader, nil
}

// Extract export
func (id *Pack) Extract() ([]string, error) {
	log.Println("Extract")
	executable := id.Container()
	basePath, _ := filepath.Split(executable)

	keys := []string{}
	zipReader, zipErr := id.Reader()
	if zipErr != nil {
		return nil, zipErr
	}
	for _, zipFile := range zipReader.File {
		keys = append(keys, zipFile.Name)
		outPath, _ := filepath.Split(zipFile.Name)
		fullPath := basePath + zipFile.Name
		log.Println("Extracting: ", fullPath)
		if len(outPath) > 0 {
			os.MkdirAll(basePath+outPath, os.ModeDir|0770)
		}
		dest, destErr := os.Create(fullPath)
		if destErr != nil {
			log.Println("Failed to extract", fullPath, destErr)
			return nil, destErr
		}
		defer dest.Close()
		src, _ := zipFile.Open()
		defer src.Close()
		io.Copy(dest, src)
	}
	return keys, nil
}

// Load export
func (id *Pack) Load() ([]string, error) {
	log.Println("Load")
	zipReader, zipErr := id.Reader()
	if zipErr != nil {
		return nil, zipErr
	}

	keys := []string{}
	for _, zipFile := range zipReader.File {
		keys = append(keys, zipFile.Name)
		id.ZipFiles[zipFile.Name] = zipFile
	}
	id.LoadedPaths()
	return keys, nil
}

// LoadedPaths export
func (id *Pack) LoadedPaths() []string {
	log.Println("LoadedPaths")
	keys := []string{}
	for k := range id.ZipFiles {
		keys = append(keys, k)
	}
	return keys
}

// Pipe export
func (id *Pack) Pipe(filePath string) (io.Reader, error) {
	log.Println("Pipe")
	zipFile := id.ZipFiles[filePath]
	if zipFile == nil {
		err := errors.New("File not found in zip set " + filePath)
		log.Println(err)
		return nil, err
	}
	reader, err := zipFile.Open()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return reader, nil
}

// Bytes export
func (id *Pack) Bytes(filePath string) ([]byte, error) {
	log.Println("Bytes")
	pipe, pipeErr := id.Pipe(filePath)
	if pipeErr != nil {
		return nil, pipeErr
	}
	raw, err := ioutil.ReadAll(pipe)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return raw, nil
}

// String export
func (id *Pack) String(filePath string) (*string, error) {
	log.Println("String")
	raw, err := id.Bytes(filePath)
	if err != nil {
		return nil, err
	}
	result := string(raw)
	return &result, nil
}

// File export
func (id *Pack) File(filePath string) (*string, error) {
	data, err := id.Pipe(filePath)
	if err != nil {
		return nil, err
	}
	_, fileName := filepath.Split(filePath)

	file, err := ioutil.TempFile("", fileName)
	if err != nil {
		return nil, err
	}
	tempPath := file.Name()

	io.Copy(file, data)
	return &tempPath, nil
}
