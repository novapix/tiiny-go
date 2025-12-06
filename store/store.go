package store

// URLStore defines the interface for a URL storage backend.
type URLStore interface {
	Save(key, url string) error
	Get(key string) (string, error)
}
