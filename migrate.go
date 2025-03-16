package main

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection2() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=point-of-sales-golang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := OpenConnection2()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{{
		// create `users` table
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type user struct {
				ID   uuid.UUID `gorm:"type:uuid;primaryKey;uniqueIndex"`
				Name string
			}
			return tx.Migrator().CreateTable(&user{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("users")
		},
	}, {
		// add `age` column to `users` table
		ID: "201608301415",
		Migrate: func(tx *gorm.DB) error {
			// when table already exists, define only columns that are about to change
			type user struct {
				Age int
			}
			return tx.Migrator().AddColumn(&user{}, "Age")
		},
		Rollback: func(tx *gorm.DB) error {
			type user struct {
				Age int
			}
			return db.Migrator().DropColumn(&user{}, "Age")
		},
	}, {
		// create `organizations` table where users belong to
		ID: "201608301430",
		Migrate: func(tx *gorm.DB) error {
			type organization struct {
				ID      uuid.UUID `gorm:"type:uuid;primaryKey;uniqueIndex"`
				Name    string
				Address string
			}
			if err := tx.Migrator().CreateTable(&organization{}); err != nil {
				return err
			}
			type user struct {
				OrganizationID uuid.UUID `gorm:"type:uuid"`
			}
			return tx.Migrator().AddColumn(&user{}, "OrganizationID")
		},
		Rollback: func(tx *gorm.DB) error {
			type user struct {
				OrganizationID uuid.UUID `gorm:"type:uuid"`
			}
			if err := db.Migrator().DropColumn(&user{}, "OrganizationID"); err != nil {
				return err
			}
			return tx.Migrator().DropTable("organizations")
		},
	}})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration did run successfully")
}
