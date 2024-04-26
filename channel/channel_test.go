package channel

import (
	"net/http"
	"testing"

	"github.com/ksysoev/wasabi/mocks"
)

func TestNewChannel(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	if channel.path != path {
		t.Errorf("Unexpected path: got %q, expected %q", channel.path, path)
	}

	if channel.disptacher != dispatcher {
		t.Errorf("Unexpected dispatcher: got %v, expected %v", channel.disptacher, dispatcher)
	}

	if len(channel.middlewares) != 0 {
		t.Errorf("Unexpected number of middlewares: got %d, expected %d", len(channel.middlewares), 0)
	}
}
func TestChannel_Path(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	if channel.Path() != path {
		t.Errorf("Unexpected path: got %q, expected %q", channel.Path(), path)
	}
}
func TestChannel_Handler(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	// Call the Handler method
	handler := channel.Handler()

	if handler == nil {
		t.Errorf("Unexpected nil handler")
	}
}

func TestChannel_Use(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Custom middleware logic
		})
	}

	channel.Use(middleware)

	if len(channel.middlewares) != 1 {
		t.Errorf("Unexpected number of middlewares: got %d, expected %d", len(channel.middlewares), 1)
	}
}

func TestChannel_wrapMiddleware(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	// Create a mock handler
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mock handler logic
	})

	// Create mock middlewares
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Mock middleware 1 logic
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Mock middleware 2 logic
		})
	}

	// Add the mock middlewares to the channel
	channel.middlewares = append(channel.middlewares, middleware1, middleware2)

	// Call the wrapMiddleware method
	wrappedHandler := channel.wrapMiddleware(mockHandler)

	// Assert that the wrappedHandler is the result of applying the middlewares to the mockHandler
	if wrappedHandler == nil {
		t.Errorf("Unexpected nil wrappedHandler")
	}
}

func TestChannel_WithOriginPatterns(t *testing.T) {
	path := "/test/path"
	dispatcher := mocks.NewMockDispatcher(t)

	channel := NewChannel(path, dispatcher, NewConnectionRegistry())

	if len(channel.config.originPatterns) != 1 {
		t.Errorf("Unexpected number of origin patterns: got %d, expected %d", len(channel.config.originPatterns), 1)
	}

	if channel.config.originPatterns[0] != "*" {
		t.Errorf("Unexpected to get default origin pattern: got %s, expected %s", channel.config.originPatterns[0], "*")
	}

	channel = NewChannel(path, dispatcher, NewConnectionRegistry(), WithOriginPatterns("test", "test2"))

	if len(channel.config.originPatterns) != 2 {
		t.Errorf("Unexpected number of origin patterns: got %d, expected %d", len(channel.config.originPatterns), 1)
	}

	if channel.config.originPatterns[0] != "test" {
		t.Errorf("Unexpected to get default origin pattern: got %s, expected %s", channel.config.originPatterns[0], "test")
	}

	if channel.config.originPatterns[1] != "test2" {
		t.Errorf("Unexpected to get default origin pattern: got %s, expected %s", channel.config.originPatterns[1], "test2")
	}
}
