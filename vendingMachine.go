package main

import (
	"fmt"
)

// VendingMachine represents the vending machine context in the state pattern
type VendingMachine struct {
	products []Product
	state    State
	balance  float64
}

// NewVendingMachine creates a new vending machine with initial products
func NewVendingMachine(products []Product) (*VendingMachine, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("cannot create vending machine without products")
	}

	vm := &VendingMachine{
		products: products,
		state:    &IdleState{},
		balance:  0.0,
	}

	return vm, nil
}

// GetState returns the current state of the vending machine
func (vm *VendingMachine) GetState() State {
	return vm.state
}

// SetState changes the state of the vending machine
func (vm *VendingMachine) SetState(state State) {
	if state != nil {
		fmt.Printf("State changed to: %s\n", state.String())
		vm.state = state
	}
}

// ProcessRequest handles a request in the current state
func (vm *VendingMachine) ProcessRequest() error {
	return vm.state.HandleRequest(vm)
}

// GetBalance returns the current balance of coins in the machine
func (vm *VendingMachine) GetBalance() float64 {
	return vm.balance
}

// InsertCoins adds coins to the machine
func (vm *VendingMachine) InsertCoins(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("coin amount must be positive")
	}
	vm.balance += amount
	fmt.Printf("Inserted $%.2f. Current balance: $%.2f\n", amount, vm.balance)
	return nil
}

// GetProducts returns a copy of the product list
func (vm *VendingMachine) GetProducts() []Product {
	return vm.products
}

// SelectProduct selects a product by ID and attempts to dispense it
func (vm *VendingMachine) SelectProduct(productID string) error {
	product, index, err := vm.findProduct(productID)
	if err != nil {
		return err
	}

	// Check if product is in stock
	if product.Quantity <= 0 {
		vm.SetState(&OutOfStockState{})
		return fmt.Errorf("product %s is out of stock", product.Name)
	}

	// Check if user has enough balance
	if vm.balance < product.Price {
		return fmt.Errorf("insufficient funds: need $%.2f, have $%.2f", product.Price, vm.balance)
	}

	// Dispense the product
	vm.dispenseProduct(index)
	change := vm.balance - product.Price
	fmt.Printf("Dispensing %s...\n", product.Name)

	// Handle change
	if change > 0 {
		fmt.Printf("Change: $%.2f\n", change)
	}

	vm.balance = 0
	vm.SetState(&IdleState{})

	return nil
}

// RefundBalance returns all coins to the user
func (vm *VendingMachine) RefundBalance() float64 {
	refund := vm.balance
	vm.balance = 0
	if refund > 0 {
		fmt.Printf("Refunding $%.2f\n", refund)
	}
	return refund
}

// AddProducts adds new products to the vending machine
func (vm *VendingMachine) AddProducts(newProducts []Product) error {
	if len(newProducts) == 0 {
		return fmt.Errorf("must add at least one product")
	}

	for _, newProduct := range newProducts {
		if existing, index, err := vm.findProduct(newProduct.ID); err == nil {
			// Product exists, increase quantity
			vm.products[index].Quantity += newProduct.Quantity
			fmt.Printf("Updated %s quantity to %d\n", existing.Name, vm.products[index].Quantity)
		} else {
			// New product, add it
			vm.products = append(vm.products, newProduct)
			fmt.Printf("Added new product: %s\n", newProduct.Name)
		}
	}

	return nil
}

// RemoveProduct removes a product from the vending machine
func (vm *VendingMachine) RemoveProduct(productID string) error {
	_, index, err := vm.findProduct(productID)
	if err != nil {
		return err
	}

	product := vm.products[index]
	vm.products = append(vm.products[:index], vm.products[index+1:]...)
	fmt.Printf("Removed product: %s\n", product.Name)

	return nil
}

// findProduct finds a product by ID and returns the product, index, and error
func (vm *VendingMachine) findProduct(productID string) (*Product, int, error) {
	for i, product := range vm.products {
		if product.ID == productID {
			return &product, i, nil
		}
	}
	return nil, -1, fmt.Errorf("product with ID %s not found", productID)
}

// dispenseProduct reduces the quantity of a product (internal helper)
func (vm *VendingMachine) dispenseProduct(index int) {
	if index >= 0 && index < len(vm.products) {
		vm.products[index].Quantity--
	}
}

// Reset clears all products and balance
func (vm *VendingMachine) Reset() {
	vm.products = []Product{}
	vm.balance = 0
	vm.state = &IdleState{}
	fmt.Println("Machine reset")
}

// DisplayProducts shows all available products
func (vm *VendingMachine) DisplayProducts() {
	fmt.Println("\n=== Available Products ===")
	for _, product := range vm.products {
		status := "Available"
		if product.Quantity == 0 {
			status = "Out of Stock"
		}
		fmt.Printf("ID: %s | Name: %s | Price: $%.2f | Quantity: %d | %s\n",
			product.ID, product.Name, product.Price, product.Quantity, status)
	}
	fmt.Println("========================")
}
