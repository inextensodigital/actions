package printer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/actions/workflow-parser/parser"

	"github.com/actions/workflow-parser/model"
)

// Encode HCL configuration to bytes
func Encode(in *model.Configuration) ([]byte, error) {
	b := &bytes.Buffer{}

	for _, workflow := range in.Workflows {
		b.WriteString(
			fmt.Sprintf("workflow \"%s\" {\n  on = \"%s\"\n",
				workflow.Identifier,
				workflow.On,
			),
		)

		if len(workflow.Resolves) > 0 {
			b.WriteString(
				fmt.Sprintf("  resolves = [\n    \"%s\"\n  ]\n", strings.Join(workflow.Resolves, "\",\n    \"")),
			)
		}
		b.WriteString("}\n\n")
	}

	for _, action := range in.Actions {
		b.WriteString(
			fmt.Sprintf("action \"%s\" {\n  uses = \"%s\"\n",
				action.Identifier,
				action.Uses.String(),
			),
		)

		if len(action.Needs) > 0 {
			b.WriteString(
				fmt.Sprintf("  needs = [\n    \"%s\"\n  ]\n", strings.Join(action.Needs, "\",\n    \"")),
			)
		}

		if action.Args != nil && len(action.Args.Split()) > 0 {
			b.WriteString(
				fmt.Sprintf("  args = \"%s\"\n", strings.Join(action.Args.Split(), " ")),
			)
		}

		if len(action.Secrets) > 0 {
			b.WriteString("  secrets = [\n")
			for _, secret := range action.Secrets {
				b.WriteString(
					fmt.Sprintf("    \"%s\",\n", secret),
				)
			}
			b.WriteString("  ]\n")
		}

		if len(action.Env) > 0 {
			b.WriteString("  env = {\n")
			for key, value := range action.Env {
				b.WriteString(
					fmt.Sprintf("    %s = \"%s\"\n", key, value),
				)
			}
			b.WriteString("  }\n")
		}
		b.WriteString("}\n\n")
	}

	return b.Bytes(), nil
}

// Write bytes content
func Write(c []byte, s string) {
	r := bytes.NewReader(c)
	_, err := parser.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if s == "-" {
		s = "/dev/stdout"
	}

	err = ioutil.WriteFile(s, c, 0644)
	if err != nil {
		panic(err)
	}
}
