package testkit

import (
	"errors"
	"fmt"
)

// TestUser represents a test user entity for demonstration purposes.
type TestUser struct {
	ID       int
	Name     string
	Email    string
	Age      int
	Active   bool
	Tags     map[string]string
	Metadata map[string]interface{}
}

// UserBuilder builds TestUser instances for testing.
// It embeds BaseBuilder to inherit common functionality.
type UserBuilder struct {
	*BaseBuilder
	user *TestUser
}

// NewUserBuilder creates a new UserBuilder instance.
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		BaseBuilder: NewBaseBuilder(),
		user: &TestUser{
			Tags:     make(map[string]string),
			Metadata: make(map[string]interface{}),
		},
	}
}

// WithID sets the user ID.
func (b *UserBuilder) WithID(id int) *UserBuilder {
	if b.IsValidationEnabled() && id < 0 {
		b.AddError(errors.New("user ID must be non-negative"))
		return b
	}
	b.user.ID = id
	return b
}

// WithName sets the user name.
func (b *UserBuilder) WithName(name string) *UserBuilder {
	if b.IsValidationEnabled() && name == "" {
		b.AddError(errors.New("user name cannot be empty"))
		return b
	}
	b.user.Name = name
	return b
}

// WithEmail sets the user email.
func (b *UserBuilder) WithEmail(email string) *UserBuilder {
	if b.IsValidationEnabled() && email == "" {
		b.AddError(errors.New("user email cannot be empty"))
		return b
	}
	b.user.Email = email
	return b
}

// WithAge sets the user age.
func (b *UserBuilder) WithAge(age int) *UserBuilder {
	if b.IsValidationEnabled() && age < 0 {
		b.AddError(errors.New("user age must be non-negative"))
		return b
	}
	b.user.Age = age
	return b
}

// WithActive sets the user active status.
func (b *UserBuilder) WithActive(active bool) *UserBuilder {
	b.user.Active = active
	return b
}

// WithUserTag adds a tag specific to the user entity.
func (b *UserBuilder) WithUserTag(key, value string) *UserBuilder {
	if b.user.Tags == nil {
		b.user.Tags = make(map[string]string)
	}
	b.user.Tags[key] = value
	return b
}

// WithMetadata adds metadata to the user.
func (b *UserBuilder) WithMetadata(key string, value interface{}) *UserBuilder {
	if b.user.Metadata == nil {
		b.user.Metadata = make(map[string]interface{})
	}
	b.user.Metadata[key] = value
	return b
}

// Build creates the TestUser instance.
// It performs final validation and returns the user or an error.
func (b *UserBuilder) Build() interface{} {
	if b.HasErrors() {
		return fmt.Errorf("cannot build user due to validation errors: %v", b.GetErrors())
	}

	// Perform final validation
	if b.IsValidationEnabled() {
		if b.user.Name == "" {
			return errors.New("user name is required")
		}
		if b.user.Email == "" {
			return errors.New("user email is required")
		}
	}

	// Create a copy to avoid mutation
	result := &TestUser{
		ID:       b.user.ID,
		Name:     b.user.Name,
		Email:    b.user.Email,
		Age:      b.user.Age,
		Active:   b.user.Active,
		Tags:     make(map[string]string),
		Metadata: make(map[string]interface{}),
	}

	// Deep copy tags
	for k, v := range b.user.Tags {
		result.Tags[k] = v
	}

	// Deep copy metadata
	for k, v := range b.user.Metadata {
		result.Metadata[k] = v
	}

	return result
}

// Reset clears the builder state for reuse.
func (b *UserBuilder) Reset() Builder {
	b.BaseBuilder.Reset()
	b.user = &TestUser{
		Tags:     make(map[string]string),
		Metadata: make(map[string]interface{}),
	}
	return b
}

// Clone creates a deep copy of the UserBuilder.
func (b *UserBuilder) Clone() Builder {
	clone := &UserBuilder{
		BaseBuilder: b.BaseBuilder.Clone().(*BaseBuilder),
		user: &TestUser{
			ID:       b.user.ID,
			Name:     b.user.Name,
			Email:    b.user.Email,
			Age:      b.user.Age,
			Active:   b.user.Active,
			Tags:     make(map[string]string),
			Metadata: make(map[string]interface{}),
		},
	}

	// Deep copy user tags
	for k, v := range b.user.Tags {
		clone.user.Tags[k] = v
	}

	// Deep copy user metadata
	for k, v := range b.user.Metadata {
		clone.user.Metadata[k] = v
	}

	return clone
}

// ApplyConfig implements ConfigurableBuilder interface.
func (b *UserBuilder) ApplyConfig(config *BuilderConfig) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}

	// Apply base configuration
	b.WithValidation(config.ValidationEnabled)

	// Apply tags
	for key, value := range config.Tags {
		b.WithTag(key, value)
	}

	// Apply default values specific to UserBuilder
	if defaults := config.DefaultValues; defaults != nil {
		if id, ok := defaults["id"].(int); ok {
			b.WithID(id)
		}
		if name, ok := defaults["name"].(string); ok {
			b.WithName(name)
		}
		if email, ok := defaults["email"].(string); ok {
			b.WithEmail(email)
		}
		if age, ok := defaults["age"].(int); ok {
			b.WithAge(age)
		}
		if active, ok := defaults["active"].(bool); ok {
			b.WithActive(active)
		}
	}

	return nil
}

// Factory function for UserBuilder
func createUserBuilder() Builder {
	return NewUserBuilder()
}

// Register UserBuilder in the default factory
func init() {
	RegisterBuilder("user", createUserBuilder)
}
