package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "file-encryption",
	Short: "file-encryption",
	Long:  `file-encryption`,
}

func init() {
	rootCmd.AddCommand()
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
