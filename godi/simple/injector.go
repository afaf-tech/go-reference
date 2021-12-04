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

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository,
	)

	return nil
}

// advantage: just to simplify when creating provider and its injectors
var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}
