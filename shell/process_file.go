package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProcessFile reads the file commands and execute one by one
// prints output in stdout
func ProcessFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	if shell == nil {
		Initshell()
	}
	var cmdInputStr string
	for fileScanner.Scan() {
		cmdInputStr = fileScanner.Text()
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
		}
		fmt.Println(response)
	}

	return
}
