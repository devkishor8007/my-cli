package main

import (
	"fmt"
	"os"
	"strings"

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

var cmdEcho = &cobra.Command{
	Use:   "echo [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Print: " + strings.Join(args, " "))
	},
}

var uppercaseCmd = &cobra.Command{
	Use:     "uppercase",
	Short:   "Uppercase a string",
	Aliases: []string{"upper"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := strings.ToUpper(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(cmdEcho)
	rootCmd.AddCommand(uppercaseCmd)

	rootCmd.PersistentFlags().StringVarP(&userName, "msg", "m", "", "please enter a valid userName")
	rootCmd.PersistentFlags().BoolVarP(&showUser, "showUser", "i", false, "Display the userName if valid")
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
