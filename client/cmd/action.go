package cmd

import (
	"fmt"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/client/parser"
	"github.com/inextensodigital/actions/client/printer"
	"github.com/spf13/cobra"
)

var Filter string

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Actions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("action called")
	},
}

var actionLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List actions",
	Run: func(cmd *cobra.Command, args []string) {
		for _, action := range parser.LoadData().Actions {
			str := fmt.Sprintf("%s", action.Uses)
			if len(args) >= 1 {
				if args[0] == str {
					fmt.Printf("%s\n", action.Identifier)
				}
			} else {
				fmt.Printf("%s\n", action.Identifier)
			}
		}
	},
}

var actionRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename actions",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()
		for _, action := range conf.Actions {
			if args[0] == action.Identifier {
				action.Identifier = args[1]
			}
		}

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var actionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new action",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		u := model.UsesDockerImage{Image: args[1]}
		uh := &u

		env := make(map[string]string)
		env[args[2]] = args[2]

		ghaction := model.Action{Identifier: args[0], Uses: uh, Env: env}

		if len(args) == 4 {
			ghaction.Secrets[0] = args[3]
		}

		conf := parser.LoadData()
		conf.Actions = append(conf.Actions, &ghaction)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func removeAction(slice []*model.Action, s int) []*model.Action {
	return append(slice[:s], slice[s+1:]...)
}

func removeResolver(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

var actionRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove actions",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()

		var ia int
		for k, action := range conf.Actions {
			if action.Identifier == args[0] {
				ia = k
			}
		}

		listAction := removeAction(conf.Actions, ia)

		conf.Actions = listAction
		content, _ := printer.Encode(conf)
		printer.Write(content)

		listWorkflow := make([]*model.Workflow, 0)
		for _, workflow := range conf.Workflows {
			for kr, resolver := range workflow.Resolves {
				if resolver == args[0] {
					workflow.Resolves = removeResolver(workflow.Resolves, kr)
				}
			}
			listWorkflow = append(listWorkflow, workflow)
		}

		conf.Workflows = listWorkflow
		for _, workflow := range conf.Workflows {
			fmt.Printf("%s\n", workflow.Resolves)
		}

		content, _ = printer.Encode(conf)
		printer.Write(content)
	},
}

func init() {
	actionLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")

	actionCmd.AddCommand(actionLsCmd)
	actionCmd.AddCommand(actionCreateCmd)
	actionCmd.AddCommand(actionRenameCmd)
	actionCmd.AddCommand(actionRemoveCmd)

	rootCmd.AddCommand(actionCmd)
}
