package commands

import "github.com/spf13/cobra"

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt file",
	Long:  "decrypt file",
	Run:   DecryptFile,
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}

func DecryptFile(cmd *cobra.Command, args []string) {

}