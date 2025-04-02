package lib

import (
	"bufio"
	"os"
)

type StdinInput struct {
	scanner *bufio.Scanner
}

func NewStdinInput() *StdinInput {
	scanner := bufio.NewScanner(os.Stdin)
	return &StdinInput{scanner: scanner}
}

func (s *StdinInput) GetNextLine() *string {
	if s.scanner.Scan() {
		line := s.scanner.Text()
		return &line
	}
	return nil
}

func (s *StdinInput) Close() {}
