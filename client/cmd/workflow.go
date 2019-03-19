package cmd

import (
	"fmt"
	"os"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/github-workflow/parser"
	"github.com/inextensodigital/actions/github-workflow/printer"
	"github.com/spf13/cobra"
)

// Action list
var Action []string

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Workflow",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var workflowLsCmd = &cobra.Command{
	Use:   "ls [ID]",
	Short: "List workflow by name or filtered on event",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()

		var workflows = make([]*model.Workflow, 0)
		if On != "" {
			workflows = conf.GetWorkflows(On)
		} else {
			workflows = conf.Workflows
		}

		iW := 0
		if len(args) == 1 {

		}
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
	Use:   "create ID ON",
	Short: "Create a new workflow",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, on := args[0], args[1]

		conf := parser.LoadData()
		if conf.GetWorkflow(id) != nil {
			fmt.Fprintf(os.Stderr, "Workflow '%s' already exists, can't create the same one more than once.\n", id)
			os.Exit(1)
		}

		w := model.Workflow{
			Identifier: id,
			On:         on,
		}

		resolve := Action
		w.Resolves = resolve

		conf.Workflows = append(conf.Workflows, &w)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var workflowAddCmd = &cobra.Command{
	Use:   "add ID",
	Short: "Add action to a workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		conf := parser.LoadData()
		workflow := conf.GetWorkflow(name)

		for _, action := range Action {
			if conf.GetAction(action) == nil {
				fmt.Fprintf(os.Stderr, "Unknown action '%s'. You have to create it first or to check if you have made a typo.\n", action)
				os.Exit(1)
			}

			if contains(workflow.Resolves, action) {
				fmt.Fprintf(os.Stderr, "Action '%s' already resolved by this workflow.\n", action)
				os.Exit(1)
			}

			workflow.Resolves = append(workflow.Resolves, action)
		}

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

// contains check if an array of strings contains a string
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

var workflowRenameCmd = &cobra.Command{
	Use:   "rename SOURCE TARGET",
	Short: "Rename a workflow",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source, target := args[0], args[1]

		conf := parser.LoadData()

		workflow := conf.GetWorkflow(target)
		if workflow != nil {
			fmt.Fprintf(os.Stderr, "Workflow '%s' already exists, you should find another name to rename your workflow.\n", target)
			os.Exit(1)
		}

		workflow = conf.GetWorkflow(source)
		if workflow == nil {
			fmt.Fprintf(os.Stderr, "Unknown workflow '%s'. Please check if you have made a typo.\n", source)
			os.Exit(1)
		}

		workflow.Identifier = target

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func init() {
	workflowCmd.AddCommand(workflowLsCmd)
	workflowLsCmd.Flags().StringVarP(&On, "on", "o", "", "Filter on")

	workflowCmd.AddCommand(workflowAddCmd)
	workflowAddCmd.Flags().StringSliceVarP(&Action, "resolve", "r", []string{}, "action to be resolved by this workflow")
	workflowAddCmd.MarkFlagRequired("resolve")

	workflowCmd.AddCommand(workflowCreateCmd)
	workflowCreateCmd.Flags().StringSliceVarP(&Action, "resolve", "r", []string{}, "action to be resolved by this workflow")

	workflowCmd.AddCommand(workflowRenameCmd)

	rootCmd.AddCommand(workflowCmd)
}
