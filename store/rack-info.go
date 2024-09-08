package store

import (
	"strconv"
	"strings"

	"warehouse_rack/errors"
	"warehouse_rack/schema"
)

type rackSkuNumberByExpDateStore struct {
	*store
}

// NewStatusStore returns new store object
func NewSkuNumberByExpDate(st *store) *rackSkuNumberByExpDateStore {
	pl := &rackSkuNumberByExpDateStore{st}
	return pl
}

// Execute will returns the current status of all the slots.
func (pl *rackSkuNumberByExpDateStore) Execute(cmd *schema.Command) (string, error) {
	if RackData == nil {
		return "", errors.ErrRackEmpty
	}

	var slotStatus = []string{}
	for _, slot := range RackData.Slots {
		if slot.Exp == cmd.Arguments[0] {
			slotStatus = append(slotStatus, slot.Product.GetSku())
		}
	}

	return strings.Join(slotStatus, ","), nil
}

type rackSlotNumberByExpDateStore struct {
	*store
}

// NewStatusStore returns new store object
func NewSlotNumberByExpDate(st *store) *rackSlotNumberByExpDateStore {
	pl := &rackSlotNumberByExpDateStore{st}
	return pl
}

// Execute will returns the current status of all the slots.
func (pl *rackSlotNumberByExpDateStore) Execute(cmd *schema.Command) (string, error) {
	if RackData == nil {
		return "", errors.ErrRackEmpty

	}
	var slotStatus = []string{}
	for _, slot := range RackData.Slots {
		if slot.Exp == cmd.Arguments[0] {
			slotStatus = append(slotStatus, strconv.Itoa(slot.GetID()))
		}
	}

	return strings.Join(slotStatus, ","), nil
}

type rackSlotNumberBySkuNumber struct {
	*store
}

// NewStatusStore returns new store object
func NewSlotNumberBySkuNumber(st *store) *rackSlotNumberBySkuNumber {
	pl := &rackSlotNumberBySkuNumber{st}
	return pl
}

// Execute will returns the current status of all the slots.
func (pl *rackSlotNumberBySkuNumber) Execute(cmd *schema.Command) (string, error) {
	if RackData == nil {
		return "", errors.ErrRackEmpty

	}
	var slotStatus = []string{}
	for _, slot := range RackData.Slots {
		if slot.Product.GetSku() == cmd.Arguments[0] {
			slotStatus = append(slotStatus, strconv.Itoa(slot.GetID()))
		}
	}

	if len(slotStatus) == 0 {
		return "Not found", nil
	}

	return strings.Join(slotStatus, ","), nil
}
