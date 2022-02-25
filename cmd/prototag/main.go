package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wirekang/prototag/pkg/prototag"
)

func main() {
	flag.Usage = func() {
		flag.CommandLine.Output()
		fmt.Fprintf(flag.CommandLine.Output(), "prototag [flags] file\n")
		flag.PrintDefaults()
	}

	outputFile := flag.String("o", "", "output to file")
	isJson := flag.Bool("j", false, "json output")
	isPretty := flag.Bool("p", false, "pretty output")
	isDebug := flag.Bool("d", false, "debug mode")
	flag.Parse()

	var err error
	defer func() {
		var reason interface{}
		if err == nil {
			reason = recover()
		} else {
			reason = err
		}
		if reason != nil {
			if *isDebug {
				panic(reason)
			} else {
				fmt.Printf("%+v\n", reason)
				os.Exit(1)
			}
		}
	}()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	targetFile := flag.Arg(0)
	m, err := prototag.ParseFile(targetFile)
	if err != nil {
		return
	}

	var out []byte
	if *isJson {
		if *isPretty {
			out, err = json.MarshalIndent(*m, "", "    ")
		} else {
			out, err = json.Marshal(*m)
		}
		if err != nil {
			return
		}
	}

	fmt.Printf("%s\n", out)
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, out, 0)
		if err != nil {
			return
		}
	}

}
