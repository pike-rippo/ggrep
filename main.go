package main

import (
	"fmt"
	"log"
	"main/lib"
	"os"
	"path/filepath"
	"strings"
)

func grep(input *lib.InputInterface, data *lib.Data, fileName string) {
	var sp *string
	var isPrinted bool
	isSeparate := data.Args.IsAdaptFileName || data.Args.IsAdaptLineNumber
	for i := 1; ; i++ {
		sp = lib.GetNextLine(*input)
		if sp == nil {
			break
		}

		indexArr := data.Regex.FindAllStringIndex(*sp, -1)

		if !data.Args.IsPrintAllLine && ((data.Args.IsInvertMatch && len(indexArr) != 0) || (!data.Args.IsInvertMatch && len(indexArr) == 0)) {
			continue
		}

		var sb strings.Builder
		if data.Args.IsAdaptFileName {
			sb.WriteString(fileName + " ")
		}
		if data.Args.IsAdaptLineNumber {
			fmt.Fprintf(&sb, "%4d ", i)
		}
		if isSeparate {
			sb.WriteString("| ")
		}
		sb.WriteString(getHighlightedLine(*sp, indexArr))
		os.Stdout.WriteString(sb.String() + "\n")
		isPrinted = true
	}
	if isPrinted {
		fmt.Println()
	}
}

func getHighlightedLine(line string, indexArr [][]int) string {
	var sb strings.Builder
	currentIndex := 0

	for _, match := range indexArr {
		sb.WriteString(line[currentIndex:match[0]])
		sb.WriteString("\033[032m")
		sb.WriteString(line[match[0]:match[1]])
		sb.WriteString("\033[0m")
		currentIndex = match[1]
	}
	sb.WriteString(line[currentIndex:])
	return sb.String()
}

func main() {
	args := lib.GetArgs()
	data := lib.NewData(&args)

	var input lib.InputInterface
	if args.FilePath == "" {
		input = lib.NewStdinInput()
		grep(&input, &data, "stdin")
		return
	}

	files, err := getMatchingFiles(args.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, filePath := range files {
		input, err = lib.NewFileInput(filePath)
		if err != nil {
			log.Fatal(err)
		}
		grep(&input, &data, filePath)
		input.Close()
	}
}

func getMatchingFiles(pattern string) ([]string, error) {
	var files []string

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	files = append(files, matches...)

	return files, nil
}
