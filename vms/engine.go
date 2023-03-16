package vms

import (
	"context"
	"io"
)

// Engine is an interface to the game engine.
type Engine interface {
	TimeSource

	// Console returns script console.
	Console(error bool) Writer
}

// Writer is an interface for writing text messages.
type Writer interface {
	io.Writer
	io.StringWriter
}

type engineKey struct{}

// WithEngine stores Engine reference in the context.
func WithEngine(ctx context.Context, g Engine) context.Context {
	return context.WithValue(ctx, engineKey{}, g)
}

// GetEngine loads Game from the context.
func GetEngine(ctx context.Context) Engine {
	return ctx.Value(engineKey{}).(Engine)
}
