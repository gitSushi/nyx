/*
Copyright © 2023 gitSushi <lossushi37@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nyx command [-l language] [-d database] directory controller table",
	Short: "Nyx : the CLI library from the shadows.",
	Long: `
  .             *        .     .       .           .      
        _   _           .               ,-,    .      *   
.    . | \ | |    .          .  *      /.(         .      
       |  \| |  _   _  __  __          \ {              . 
 *     |   ' | | | | | \ \/ /     .     '-'      *        
       | |\  | | |_| |  >  <    .              .         .
    .  \_| \_/  \__, | /_/\_\             *               
                 __/ |        .                        .  
 .        .     |___/	.         .        .      .       

Nyx is the Greek goddess and personification of the
night.
	
As such this CLI library helps you build your
application from the shadows.

nyx command [-l language] [-d database] directory controller table

Command __________________________________________________
add                 Adds a controller and a model.

Fields ___________________________________________________
directory           The directory that holds the
                    controller and the model folders.
controller          The stem of your controller and
                    model's name.
table               The name of the related table in your
                    database.

Flags ____________________________________________________

-l | --language     English and French are available.
-d | --database     Specify the database's name.

Example : nyx -d WEBCARE configuration VilleV4 VILLE
	
The name of her Roman equivalent is Nox.
nox will be an alias as of version 2.0.`,

	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nyx.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


