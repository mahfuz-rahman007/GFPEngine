package processor

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type Processor interface {
	Name() string
	Process(path string) (string, error)
}

type SizeProcessor struct{}

func (s SizeProcessor) Name() string {
	return "size"
}

func (s SizeProcessor) Process(path string) (string, error) {
	info,err := os.Stat(path)

	if err!=nil { return "", err}

	return fmt.Sprintf("%d bytes", info.Size()), nil
}

type HashProcessor struct{}

func (h HashProcessor) Name() string {
	return "sha256"
}

func (h HashProcessor) Process(path string) (string, error) {
	file,err := os.Open(path)
	if err!=nil {return "", err}

	defer file.Close()

	hasher := sha256.New()
	if _,err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}