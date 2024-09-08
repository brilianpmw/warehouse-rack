package errors

import (
	"errors"
	"fmt"
)

var (
	ErrRackSlotsFull = errors.New("Sorry, rack is full")

	ErrInvalidCMD      = "Command '%s' is invalid!. Please type 'help' see available commands"
	ErrInsuffArguments = "Insufficient arguments for '%s'. Expected %d, Got %d"
	ErrRackEmpty       = errors.New("rack is empty")

	ErrSlotAlreadyOccupied  = errors.New("Slot: Slot already occupied")
	ErrSlotAlreadyAvailable = errors.New("Slot: Slot already available")
	ErrInvalidSlotID        = errors.New("Slot: Invalid slot number")
	ErrInvalidInputSlot     = errors.New("Slot: Invalid inputr")
	ErrRackAlreadyCreated   = errors.New("Rack Already created")
	ErrInvalidTabSpace      = errors.New("Don't use Tab spaces in the commands")
)

// ErrInvalidCommand err wrapper
func ErrInvalidCommand(cmd string) error {
	return fmt.Errorf(ErrInvalidCMD, cmd)
}

// ErrInvalidArguments err wrapper
func ErrInvalidArguments(cmd string, expect, got int) error {
	return fmt.Errorf(ErrInsuffArguments, cmd, expect, got)
}
