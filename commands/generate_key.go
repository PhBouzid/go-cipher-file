package commands

import "github.com/spf13/cobra"

var generateKey = &cobra.Command{
	Use:   "gk",
	Short: "generate key",
	Long:  "genereate key",
	Run:   GenerateKey,
}

func init() {
	rootCmd.AddCommand(generateKey)
}

func GenerateKey(cmd *cobra.Command, args []string) {

}
