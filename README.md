# testkit

Go testing utility library offering a modular builder framework to streamline test environment setup. It provides interfaces, ready-to-use structs, and patterns for creating reusable builders, fixtures, and mocks. Extensible and ideal for consistent, maintainable tests across repos.

## Features

- **Modular Builder Pattern**: Extensible base builder with common functionality
- **Factory System**: Register and create builders dynamically
- **Configuration Management**: Apply default values and settings to builders
- **Validation Framework**: Built-in validation with error accumulation
- **Tag System**: Metadata support for builders
- **Clone & Reset**: Deep copy and state management capabilities

## Installation

```bash
go get github.com/rios0rios0/testkit
```

## Quick Start

### Basic Builder Usage

```go
package main

import (
    "fmt"
    "github.com/rios0rios0/testkit"
)

func main() {
    // Create a base builder
    builder := testkit.NewBaseBuilder()
    builder.WithTag("env", "test").WithValidation(true)
    
    // Use the pre-built UserBuilder example
    userBuilder := testkit.NewUserBuilder()
    userBuilder.WithName("John Doe").
        WithEmail("john@example.com").
        WithAge(30).
        WithActive(true)
    
    result := userBuilder.Build()
    if user, ok := result.(*testkit.TestUser); ok {
        fmt.Printf("Created user: %+v\n", user)
    }
}
```

### Factory Pattern

```go
// Register a custom builder
testkit.RegisterBuilder("mybuilder", func() testkit.Builder {
    return testkit.NewUserBuilder()
})

// Create from factory
builder, err := testkit.CreateBuilder("mybuilder")
if err != nil {
    panic(err)
}

// Or use a custom factory
factory := testkit.NewBuilderFactory()
factory.Register("custom", func() testkit.Builder {
    return testkit.NewBaseBuilder()
})
```

### Configuration System

```go
// Create configuration with defaults
config := testkit.NewBuilderConfig()
config.WithValidation(false).
    WithTag("env", "test").
    WithDefault("name", "Default Name").
    WithDefault("age", 25)

// Apply to any builder
builder := testkit.NewUserBuilder()
err := config.ApplyTo(builder)
if err != nil {
    panic(err)
}
```

## Creating Custom Builders

### Step 1: Define Your Entity

```go
type Product struct {
    ID          int
    Name        string
    Price       float64
    Category    string
    InStock     bool
    Tags        map[string]string
}
```

### Step 2: Create the Builder

```go
type ProductBuilder struct {
    *testkit.BaseBuilder
    product *Product
}

func NewProductBuilder() *ProductBuilder {
    return &ProductBuilder{
        BaseBuilder: testkit.NewBaseBuilder(),
        product: &Product{
            Tags: make(map[string]string),
        },
    }
}

func (b *ProductBuilder) WithID(id int) *ProductBuilder {
    if b.IsValidationEnabled() && id <= 0 {
        b.AddError(errors.New("product ID must be positive"))
        return b
    }
    b.product.ID = id
    return b
}

func (b *ProductBuilder) WithName(name string) *ProductBuilder {
    if b.IsValidationEnabled() && name == "" {
        b.AddError(errors.New("product name cannot be empty"))
        return b
    }
    b.product.Name = name
    return b
}

func (b *ProductBuilder) WithPrice(price float64) *ProductBuilder {
    if b.IsValidationEnabled() && price < 0 {
        b.AddError(errors.New("product price cannot be negative"))
        return b
    }
    b.product.Price = price
    return b
}

func (b *ProductBuilder) Build() interface{} {
    if b.HasErrors() {
        return fmt.Errorf("validation errors: %v", b.GetErrors())
    }
    
    // Return a copy to avoid mutation
    return &Product{
        ID:       b.product.ID,
        Name:     b.product.Name,
        Price:    b.product.Price,
        Category: b.product.Category,
        InStock:  b.product.InStock,
        Tags:     copyMap(b.product.Tags),
    }
}

func (b *ProductBuilder) Reset() testkit.Builder {
    b.BaseBuilder.Reset()
    b.product = &Product{Tags: make(map[string]string)}
    return b
}

func (b *ProductBuilder) Clone() testkit.Builder {
    clone := &ProductBuilder{
        BaseBuilder: b.BaseBuilder.Clone().(*testkit.BaseBuilder),
        product: &Product{
            ID:       b.product.ID,
            Name:     b.product.Name,
            Price:    b.product.Price,
            Category: b.product.Category,
            InStock:  b.product.InStock,
            Tags:     copyMap(b.product.Tags),
        },
    }
    return clone
}

func copyMap(m map[string]string) map[string]string {
    result := make(map[string]string)
    for k, v := range m {
        result[k] = v
    }
    return result
}
```

### Step 3: Register with Factory (Optional)

```go
func init() {
    testkit.RegisterBuilder("product", func() testkit.Builder {
        return NewProductBuilder()
    })
}
```

## Advanced Features

### Error Handling

```go
builder := NewProductBuilder()
builder.WithID(-1).WithName("") // Both will add validation errors

result := builder.Build()
if err, ok := result.(error); ok {
    fmt.Printf("Build failed: %v\n", err)
}

// Check errors during building
if builder.HasErrors() {
    for _, err := range builder.GetErrors() {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Builder State Management

```go
// Clone for independent copies
original := NewProductBuilder().WithName("Original")
clone := original.Clone().(*ProductBuilder)
clone.WithName("Clone") // Doesn't affect original

// Reset for reuse
builder := NewProductBuilder().WithName("First")
builder.Reset()
builder.WithName("Second") // Fresh state
```

### Metadata and Tags

```go
builder := NewProductBuilder()
builder.WithTag("test_type", "integration").
    WithTag("owner", "team_a")

if builder.HasTag("test_type") {
    fmt.Printf("Test type: %s\n", builder.GetTag("test_type"))
}
```

## Best Practices

1. **Always validate inputs** when validation is enabled
2. **Return copies** from Build() to prevent mutation
3. **Use meaningful error messages** for validation failures
4. **Implement Clone() and Reset()** for proper state management
5. **Register builders** with descriptive names in the factory
6. **Use tags** for test categorization and metadata
7. **Apply configurations** consistently across related tests

## API Reference

### Core Interfaces

- `Builder`: Main interface all builders must implement
- `ConfigurableBuilder`: Optional interface for configuration support

### Main Types

- `BaseBuilder`: Common functionality for all builders
- `BuilderFactory`: Factory for creating and managing builders
- `BuilderConfig`: Configuration container for builders

### Example Builders

- `UserBuilder`: Complete example of a custom builder
- `TestUser`: Example entity for testing

## Contributing

This library is designed to be extended. Feel free to:

1. Create additional builder types
2. Add new validation patterns
3. Extend the configuration system
4. Improve error handling

## License

MIT License - see LICENSE file for details.
