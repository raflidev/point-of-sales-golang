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

	entityMember "golang-point-of-sales-system/modules/members/domain/entity"
	responseMember "golang-point-of-sales-system/modules/members/dto/response"

	entityPembelian "golang-point-of-sales-system/modules/pembelian/domain/entity"
	responsePembelian "golang-point-of-sales-system/modules/pembelian/dto/response"
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

func ToMemberResponse(member entityMember.Member) responseMember.MemberResponse {
	return responseMember.MemberResponse{
		Id:          member.Id,
		Kode_member: member.Kode_member,
		Nama:        member.Nama,
		Telepon:     member.Telepon,
		Alamat:      member.Alamat,
		Keterangan:  member.Keterangan,
	}
}

func ToMemberResponses(members []entityMember.Member) []responseMember.MemberResponse {
	var memberResponses []responseMember.MemberResponse
	for _, member := range members {
		memberResponses = append(memberResponses, ToMemberResponse(member))
	}

	return memberResponses
}

func ToPembelianResponse(pembelian entityPembelian.Pembelian) responsePembelian.PembelianResponse {
	return responsePembelian.PembelianResponse{
		Id:          pembelian.Id,
		Supplier_id: pembelian.Supplier_id,
		Total_item:  pembelian.Total_item,
		Total_harga: pembelian.Total_harga,
		Diskon:      pembelian.Diskon,
		Bayar:       pembelian.Bayar,
	}
}

func ToPembelianResponses(Pembelians []entityPembelian.Pembelian) []responsePembelian.PembelianResponse {
	var pembelianResponses []responsePembelian.PembelianResponse
	for _, pembelian := range Pembelians {
		pembelianResponses = append(pembelianResponses, ToPembelianResponse(pembelian))
	}

	return pembelianResponses
}
