package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/actions/workflow-parser/model"
	ghparser "github.com/actions/workflow-parser/parser"
)

// LoadData load workflow file and parse it
func LoadData() *model.Configuration {
	var r io.Reader
	var err error
	r, err = os.Open(".github/main.workflow")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	data, err := ghparser.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	return data
}
