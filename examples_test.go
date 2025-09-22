package testkit

import (
	"testing"
)

func TestUserBuilder_NewUserBuilder(t *testing.T) {
	builder := NewUserBuilder()

	if builder == nil {
		t.Fatal("NewUserBuilder() returned nil")
	}

	if builder.BaseBuilder == nil {
		t.Error("Expected BaseBuilder to be embedded")
	}

	if builder.user == nil {
		t.Error("Expected user to be initialized")
	}

	if builder.user.Tags == nil {
		t.Error("Expected user tags to be initialized")
	}

	if builder.user.Metadata == nil {
		t.Error("Expected user metadata to be initialized")
	}
}

func TestUserBuilder_WithMethods(t *testing.T) {
	builder := NewUserBuilder()

	// Test WithID
	result := builder.WithID(123)
	if result != builder {
		t.Error("WithID should return the same builder instance")
	}
	if builder.user.ID != 123 {
		t.Error("Expected ID to be set to 123")
	}

	// Test WithName
	builder.WithName("John Doe")
	if builder.user.Name != "John Doe" {
		t.Error("Expected name to be set to 'John Doe'")
	}

	// Test WithEmail
	builder.WithEmail("john@example.com")
	if builder.user.Email != "john@example.com" {
		t.Error("Expected email to be set")
	}

	// Test WithAge
	builder.WithAge(30)
	if builder.user.Age != 30 {
		t.Error("Expected age to be set to 30")
	}

	// Test WithActive
	builder.WithActive(true)
	if !builder.user.Active {
		t.Error("Expected active to be true")
	}

	// Test WithUserTag
	builder.WithUserTag("role", "admin")
	if builder.user.Tags["role"] != "admin" {
		t.Error("Expected user tag to be set")
	}

	// Test WithMetadata
	builder.WithMetadata("created_by", "test")
	if builder.user.Metadata["created_by"] != "test" {
		t.Error("Expected metadata to be set")
	}
}

func TestUserBuilder_Validation(t *testing.T) {
	builder := NewUserBuilder()

	// Test negative ID validation
	builder.WithID(-1)
	if !builder.HasErrors() {
		t.Error("Expected error for negative ID")
	}

	// Reset and test empty name validation
	builder = NewUserBuilder()
	builder.WithName("")
	if !builder.HasErrors() {
		t.Error("Expected error for empty name")
	}

	// Reset and test empty email validation
	builder = NewUserBuilder()
	builder.WithEmail("")
	if !builder.HasErrors() {
		t.Error("Expected error for empty email")
	}

	// Test negative age validation
	builder = NewUserBuilder()
	builder.WithAge(-1)
	if !builder.HasErrors() {
		t.Error("Expected error for negative age")
	}

	// Test that validation can be disabled
	builder = NewUserBuilder()
	builder.WithValidation(false)
	builder.WithID(-1)
	builder.WithName("")
	builder.WithEmail("")
	builder.WithAge(-1)
	if builder.HasErrors() {
		t.Error("Expected no errors when validation is disabled")
	}
}

func TestUserBuilder_Build(t *testing.T) {
	// Test successful build
	builder := NewUserBuilder()
	builder.WithID(123)
	builder.WithName("John Doe")
	builder.WithEmail("john@example.com")
	builder.WithAge(30)
	builder.WithActive(true)
	builder.WithUserTag("role", "admin")
	builder.WithMetadata("created_by", "test")

	result := builder.Build()
	user, ok := result.(*TestUser)
	if !ok {
		t.Fatalf("Expected *TestUser, got %T", result)
	}

	if user.ID != 123 {
		t.Error("Expected ID to be 123")
	}
	if user.Name != "John Doe" {
		t.Error("Expected name to be 'John Doe'")
	}
	if user.Email != "john@example.com" {
		t.Error("Expected email to be set")
	}
	if user.Age != 30 {
		t.Error("Expected age to be 30")
	}
	if !user.Active {
		t.Error("Expected active to be true")
	}
	if user.Tags["role"] != "admin" {
		t.Error("Expected role tag to be copied")
	}
	if user.Metadata["created_by"] != "test" {
		t.Error("Expected metadata to be copied")
	}

	// Test that returned user is independent copy
	user.Name = "Modified"
	if builder.user.Name == "Modified" {
		t.Error("Returned user should be independent copy")
	}
}

