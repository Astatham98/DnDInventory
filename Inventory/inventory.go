package inventory

import (
	"errors"
	"fmt"
)

type InventorySlot struct {
	ItemName   string  // name of the item in slot
	StackLimit int     // max amount of items in a slot
	StackSize  int     // amount of items in the stack
	ImagePath  string  // path to the image for the item
	ItemWeight float64 // weight of the item
	Category   string  // category of the item
}

type Inventory struct {
	WeightLimit int             // Max Weight of the inventory
	Name        string          // name of the inventory
	ItemArray   []InventorySlot // Array of items in the inventory
}

func (s InventorySlot) ItemInfo() {
	fmt.Printf("The item in this slot is named %s and has a stack size of %d and a weight of %f\n", s.ItemName, s.StackSize, s.ItemWeight)
}

func (s InventorySlot) GetTotalWeight() float64 {
	return float64(s.StackSize) * s.ItemWeight
}

func (s *InventorySlot) SetStackSize(sizeDifference int) {
	new_size := s.StackSize + sizeDifference
	if -1 < new_size && new_size < s.StackLimit {
		s.StackSize = new_size
	}
}

func (inv *Inventory) IncreaseInventoryWeight(weight int) {
	inv.WeightLimit += weight
}

func (inv *Inventory) RemoveItem(index int) {
	slice := inv.ItemArray
	inv.ItemArray = append(slice[:index], slice[index+1:]...)
}

func (inv Inventory) CanUpdateItemStackSize(size int, slot InventorySlot) (bool, error) {
	old_inventory := slot

	slot.SetStackSize(size)
	if old_inventory.GetTotalWeight() == slot.GetTotalWeight() || slot.GetTotalWeight() > float64(inv.WeightLimit) {
		return false, errors.New("weight too large")
	} else {
		return true, nil
	}
}

func (inv *Inventory) UpdateItemStackSize(size int, index int) error {
	slot := inv.ItemArray[index]

	canUpdate, err := inv.CanUpdateItemStackSize(size, slot)
	if !canUpdate {
		return err
	}
	slot.SetStackSize(size)
	inv.ItemArray[index] = slot
	return nil
}

func (inv *Inventory) AddItem(item InventorySlot) {
	inv.ItemArray = append(inv.ItemArray, item)
}

func (inv Inventory) IsEmpty() bool {
	arr := inv.ItemArray
	for _, slot := range arr {
		if slot.ItemName != "" {
			return false
		}
	}
	return true
}

func (inv Inventory) IsFull() bool {
	arr := inv.ItemArray
	for _, slot := range arr {
		if slot.ItemName == "" {
			return false
		}
	}
	return true
}
