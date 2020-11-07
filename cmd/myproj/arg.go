package main

import "flag"

func parseArgs() string {
	const defaultTarget = "myproj"
	var target string
	flag.StringVar(&target, "target", defaultTarget, "hello target")
	flag.Parse()

	return target
}
