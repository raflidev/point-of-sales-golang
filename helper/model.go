package helper

import (
	entityProduct "golang-point-of-sales-system/modules/products/domain/entity"
	responseProduct "golang-point-of-sales-system/modules/products/dto/response"

	entitySupplier "golang-point-of-sales-system/modules/suppliers/domain/entity"
	responseSupplier "golang-point-of-sales-system/modules/suppliers/dto/response"
)

func ToProductResponse(category entityProduct.Product) responseProduct.ProductResponse {
	return responseProduct.ProductResponse{
		Id:          category.Id,
		Kode_produk: category.Kode_produk,
		Nama_produk: category.Nama_produk,
		Merk:        category.Merk,
		Harga_beli:  category.Harga_beli,
		Harga_jual:  category.Harga_jual,
		Stok:        category.Stok,
	}
}

func ToProductResponses(products []entityProduct.Product) []responseProduct.ProductResponse {
	var productResponses []responseProduct.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}

func ToSupplierResponse(supplier entitySupplier.Supplier) responseSupplier.SupplierResponse {
	return responseSupplier.SupplierResponse{
		Id:      supplier.Id,
		Nama:    supplier.Nama,
		Alamat:  supplier.Alamat,
		Telepon: supplier.Telepon,
	}
}

func ToSupplierResponses(suppliers []entitySupplier.Supplier) []responseSupplier.SupplierResponse {
	var supplierResponses []responseSupplier.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, ToSupplierResponse(supplier))
	}

	return supplierResponses
}
