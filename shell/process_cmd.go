package shell

import (
	"strings"

	"warehouse_rack/schema"
	"warehouse_rack/utils"
)

// Process attempts to validate the command input string.
// Checks the command and arguments list are valid or not.
func Process(cmdInputStr string) (*schema.Command, error) {
	var cmd = new(schema.Command)
	cmdInputList, err := utils.SplitCmdArguments(cmdInputStr)
	if err != nil {
		return nil, err
	}
	cmd.Command = cmdInputList[0]
	if len(cmdInputList) > 1 {
		cmd.Arguments = strings.Split(cmdInputList[1], utils.Space)
	}
	// validate the inputs
	if err := cmd.Ok(); err != nil {
		return nil, err
	}

	// set cmd store connection
	setStoreConnection(cmd)

	return cmd, nil
}

func setStoreConnection(command *schema.Command) {
	switch command.Command {
	case string(schema.CMDCreateRack):
		command.Connection = Store.CreateRack()
	case string(schema.CMDRackIn):
		command.Connection = Store.RackIn()
	case string(schema.CMDRackOut):
		command.Connection = Store.RackOut()
	case string(schema.CMDStatus):
		command.Connection = Store.Status()
	case string(schema.CMDSkuNumbersForProductWithExpDate):
		command.Connection = Store.SkuNumberByExpDate()
	case string(schema.CMDSlotNumbersForProductWithExpDate):
		command.Connection = Store.SlotNumbersByExpDate()
	case string(schema.CMDSlotNumberBySkuNumber):
		command.Connection = Store.SlotNumberBySku()
	}

}
