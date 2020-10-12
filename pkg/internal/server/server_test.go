package server

import (
	"context"
	"testing"
)

func TestCreateServer(t *testing.T) {
	ctx := context.Background()
	connection := NewAuthServer(ctx, []string{"localhost:6379"}, "1234")
	if connection == nil {
		t.Error("can't create server")
	}
}
