package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rrawrriw/life-ctrl"
)

func main() {
	dir := flag.String("dir", "", "Path to stage files")
	out := flag.String("out", "", "Output file")
	flag.Parse()

	err := lifectrl.NewStageFile(*dir, *out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
