package lib

import (
	"fmt"

	"os/exec"

	"github.com/xfrr/goffmpeg/transcoder"

	"github.com/google/uuid"
)

const (
	PALETTEGEN_ARGS = "fps=10,scale='min(320,iw)':'min(240,ih)',crop=ih:ih,setsar=1,palettegen"
	PALETTEUSE_ARGS = "scale='min(320,iw)':-1,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"
	OUTPUT_GIT      = "output.gif"
)

var (
	palette_path      = ""
	intermediate_path = ""
)

// Generates one palette for a whole video stream.
func PaletteGen(trans *transcoder.Transcoder, input_path string) error {
	err := trans.Initialize(input_path, palette_path)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	trans.MediaFile().SetVideoFilter(PALETTEGEN_ARGS)

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println(msg)
	}

	err = <-done

	return err
}

// Uses a palette to downsample an input video stream.
func PaletteUse(input_path string) error {
	cmd := exec.Command("ffmpeg", []string{"-y", "-i", input_path, "-i", palette_path, "-filter_complex", PALETTEUSE_ARGS, intermediate_path}...)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}

// Attempts lossless reduction of an image's file size.
func CompressGit() {
	cmd := exec.Command("gifsicle", []string{"-i", intermediate_path, "-O3", "--colors", "256", "-o", OUTPUT_GIT}...)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	palette_path = fmt.Sprintf("/tmp/%s.png", uuid.New())
	intermediate_path = fmt.Sprintf("/tmp/%s.gif", uuid.New())
}
