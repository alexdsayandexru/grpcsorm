package security

import (
	"context"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_EmptySecurity(t *testing.T) {
	if ok, err := NewEmptySecurity().Auth(nil); !ok {
		t.Error(err)
	}
}

func Test_SecurityImpl(t *testing.T) {
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("key", "keyValue"))
	if ok, err := NewSecurity("key", "keyValue").Auth(&ctx); !ok {
		t.Error(err)
	}
	ctx2 := context.Background()
	if ok, err := NewSecurity("key", "keyValue").Auth(&ctx2); ok {
		t.Error(err)
	}

	ctx3 := metadata.NewIncomingContext(context.Background(), metadata.Pairs("key", "keyValue"))
	if ok, err := NewSecurity("key", "keyValue1").Auth(&ctx3); ok {
		t.Error(err)
	}
}
