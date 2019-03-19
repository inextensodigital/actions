package printer

import (
	// "fmt"
	"github.com/actions/workflow-parser/model"
	"github.com/rodaine/hclencoder"
	"io/ioutil"
	"log"
)

type ConfigurationPrinter struct {
	Actions   []*ActionPrinter   `hcl:"action"`
	Workflows []*WorkflowPrinter `hcl:"workflow"`
}

type ActionPrinter struct {
	Identifier string            `hcl:",key"`
	Uses       string            `hcl:"uses"`
	Runs       []string          `hcl:"runs"`
	Args       []string          `hcl:"args"`
	Needs      []string          `hcl:"needs"`
	Env        map[string]string `hcl:"env"`
	Secrets    []string          `hcl:"secrets"`
}

type WorkflowPrinter struct {
	Identifier string   `hcl:",key"`
	On         string   `hcl:"on"`
	Resolves   []string `hcl:"resolves"`
}

// func Encode(in interface{}) ([]byte, error) {
func Encode(in *model.Configuration) ([]byte, error) {
	cp := ConfigurationPrinter{}

	for _, action := range in.Actions {
		var runs []string
		var args []string
		if action.Runs != nil {
			runs = action.Runs.Split()
		}
		if action.Runs != nil {
			args = action.Args.Split()
		}

		ap := ActionPrinter{
			Identifier: action.Identifier,
			Uses:       action.Uses.String(),
			Runs:       runs,
			Args:       args,
			Needs:      action.Needs,
			Env:        action.Env,
			Secrets:    action.Secrets,
		}

		cp.Actions = append(cp.Actions, &ap)
	}

	for _, workflow := range in.Workflows {
		wp := WorkflowPrinter{
			Identifier: workflow.Identifier,
			On:         workflow.On,
			Resolves:   workflow.Resolves,
		}
		cp.Workflows = append(cp.Workflows, &wp)
	}

	hcl, err := hclencoder.Encode(cp)
	if err != nil {
		log.Fatal("unable to encode: ", err)
	}

	return hcl, err
}

func Write(c []byte) {
	err := ioutil.WriteFile("/tmp/test.workflow", c, 0644)
	if err != nil {
		panic(err)
	}
}
