package main

import (
	"fmt"
	"os"

	"github.com/outillage/oto-tools/internal/generaterunner"
	"github.com/outillage/oto-tools/internal/npm"
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

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generates package json",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			runner := generaterunner.Runner{}

			path := cmd.Flag("path").Value.String()
			packageName := cmd.Flag("package-name").Value.String()
			packageVersion := cmd.Flag("package-version").Value.String()

			return runner.Run(path, npm.Package{Name: packageName, Version: packageVersion})
		},
	}
	generateCmd.PersistentFlags().String("path", "./js", "set path for output of package.json")
	generateCmd.PersistentFlags().String("package-name", "", "set name of project in package.json")
	generateCmd.PersistentFlags().String("package-version", "", "set version of project in package.json")

	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
