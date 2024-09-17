package main

import (
	"fmt"
	inventory "inventory/Inventory"
)

func main() {
	slot1 := inventory.InventorySlot{
		ItemName:   "coins",
		StackLimit: 100,
		StackSize:  5,
		ImagePath:  "/images/coins",
		ItemWeight: 0.1,
		Category:   "money",
	}
	slot1.ItemInfo()
	fmt.Println(slot1.GetTotalWeight())
	slot1.SetStackSize(100)
	fmt.Println(slot1.GetTotalWeight())
	slot1.SetStackSize(50)
	fmt.Println(slot1.GetTotalWeight())
	slot1.SetStackSize(-100)
	fmt.Println(slot1.GetTotalWeight())
	slot1.SetStackSize(-10)
	fmt.Println(slot1.GetTotalWeight())

	backpack := inventory.Inventory{
		WeightLimit: 100,
		Name:        "Backpack",
		ItemArray:   []inventory.InventorySlot{},
	}

	backpack.AddItem(slot1)
	fmt.Println(backpack)

	err := backpack.UpdateItemStackSize(100, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(backpack)

	backpack.RemoveItem(0)
	fmt.Println(backpack)

}
