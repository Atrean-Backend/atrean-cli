/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Atrean services",
	Long: `Start the Atrean services using docker-compose.
This command runs docker-compose up to start all services defined
in the docker-compose.yml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Atrean containers...")

		// Run docker-compose up -d
		dockerCmd := exec.Command("docker-compose", "up", "-d")
		dockerCmd.Stdout = os.Stdout
		dockerCmd.Stderr = os.Stderr
		err := dockerCmd.Run()
		if err != nil {
			fmt.Printf("Error starting containers: %s\n", err)
			return
		}

		fmt.Println("Atrean containers started successfully!")
	},
}

func init() {
	setupCmd.AddCommand(startCmd)
}
