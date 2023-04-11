package internal

import "context"

//go:generate mockery --name Repository
type Repository interface {
	Save(context.Context, *Port) error
}
