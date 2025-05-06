package helper

import (
	entityProduct "golang-point-of-sales-system/modules/products/domain/entity"
	responseProduct "golang-point-of-sales-system/modules/products/dto/response"
	"golang-point-of-sales-system/modules/users/domain/entity"
	"golang-point-of-sales-system/modules/users/dto/response"

	entitySupplier "golang-point-of-sales-system/modules/suppliers/domain/entity"
	responseSupplier "golang-point-of-sales-system/modules/suppliers/dto/response"

	entityCategory "golang-point-of-sales-system/modules/categories/domain/entity"
	responseCategory "golang-point-of-sales-system/modules/categories/dto/response"
)

func ToProductResponse(product entityProduct.Product) responseProduct.ProductResponse {
	return responseProduct.ProductResponse{
		Id:          product.Id,
		Kode_produk: product.Kode_produk,
		Nama_produk: product.Nama_produk,
		Merk:        product.Merk,
		Harga_beli:  product.Harga_beli,
		Harga_jual:  product.Harga_jual,
		Stok:        product.Stok,
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

func ToCategoryResponse(category entityCategory.Category) responseCategory.CategoryResponse {
	return responseCategory.CategoryResponse{
		Id:            category.Id,
		Nama_kategori: category.Nama_kategori,
	}
}

func ToCategoryResponses(categories []entityCategory.Category) []responseCategory.CategoryResponse {
	var categoryResponses []responseCategory.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToUserResponse(user entity.User) response.UserResponse {
	return response.UserResponse{
		Id:    user.Id,
		Nama:  user.Nama,
		Email: user.Email,
		Foto:  user.Foto,
		Role:  user.Role,
	}
}
