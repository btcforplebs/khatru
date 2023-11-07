package khatru

import (
	"context"
	"hash/maphash"
	"regexp"
	"unsafe"

	"github.com/nbd-wtf/go-nostr"
)

const (
	AUTH_CONTEXT_KEY = iota
	WS_KEY           = iota
)

var nip20prefixmatcher = regexp.MustCompile(`^\w+: `)

func GetConnection(ctx context.Context) *WebSocket {
	return ctx.Value(WS_KEY).(*WebSocket)
}

func GetAuthed(ctx context.Context) string {
	authedPubkey := ctx.Value(AUTH_CONTEXT_KEY)
	if authedPubkey == nil {
		return ""
	}
	return authedPubkey.(string)
}

func pointerHasher[V any](_ maphash.Seed, k *V) uint64 {
	return uint64(uintptr(unsafe.Pointer(k)))
}

func isOlder(previous, next *nostr.Event) bool {
	return previous.CreatedAt < next.CreatedAt ||
		(previous.CreatedAt == next.CreatedAt && previous.ID > next.ID)
}
