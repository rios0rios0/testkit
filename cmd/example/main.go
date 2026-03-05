package main

import (
	"log"

	test "github.com/rios0rios0/testkit/pkg/test"
)

const (
	exampleAge       = 28
	factoryAge       = 35
	configDefaultAge = 30
)

func main() {
	log.Println("=== Testkit Library Demo ===")

	basicUserBuilderExample()
	factoryPatternExample()
	configurationSystemExample()
	validationExample()
	builderStateManagementExample()
	customFactoryExample()

	log.Println("\n=== Demo Complete ===")
}

func basicUserBuilderExample() {
	log.Println("1. Basic UserBuilder Usage:")
	userBuilder := test.NewUserBuilder()
	userBuilder.WithName("Alice Smith").
		WithEmail("alice@example.com").
		WithAge(exampleAge).
		WithActive(true).
		WithUserTag("department", "engineering").
		WithMetadata("hire_date", "2023-01-15")

	result := userBuilder.Build()
	switch v := result.(type) {
	case *test.TestUser:
		log.Printf("   Created user: %+v\n", v)
	case error:
		log.Printf("   Error: %v\n", v)
	}
}

func factoryPatternExample() {
	log.Println("\n2. Factory Pattern Usage:")
	builder, err := test.CreateBuilder("user")
	if err != nil {
		log.Fatal(err)
	}

	userBuilder, isUserBuilder := builder.(*test.UserBuilder)
	if !isUserBuilder {
		log.Fatal("unexpected builder type")
	}

	userBuilder.WithName("Bob Wilson").
		WithEmail("bob@example.com").
		WithAge(factoryAge)

	result := userBuilder.Build()
	if user, isUser := result.(*test.TestUser); isUser {
		log.Printf("   Factory-created user: %+v\n", user)
	}
}

func configurationSystemExample() {
	log.Println("\n3. Configuration System:")
	config := test.NewBuilderConfig()
	config.WithValidation(true).
		WithTag("env", "test").
		WithTag("team", "qa").
		WithDefault("name", "Test User").
		WithDefault("email", "test@example.com").
		WithDefault("age", configDefaultAge).
		WithDefault("active", true)

	configuredBuilder := test.NewUserBuilder()
	if err := config.ApplyTo(configuredBuilder); err != nil {
		log.Fatal(err)
	}

	result := configuredBuilder.Build()
	if user, ok := result.(*test.TestUser); ok {
		log.Printf("   Configured user: %+v\n", user)
		log.Printf("   Builder tags: env=%s, team=%s\n",
			configuredBuilder.GetTag("env"),
			configuredBuilder.GetTag("team"))
	}
}

func validationExample() {
	log.Println("\n4. Validation and Error Handling:")
	invalidBuilder := test.NewUserBuilder()
	invalidBuilder.WithName(""). // Invalid empty name
					WithEmail(""). // Invalid empty email
					WithAge(-5)    // Invalid negative age

	result := invalidBuilder.Build()
	if buildErr, ok := result.(error); ok {
		log.Printf("   Validation failed as expected: %v\n", buildErr)
	}
}

func builderStateManagementExample() {
	log.Println("\n5. Builder State Management:")
	originalBuilder := test.NewUserBuilder()
	originalBuilder.WithName("Original User").
		WithEmail("original@example.com").
		WithTag("version", "v1")

	// Clone the builder
	clonedBuilder, ok := originalBuilder.Clone().(*test.UserBuilder)
	if !ok {
		log.Fatal("unexpected clone type")
	}

	clonedBuilder.WithName("Cloned User").
		WithTag("version", "v2")

	log.Printf("   Original builder name: %s, tag: %s\n",
		originalBuilder.GetTag("version"),
		getUserName(originalBuilder))

	log.Printf("   Cloned builder name: %s, tag: %s\n",
		clonedBuilder.GetTag("version"),
		getUserName(clonedBuilder))

	// Reset the original builder
	originalBuilder.Reset()
	log.Printf("   After reset, original builder has errors: %v\n",
		originalBuilder.HasErrors())
}

func customFactoryExample() {
	log.Println("\n6. Custom Factory Usage:")
	customFactory := test.NewBuilderFactory()

	if err := customFactory.Register("admin_user", func() test.Builder {
		b := test.NewUserBuilder()
		b.WithUserTag("role", "admin").
			WithActive(true)
		return b
	}); err != nil {
		log.Fatal(err)
	}

	adminBuilder, err := customFactory.Create("admin_user")
	if err != nil {
		log.Fatal(err)
	}

	adminUserBuilder, isUserBuilder := adminBuilder.(*test.UserBuilder)
	if !isUserBuilder {
		log.Fatal("unexpected builder type")
	}

	adminUserBuilder.WithName("Admin User").
		WithEmail("admin@example.com")

	result := adminUserBuilder.Build()
	if user, isUser := result.(*test.TestUser); isUser {
		log.Printf("   Admin user: %+v\n", user)
	}
}

func getUserName(builder *test.UserBuilder) string {
	switch v := builder.Build().(type) {
	case *test.TestUser:
		return v.Name
	default:
		return "unknown"
	}
}
