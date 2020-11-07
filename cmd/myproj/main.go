package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	myproj "github.com/k28/go-myproj"
)

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

// リターンコードの定義
const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

// CLI テスト可能なCLIツール作成のためのIO定義
type CLI struct {
	outStream, errStream io.Writer
}

// Run main
func (c *CLI) Run(args []string) int {

	// オプション引数のパース
	var target string
	flags := flag.NewFlagSet("myproj", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.StringVar(&target, "target", "myproj", "Greeting target")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	greeting := myproj.Hello(target)
	fmt.Fprintf(c.outStream, greeting)
	return ExitCodeOK
}
