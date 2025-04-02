package lib

import (
	"fmt"
	"os"
	"regexp"
)

type Data struct {
	Args  *Args
	Regex *regexp.Regexp
}

func NewData(args *Args) Data {
	if !args.IsExtendedRegex {
		args.Target = regexp.QuoteMeta(args.Target)
	}
	var re *regexp.Regexp
	var err error

	if args.IsIgnoreCase {
		re, err = regexp.Compile("(?i)" + args.Target)
	} else {
		re, err = regexp.Compile(args.Target)
	}

	if err != nil {
		fmt.Println("invalid regex pattern")
		os.Exit(0)
	}

	return Data{args, re}
}
