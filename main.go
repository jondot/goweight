package main

import (
	"encoding/json"
	"fmt"

	"github.com/jondot/goweight/pkg"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	jsonOutput = kingpin.Flag("json", "Output json").Short('j').Bool()
	buildTags  = kingpin.Flag("tags", "Build tags").String()
)

func main() {
	kingpin.Version(fmt.Sprintf("%s (%s)", version, commit))
	kingpin.Parse()
	weight := pkg.NewGoWeight()
	if *buildTags != "" {
		weight.BuildArgs = "-tags " + *buildTags
	}

	work := weight.BuildCurrent()
	modules := weight.Process(work)

	if *jsonOutput {
		m, _ := json.Marshal(modules)
		fmt.Print(string(m))
	} else {
		for _, module := range modules {
			fmt.Printf("%8s %s\n", module.SizeHuman, module.Name)
		}
	}
}
