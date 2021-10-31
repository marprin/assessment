package cache

type (
	Cache interface {
		Get(path string) ([]string, error)
		Set(path string, value string) error
	}
)
