package cmd

import (
	"fmt"

	"github.com/inextensodigital/actions/client/parser"
	"github.com/spf13/cobra"
)

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workflow called")
	},
}

var workflowLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List workflow",
	Run: func(cmd *cobra.Command, args []string) {
		for _, workflow := range parser.LoadData().Workflows {
			str := workflow.On
			if len(args) >= 1 {
				if args[0] == str {
					fmt.Printf("%s\n", workflow.Identifier)
				}
			} else {
				fmt.Printf("%s\n", workflow.Identifier)
			}
		}
	},
}

func init() {
	workflowLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")
	workflowCmd.AddCommand(workflowLsCmd)

	rootCmd.AddCommand(workflowCmd)
}
