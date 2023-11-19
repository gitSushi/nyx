/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CityDescription struct {
    Field string `json:"Field"`
    Type string `json:"Type"`
    Null string `json:"Null"`
    Key sql.NullString
    Default sql.NullString
    Extra sql.NullString
}

func chooseType(str string) string {
	if len(str) > 7 && str[:7] == "varchar" {
		return "string"
	} else {
		return str
	}
}

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

			// Obviously those infos should be hidden from prying eyes
			db, err := sql.Open("mysql", "root:root@tcp(172.20.0.2:3306)/insee")
			if err != nil {
				// panic(err.Error())
				fmt.Println("Could not connect to DB !", err)
				return
			}
			defer db.Close()

			results, err := db.Query("DESC CITY")
			if err !=nil {
				fmt.Println("Could not query !", err)
				return
			}

			fmt.Fprintf(file, "<?php\n\n")
			fmt.Fprintf(file, "class %s {\n\n", controllerName)
			fmt.Fprintf(file, "\t// properties\n")

			var cityDescriptions []CityDescription
			for results.Next() {
				var cityDescription CityDescription
				err = results.Scan(&cityDescription.Field, &cityDescription.Type, &cityDescription.Null, &cityDescription.Key, &cityDescription.Default, &cityDescription.Extra)
				if err !=nil {
					fmt.Println("Could not scan result !", err)
					return
				}
				cityDescriptions = append(cityDescriptions, cityDescription)
				fmt.Fprintf(file, "\tprivate %s;\n", cityDescription.Field)
			}

			fmt.Fprintf(file, "\n\tpublic function __constructor() {}\n")

			fmt.Fprintf(file, "\n\t// Some methods begining with letter a to f\n")
			fmt.Fprintf(file, "\n\t// Getters\n")
			
			for _, citydesc := range cityDescriptions {
				splitStrs := strings.Split(citydesc.Field, "_")
				var capitalizeds []string
				for _, str := range splitStrs {
					capitalizeds = append(capitalizeds, cases.Title(language.Und).String(str))
				}

				fmt.Fprintf(file, "\tpublic function get%s() { return (%s) $this->%s; }\n", strings.Join(capitalizeds, ""), chooseType(citydesc.Type), citydesc.Field)
			}

			fmt.Fprintf(file, "\n\t// Some methods begining with letter h to z\n")

			fmt.Fprintf(file, "\n\t// TODO setters\n")
			
			fmt.Fprintf(file, "\n}\n")

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
