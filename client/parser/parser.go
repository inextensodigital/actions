package parser

import (
	"github.com/actions/workflow-parser/model"
	ghparser "github.com/actions/workflow-parser/parser"
	"io"
	"log"
	"os"
)

func LoadData() *model.Configuration {
	var r io.Reader
	var err error
	r, err = os.Open(".github/main.workflow")
	if err != nil {
		log.Fatalln(err)
	}

	data, err := ghparser.Parse(r)
	if err != nil {
		log.Fatalln(err)
	}

	return data
}
