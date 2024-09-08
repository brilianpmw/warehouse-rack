package schema

import "warehouse_rack/errors"

type Slot struct {
	No      int      `json:"no"`
	Exp     string   `json:"expired"`
	IsFree  bool     `json:"is_free"`
	Product *Product `json:"product"`
}

func (s *Slot) MakeSlotFree() {
	s.IsFree = true
	return
}

// SetSlotOccupied updates the slot as occupied
func (s *Slot) SetSlotOccupied() {
	s.IsFree = false
	return
}

// IsSlotAvailable checks if the slot is occupied or not
func (s *Slot) IsSlotAvailable() bool {
	return (s.IsFree && s.Product == nil)
}

// ParkVehicle assign/park the Vehicle in the slot
func (s *Slot) ParkVehicle(product *Product) error {
	if s.Product != nil {
		return errors.ErrSlotAlreadyOccupied
	}
	// park vehicle here, make slot occupied
	s.SetSlotOccupied()
	s.Product = product
	return nil
}

func (s *Slot) GetID() int {
	return s.No
}

// ExitPark makes the slot available to park next Vehicle
func (s *Slot) ExecOut() error {
	if s.Product == nil {
		return errors.ErrSlotAlreadyAvailable
	}
	s.MakeSlotFree()
	s.Product = nil
	return nil
}

func (s *Slot) GetExpired() string {
	return s.Exp
}
