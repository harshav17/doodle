package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(imgCommand)
}

var imgCommand = &cobra.Command{
	Use:   "img",
	Short: "various functionality to process images according to my specific needs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Does nothing")
	},
}
