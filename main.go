package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var (
	userName       string
	showUser       bool
	fileName       string
	isDefaultFile  bool
	notepadContent string
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

var notePadCmd = &cobra.Command{
	Use:   "notepad",
	Short: "Create a simple notepad",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case isDefaultFile && fileName == "":
			createNotePad("default.txt", notepadContent)
		case !isDefaultFile && fileName != "" && notepadContent != "":
			createNotePad(fileName, notepadContent)
		case notepadContent == "":
			fmt.Println("Please provide notepad content using the -n flag")
		default:
			fmt.Println("Please provide a valid filename using the -f flag.")
		}
	},
}

func createNotePad(filename string, content string) {
	nfilename := filename
	err := ioutil.WriteFile(nfilename, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func init() {
	rootCmd.AddCommand(cmdEcho)
	rootCmd.AddCommand(uppercaseCmd)
	rootCmd.AddCommand(notePadCmd)

	rootCmd.PersistentFlags().StringVarP(&userName, "msg", "m", "", "please enter a valid userName")
	rootCmd.PersistentFlags().BoolVarP(&showUser, "showUser", "i", false, "Display the userName if valid")

	// notepad
	rootCmd.PersistentFlags().BoolVarP(&isDefaultFile, "isDefaultFile", "d", false, "Use for default filename")
	rootCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "Please enter a filename")
	rootCmd.PersistentFlags().StringVarP(&notepadContent, "fileinput", "n", "", "Please enter the notepad content")
}

func main() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
