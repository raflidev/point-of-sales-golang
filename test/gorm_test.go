package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "user=postgres dbname=point-of-sales-golang sslmode=disable password=postgres host=localhost"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func truncateProductGorm(db *gorm.DB) {
	db.Exec("TRUNCATE product")
}

func TestDB(t *testing.T) {
	assert.NotNil(t, db)
}

type Product struct {
	Id          uuid.UUID
	Kode_produk string
	Nama_produk string
	Merk        string
	Harga_beli  int
	Harga_jual  int
	Stok        int
}

func TestExecuteSQL(t *testing.T) {
	truncateProductGorm(db)
	var id string
	var product Product
	err := db.Raw("insert into product (kode_produk, nama_produk, merk, harga_beli, harga_jual, stok) values (?, ?, ?, ?, ?, ?) RETURNING id", "P-01", "Momogi", "Momogi", "1000", "1500", "10").Scan(&id).Error

	fmt.Println(id)
	assert.Nil(t, err)

	err = db.Raw("select * from product where id = ?", id).Scan(&product).Error
	fmt.Println(product)
	assert.Nil(t, err)
	assert.Equal(t, "Momogi", product.Nama_produk)
}

func TestSqlRow(t *testing.T) {
	truncateProductGorm(db)
	var id string

	err := db.Raw("insert into product (kode_produk, nama_produk, merk, harga_beli, harga_jual, stok) values (?, ?, ?, ?, ?, ?) RETURNING id", "P-01", "Momogi", "Momogi", "1000", "1500", "10").Scan(&id).Error

	fmt.Println(id)
	assert.Nil(t, err)

	rows, err := db.Raw("select id, kode_produk, nama_produk, merk, harga_beli, harga_jual, stok from product").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var Id uuid.UUID
		var Kode_produk string
		var Nama_produk string
		var Merk string
		var Harga_beli int
		var Harga_jual int
		var Stok int

		fmt.Println(rows)
		err := rows.Scan(&Id, &Kode_produk, &Nama_produk, &Merk, &Harga_beli, &Harga_jual, &Stok)
		assert.Nil(t, err)

		products = append(products, Product{
			Id:          Id,
			Kode_produk: Kode_produk,
			Nama_produk: Nama_produk,
			Merk:        Merk,
			Harga_beli:  Harga_beli,
			Harga_jual:  Harga_jual,
			Stok:        Stok,
		})
	}

	assert.Equal(t, 1, len(products))

}
