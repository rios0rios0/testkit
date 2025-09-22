package testkit

import (
	"fmt"
	"reflect"
)

// BuilderFactory provides a way to register and create different types of builders.
type BuilderFactory struct {
	builders map[string]func() Builder
}

// NewBuilderFactory creates a new BuilderFactory instance.
func NewBuilderFactory() *BuilderFactory {
	return &BuilderFactory{
		builders: make(map[string]func() Builder),
	}
}

// Register registers a builder creation function with a given name.
func (f *BuilderFactory) Register(name string, createFunc func() Builder) error {
	if name == "" {
		return fmt.Errorf("builder name cannot be empty")
	}
	if createFunc == nil {
		return fmt.Errorf("builder creation function cannot be nil")
	}
	f.builders[name] = createFunc
	return nil
}

// Create creates a new builder instance by name.
func (f *BuilderFactory) Create(name string) (Builder, error) {
	createFunc, exists := f.builders[name]
	if !exists {
		return nil, fmt.Errorf("builder '%s' not registered", name)
	}
	return createFunc(), nil
}

// IsRegistered checks if a builder is registered with the given name.
func (f *BuilderFactory) IsRegistered(name string) bool {
	_, exists := f.builders[name]
	return exists
}

// GetRegisteredNames returns all registered builder names.
func (f *BuilderFactory) GetRegisteredNames() []string {
	names := make([]string, 0, len(f.builders))
	for name := range f.builders {
		names = append(names, name)
	}
	return names
}

// DefaultFactory is a global factory instance for convenience.
var DefaultFactory = NewBuilderFactory()

// RegisterBuilder registers a builder in the default factory.
func RegisterBuilder(name string, createFunc func() Builder) error {
	return DefaultFactory.Register(name, createFunc)
}

// CreateBuilder creates a builder from the default factory.
func CreateBuilder(name string) (Builder, error) {
	return DefaultFactory.Create(name)
}

// BuilderConfig provides configuration options for builders.
type BuilderConfig struct {
	ValidationEnabled bool
	Tags              map[string]string
	DefaultValues     map[string]interface{}
}

// NewBuilderConfig creates a new BuilderConfig with default settings.
func NewBuilderConfig() *BuilderConfig {
	return &BuilderConfig{
		ValidationEnabled: true,
		Tags:              make(map[string]string),
		DefaultValues:     make(map[string]interface{}),
	}
}

// WithValidation sets the validation enabled flag.
func (c *BuilderConfig) WithValidation(enabled bool) *BuilderConfig {
	c.ValidationEnabled = enabled
	return c
}

// WithTag adds a tag to the configuration.
func (c *BuilderConfig) WithTag(key, value string) *BuilderConfig {
	if c.Tags == nil {
		c.Tags = make(map[string]string)
	}
	c.Tags[key] = value
	return c
}

// WithDefault sets a default value for a field.
func (c *BuilderConfig) WithDefault(key string, value interface{}) *BuilderConfig {
	if c.DefaultValues == nil {
		c.DefaultValues = make(map[string]interface{})
	}
	c.DefaultValues[key] = value
	return c
}

// ApplyTo applies the configuration to a builder.
func (c *BuilderConfig) ApplyTo(builder Builder) error {
	if builder == nil {
		return fmt.Errorf("builder cannot be nil")
	}

	// Use reflection to check if the builder has BaseBuilder methods
	builderValue := reflect.ValueOf(builder)

	// Check if builder has WithValidation method
	if method := builderValue.MethodByName("WithValidation"); method.IsValid() {
		method.Call([]reflect.Value{reflect.ValueOf(c.ValidationEnabled)})
	}

	// Apply tags if the builder supports them
	if c.Tags != nil {
		for key, value := range c.Tags {
			if method := builderValue.MethodByName("WithTag"); method.IsValid() {
				method.Call([]reflect.Value{reflect.ValueOf(key), reflect.ValueOf(value)})
			}
		}
	}

	// For more complex default value application, builders should implement
	// a ConfigurableBuilder interface if they need this functionality
	if configurableBuilder, ok := builder.(ConfigurableBuilder); ok {
		return configurableBuilder.ApplyConfig(c)
	}

	return nil
}

// ConfigurableBuilder interface for builders that can accept configuration.
type ConfigurableBuilder interface {
	Builder
	ApplyConfig(config *BuilderConfig) error
}
