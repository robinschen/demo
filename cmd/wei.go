/*
Copyright © 2022 Robins Chen <robinschen1989@gmail.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// weiCmd represents the wei command
var weiCmd = &cobra.Command{
	Use:   "wei",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wei called")
	},
}

func init() {
	rootCmd.AddCommand(weiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	weiCmd.PersistentFlags().String("foo", "666", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	weiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
