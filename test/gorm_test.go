package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dsn := "user=postgres dbname=point-of-sales-golang-test sslmode=disable password=postgres host=localhost"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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
	Id          uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Kode_produk string    `gorm:"column:kode_produk;type:varchar(50);unique;not null"`
	Nama_produk string    `gorm:"column:nama_produk;type:varchar(50);not null"`
	Merk        string    `gorm:"column:merk;type:varchar(50);not null"`
	Harga_beli  int       `gorm:"column:harga_beli;type:int;not null"`
	Harga_jual  int       `gorm:"column:harga_jual;type:int;not null"`
	Stok        int       `gorm:"column:stok;type:int;not null"`
	CreateAt    time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt    time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductLog struct {
	ID        uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	ProductID uuid.UUID `gorm:"column:product_id;type:uuid;not null"`
	Action    string    `gorm:"column:action;type:varchar(50);not null"`
	CreateAt  time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt  time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
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
		var CreateAt time.Time
		var UpdateAt time.Time

		fmt.Println(rows)
		err := rows.Scan(&Id, &Kode_produk, &Nama_produk, &Merk, &Harga_beli, &Harga_jual, &Stok, &CreateAt, &UpdateAt)
		assert.Nil(t, err)

		products = append(products, Product{
			Id:          Id,
			Kode_produk: Kode_produk,
			Nama_produk: Nama_produk,
			Merk:        Merk,
			Harga_beli:  Harga_beli,
			Harga_jual:  Harga_jual,
			Stok:        Stok,
			CreateAt:    CreateAt,
			UpdateAt:    UpdateAt,
		})
	}

	assert.Equal(t, 1, len(products))

}

func TestCreateUser(t *testing.T) {
	product := Product{
		Kode_produk: "P-01",
		Nama_produk: "Momogi",
		Merk:        "Momogi",
		Harga_beli:  1000,
		Harga_jual:  1500,
		Stok:        10,
	}

	response := db.Create(&product)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestUpdateUser(t *testing.T) {
	product := Product{}
	err := db.First(&product, "id = ?", "0853420c-e23e-439a-a956-3df32204e8b6").Error
	assert.Nil(t, err)

	product.Nama_produk = "Momogi Baru"
	product.Harga_beli = 3000

	err = db.Save(&product).Error
	assert.Nil(t, err)

	result := db.Where("id = ?", "0853420c-e23e-439a-a956-3df32204e8b6").Updates(Product{
		Nama_produk: "Momogi Baru",
		Harga_beli:  3001,
	})
	assert.Nil(t, result.Error)
}
