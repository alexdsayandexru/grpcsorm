package security

import "context"

type EmptySecurity struct{}

func (p *EmptySecurity) Auth(*context.Context) (bool, error) {
	return true, nil
}

func NewEmptySecurity() ISecurity {
	return &EmptySecurity{}
}
