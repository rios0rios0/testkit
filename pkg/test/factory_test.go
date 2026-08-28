package testkit //nolint:testpackage // tests require access to unexported fields for thorough verification

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBuilderFactory_NewBuilderFactory(t *testing.T) {
	factory := NewBuilderFactory()

	if factory == nil {
		t.Fatal("NewBuilderFactory() returned nil")
	}

	if factory.builders == nil {
		t.Error("Expected builders map to be initialized")
	}

	if len(factory.builders) != 0 {
		t.Error("Expected builders map to be empty initially")
	}
}

func TestBuilderFactory_Register(t *testing.T) {
	factory := NewBuilderFactory()

	// Test successful registration
	createFunc := func() Builder { return NewBaseBuilder() }
	err := factory.Register("test", createFunc)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !factory.IsRegistered("test") {
		t.Error("Expected builder to be registered")
	}

	// Test empty name error
	err = factory.Register("", createFunc)
	if err == nil {
		t.Error("Expected error for empty name")
	}

	// Test nil function error
	err = factory.Register("test2", nil)
	if err == nil {
		t.Error("Expected error for nil function")
	}
}

func TestBuilderFactory_Create(t *testing.T) {
	factory := NewBuilderFactory()

	// Test creating non-existent builder
	_, err := factory.Create("nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent builder")
	}

	// Register a builder and test creation
	createFunc := func() Builder { return NewBaseBuilder() }
	factory.Register("test", createFunc)

	builder, err := factory.Create("test")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if builder == nil {
		t.Error("Expected builder instance, got nil")
	}

	// Verify it's the correct type
	_, ok := builder.(*BaseBuilder)
	if !ok {
		t.Error("Expected BaseBuilder instance")
	}
}

func TestBuilderFactory_GetRegisteredNames(t *testing.T) {
	factory := NewBuilderFactory()

	// Test empty factory
	names := factory.GetRegisteredNames()
	if len(names) != 0 {
		t.Error("Expected no registered names in empty factory")
	}

	// Register some builders
	createFunc := func() Builder { return NewBaseBuilder() }
	factory.Register("builder1", createFunc)
	factory.Register("builder2", createFunc)

	names = factory.GetRegisteredNames()
	if len(names) != 2 {
		t.Errorf("Expected 2 registered names, got %d", len(names))
	}

	// Check that both names are present (order doesn't matter)
	nameMap := make(map[string]bool)
	for _, name := range names {
		nameMap[name] = true
	}

	if !nameMap["builder1"] || !nameMap["builder2"] {
		t.Error("Expected both builder1 and builder2 to be in registered names")
	}
}

func TestDefaultFactory(t *testing.T) {
	// Test that default factory exists
	if DefaultFactory == nil {
		t.Error("Expected DefaultFactory to be initialized")
	}

	// Test global registration and creation functions
	createFunc := func() Builder { return NewBaseBuilder() }
	err := RegisterBuilder("global_test", createFunc)
	if err != nil {
		t.Errorf("Expected no error registering with global function, got %v", err)
	}

	builder, err := CreateBuilder("global_test")
	if err != nil {
		t.Errorf("Expected no error creating with global function, got %v", err)
	}

	if builder == nil {
		t.Error("Expected builder instance from global function")
	}
}

func TestBuilderConfig_NewBuilderConfig(t *testing.T) {
	config := NewBuilderConfig()

	if config == nil {
		t.Fatal("NewBuilderConfig() returned nil")
	}

	if !config.ValidationEnabled {
		t.Error("Expected validation to be enabled by default")
	}

	if config.Tags == nil {
		t.Error("Expected tags to be initialized")
	}

	if config.DefaultValues == nil {
		t.Error("Expected default values to be initialized")
	}
}

func TestBuilderConfig_With_Methods(t *testing.T) {
	config := NewBuilderConfig()

	// Test WithValidation
	result := config.WithValidation(false)
	if result != config {
		t.Error("WithValidation should return the same config instance")
	}
	if config.ValidationEnabled {
		t.Error("Expected validation to be disabled")
	}

	// Test WithTag
	result = config.WithTag("env", "test")
	if result != config {
		t.Error("WithTag should return the same config instance")
	}
	if config.Tags["env"] != "test" {
		t.Error("Expected tag to be set")
	}

	// Test WithDefault
	result = config.WithDefault("name", "test_name")
	if result != config {
		t.Error("WithDefault should return the same config instance")
	}
	if config.DefaultValues["name"] != "test_name" {
		t.Error("Expected default value to be set")
	}
}

