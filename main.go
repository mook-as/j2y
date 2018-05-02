// Package main implements a command line utility to convert JSON on stdin to YAML
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	var data interface{}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		os.Exit(1)
	}
	if err = json.Unmarshal(input, &data); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling JSON input: %s\n", err)
		os.Exit(1)
	}
	encoder := yaml.NewEncoder(os.Stdout)
	if err = encoder.Encode(data); err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling YAML output: %s\n", err)
		os.Exit(1)
	}
	if err = encoder.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error closing output: %s\n", err)
		os.Exit(1)
	}
}
