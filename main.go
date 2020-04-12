package main

import (
	"fmt"
	"os"

	"github.com/outillage/oto-tools/internal/generaterunner"
	"github.com/outillage/oto-tools/internal/npm"
	"github.com/outillage/oto-tools/internal/publishrunner"
	"github.com/outillage/oto-tools/internal/versioning"
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
			otoTemplate := cmd.Flag("oto-template").Value.String()
			otoDefinitions := cmd.Flag("oto-definitions").Value.String()

			if packageVersion == "" {
				packageVersion = versioning.FindVersion()
			}

			return runner.Run(
				path,
				npm.Package{Name: packageName, Version: packageVersion},
				generaterunner.OtoOptions{TemplatePath: otoTemplate, DefinitionPath: otoDefinitions},
			)
		},
	}
	generateCmd.PersistentFlags().String("path", "./js", "set path for output of package.json")
	generateCmd.PersistentFlags().String("package-name", "", "set name of project in package.json")
	generateCmd.PersistentFlags().String("package-version", "", "set version of project in package.json")
	generateCmd.PersistentFlags().String("oto-template", "", "path to oto template")
	generateCmd.PersistentFlags().String("oto-definitions", "", "path to oto definition")

	rootCmd.AddCommand(generateCmd)

	var NPMPublishCmd = &cobra.Command{
		Use:   "publish-npm",
		Short: "Publishes provided package to NPM",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			runner := publishrunner.Runner{}

			path := cmd.Flag("path").Value.String()
			token := cmd.Flag("token").Value.String()
			registry := cmd.Flag("registry").Value.String()
			owner := cmd.Flag("owner").Value.String()

			return runner.Run(
				path,
				npm.PublishOptions{
					Token:       token,
					RegistryURL: registry,
					Owner:       owner,
				},
			)
		},
	}

	NPMPublishCmd.PersistentFlags().String("path", "./js", "set path for js library to publish")
	NPMPublishCmd.PersistentFlags().String("token", "", "token for authenticating against NPM")
	NPMPublishCmd.PersistentFlags().String("registry", "", "NPM compatible registry to publish to, accepts npm, github, or url")
	NPMPublishCmd.PersistentFlags().String("owner", "", "for Github owner is required to build the registry url")

	rootCmd.AddCommand(NPMPublishCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
