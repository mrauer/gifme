package main

import (
	"github.com/mrauer/gifme/lib"
	"github.com/xfrr/goffmpeg/transcoder"
)

const (
	INPUT_PATH   = "example/input.mp4"
)

func main() {
	trans := new(transcoder.Transcoder)

	err := lib.PaletteGen(trans, INPUT_PATH)

	if err == nil {
		err = lib.PaletteUse(INPUT_PATH)
	}

	if err == nil {
		lib.CompressGit()
	}

	return
}
