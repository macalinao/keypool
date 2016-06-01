# keypool

Key pooling library to circumvent rate limits.

## Install

```bash
$ go get github.com/simplyianm/keypool
```

## Usage

```go
keys := []string{"a", "b", "c"}
// Create a key pool with the given keys
// with no less than 10ms gap between using a key
pool := keypool.New(keys, 10*time.Millisecond)

key := pool.Fetch()
key.Value // "a"
pool.Fetch()
pool.Fetch()

key.Return() // Returns "a" back to the key pool
pool.Fetch() // returns "a" after at least 10ms has elapsed since "a" was first retrieved
```

## License

MIT
