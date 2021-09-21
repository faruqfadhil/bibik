package entity

import "fmt"

type Command struct {
	Key     string
	Value   string
	Options *Options
}

type Options struct {
	Dir string
}

func (c *Command) SetDirKey() string {
	return fmt.Sprintf("dir-%s", c.Key)
}
