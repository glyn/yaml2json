package main

import (
	"fmt"
	"os"

	"io/ioutil"

	"github.com/peter-edge/pkg-go/yaml"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "YAML file not specified")
		os.Exit(1)
	}

	y, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "YAML file %s could not be read: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	opts := pkgyaml.ToJSONOptions{
		Pretty: true,
	}

	j, err := pkgyaml.ToJSON(y, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Conversion failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(j))
}
