<h1 align="center">testkit</h1>
<p align="center">
    <a href="https://github.com/rios0rios0/testkit/releases/latest">
        <img src="https://img.shields.io/github/release/rios0rios0/testkit.svg?style=for-the-badge&logo=github" alt="Latest Release"/></a>
    <a href="https://github.com/rios0rios0/testkit/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/rios0rios0/testkit.svg?style=for-the-badge&logo=github" alt="License"/></a>
    <a href="https://github.com/rios0rios0/testkit/actions/workflows/default.yaml">
        <img src="https://img.shields.io/github/actions/workflow/status/rios0rios0/testkit/default.yaml?branch=main&style=for-the-badge&logo=github" alt="Build Status"/></a>
    <a href="https://sonarcloud.io/summary/overall?id=rios0rios0_testkit">
        <img src="https://img.shields.io/sonar/coverage/rios0rios0_testkit?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonarqubecloud" alt="Coverage"/></a>
    <a href="https://sonarcloud.io/summary/overall?id=rios0rios0_testkit">
        <img src="https://img.shields.io/sonar/quality_gate/rios0rios0_testkit?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonarqubecloud" alt="Quality Gate"/></a>
    <a href="https://www.bestpractices.dev/projects/12032">
        <img src="https://img.shields.io/cii/level/12032?style=for-the-badge&logo=opensourceinitiative" alt="OpenSSF Best Practices"/></a>
</p>

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

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
