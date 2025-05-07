package encryption

type Key struct {
	Value     []byte
	ExpiresAt int64 // Unix timestamp
}

type KeyProvider interface {
	FetchKey() (*Key, error)
}

type CachedKey struct {
	Value     []byte
	ExpiresAt int64 // Unix timestamp
}

type KeyCache interface {
	GetKey() (*CachedKey, bool, error)
	SetKey(*Key) error
}
