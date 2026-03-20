package main

import (
	"fmt"
)

func main() {
	// Initialize products for the vending machine
	initialProducts := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
		{ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 8},
		{ID: "3", Name: "Water", Price: 1.00, Quantity: 15},
		{ID: "4", Name: "Juice", Price: 2.00, Quantity: 5},
	}

	// Create a new vending machine
	machine, err := NewVendingMachine(initialProducts)
	if err != nil {
		fmt.Printf("Error creating vending machine: %v\n", err)
		return
	}

	fmt.Println("🤖 Welcome to the Vending Machine Demo!")
	fmt.Println("=====================================")

	// Display initial products
	machine.DisplayProducts()

	// Example 1: Insert coins and buy a product
	fmt.Println("📌 Example 1: Buying Coke")
	fmt.Println("---")
	_ = machine.InsertCoins(1.50)
	err = machine.SelectProduct("1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// Example 2: Try to buy something without enough money
	fmt.Println("📌 Example 2: Insufficient funds")
	fmt.Println("---")
	_ = machine.InsertCoins(0.50)
	err = machine.SelectProduct("4") // Juice costs $2.00
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Refund the coins
	refund := machine.RefundBalance()
	fmt.Printf("Remaining balance: $%.2f\n", refund)
	fmt.Println()

	// Example 3: Add more products
	fmt.Println("📌 Example 3: Restocking products")
	fmt.Println("---")
	newProducts := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 5},   // Restock Coke
		{ID: "5", Name: "Soda", Price: 1.75, Quantity: 10}, // New product
	}
	_ = machine.AddProducts(newProducts)
	fmt.Println()

	// Example 4: Check state and process request
	fmt.Println("📌 Example 4: Check machine state")
	fmt.Println("---")
	fmt.Printf("Current state: %s\n", machine.GetState().String())
	err = machine.ProcessRequest()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println()

	// Example 5: Try to buy an out of stock product (after all are bought)
	fmt.Println("📌 Example 5: Out of stock scenario")
	fmt.Println("---")
	_ = machine.InsertCoins(3.00)
	err = machine.SelectProduct("2") // Pepsi
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Display final state
	fmt.Println("\n=== Final Vending Machine State ===")
	machine.DisplayProducts()
}

