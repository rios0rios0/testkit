/*
Package testkit provides utilities for testing setups including a modular builder framework.

The testkit library offers a comprehensive set of tools for creating consistent, maintainable
test environments. It provides interfaces, ready-to-use structs, and patterns for creating
reusable builders, fixtures, and mocks.

# Core Components

The library is built around three main components:

1. Builder Interface and BaseBuilder: A flexible builder pattern implementation
2. BuilderFactory: Factory pattern for creating and managing different builder types
3. BuilderConfig: Configuration system for setting up builders with default values

# Basic Usage

Creating a simple builder:

	builder := NewBaseBuilder()
	builder.WithTag("env", "test").WithValidation(true)
	result := builder.Build() // Returns nil for base builder

Using the factory pattern:

	factory := NewBuilderFactory()
	factory.Register("mybuilder", func() Builder {
		return NewBaseBuilder()
	})
	
	builder, err := factory.Create("mybuilder")
	if err != nil {
		// handle error
	}

# Creating Custom Builders

To create a custom builder, embed BaseBuilder and implement the Builder interface:

	type MyObjectBuilder struct {
		*BaseBuilder
		obj *MyObject
	}
	
	func NewMyObjectBuilder() *MyObjectBuilder {
		return &MyObjectBuilder{
			BaseBuilder: NewBaseBuilder(),
			obj:         &MyObject{},
		}
	}
	
	func (b *MyObjectBuilder) WithName(name string) *MyObjectBuilder {
		if b.IsValidationEnabled() && name == "" {
			b.AddError(errors.New("name cannot be empty"))
			return b
		}
		b.obj.Name = name
		return b
	}
	
	func (b *MyObjectBuilder) Build() interface{} {
		if b.HasErrors() {
			return fmt.Errorf("validation errors: %v", b.GetErrors())
		}
		// Return a copy to avoid mutation
		return &MyObject{Name: b.obj.Name}
	}

# Configuration System

Use BuilderConfig for setting up builders with defaults:

	config := NewBuilderConfig()
	config.WithValidation(false)
	config.WithTag("env", "test")
	config.WithDefault("name", "default_name")
	
	builder := NewMyObjectBuilder()
	config.ApplyTo(builder)

# Thread Safety

The builders in this library are not thread-safe by design. Each goroutine should
use its own builder instance. Use Clone() to create independent copies when needed.

# Error Handling

Builders accumulate errors during configuration and report them at build time:

	builder := NewMyObjectBuilder()
	builder.WithName("") // Adds validation error
	
	result := builder.Build()
	if err, ok := result.(error); ok {
		// Handle build errors
	}
*/
package testkit