// Package testkit provides utilities for testing setups including a modular builder framework.
package testkit

// Builder defines the interface that all builders must implement.
// This provides a common contract for all test builders in the library.
type Builder interface {
	// Build performs the final construction step and returns the built object.
	// Implementations should return the appropriate type for their specific builder.
	Build() interface{}

	// Reset clears the builder state, allowing it to be reused.
	Reset() Builder

	// Clone creates a deep copy of the builder with the same configuration.
	Clone() Builder
}

// BaseBuilder provides common functionality for all builders.
// It implements the Builder interface and can be embedded in specific builders.
type BaseBuilder struct {
	// tags holds metadata tags for the builder
	tags map[string]string
	// validationEnabled controls whether validation should be performed
	validationEnabled bool
	// errors holds any validation or configuration errors
	errors []error
}

// NewBaseBuilder creates a new BaseBuilder instance with default settings.
func NewBaseBuilder() *BaseBuilder {
	return &BaseBuilder{
		tags:              make(map[string]string),
		validationEnabled: true,
		errors:            make([]error, 0),
	}
}

// WithTag adds a metadata tag to the builder.
// Tags can be used for identification, debugging, or conditional logic.
func (b *BaseBuilder) WithTag(key, value string) *BaseBuilder {
	if b.tags == nil {
		b.tags = make(map[string]string)
	}
	b.tags[key] = value
	return b
}

// GetTag retrieves a metadata tag value by key.
// Returns empty string if the tag doesn't exist.
func (b *BaseBuilder) GetTag(key string) string {
	if b.tags == nil {
		return ""
	}
	return b.tags[key]
}

// HasTag checks if a metadata tag exists.
func (b *BaseBuilder) HasTag(key string) bool {
	if b.tags == nil {
		return false
	}
	_, exists := b.tags[key]
	return exists
}

// WithValidation enables or disables validation for this builder.
func (b *BaseBuilder) WithValidation(enabled bool) *BaseBuilder {
	b.validationEnabled = enabled
	return b
}

// IsValidationEnabled returns whether validation is enabled for this builder.
func (b *BaseBuilder) IsValidationEnabled() bool {
	return b.validationEnabled
}

// AddError adds an error to the builder's error collection.
func (b *BaseBuilder) AddError(err error) *BaseBuilder {
	if err != nil {
		b.errors = append(b.errors, err)
	}
	return b
}

// GetErrors returns all errors accumulated by the builder.
func (b *BaseBuilder) GetErrors() []error {
	return b.errors
}

// HasErrors returns true if the builder has any errors.
func (b *BaseBuilder) HasErrors() bool {
	return len(b.errors) > 0
}

// ClearErrors removes all errors from the builder.
func (b *BaseBuilder) ClearErrors() *BaseBuilder {
	b.errors = make([]error, 0)
	return b
}

// Build is a default implementation that returns nil.
// Specific builders should override this method.
func (b *BaseBuilder) Build() interface{} {
	return nil
}

// Reset clears the builder state, allowing it to be reused.
func (b *BaseBuilder) Reset() Builder {
	b.tags = make(map[string]string)
	b.validationEnabled = true
	b.errors = make([]error, 0)
	return b
}

// Clone creates a deep copy of the BaseBuilder.
func (b *BaseBuilder) Clone() Builder {
	clone := &BaseBuilder{
		tags:              make(map[string]string),
		validationEnabled: b.validationEnabled,
		errors:            make([]error, len(b.errors)),
	}

	// Deep copy tags
	for k, v := range b.tags {
		clone.tags[k] = v
	}

	// Deep copy errors
	copy(clone.errors, b.errors)

	return clone
}
