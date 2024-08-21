/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gha-docs.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&actionFile, "action-file", "f", "action.yml", "Github action file")
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "render the markdown docs and print to stdout")
	rootCmd.Flags().StringVarP(&outputFile, "output-file", "o", "README.md", "markdown file path relative to each action directory where rendered documentation will be written")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gha-docs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gha-docs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
