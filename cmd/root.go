/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version information
var (
	Version   = "0.1.0"
	BuildDate = "2023-05-09"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "atrean",
	Short: "Atrean CLI - Manage your Atrean environment",
	Long: `Atrean CLI is a command line tool for managing your Atrean environment.
It allows you to easily set up and manage the Docker containers needed to run Atrean.

Get started with 'atrean setup' to create your environment files, then use
'atrean setup start' to initialize your containers.`,
	// Handle version flag
	Run: func(cmd *cobra.Command, args []string) {
		// Check if version flag is set
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("Atrean CLI version %s (built on %s)\n", Version, BuildDate)
			return
		}

		// If no flags, show help
		cmd.Help()
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.atrean-cli.yaml)")

	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number of Atrean CLI")
}
