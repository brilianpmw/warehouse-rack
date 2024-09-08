package store

import (
	"warehouse_rack/schema"
)

// Store interface holds all the available cmd exc methods
type Store interface {
	CreateRack() schema.CMDStore
	RackIn() schema.CMDStore
	RackOut() schema.CMDStore
	Status() schema.CMDStore
	SlotNumbersByExpDate() schema.CMDStore
	SkuNumberByExpDate() schema.CMDStore
	SlotNumberBySku() schema.CMDStore
}
