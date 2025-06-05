package helloapi

import (
	"context"
	hello "hello/gen/hello"
	"fmt"
	"goa.design/clue/log"
)

// hello service example implementation.
// The example methods log the requests and return zero values.
type hellosrvc struct{}

// NewHello returns the hello service implementation.
func NewHello() hello.Service {
	return &hellosrvc{}
}

// SayHello implements sayHello.
func (s *hellosrvc) SayHello(ctx context.Context, name string) (string, error) {
	log.Printf(ctx, "hello.sayHello")
	return fmt.Sprintf("こんにちは、%sさん!", name), nil
}
