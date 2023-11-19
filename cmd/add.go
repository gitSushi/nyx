/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Aliases: []string{"a"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		arguments := os.Args
		if len(arguments) < 3 {
			fmt.Println("nyx command directory controller table [-d database]")
			return
		}
		folderName := arguments[2]
		path := "/home/" + folderName + "/"
		controllerName := "C" + arguments[3]
		filePath := path + controllerName + ".php"

		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Println(path + " does not exist !", err)
			return
		}

		mode := fileInfo.Mode()
		if mode.IsDir() {
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("File could not be created !", err)
				return
			}
			defer file.Close()

			fmt.Fprintf(file, "<?php\n\n")
			fmt.Fprintf(file, "class %s {\n\n", controllerName)
			fmt.Fprintf(file, "\tpublic function __constructor() {}\n\n")
			fmt.Fprintf(file, "}\n")

			fmt.Println("CController.php was created")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
