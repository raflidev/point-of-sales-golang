//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)

	return nil, nil
}

func InitDatabaseRepo() *DatabaseRepository {
	wire.Build(NewDatabaseMySQL, NewDatabasePostgreSQL, NewDatabaseRepository)

	return nil
}

var fooSet = wire.NewSet(NewFooService, NewFooRepository)
var barSet = wire.NewSet(NewBarService, NewBarRepository)

func initFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}
