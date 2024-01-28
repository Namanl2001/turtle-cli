package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	encrypt "turtle/pkg/encrypt"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "encrypt your api keys",
	Long:  `encypt and secure your private keys`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Lookup("name").Value.String()
		key := cmd.Flags().Lookup("key").Value.String()
		password := cmd.Flags().Lookup("password").Value.String()

		fmt.Println(key, name, password)

		err := encrypt.Encryptsave(name, []byte(password), []byte(key))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("api-key named %s saved successfully \n", name)
		}
	},
}

func init() {
	encryptCmd.Flags().StringP("name", "n", "", "name of the api-key")
	encryptCmd.MarkFlagRequired("name")

	encryptCmd.Flags().StringP("key", "k", "", "your api-key")
	encryptCmd.MarkFlagRequired("key")

	encryptCmd.Flags().StringP("password", "p", "", "password to encypt it with")
	encryptCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(encryptCmd)

	// turtle encrypt --name="facebook" --key="cnkkn4" --password="naman@12345"
}
