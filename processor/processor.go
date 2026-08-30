package processor

import (
	"fmt"
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