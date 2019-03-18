package cmd

import (
	"fmt"

	"github.com/actions/workflow-parser/model"
	"github.com/inextensodigital/actions/client/parser"
	"github.com/inextensodigital/actions/client/printer"
	"github.com/spf13/cobra"
	// "reflect"
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

func init() {
	actionLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")

	actionCmd.AddCommand(actionLsCmd)
	actionCmd.AddCommand(actionCreateCmd)
	actionCmd.AddCommand(actionRenameCmd)

	rootCmd.AddCommand(actionCmd)
}
