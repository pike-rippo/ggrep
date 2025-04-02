package lib

import (
	"bufio"
	"os"
)

type FileInput struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileInput(filename string) (*FileInput, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	return &FileInput{file: file, scanner: scanner}, nil
}

func (f *FileInput) GetNextLine() *string {
	if f.scanner.Scan() {
		line := f.scanner.Text()
		return &line
	}
	return nil
}

func (f *FileInput) Close() {
	f.file.Close()
}
