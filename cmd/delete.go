/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"turtle/pkg/delete"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete your api key",
	Long:  `delete existing api-key`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flags().Lookup("name").Value.String()
		err := delete.Delete(name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("successfully deleted api-key %s \n", name)
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("name", "n", "", "name of the api-key")
	deleteCmd.MarkFlagRequired("name")
	rootCmd.AddCommand(deleteCmd)
}
