package main

import (
	"fmt"
	"os"
)

const (
	helpMsg = `hcl2json

Converts Hashicorp Configuration Langauge (HCL) to JavaScript Object Notation (JSON).
Can also output YAML and TOML. If multiple output format command line flags and/or
filename arguments are given, the rightmost wins. If no filename or - is given, reads
from stdin.

Usage:
  hcl2json [FLAGS] [FILENAME]

Examples:
  Concatenate all Terraform files in a directory convert the result to JSON via stdin
  > cat *.tf | hcl2json

  Convert single HCL file to YAML
  > hcl2json -y example.hcl

Flags:
  -h, --help      help for hcl2json
  -v, --version   print program version
  -j, --json      output JSON (default)
  -y, --yaml      output YAML
  -t, --toml      output TOML`
)

var (
	ver  string
	port = 8080
)

func main() {
	server := NewServer()
	server.Listen(port)
}

func check(err error, msg string) {
	if err != nil {
		if _, err := fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
