package database

type Model struct {
	Key   string
	Value string
}
type Database interface {
	Save(req *Model) error
	FindByKey(key string) (*Model, error)
}
