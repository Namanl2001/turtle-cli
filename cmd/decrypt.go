/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"turtle/pkg/decrypt"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt your api keys",
	Long:  `decrypt your secured api keys`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Lookup("name").Value.String()
		password := cmd.Flags().Lookup("password").Value.String()

		fmt.Println(name, password)

		err := decrypt.Decryptshow(name, []byte(password))
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {

	decryptCmd.Flags().StringP("name", "n", "", "name of the api-key")
	decryptCmd.MarkFlagRequired("name")

	decryptCmd.Flags().StringP("password", "p", "", "password to decrypt it with")
	decryptCmd.MarkFlagRequired("password")

	rootCmd.AddCommand(decryptCmd)

	// turtle decrypt --name="naman" --password="@123"
}
