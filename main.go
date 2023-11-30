package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	userName string
	showUser bool
)

var rootCmd = &cobra.Command{
	Use:   "simple",
	Short: "simple cli with cobra",
	Long: `Creating a simple CLI with Cobra and Golang
	to explore and create something useful.`,
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if userName != "" && showUser {
			fmt.Println("Username: " + userName)
		} else if userName != "" {
			fmt.Println("Hello", userName)
		} else {
			fmt.Println("Please provide a valid username using the -m flag.")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&userName, "msg", "m", "", "please enter a valid userName")
	rootCmd.PersistentFlags().BoolVarP(&showUser, "showUser", "i", false, "Display the userName if valid")
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
