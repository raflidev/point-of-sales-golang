package helper

import (
	"golang-point-of-sales-system/modules/products/domain/entity"
	"golang-point-of-sales-system/modules/products/dto/response"
)

func ToProductResponse(category entity.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:          category.Id,
		Kode_produk: category.Kode_produk,
		Nama_produk: category.Nama_produk,
		Merk:        category.Merk,
		Harga_beli:  category.Harga_beli,
		Harga_jual:  category.Harga_jual,
		Stok:        category.Stok,
	}
}

func ToProductResponses(products []entity.Product) []response.ProductResponse {
	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
