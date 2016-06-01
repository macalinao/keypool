package keypool

import "time"

// Key represents an API key
type Key struct {
	Value    string
	lastUsed time.Time
	pool     *Keypool
}

// Return puts a key back into the pool
func (k *Key) Return() {
	k.lastUsed = time.Now()
	k.pool.keys <- k
}

// Keypool stores keys
type Keypool struct {
	keys chan *Key
	rate time.Duration
}

// New creates a new key
func New(keys []string, rate time.Duration) *Keypool {
	pool := &Keypool{
		rate: rate,
	}
	kc := make(chan *Key, len(keys))
	for _, key := range keys {
		kc <- &Key{
			Value: key,
			pool:  pool,
		}
	}
	pool.keys = kc
	return pool
}

// Fetch gets a new key
func (k *Keypool) Fetch() *Key {
	next := <-k.keys
	dur := time.Since(next.lastUsed)
	if dur < k.rate {
		time.Sleep(k.rate - dur)
	}
	return next
}
