package lib

import (
	"fmt"

	"os/exec"

	"github.com/xfrr/goffmpeg/transcoder"

	"github.com/google/uuid"
)

const (
	palettegenArgs = "fps=10,scale='min(320,iw)':'min(240,ih)',crop=ih:ih,setsar=1,palettegen"
	paletteuseArgs = "scale='min(320,iw)':-1,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"
	outputGit      = "output.gif"
)

var (
	palettePath      = ""
	intermediatePath = ""
)

// Generates one palette for a whole video stream.
func PaletteGen(trans *transcoder.Transcoder, inputPath string) error {
	err := trans.Initialize(inputPath, palettePath)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	trans.MediaFile().SetVideoFilter(palettegenArgs)

	done := trans.Run(true)

	progress := trans.Output()

	for msg := range progress {
		fmt.Println(msg)
	}

	err = <-done

	return err
}

// Uses a palette to downsample an input video stream.
func PaletteUse(inputPath string) error {
	cmd := exec.Command("ffmpeg", []string{"-y", "-i", inputPath, "-i", palettePath, "-filter_complex", paletteuseArgs, intermediatePath}...) //nolint:gosec

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return err
}

// Attempts lossless reduction of an image's file size.
func CompressGit() {
	cmd := exec.Command("gifsicle", []string{"-i", intermediatePath, "-O3", "--colors", "256", "-o", outputGit}...)

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	palettePath = fmt.Sprintf("/tmp/%s.png", uuid.New())
	intermediatePath = fmt.Sprintf("/tmp/%s.gif", uuid.New())
}
