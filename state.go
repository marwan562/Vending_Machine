package main

import (
	"fmt"
)

// State interface defines the behavior for different vending machine states
type State interface {
	// HandleRequest processes a request based on the current state
	HandleRequest(machine *VendingMachine) error
	// String returns the name of the state
	String() string
}

// IdleState represents the idle state of the vending machine
type IdleState struct{}

func (s *IdleState) HandleRequest(machine *VendingMachine) error {
	fmt.Println("Machine is idle. Insert coins to continue.")
	return nil
}

func (s *IdleState) String() string {
	return "Idle"
}

// AcceptingCoinsState represents the state when the machine is accepting coins
type AcceptingCoinsState struct{}

func (s *AcceptingCoinsState) HandleRequest(machine *VendingMachine) error {
	if machine.balance < 0.01 {
		return fmt.Errorf("please insert coins")
	}
	fmt.Printf("Balance: $%.2f\n", machine.balance)
	return nil
}

func (s *AcceptingCoinsState) String() string {
	return "AcceptingCoins"
}

// SelectingProductState represents the state when user is selecting a product
type SelectingProductState struct{}

func (s *SelectingProductState) HandleRequest(machine *VendingMachine) error {
	return nil
}

func (s *SelectingProductState) String() string {
	return "SelectingProduct"
}

// DispensingState represents the state when the machine is dispensing a product
type DispensingState struct{}

func (s *DispensingState) HandleRequest(machine *VendingMachine) error {
	fmt.Println("Dispensing product...")
	return nil
}

func (s *DispensingState) String() string {
	return "Dispensing"
}

// OutOfStockState represents the state when a product is out of stock
type OutOfStockState struct{}

func (s *OutOfStockState) HandleRequest(machine *VendingMachine) error {
	return fmt.Errorf("product out of stock")
}

func (s *OutOfStockState) String() string {
	return "OutOfStock"
}
