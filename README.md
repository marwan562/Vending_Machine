# Vending Machine - State Design Pattern

A well-structured Go implementation of the State Design Pattern using a vending machine as an example. This is an educational project demonstrating GoLang best practices and design patterns.

## 📋 Project Structure

```
├── main.go              # Entry point with demo examples
├── state.go             # State interface and concrete state implementations
├── vendingMachine.go    # VendingMachine context and core logic
├── product.go           # Product model with helper methods
├── go.mod              # Go module definition
└── README.md           # This file
```

## 🎯 Key Improvements in Refactoring

### 1. **Proper State Pattern Implementation**
- **Before**: States were simple string constants
- **After**: State is an interface with concrete implementations (`IdleState`, `AcceptingCoinsState`, etc.)
- **Benefit**: Type-safe, extensible, and follows Go idioms

### 2. **Better Error Handling**
- **Before**: Silent failures with no error feedback
- **After**: All operations return `error` for proper error handling
- **Benefit**: Callers can detect and handle errors appropriately

### 3. **Cleaner Naming Conventions**
- **Before**: `VendingMachineContext`, `GetTotalCoins()`, `Request()`
- **After**: `VendingMachine`, `GetBalance()`, `ProcessRequest()`
- **Benefit**: More intuitive and idiomatic Go naming

### 4. **Improved Method Names**
- **Before**: `InsertCoins()` took no return, `RefundCoins()` returned void
- **After**: `InsertCoins()` returns error, `RefundBalance()` returns the refund amount
- **Benefit**: Better information flow and functional clarity

### 5. **Helper Methods on Model**
- **Before**: No methods on `Product`
- **After**: `IsAvailable()`, `IsAffordable()` methods on `Product`
- **Benefit**: Encapsulation and code reusability

### 6. **Better Separation of Concerns**
- **Before**: Mixed logic for finding/removing products
- **After**: Private `findProduct()` and `dispenseProduct()` helpers
- **Benefit**: Internal implementation details are private, API is clean

### 7. **Comprehensive Documentation**
- Added godoc-style comments on all public types and methods
- Better variable naming for clarity
- Inline comments for complex logic

### 8. **Utility Methods**
- `DisplayProducts()`: Shows all products in a formatted table
- Improved `Reset()` method that properly reinitializes state

## 📚 Architecture

### State Interface
```go
type State interface {
    HandleRequest(machine *VendingMachine) error
    String() string
}
```

### States
- **IdleState**: Machine is idle, waiting for coins
- **AcceptingCoinsState**: Machine accepts coins from user
- **SelectingProductState**: User is selecting a product
- **DispensingState**: Machine is dispensing a product
- **OutOfStockState**: Selected product is out of stock

### VendingMachine (Context)
The main context that holds the current state and manages transitions between states.

## 🚀 Usage Example

```go
// Create a vending machine
products := []Product{
    {ID: "1", Name: "Coke", Price: 1.50, Quantity: 10},
    {ID: "2", Name: "Pepsi", Price: 1.50, Quantity: 8},
}

machine, err := NewVendingMachine(products)
if err != nil {
    log.Fatal(err)
}

// Insert coins
machine.InsertCoins(1.50)

// Select a product
err = machine.SelectProduct("1")
if err != nil {
    fmt.Printf("Error: %v\n", err)
}

// Display products
machine.DisplayProducts()
```

## 🏃 Running the Project

```bash
# Run the demo
go run .

# Or run with more verbose output
go run . -v
```

## 📖 Go Best Practices Used

1. **Interfaces**: State is defined as an interface for flexibility
2. **Error Returns**: Functions return errors instead of using panic
3. **Receiver Methods**: Methods on types (not pointers unless needed)
4. **Private/Public**: Lowercase for private, uppercase for public
5. **Documentation**: Comments on all exported types and functions
6. **Idiomatic Naming**: Meaningful names that follow Go conventions
7. **Encapsulation**: Helper functions are private with underscore prefix

## 🔧 Extending the System

To add new features:

1. **New States**: Create a new struct implementing the `State` interface
2. **New Operations**: Add methods to `VendingMachine` 
3. **New Products**: Just initialize `Product` structs with different data

Example - Adding a ReturningChangeState:
```go
type ReturningChangeState struct{}

func (s *ReturningChangeState) HandleRequest(machine *VendingMachine) error {
    // Return change logic
    return nil
}

func (s *ReturningChangeState) String() string {
    return "ReturningChange"
}
```

## 📝 Learning Outcomes

This refactored project demonstrates:
- ✅ State Design Pattern implementation
- ✅ Interface-driven design in Go
- ✅ Error handling best practices
- ✅ Clean code principles
- ✅ Go idioms and conventions
- ✅ Proper encapsulation
- ✅ Documentation and comments

## 🎓 Original vs Refactored Comparison

| Aspect | Original | Refactored |
|--------|----------|-----------|
| States | String constants | Interface-based |
| Error Handling | None | Comprehensive |
| Method Names | Ambiguous | Clear and idiomatic |
| Documentation | None | Full godoc comments |
| Validation | Minimal | Comprehensive |
| Code Organization | Basic | Well-structured |
| Extensibility | Limited | Highly extensible |

---

**Happy Learning! 🚀**
