package store

import (
	"fmt"
	"strconv"
	"warehouse_rack/errors"
	"warehouse_rack/schema"
)

type createRackStore struct {
	*store
}

func NewCreateRack(st *store) *createRackStore {
	pl := &createRackStore{st}
	return pl
}

func (pl *createRackStore) Execute(cmd *schema.Command) (string, error) {

	totalSlots, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return "", errors.ErrInvalidInputSlot
	}
	if totalSlots <= 0 {
		return "", errors.ErrInvalidInputSlot
	}

	if RackData != nil {
		return "", errors.ErrRackAlreadyCreated
	}

	newRack := &schema.Rack{
		Slots: make([]*schema.Slot, totalSlots),
	}

	for i := range newRack.Slots {
		newRack.Slots[i] = new(schema.Slot)
		newRack.Slots[i].No = i + 1
		newRack.Slots[i].MakeSlotFree()
		newRack.Slots[i].Product = nil
	}

	RackData = newRack
	return fmt.Sprintf(RackCreatedInfo, totalSlots), nil
}
