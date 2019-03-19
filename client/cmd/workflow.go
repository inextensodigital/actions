package cmd

import (
	"fmt"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/client/parser"
	"github.com/inextensodigital/actions/client/printer"
	"github.com/spf13/cobra"
)

var Action string

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

var workflowCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new workflow",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		w := model.Workflow{
			Identifier: args[0],
			On:         args[1],
		}

		resolve := []string{args[2]}
		w.Resolves = resolve

		conf := parser.LoadData()
		conf.Workflows = append(conf.Workflows, &w)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var workflowAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add action to a workflow",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()
		workflow := conf.GetWorkflow(args[0])

		rs := args[1]
		workflow.Resolves = append(workflow.Resolves, rs)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var workflowRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a workflow",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()
		workflow := conf.GetWorkflow(args[0])

		workflow.Identifier = args[1]

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func init() {
	workflowLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")
	workflowAddCmd.Flags().StringVarP(&Action, "action", "a", "", "action")
	workflowAddCmd.MarkFlagRequired("action")

	workflowCmd.AddCommand(workflowLsCmd)
	workflowCmd.AddCommand(workflowAddCmd)
	workflowCmd.AddCommand(workflowCreateCmd)
	workflowCmd.AddCommand(workflowRenameCmd)

	rootCmd.AddCommand(workflowCmd)
}
