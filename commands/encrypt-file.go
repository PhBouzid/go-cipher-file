package commands

import "github.com/spf13/cobra"

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt file",
	Long:  "",
	Run:   EncryptFile,
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}

func EncryptFile(cmd *cobra.Command, args []string) {

}

