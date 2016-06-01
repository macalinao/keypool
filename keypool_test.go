package keypool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeypoolNoRateLimit(t *testing.T) {
	keys := []string{"a", "b", "c", "d"}
	k := New(keys, 0*time.Millisecond)
	key := k.Fetch()
	assert.Equal(t, key.Value, "a")
	assert.Equal(t, k.Fetch().Value, "b")
	assert.Equal(t, k.Fetch().Value, "c")
	assert.Equal(t, k.Fetch().Value, "d")
	key.Return()
	assert.Equal(t, k.Fetch().Value, "a")
}

func TestKeypoolWithRateLimit(t *testing.T) {
	keys := []string{"a", "b", "c", "d"}
	k := New(keys, 10*time.Millisecond)
	start := time.Now()
	key := k.Fetch()
	assert.Equal(t, key.Value, "a")
	assert.Equal(t, k.Fetch().Value, "b")
	assert.Equal(t, k.Fetch().Value, "c")
	assert.Equal(t, k.Fetch().Value, "d")
	assert.True(t, time.Since(start) < 1*time.Millisecond, "No sleep if no dupes")
	key.Return()
	assert.Equal(t, k.Fetch().Value, "a")
	assert.False(t, time.Since(start) < 1*time.Millisecond, "Sleep if duration")
}
