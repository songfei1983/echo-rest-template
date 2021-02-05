package cache

type Cache interface {
	GET(k string) (v interface{}, err error)
	PUT(k string, v interface{}) error
	DEL(k string) error
}
