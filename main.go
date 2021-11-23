package main

import (
	"flag"

	"github.com/mrauer/gifme/lib"
	"github.com/xfrr/goffmpeg/transcoder"
)

func main() {
	trans := new(transcoder.Transcoder)

	input_path := flag.String("input", "example/input.mp4", "The file we want to convert to git")

	flag.Parse()

	err := lib.PaletteGen(trans, *input_path)

	if err == nil {
		err = lib.PaletteUse(*input_path)
	}

	if err == nil {
		lib.CompressGit()
	}
}
