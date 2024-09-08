package schema

import (
	"warehouse_rack/errors"
)

type Rack struct {
	Slots []*Slot `json:"slots"`
}

// FirstAvailableSlot returns the first available slot to park Vehicle
func (pl *Rack) FirstAvailableSlot() (*Slot, error) {
	for _, slot := range pl.Slots {
		if slot.IsSlotAvailable() {
			return slot, nil
		}
	}

	return nil, errors.ErrRackSlotsFull
}
