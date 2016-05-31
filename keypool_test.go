package keypool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeypoolNoRateLimit(t *testing.T) {
	keys := []string{"a", "b", "c", "d"}
	k := New(keys, 0*time.Millisecond)
	assert.Equal(t, k.Fetch(), "a")
	assert.Equal(t, k.Fetch(), "b")
	assert.Equal(t, k.Fetch(), "c")
	assert.Equal(t, k.Fetch(), "d")
	assert.Equal(t, k.Fetch(), "a")
}

func TestKeypoolWithRateLimit(t *testing.T) {
	keys := []string{"a", "b", "c", "d"}
	k := New(keys, 10*time.Millisecond)
	start := time.Now()
	assert.Equal(t, k.Fetch(), "a")
	assert.Equal(t, k.Fetch(), "b")
	assert.Equal(t, k.Fetch(), "c")
	assert.Equal(t, k.Fetch(), "d")
	assert.True(t, time.Since(start) < 1*time.Millisecond, "No sleep if no dupes")
	assert.Equal(t, k.Fetch(), "a")
	assert.False(t, time.Since(start) < 1*time.Millisecond, "Sleep if duration")
}
