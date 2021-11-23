package main

import (
	"flag"

	"github.com/mrauer/gifme/lib"
	"github.com/xfrr/goffmpeg/transcoder"
)

var (
	inputPath *string
)

func main() {
	trans := new(transcoder.Transcoder)

	flag.Parse()

	err := lib.PaletteGen(trans, *inputPath)

	if err == nil {
		err = lib.PaletteUse(*inputPath)
	}

	if err == nil {
		lib.CompressGit()
	}
}

func init() {
	inputPath = flag.String("input", "example/input.mp4", "The file we want to convert to git")
}
