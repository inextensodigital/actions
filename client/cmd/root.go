package cmd

import (
	"fmt"
	"os"

	"github.com/inextensodigital/actions/client/parser"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Command your github action in a cli",
	Long:  ``,
}

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Check file integrity",
	Run: func(cmd *cobra.Command, args []string) {
		parser.LoadData()
		fmt.Println("Configuration ok")
	},
}

var initCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize file integrity",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(".github/main.workflow"); os.IsNotExist(err) {
			emptyFile, err := os.Create(".github/main.workflow")
			if err != nil {
				fmt.Println(err)
			}
			emptyFile.Close()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(lintCmd)
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
}
