package main

import (
	"fmt"
	"log"

	test "github.com/rios0rios0/testkit/pkg/test"
)

func main() {
	fmt.Println("=== Testkit Library Demo ===")

	// Example 1: Basic UserBuilder usage
	fmt.Println("1. Basic UserBuilder Usage:")
	userBuilder := test.NewUserBuilder()
	userBuilder.WithName("Alice Smith").
		WithEmail("alice@example.com").
		WithAge(28).
		WithActive(true).
		WithUserTag("department", "engineering").
		WithMetadata("hire_date", "2023-01-15")

	result := userBuilder.Build()
	if user, ok := result.(*test.TestUser); ok {
		fmt.Printf("   Created user: %+v\n", user)
	} else if err, ok := result.(error); ok {
		fmt.Printf("   Error: %v\n", err)
	}

	// Example 2: Factory Pattern
	fmt.Println("\n2. Factory Pattern Usage:")
	builder, err := test.CreateBuilder("user")
	if err != nil {
		log.Fatal(err)
	}
	
	userBuilder2 := builder.(*test.UserBuilder)
	userBuilder2.WithName("Bob Wilson").
		WithEmail("bob@example.com").
		WithAge(35)

	result2 := userBuilder2.Build()
	if user, ok := result2.(*test.TestUser); ok {
		fmt.Printf("   Factory-created user: %+v\n", user)
	}

	// Example 3: Configuration System
	fmt.Println("\n3. Configuration System:")
	config := test.NewBuilderConfig()
	config.WithValidation(true).
		WithTag("env", "test").
		WithTag("team", "qa").
		WithDefault("name", "Test User").
		WithDefault("email", "test@example.com").
		WithDefault("age", 30).
		WithDefault("active", true)

	configuredBuilder := test.NewUserBuilder()
	err = config.ApplyTo(configuredBuilder)
	if err != nil {
		log.Fatal(err)
	}

	result3 := configuredBuilder.Build()
	if user, ok := result3.(*test.TestUser); ok {
		fmt.Printf("   Configured user: %+v\n", user)
		fmt.Printf("   Builder tags: env=%s, team=%s\n", 
			configuredBuilder.GetTag("env"), 
			configuredBuilder.GetTag("team"))
	}

	// Example 4: Validation and Error Handling
	fmt.Println("\n4. Validation and Error Handling:")
	invalidBuilder := test.NewUserBuilder()
	invalidBuilder.WithName("").  // Invalid empty name
		WithEmail("").              // Invalid empty email
		WithAge(-5)                 // Invalid negative age

	result4 := invalidBuilder.Build()
	if err, ok := result4.(error); ok {
		fmt.Printf("   Validation failed as expected: %v\n", err)
	}

	// Example 5: Builder State Management
	fmt.Println("\n5. Builder State Management:")
	originalBuilder := test.NewUserBuilder()
	originalBuilder.WithName("Original User").
		WithEmail("original@example.com").
		WithTag("version", "v1")

	// Clone the builder
	clonedBuilder := originalBuilder.Clone().(*test.UserBuilder)
	clonedBuilder.WithName("Cloned User").
		WithTag("version", "v2")

	fmt.Printf("   Original builder name: %s, tag: %s\n", 
		originalBuilder.GetTag("version"),
		func() string {
			if result := originalBuilder.Build(); result != nil {
				if user, ok := result.(*test.TestUser); ok {
					return user.Name
				}
			}
			return "unknown"
		}())

	fmt.Printf("   Cloned builder name: %s, tag: %s\n", 
		clonedBuilder.GetTag("version"),
		func() string {
			if result := clonedBuilder.Build(); result != nil {
				if user, ok := result.(*test.TestUser); ok {
					return user.Name
				}
			}
			return "unknown"
		}())

	// Reset the original builder
	originalBuilder.Reset()
	fmt.Printf("   After reset, original builder has errors: %v\n", 
		originalBuilder.HasErrors())

	// Example 6: Custom Factory
	fmt.Println("\n6. Custom Factory Usage:")
	customFactory := test.NewBuilderFactory()
	customFactory.Register("admin_user", func() test.Builder {
		builder := test.NewUserBuilder()
		builder.WithUserTag("role", "admin").
			WithActive(true)
		return builder
	})

	adminBuilder, err := customFactory.Create("admin_user")
	if err != nil {
		log.Fatal(err)
	}

	adminUserBuilder := adminBuilder.(*test.UserBuilder)
	adminUserBuilder.WithName("Admin User").
		WithEmail("admin@example.com")

	result6 := adminUserBuilder.Build()
	if user, ok := result6.(*test.TestUser); ok {
		fmt.Printf("   Admin user: %+v\n", user)
	}

	fmt.Println("\n=== Demo Complete ===")
}