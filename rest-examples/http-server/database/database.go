package database

import (
	"github.com/bygui86/go-testing/rest-examples/http-server/logging"
)

func New() InMemoryDb {
	logging.Log.Info("Create new in-memory DB")

	return &DefaultInMemoryDb{
		products: make(map[string]*Product, 1000),
	}
}
