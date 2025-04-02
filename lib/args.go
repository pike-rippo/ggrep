package lib

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	IsIgnoreCase      bool
	IsAdaptLineNumber bool
	IsAdaptFileName   bool
	IsInvertMatch     bool
	IsExtendedRegex   bool
	IsPrintAllLine    bool
	FilePath          string
	Target            string
}

func (a *Args) Debug() {
	fmt.Println("Args:")
	fmt.Printf("  -i: %t\n", a.IsIgnoreCase)
	fmt.Printf("  -l: %t\n", a.IsAdaptLineNumber)
	fmt.Printf("  -n: %t\n", a.IsAdaptFileName)
	fmt.Printf("  -v: %t\n", a.IsInvertMatch)
	fmt.Printf("  -E: %t\n", a.IsExtendedRegex)
	fmt.Printf("  -a: %t\n", a.IsPrintAllLine)
}

func GetArgs() Args {
	var (
		help              = flag.Bool("h", false, "show help")
		IsIgnoreCase      = flag.Bool("i", false, "Ignore case")
		IsAdaptLineNumber = flag.Bool("l", false, "line number")
		IsAdaptFileName   = flag.Bool("n", false, "file name")
		IsInvertMatch     = flag.Bool("v", false, "iInvert match")
		IsExtendedRegex   = flag.Bool("r", false, "regular expressions")
		IsPrintAllLine    = flag.Bool("a", false, "print all line")
	)
	flag.Parse()

	var FilePath, Target string

	switch len(flag.Args()) {
	case 0:
		*help = true
	case 1:
		Target = flag.Args()[0]
	case 2:
		Target = flag.Args()[0]
		FilePath = flag.Args()[1]
	default:
		Target = flag.Args()[0]
		FilePath = flag.Args()[1]
	}

	if *help {
		filepath, _ := os.Executable()
		fmt.Printf("%s [OPTIONS] <Target> [FilePath]\n", filepath)
		flag.Usage()
		os.Exit(0)
	}

	return Args{
		*IsIgnoreCase,
		*IsAdaptLineNumber,
		*IsAdaptFileName,
		*IsInvertMatch,
		*IsExtendedRegex,
		*IsPrintAllLine,
		FilePath,
		Target,
	}
}
