package cmd

import (
	"fmt"
	"os"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/client/parser"
	"github.com/inextensodigital/actions/client/printer"
	"github.com/spf13/cobra"
)

var Action string
var Filter string

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("workflow called")
	},
}

var workflowLsCmd = &cobra.Command{
	Use:   "ls NAME",
	Short: "List workflow by name or filtered on event",
	Run: func(cmd *cobra.Command, args []string) {
		workflows := parser.LoadData().Workflows

		if Filter != "" {
			workflows = parser.LoadData().GetWorkflows(Filter)
		}

		iW := 0
		for _, workflow := range workflows {
			if len(args) >= 1 {
				if args[0] == workflow.Identifier {
					fmt.Printf("%s\n", workflow.Identifier)
					iW++
				}
			} else {
				fmt.Printf("%s\n", workflow.Identifier)
				iW++
			}
		}

		if iW == 0 {
			os.Exit(1)
		}
	},
}

var workflowCreateCmd = &cobra.Command{
	Use:   "create NAME ON ACTION",
	Short: "Create a new workflow",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		name, on, action := args[0], args[1], args[2]

		w := model.Workflow{
			Identifier: name,
			On:         on,
		}

		resolve := []string{action}
		w.Resolves = resolve

		conf := parser.LoadData()
		conf.Workflows = append(conf.Workflows, &w)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var workflowAddCmd = &cobra.Command{
	Use:   "add NAME ACTION",
	Short: "Add action to a workflow",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name, action := args[0], args[1]

		conf := parser.LoadData()
		workflow := conf.GetWorkflow(name)

		rs := action
		workflow.Resolves = append(workflow.Resolves, rs)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var workflowRenameCmd = &cobra.Command{
	Use:   "rename SOURCE TARGET",
	Short: "Rename a workflow",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source, target := args[0], args[1]

		conf := parser.LoadData()
		workflow := conf.GetWorkflow(source)

		workflow.Identifier = target

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func init() {
	workflowLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")
	workflowAddCmd.Flags().StringVarP(&Action, "action", "a", "", "action")

	workflowCmd.AddCommand(workflowLsCmd)
	workflowCmd.AddCommand(workflowAddCmd)
	workflowCmd.AddCommand(workflowCreateCmd)
	workflowCmd.AddCommand(workflowRenameCmd)

	rootCmd.AddCommand(workflowCmd)
}
