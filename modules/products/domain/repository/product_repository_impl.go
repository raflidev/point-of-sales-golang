package repository

import (
	"context"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/domain/entity"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product entity.Product) entity.Product {

	tx := repository.DB.MustBegin()
	defer tx.Commit()

	SQL := "insert into product (kode_produk, nama_produk, merk, harga_beli, harga_jual, stok) values($1, $2, $3, $4, $5, $6)"
	result := tx.MustExec(SQL, product.Kode_produk, product.Nama_produk, product.Merk, product.Harga_beli, product.Harga_jual, product.Stok)
	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		log.Println(err)
	}

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product entity.Product) entity.Product {
	tx := repository.DB.MustBegin()
	defer tx.Commit()

	SQL := "update product set kode_produk=$2, nama_produk=$3, merk=$4, harga_beli=$5, harga_jual=$6, stok=$7 where id=$1"
	result := tx.MustExec(SQL, product.Id, product.Kode_produk, product.Nama_produk, product.Merk, product.Harga_beli, product.Harga_jual, product.Stok)
	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		log.Println(err)
	}

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product entity.Product) {
	tx := repository.DB.MustBegin()
	defer tx.Commit()

	SQL := "delete from product where id=$1"
	result := tx.MustExec(SQL, product.Id)
	// helper.PanicIfError(result.Error())
	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		log.Println(err)
	}
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId uuid.UUID) (entity.Product, error) {
	tx := repository.DB.MustBegin()
	defer tx.Commit()
	SQL := "select id, kode_produk, nama_produk, merk, harga_beli, harga_jual, stok from product where id=$1"
	result := tx.MustExec(SQL, productId)
	n, err := result.RowsAffected()
	if err != nil || n == 0 {
		log.Println(err)
	}

	product := entity.Product{}
	err = tx.Get(&product, SQL, productId)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []entity.Product {
	tx := repository.DB.MustBegin()
	defer tx.Commit()

	SQL := "select id, kode_produk, nama_produk, merk, harga_beli, harga_jual, stok from product"
	rows, err := tx.QueryxContext(ctx, SQL)
	helper.PanicIfError(err)

	var products []entity.Product
	for rows.Next() {
		product := entity.Product{}
		err = rows.Scan(&product.Id, &product.Kode_produk, &product.Nama_produk, &product.Merk, &product.Harga_beli, &product.Harga_jual, &product.Stok)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}
