package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/github-workflow/parser"
	"github.com/inextensodigital/actions/github-workflow/printer"
	"github.com/spf13/cobra"
)

// On filter
var On string

// Env passed at creation
var Env []string

// Secret passed at creation
var Secret []string

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Actions",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var actionLsCmd = &cobra.Command{
	Use:   "ls [ID]",
	Short: "List actions",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf := parser.LoadData()

		iA := 0
		if len(args) == 1 {
			action := conf.GetAction(args[0])
			fmt.Printf("%s\n", action.Identifier)
			iA++
		} else {
			for _, action := range conf.Actions {
				fmt.Printf("%s\n", action.Identifier)
				iA++
			}
		}

		if iA == 0 {
			os.Exit(1)
		}
	},
}

var actionRenameCmd = &cobra.Command{
	Use:   "rename SOURCE TARGET",
	Short: "Rename action",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source, target := args[0], args[1]
		conf := parser.LoadData()

		action := conf.GetAction(target)
		if action != nil {
			fmt.Fprintf(os.Stderr, "Action '%s' already exists, you should find another name to rename your action.\n", target)
			os.Exit(1)
		}

		action = conf.GetAction(source)
		if action == nil {
			fmt.Fprintf(os.Stderr, "Unknown action '%s'. Please check if you have made a typo.\n", source)
			os.Exit(1)
		}
		action.Identifier = target

		for _, action := range conf.Actions {
			for index, need := range action.Needs {
				if need == source {
					action.Needs[index] = target
				}
			}
		}

		for _, workflow := range conf.Workflows {
			for index, resolve := range workflow.Resolves {
				if resolve == source {
					workflow.Resolves[index] = target
				}
			}
		}

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

var actionCreateCmd = &cobra.Command{
	Use:   "create ID USE",
	Short: "Create new action",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, use := args[0], args[1]

		conf := parser.LoadData()
		action := conf.GetAction(id)
		if action != nil {
			fmt.Fprintf(os.Stderr, "Action '%s' already exists, can't create the same one more than once.\n", id)
			os.Exit(1)
		}

		u := model.UsesDockerImage{Image: use}
		uh := &u

		envList := make(map[string]string, 0)
		for _, v := range Env {
			env := strings.SplitN(v, "=", 2)
			if len(env) != 2 {
				fmt.Fprintf(os.Stderr, "Invalid env value '%s'. You should comply to the format: 'key=value'\n", v)
				os.Exit(1)
			}

			key, value := env[0], env[1]
			envList[key] = value
		}

		ghaction := model.Action{Identifier: id, Uses: uh, Env: envList}

		ghaction.Secrets = Secret

		conf.Actions = append(conf.Actions, &ghaction)

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func removeAction(slice []*model.Action, s int) []*model.Action {
	return append(slice[:s], slice[s+1:]...)
}

func removeFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

var actionRemoveCmd = &cobra.Command{
	Use:   "rm NAME",
	Short: "Remove action",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		conf := parser.LoadData()

		var ia int
		for k, action := range conf.Actions {
			if action.Identifier == name {
				ia = k
			}
		}

		listAction := removeAction(conf.Actions, ia)
		conf.Actions = listAction

		// remove the action from workflow resolves
		listWorkflow := make([]*model.Workflow, 0)
		for _, workflow := range conf.Workflows {
			for kr, resolver := range workflow.Resolves {
				if resolver == name {
					workflow.Resolves = removeFromSlice(workflow.Resolves, kr)
				}
			}
			listWorkflow = append(listWorkflow, workflow)
		}

		conf.Workflows = listWorkflow

		listAction = make([]*model.Action, 0)
		for _, action := range conf.Actions {
			for kr, need := range action.Needs {
				if need == name {
					action.Needs = removeFromSlice(action.Needs, kr)
				}
			}
			listAction = append(listAction, action)
		}
		conf.Actions = listAction

		content, _ := printer.Encode(conf)
		printer.Write(content)
	},
}

func init() {
	actionCmd.AddCommand(actionLsCmd)
	actionCmd.AddCommand(actionCreateCmd)
	actionCreateCmd.Flags().StringArrayVarP(&Env, "env", "e", []string{}, "")
	actionCreateCmd.Flags().StringArrayVarP(&Secret, "secret", "s", []string{}, "")
	actionCmd.AddCommand(actionRenameCmd)
	actionCmd.AddCommand(actionRemoveCmd)

	rootCmd.AddCommand(actionCmd)
}
