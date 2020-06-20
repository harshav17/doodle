package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	xoodle "github.com/harshav17/doodle/xoodle"
)

var exit = os.Exit

func usage() {
	fmt.Fprintf(os.Stderr, "usage: doodle [options]\n")
	fmt.Fprintf(os.Stderr, "options:\n")
	flag.PrintDefaults()
	exit(2)
}

var (
	runCommand = flag.NewFlagSet("run", flag.ExitOnError)
	flagMeth   = runCommand.String("method", "conc", "which method to run")
)

func main() {
	log.SetPrefix("doodle: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	switch *flagMeth {
	case "conc":
		fmt.Println("concing")
		xoodle.TestCores(6)
	case "heap":
		fmt.Println("heap")
	default:
		fmt.Println("default")
	}
}
