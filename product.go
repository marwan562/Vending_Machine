package main

// Product represents an item that can be sold in the vending machine
type Product struct {
	ID       string  // Unique identifier for the product
	Name     string  // Display name of the product
	Price    float64 // Price in dollars
	Quantity int     // Number of units available
}

// IsAvailable checks if the product is in stock
func (p *Product) IsAvailable() bool {
	return p.Quantity > 0
}

// IsAffordable checks if a given balance can purchase this product
func (p *Product) IsAffordable(balance float64) bool {
	return balance >= p.Price
}
