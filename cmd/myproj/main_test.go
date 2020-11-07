package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("myproj -target k28", " ")

	flag.CommandLine.Set("target", "k28")
	cli.Run(args)
	expected := fmt.Sprintf("Hello k28")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
