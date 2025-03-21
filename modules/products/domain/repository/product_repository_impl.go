package repository

import (
	"context"
	"errors"
	"golang-point-of-sales-system/helper"
	domain "golang-point-of-sales-system/modules/products/domain/entity"

	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sqlx.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(kode_produk, nama_produk, merk, harga_beli, harga_jual, stok) values(?)"
	result, err := tx.ExecContext(ctx, SQL, product.Kode_produk, product.Nama_produk, product.Merk, product.Harga_beli, product.Harga_jual, product.Stok)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sqlx.Tx, product domain.Product) domain.Product {
	SQL := "update product set kode_produk=?, nama_produk=?, merk=?, harga_beli=?, harga_jual=?, stok=? where id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Kode_produk, product.Nama_produk, product.Merk, product.Harga_beli, product.Harga_jual, product.Stok, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sqlx.Tx, product domain.Product) {
	SQL := "delete from product where id=?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sqlx.Tx, productId int) (domain.Product, error) {
	SQL := "select id, kode_produk, nama_produk, merk, harga_beli, harga_jual, stok from product where id=?"
	rows, err := tx.QueryxContext(ctx, SQL, productId)
	helper.PanicIfError(err)

	product := domain.Product{}
	if rows.Next() {
		err = rows.Scan(&product.Id, &product.Kode_produk, &product.Nama_produk, &product.Merk, &product.Harga_beli, &product.Harga_jual, &product.Stok)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product not found")
	}

}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Product {
	SQL := "select id, kode_produk, nama_produk, merk, harga_beli, harga_jual, stok from products"
	rows, err := tx.QueryxContext(ctx, SQL)
	helper.PanicIfError(err)

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err = rows.Scan(&product.Id, &product.Kode_produk, &product.Nama_produk, &product.Merk, &product.Harga_beli, &product.Harga_jual, &product.Stok)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}
