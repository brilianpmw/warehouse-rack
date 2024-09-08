package store

import (
	"fmt"
	"strconv"
	"warehouse_rack/schema"
)

type rackOutStore struct {
	*store
}

func NewRackOut(st *store) *rackOutStore {
	pl := &rackOutStore{st}
	return pl
}

func (pl *rackOutStore) Execute(cmd *schema.Command) (string, error) {
	intParams, _ := strconv.Atoi(cmd.Arguments[0])
	// want to releases the slot byID
	for _, lot := range RackData.Slots {
		if int(lot.No) == intParams {
			lot.ExecOut()
			break
		}
	}

	return fmt.Sprintf(SlotLeaveInfo, cmd.Arguments[0]), nil
}
