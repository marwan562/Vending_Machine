package main

import (
	"testing"
)

func TestNewVendingMachine(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}

	// Test successful creation
	machine, err := NewVendingMachine(products)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if machine == nil {
		t.Fatal("Expected machine, got nil")
	}

	// Test error on empty products
	_, err = NewVendingMachine([]Product{})
	if err == nil {
		t.Fatal("Expected error for empty products, got nil")
	}
}

func TestInsertCoins(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)

	// Test valid coin insertion
	err := machine.InsertCoins(1.50)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if machine.GetBalance() != 1.50 {
		t.Errorf("Expected balance 1.50, got %f", machine.GetBalance())
	}

	// Test invalid coin insertion
	err = machine.InsertCoins(-1.0)
	if err == nil {
		t.Fatal("Expected error for negative coins, got nil")
	}

	err = machine.InsertCoins(0)
	if err == nil {
		t.Fatal("Expected error for zero coins, got nil")
	}
}

func TestSelectProduct(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
		{ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 0}, // Out of stock
	}
	machine, _ := NewVendingMachine(products)

	// Test insufficient funds
	machine.InsertCoins(0.50)
	err := machine.SelectProduct("1")
	if err == nil {
		t.Fatal("Expected error for insufficient funds, got nil")
	}

	// Clear balance and try again
	machine.RefundBalance()

	// Test successful purchase
	machine.InsertCoins(1.50)
	err = machine.SelectProduct("1")
	if err != nil {
		t.Fatalf("Expected successful purchase, got error: %v", err)
	}

	// Verify quantity decreased
	prods := machine.GetProducts()
	if prods[0].Quantity != 9 {
		t.Errorf("Expected quantity 9, got %d", prods[0].Quantity)
	}

	// Verify balance was reset
	if machine.GetBalance() != 0 {
		t.Errorf("Expected balance 0, got %f", machine.GetBalance())
	}

	// Test out of stock product
	machine.InsertCoins(1.50)
	err = machine.SelectProduct("2")
	if err == nil {
		t.Fatal("Expected error for out of stock, got nil")
	}
}

func TestSelectNonexistentProduct(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)
	machine.InsertCoins(5.00)

	err := machine.SelectProduct("999")
	if err == nil {
		t.Fatal("Expected error for nonexistent product, got nil")
	}
}

func TestAddProducts(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)

	// Test adding existing product (restock)
	newProducts := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 5},
	}
	err := machine.AddProducts(newProducts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	products = machine.GetProducts()
	if products[0].Quantity != 15 {
		t.Errorf("Expected quantity 15, got %d", products[0].Quantity)
	}

	// Test adding new product
	newProducts = []Product{
		{ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 8},
	}
	err = machine.AddProducts(newProducts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	products = machine.GetProducts()
	if len(products) != 2 {
		t.Errorf("Expected 2 products, got %d", len(products))
	}

	// Test adding empty products list
	err = machine.AddProducts([]Product{})
	if err == nil {
		t.Fatal("Expected error for empty products, got nil")
	}
}

func TestRemoveProduct(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
		{ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 8},
	}
	machine, _ := NewVendingMachine(products)

	// Test removing existing product
	err := machine.RemoveProduct("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	products = machine.GetProducts()
	if len(products) != 1 {
		t.Errorf("Expected 1 product, got %d", len(products))
	}

	// Test removing nonexistent product
	err = machine.RemoveProduct("999")
	if err == nil {
		t.Fatal("Expected error for nonexistent product, got nil")
	}
}

func TestRefundBalance(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)

	machine.InsertCoins(2.50)
	refund := machine.RefundBalance()

	if refund != 2.50 {
		t.Errorf("Expected refund 2.50, got %f", refund)
	}

	if machine.GetBalance() != 0 {
		t.Errorf("Expected balance 0, got %f", machine.GetBalance())
	}
}

func TestReset(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)

	machine.InsertCoins(2.50)
	machine.Reset()

	if machine.GetBalance() != 0 {
		t.Errorf("Expected balance 0 after reset, got %f", machine.GetBalance())
	}

	if len(machine.GetProducts()) != 0 {
		t.Errorf("Expected 0 products after reset, got %d", len(machine.GetProducts()))
	}
}

func TestStateTransitions(t *testing.T) {
	products := []Product{
		{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
	}
	machine, _ := NewVendingMachine(products)

	// Initial state should be Idle
	if machine.GetState().String() != "Idle" {
		t.Errorf("Expected Idle state, got %s", machine.GetState().String())
	}

	// Set to different state
	machine.SetState(&AcceptingCoinsState{})
	if machine.GetState().String() != "AcceptingCoins" {
		t.Errorf("Expected AcceptingCoins state, got %s", machine.GetState().String())
	}
}

func TestProductHelpers(t *testing.T) {
	available := Product{ID: "1", Name: "Coke", Price: 1.50, Quantity: 10}
	outOfStock := Product{ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 0}

	// Test IsAvailable
	if !available.IsAvailable() {
		t.Fatal("Expected available product to be available")
	}

	if outOfStock.IsAvailable() {
		t.Fatal("Expected out of stock product to not be available")
	}

	// Test IsAffordable
	if !available.IsAffordable(1.50) {
		t.Fatal("Expected 1.50 to afford 1.50 product")
	}

	if available.IsAffordable(1.00) {
		t.Fatal("Expected 1.00 to not afford 1.50 product")
	}
}
