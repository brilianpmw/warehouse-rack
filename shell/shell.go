package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"warehouse_rack/schema"
	"warehouse_rack/store"
	"warehouse_rack/utils"
)

var (
	shell *schema.Shell
	// Store holds the store conn interface
	Store store.Store
)

// InitIshell inits the conn
func Initshell() {
	Store = store.NewStore()
	shell = &schema.Shell{}
}

// Start attempts to create new interactive session
// the process reads the commands until it finds `exit` command
//
func Start() {
	if shell == nil {
		Initshell()
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		cmdInputStr, err := reader.ReadString(utils.EndLineDelim)
		if err != nil {
			break
		}
		cmdInputStr = strings.TrimRight(cmdInputStr, utils.NewLineDelim)
		if strings.TrimSpace(cmdInputStr) == "" {
			continue
		}

		// process the commands
		cmd, err := Process(cmdInputStr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// map shell history with command
		if cmd.IsExit() {
			break
		}
		// Execute the command
		response, err := cmd.Connection.Execute(cmd)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(response)
	}
}
