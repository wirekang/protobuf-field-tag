package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/wirekang/prototag/pkg/prototag"
	"gopkg.in/yaml.v3"
)

func main() {
	pflag.Usage = func() {
		pflag.CommandLine.SetOutput(os.Stdout)
		fmt.Printf("prototag [flags] [file or stdin]\n")
		pflag.PrintDefaults()
	}

	outputFile := pflag.StringP("output", "o", "", "output to file")
	isJson := pflag.BoolP("json", "j", false, "json output")
	isYaml := pflag.BoolP("yaml", "y", false, "yaml output")
	isPretty := pflag.BoolP("pretty", "p", false, "pretty output")
	isDebug := pflag.BoolP("debug", "d", false, "debug mode")
	help := pflag.Bool("help", false, "")
	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}

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
				fmt.Fprintf(os.Stderr, "%+v\n", reason)
				os.Exit(1)
			}
		}
	}()

	if pflag.NArg() > 1 {
		pflag.Usage()
		return
	}

	var m *prototag.Model
	if pflag.NArg() == 1 {
		targetFile := pflag.Arg(0)
		m, err = prototag.ParseFile(targetFile)
	} else {
		m, err = prototag.Parse(os.Stdin)
	}
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
	} else if *isYaml {
		out, err = yaml.Marshal(*m)
	}
	if err != nil {
		return
	}

	fmt.Printf("%s\n", out)
	if *outputFile != "" {
		err = os.WriteFile(*outputFile, out, 0)
		if err != nil {
			return
		}
	}

}
