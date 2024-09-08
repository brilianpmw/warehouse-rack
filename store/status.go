package store

import (
	"fmt"
	"strings"

	"warehouse_rack/errors"
	"warehouse_rack/schema"
	"warehouse_rack/utils"
)

type statusStore struct {
	*store
}

// NewStatusStore returns new store object
func NewStatus(st *store) *statusStore {
	pl := &statusStore{st}
	return pl
}

// Execute will returns the current status of all the slots.
func (pl *statusStore) Execute(cmd *schema.Command) (string, error) {
	if RackData == nil {
		return "", errors.ErrRackEmpty
	}
	var slotStatus = []string{fmt.Sprintf("%-10s%-20s%-10s", "Slot No.", "SKU No", "ExpDate")}
	for _, slot := range RackData.Slots {
		if !slot.IsFree {
			slotStatus = append(slotStatus, fmt.Sprintf("%-10d%-20s%-10s", slot.GetID(),
				slot.Product.GetSku(), slot.GetExpired()))
		}
	}
	return strings.Join(slotStatus, utils.NewLineDelim), nil
}
