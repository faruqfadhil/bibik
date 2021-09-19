package repository

type CLIModel struct {
	Key   string
	Value string
}
type CLIRepository interface {
	Upsert(req *CLIModel) error
	FindByKey(key string) (*CLIModel, error)
	DeleteByKey(key string) error
}
