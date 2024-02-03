package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/noxworld-dev/noxscript/ns/v3/cmd"
	"github.com/noxworld-dev/noxscript/ns/v3/noxast"
)

func main() {
	var conf noxast.Config
	flag.BoolVar(&conf.DoNotFold, "nofold", false, "do not fold the code")
	flag.BoolVar(&conf.DoNotOptimize, "noopt", false, "do not optimize the code")
	flag.StringVar(&conf.Package, "pkg", "", "package name for the script")

	flag.Parse()
	if err := cmd.Decomp(flag.Args(), &conf); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
