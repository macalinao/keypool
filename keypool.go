package keypool

import (
	"sync"
	"time"
)

// Key represents an API key
type Key struct {
	Value    string
	LastUsed time.Time
}

// Keypool stores keys
type Keypool struct {
	Keys  chan *Key
	Rate  time.Duration
	Mutex sync.Mutex
}

// New creates a new key
func New(keys []string, rate time.Duration) *Keypool {
	kc := make(chan *Key, len(keys))
	for _, key := range keys {
		kc <- &Key{
			Value: key,
		}
	}
	return &Keypool{
		Keys: kc,
		Rate: rate,
	}
}

// Fetch gets a new key
func (k *Keypool) Fetch() string {
	next := <-k.Keys
	dur := time.Since(next.LastUsed)
	if dur < k.Rate {
		time.Sleep(k.Rate - dur)
	}
	v := next.Value
	next.LastUsed = time.Now()
	k.Keys <- next
	return v
}
