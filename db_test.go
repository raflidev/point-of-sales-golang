package main

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
	"testing"
	"github.com/stretchr/testify/assert"
)


func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=point-of-sales-golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T){
	assert.NotNil(t, db)
}