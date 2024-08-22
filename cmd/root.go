/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile    string
	actionFile string
	dryRun     bool
	outputFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gha-docs",
	Short: "Auto generate documentation for GitHub Actions",
	Example: `# Generate docs from test-action/action.yml and write to test-action/README.md
gha-docs -f test-action/action.yml`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		action, err := readActionFile(actionFile)
		if err != nil {
			return err
		}
		action.getPath(actionFile)
		b, err := action.generateDocs()
		if err != nil {
			return err
		}
		if dryRun {
			fmt.Println(string(b))
			return nil
		}
		if outputFile == "README.md" {
			outputFile = action.ActionDir + outputFile
		}
		return writeDocs(outputFile, b)
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
	rootCmd.Flags().StringVarP(&actionFile, "action-file", "f", "action.yml", "Github action file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "render the markdown docs and print to stdout")
	rootCmd.Flags().StringVarP(&outputFile, "output-file", "o", "README.md", "markdown file path relative to each action directory where rendered documentation will be written")
}
