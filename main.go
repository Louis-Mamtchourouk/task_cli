package main

import (
	"os"

	"github.com/Louis-Mamtchourouk/task_cli/functionnalities"
)

func main() {
	progArgs := os.Args

	if len(progArgs) > 1 {

		switch {
		case progArgs[1] == "add" && len(progArgs) == 3:
			functionnalities.Add(progArgs[2])
		}

	}
}
