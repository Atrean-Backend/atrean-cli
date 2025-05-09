/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Docker containers",
	Long: `Start the Atrean services using docker-compose.
This command runs docker-compose up to start all services defined
in the docker-compose.yml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Atrean containers...")

		// Run docker-compose up -d
		dockerCmd := exec.Command("docker-compose", "up", "-d")
		output, err := dockerCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error starting containers: %s\n", err)
			fmt.Println(string(output))
			return
		}

		fmt.Println("Atrean containers started successfully!")
	},
}

func init() {
	setupCmd.AddCommand(startCmd)
}
