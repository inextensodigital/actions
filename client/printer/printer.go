package printer

import (
	"fmt"
	"github.com/rodaine/hclencoder"
	"io/ioutil"
	"log"
)

func Encode(in interface{}) ([]byte, error) {
	hcl, err := hclencoder.Encode(in)
	if err != nil {
		log.Fatal("unable to encode: ", err)
	}
	fmt.Print(string(hcl))

	return hcl, err
}

func Write(c []byte, d string) {
	err := ioutil.WriteFile(d, c, 0644)
	if err != nil {
		panic(err)
	}
}
