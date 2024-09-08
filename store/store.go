package store

import "warehouse_rack/schema"

var (
	RackCreatedInfo   = "Created​ a rack with %d slots"
	SlotAllocatedInfo = "Allocated slot number: %v"
	SlotIsFreeInfo    = "Slot number %v is free"
	SlotLeaveInfo     = "Slot number %v is free"
)

var RackData *schema.Rack

type store struct {
	createRack          schema.CMDStore
	rackIn              schema.CMDStore
	rackOut             schema.CMDStore
	status              schema.CMDStore
	skuNumberByExpDate  schema.CMDStore
	slotNumberByExpDate schema.CMDStore
	slotNumberBySku     schema.CMDStore
}

func (s store) SlotNumberBySku() schema.CMDStore {
	return s.slotNumberBySku
}

func (s store) CreateRack() schema.CMDStore {
	return s.createRack
}

func (s store) RackIn() schema.CMDStore {
	return s.rackIn
}

func (s store) RackOut() schema.CMDStore {
	return s.rackOut
}

func (s store) Status() schema.CMDStore {
	return s.status
}

func (s store) SkuNumberByExpDate() schema.CMDStore {
	return s.skuNumberByExpDate
}

func (s store) SlotNumbersByExpDate() schema.CMDStore {
	return s.slotNumberByExpDate
}

// NewStore returns the store object
func NewStore() *store {
	st := InitStore()
	st.createRack = NewCreateRack(st)
	st.rackIn = NewRackIN(st)
	st.rackOut = NewRackOut(st)
	st.status = NewStatus(st)
	st.skuNumberByExpDate = NewSkuNumberByExpDate(st)
	st.slotNumberByExpDate = NewSlotNumberByExpDate(st)
	st.slotNumberBySku = NewSlotNumberBySkuNumber(st)
	return st
}

// InitStore returns the store object
func InitStore() *store {
	return new(store)
}
