package main

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	if err := run(context.Background(), &StartOptions{
		address: "ws://127.0.0.1:8080",
		user:    "dajiang",
	}); err != nil {
		panic(err)
	}
}
