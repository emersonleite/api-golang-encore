package hello

import (
	"context"
	"fmt"
)

// World responds with a tailored message
//
//encore:api public path=/hello/:name
func World(ctx context.Context, name string) (*Response, error) {
	msg := fmt.Sprintf("Hello, %s", name)

	return &Response{Message: msg}, nil
}

type Response struct {
	Message string
}
