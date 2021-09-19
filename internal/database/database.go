package database

type Model struct {
	Key   string
	Value string
}
type Database interface {
	Upsert(req *Model) error
	FindByKey(key string) (*Model, error)
	DeleteByKey(key string) error
}