func TestUserBuilder_Build_ValidationErrors(t *testing.T) {
	// Test build with validation errors
	builder := NewUserBuilder()
	builder.WithID(-1) // This will add a validation error

	result := builder.Build()
	_, isError := result.(error)
	if !isError {
		t.Error("Expected error when building with validation errors")
	}

	// Test build without required fields
	builder = NewUserBuilder()
	// Don't set name or email
	result = builder.Build()
	_, isError = result.(error)
	if !isError {
		t.Error("Expected error when building without required fields")
	}
}

func TestUserBuilder_Reset(t *testing.T) {
	builder := NewUserBuilder()
	builder.WithID(123)
	builder.WithName("John Doe")
	builder.WithTag("env", "test")
	builder.AddError(nil) // This won't add an error, but let's add a real one
	builder.WithID(-1)    // This will add an error

	result := builder.Reset()
	if result != builder {
		t.Error("Reset should return the same builder instance")
	}

	// Check that user data is reset
	if builder.user.ID != 0 {
		t.Error("Expected user ID to be reset to zero value")
	}
	if builder.user.Name != "" {
		t.Error("Expected user name to be reset to empty")
	}

	// Check that base builder is reset
	if builder.HasTag("env") {
		t.Error("Expected builder tags to be reset")
	}
	if builder.HasErrors() {
		t.Error("Expected builder errors to be reset")
	}
	if !builder.IsValidationEnabled() {
		t.Error("Expected validation to be enabled after reset")
	}
}

func TestUserBuilder_Clone(t *testing.T) {
	original := NewUserBuilder()
	original.WithID(123)
	original.WithName("John Doe")
	original.WithEmail("john@example.com")
	original.WithUserTag("role", "admin")
	original.WithMetadata("created_by", "test")
	original.WithTag("env", "test")
	original.WithValidation(false)

	clone := original.Clone()
	cloneUser, ok := clone.(*UserBuilder)
	if !ok {
		t.Fatal("Expected UserBuilder clone")
	}

	if clone == original {
		t.Error("Clone should return a different instance")
	}

	// Verify clone has same state
	if cloneUser.user.ID != 123 {
		t.Error("Clone should have same user ID")
	}
	if cloneUser.user.Name != "John Doe" {
		t.Error("Clone should have same user name")
	}
	if cloneUser.user.Tags["role"] != "admin" {
		t.Error("Clone should have same user tags")
	}
	if cloneUser.user.Metadata["created_by"] != "test" {
		t.Error("Clone should have same user metadata")
	}
	if cloneUser.GetTag("env") != "test" {
		t.Error("Clone should have same builder tags")
	}
	if cloneUser.IsValidationEnabled() {
		t.Error("Clone should have same validation setting")
	}

	// Verify independence
	cloneUser.WithName("Modified")
	if original.user.Name == "Modified" {
		t.Error("Modifying clone should not affect original")
	}

	cloneUser.WithUserTag("new", "value")
	if _, exists := original.user.Tags["new"]; exists {
		t.Error("Modifying clone tags should not affect original")
	}
}

func TestUserBuilder_ApplyConfig(t *testing.T) {
	builder := NewUserBuilder()
	config := NewBuilderConfig()
	config.WithValidation(false)
	config.WithTag("env", "test")
	config.WithDefault("id", 456)
	config.WithDefault("name", "Config Name")
	config.WithDefault("email", "config@example.com")
	config.WithDefault("age", 25)
	config.WithDefault("active", true)

	err := builder.ApplyConfig(config)
	if err != nil {
		t.Errorf("Expected no error applying config, got %v", err)
	}

	// Verify configuration was applied
	if builder.IsValidationEnabled() {
		t.Error("Expected validation to be disabled")
	}
	if builder.GetTag("env") != "test" {
		t.Error("Expected tag to be applied")
	}
	if builder.user.ID != 456 {
		t.Error("Expected default ID to be applied")
	}
	if builder.user.Name != "Config Name" {
		t.Error("Expected default name to be applied")
	}
	if builder.user.Email != "config@example.com" {
		t.Error("Expected default email to be applied")
	}
	if builder.user.Age != 25 {
		t.Error("Expected default age to be applied")
	}
	if !builder.user.Active {
		t.Error("Expected default active to be applied")
	}

	// Test with nil config
	err = builder.ApplyConfig(nil)
	if err == nil {
		t.Error("Expected error with nil config")
	}
}
