package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start Atrean containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting containers...")

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
	rootCmd.AddCommand(upCmd)
}
