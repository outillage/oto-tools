package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "oto-tools",
		Short: "Tooling for the great oto lib",
		Long:  "Tooling for the great oto lib",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("There is no root command. Please check oto-tools --help.")
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
