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

var actionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new action",
	Run: func(cmd *cobra.Command, args []string) {

		u := model.UsesDockerImage{Image: "github.com/actions/workflow-parser"}
		uh := &u
		ghaction := model.Action{Identifier: "nouvelle action", Uses: uh}

		fmt.Println("action called %v", ghaction)
		content, _ := printer.Encode(ghaction)

		printer.Write(content, "/tmp/test.workflow")
	},
}

var actionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("action called")
	},
}

func init() {
	actionLsCmd.Flags().StringVarP(&Filter, "filter", "f", "", "Filter on")

	actionCmd.AddCommand(actionLsCmd)
	actionCmd.AddCommand(actionCreateCmd)
	actionCmd.AddCommand(actionAddCmd)

	rootCmd.AddCommand(actionCmd)
}
