package store

import (
	"fmt"
	"warehouse_rack/errors"
	"warehouse_rack/schema"
)

type rackInStore struct {
	*store
}

func NewRackIN(st *store) *rackInStore {
	pl := &rackInStore{st}
	return pl
}

func (pl *rackInStore) Execute(cmd *schema.Command) (string, error) {
	if RackData == nil {
		return "", errors.ErrRackEmpty
	}

	// Checks for first available slot
	product := &schema.Product{
		Sku: cmd.Arguments[0],
	}
	availSlot, err := RackData.FirstAvailableSlot()
	if err != nil {
		return "", err
	}

	availSlot.Exp = cmd.Arguments[1]
	// park vehicle in the slot
	if err := availSlot.ParkVehicle(product); err != nil {
		return "", err
	}

	return fmt.Sprintf(SlotAllocatedInfo, availSlot.GetID()), nil
}
