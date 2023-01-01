package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/h2non/bimg"
	"github.com/spf13/cobra"
)


var size int

func init() {
	imgCommand.Flags().IntVarP(&size, "size", "s", 400, "times to echo the input")
	rootCmd.AddCommand(imgCommand)
}

var imgCommand = &cobra.Command{
	Use:   "img [filename]",
	Short: "various functionality to process images according to my specific needs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := processImage(args[0])
		if err != nil {
			fmt.Println(err)
		}
	},
}

func processImage(file string) error {
	dat, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	fName := path.Base(file)
	extName := path.Ext(file)
	bName := fName[:len(fName)-len(extName)]

	options := bimg.Options{
		NoAutoRotate: true,
		Width:        size,
		Height:       size,
		Crop:         true,
		Gravity:      bimg.GravityCentre,
		Type:         bimg.WEBP,
	}
	processed, err := bimg.NewImage(dat).Process(options)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s.webp", bName), processed, 0644)
	if err != nil {
		return err
	}

	return err
}
