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

// World responds with a tailored message
//
//encore:api public path=/hello2/:name2
func World2(ctx context.Context, name2 string) (*Response, error) {
	msg := fmt.Sprintf("Hello, %s", name2)

	return &Response{Message: msg}, nil
}