func TestBuilderConfig_ApplyTo(t *testing.T) {
	config := NewBuilderConfig()
	config.WithValidation(false)
	config.WithTag("env", "test")

	// Test with nil builder
	err := config.ApplyTo(nil)
	if err == nil {
		t.Error("Expected error when applying to nil builder")
	}

	// Test with BaseBuilder
	builder := NewBaseBuilder()
	err = config.ApplyTo(builder)
	if err != nil {
		t.Errorf("Expected no error applying to BaseBuilder, got %v", err)
	}

	// Verify configuration was applied
	if builder.IsValidationEnabled() {
		t.Error("Expected validation to be disabled after applying config")
	}

	if builder.GetTag("env") != "test" {
		t.Error("Expected tag to be applied to builder")
	}
}

func TestUserBuilderRegistration(t *testing.T) {
	// Test that UserBuilder is registered in the default factory
	if !DefaultFactory.IsRegistered("user") {
		t.Error("Expected UserBuilder to be registered as 'user'")
	}

	// Test creating a UserBuilder from the factory
	builder, err := CreateBuilder("user")
	if err != nil {
		t.Errorf("Expected no error creating user builder, got %v", err)
	}

	userBuilder, ok := builder.(*UserBuilder)
	if !ok {
		t.Error("Expected UserBuilder instance")
	}

	if userBuilder == nil {
		t.Error("Expected non-nil UserBuilder")
	}
}

func TestBuilderFactory_ConcurrentAccess(t *testing.T) {
	// given
	const goroutines = 50

	factory := NewBuilderFactory()
	createFunc := func() Builder { return NewBaseBuilder() }

	var waitGroup sync.WaitGroup

	start := make(chan struct{})

	// when
	// Every goroutine blocks on the same channel so the registrations and the reads land in
	// the same instant rather than being serialised by their own start-up cost. Without the
	// mutex in BuilderFactory this is a "fatal error: concurrent map read and map write",
	// which no recover can catch, and `go test -race` reports it even when the interleaving
	// happens not to collide.
	for i := range goroutines {
		waitGroup.Add(1)

		go func(index int) {
			defer waitGroup.Done()
			<-start

			if err := factory.Register(fmt.Sprintf("builder-%d", index), createFunc); err != nil {
				t.Errorf("Register(builder-%d) returned %v", index, err)
			}

			factory.IsRegistered("builder-0")
			factory.GetRegisteredNames()
		}(i)
	}

	close(start)
	waitGroup.Wait()

	// then
	if got := len(factory.GetRegisteredNames()); got != goroutines {
		t.Errorf("Expected %d registered builders, got %d", goroutines, got)
	}

	for i := range goroutines {
		name := fmt.Sprintf("builder-%d", i)
		if !factory.IsRegistered(name) {
			t.Errorf("Expected %s to be registered", name)
		}
	}
}

func TestBuilderFactory_CreateAllowsReentrantRegistration(t *testing.T) {
	// given
	// A creation function is caller-supplied and may register further builders -- a composite
	// builder registering its parts on first use is the ordinary case. Create must therefore
	// not hold the lock while it runs, so this asserts the absence of a deadlock rather than
	// a value. It hangs forever on a regression, so the wait is bounded and reported.
	factory := NewBuilderFactory()

	err := factory.Register("composite", func() Builder {
		if regErr := factory.Register("part", func() Builder { return NewBaseBuilder() }); regErr != nil {
			t.Errorf("nested Register returned %v", regErr)
		}

		return NewBaseBuilder()
	})
	if err != nil {
		t.Fatalf("Register(composite) returned %v", err)
	}

	done := make(chan struct{})

	// when
	go func() {
		defer close(done)

		if _, createErr := factory.Create("composite"); createErr != nil {
			t.Errorf("Create(composite) returned %v", createErr)
		}
	}()

	// then
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("Create deadlocked while its creation function registered another builder")
	}

	if !factory.IsRegistered("part") {
		t.Error("Expected the nested registration to have taken effect")
	}
}
