//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
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

// BINDING INTERFACE
var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // Who need this SayHello will be sent SayHelloImpl
)

func InitializeHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

// wrong
// func InitializeHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

// STRUCT PROVIDER
var FooBarSet = wire.NewSet(NewFoo, NewBar)

func InitializeFooBar() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "*"))
	return nil
}

// BINDING VALUE

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializeFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializeReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// STRUCT FIELD PROVIDER

// => case without wire
// func InitializeConfiguration() *Configuration {
// 	application := NewApplication()
// 	configuration := application.Configuration
// 	return configuration
// }

func InitializeddConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"), // get value of Application.Configuration and make it as a new provider
	)

	return nil

}
