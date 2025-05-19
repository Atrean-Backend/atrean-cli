package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start Atrean containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting containers...")

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
	rootCmd.AddCommand(upCmd)
}
