//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil // will be replace automatically by google wire
}
