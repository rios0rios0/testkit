package testkit

import (
	"errors"
	"testing"
)

func TestBaseBuilder_NewBaseBuilder(t *testing.T) {
	builder := NewBaseBuilder()

	if builder == nil {
		t.Fatal("NewBaseBuilder() returned nil")
	}

	if builder.tags == nil {
		t.Error("Expected tags to be initialized")
	}

	if !builder.validationEnabled {
		t.Error("Expected validation to be enabled by default")
	}

	if builder.errors == nil {
		t.Error("Expected errors slice to be initialized")
	}

	if len(builder.errors) != 0 {
		t.Error("Expected errors slice to be empty initially")
	}
}

func TestBaseBuilder_WithTag(t *testing.T) {
	builder := NewBaseBuilder()

	result := builder.WithTag("env", "test")
	if result != builder {
		t.Error("WithTag should return the same builder instance")
	}

	if builder.GetTag("env") != "test" {
		t.Error("Expected tag value to be 'test'")
	}

	if !builder.HasTag("env") {
		t.Error("Expected HasTag to return true for existing tag")
	}

	if builder.HasTag("nonexistent") {
		t.Error("Expected HasTag to return false for non-existing tag")
	}
}

func TestBaseBuilder_WithValidation(t *testing.T) {
	builder := NewBaseBuilder()

	// Test disabling validation
	result := builder.WithValidation(false)
	if result != builder {
		t.Error("WithValidation should return the same builder instance")
	}

	if builder.IsValidationEnabled() {
		t.Error("Expected validation to be disabled")
	}

	// Test enabling validation
	builder.WithValidation(true)
	if !builder.IsValidationEnabled() {
		t.Error("Expected validation to be enabled")
	}
}

func TestBaseBuilder_ErrorHandling(t *testing.T) {
	builder := NewBaseBuilder()

	if builder.HasErrors() {
		t.Error("Expected no errors initially")
	}

	testError := errors.New("test error")
	builder.AddError(testError)

	if !builder.HasErrors() {
		t.Error("Expected builder to have errors after adding one")
	}

	errors := builder.GetErrors()
	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %d", len(errors))
	}

	if errors[0] != testError {
		t.Error("Expected the same error instance")
	}

	builder.ClearErrors()
	if builder.HasErrors() {
		t.Error("Expected no errors after clearing")
	}
}

func TestBaseBuilder_Reset(t *testing.T) {
	builder := NewBaseBuilder()
	builder.WithTag("env", "test")
	builder.WithValidation(false)
	builder.AddError(errors.New("test error"))

	result := builder.Reset()
	if result != builder {
		t.Error("Reset should return the same builder instance")
	}

	if builder.HasTag("env") {
		t.Error("Expected tags to be cleared after reset")
	}

	if !builder.IsValidationEnabled() {
		t.Error("Expected validation to be enabled after reset")
	}

	if builder.HasErrors() {
		t.Error("Expected errors to be cleared after reset")
	}
}

func TestBaseBuilder_Clone(t *testing.T) {
	original := NewBaseBuilder()
	original.WithTag("env", "test")
	original.WithValidation(false)
	original.AddError(errors.New("test error"))

	clone := original.Clone()

	if clone == original {
		t.Error("Clone should return a different instance")
	}

	cloneBase, ok := clone.(*BaseBuilder)
	if !ok {
		t.Fatal("Clone should return a BaseBuilder instance")
	}

	// Verify clone has same state
	if cloneBase.GetTag("env") != "test" {
		t.Error("Clone should have the same tags")
	}

	if cloneBase.IsValidationEnabled() {
		t.Error("Clone should have the same validation setting")
	}

	if !cloneBase.HasErrors() {
		t.Error("Clone should have the same errors")
	}

	// Verify independence - modifying clone shouldn't affect original
	cloneBase.WithTag("new", "value")
	if original.HasTag("new") {
		t.Error("Modifying clone should not affect original")
	}
}

func TestBaseBuilder_Build(t *testing.T) {
	builder := NewBaseBuilder()
	result := builder.Build()

	if result != nil {
		t.Error("BaseBuilder.Build() should return nil by default")
	}
}

func TestBaseBuilder_NilSafety(t *testing.T) {
	builder := NewBaseBuilder()

	// Test adding nil error
	builder.AddError(nil)
	if builder.HasErrors() {
		t.Error("Adding nil error should not add to errors")
	}

	// Test with nil tags map (edge case)
	builder.tags = nil
	if builder.GetTag("any") != "" {
		t.Error("GetTag on nil tags should return empty string")
	}

	if builder.HasTag("any") {
		t.Error("HasTag on nil tags should return false")
	}

	// WithTag should initialize the map
	builder.WithTag("key", "value")
	if builder.tags == nil {
		t.Error("WithTag should initialize tags map")
	}
}
