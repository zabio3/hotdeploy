package cmd

import (
	"flag"
	"fmt"
	"github.com/zabio3/hotdeploy/deploy"
	"io"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK = iota + 1
	ExitCodeParseFlagsError
	ExitCodeEmptyServerProgramError
	ExitCodeServerProgramError
)

const name = "hotdeploy"

const usage = `hotdeploy - a superdaemon for hot-deploying server program

Usage: hotdeploy --port=<port> --server=<server program>
Other Commands:
  --help	-h	Help about any command
`

// CLI represents CLI interface.
type CLI struct {
	OutStream, ErrStream io.Writer
}

var server string
var port int

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprint(cli.OutStream, usage)
	}

	flags.IntVar(&port, "port", 8080, "Set server port")
	flags.StringVar(&server, "server", "", "Set start sever program")

	if err := flags.Parse(args[1:]); err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeParseFlagsError
	}

	if server == "" {
		fmt.Fprint(cli.ErrStream, "empty start server program\n")
		return ExitCodeEmptyServerProgramError
	}

	if err := deploy.HotDeploy(server, port); err != nil {
		fmt.Fprint(cli.ErrStream, err)
		return ExitCodeServerProgramError
	}

	return ExitCodeOK
}
