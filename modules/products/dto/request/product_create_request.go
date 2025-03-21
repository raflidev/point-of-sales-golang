package request

type ProductCreateRequest struct {
	Kode_produk string `validate:"required"`
	Nama_produk string `validate:"required"`
	Merk        string `validate:"required"`
	Harga_beli  int    `validate:"required"`
	Harga_jual  int    `validate:"required"`
	Stok        int    `validate:"required"`
}

type ProductUpdateRequest struct {
	Id          int    `validate:"required"`
	Kode_produk string `validate:"required"`
	Nama_produk string `validate:"required"`
	Merk        string `validate:"required"`
	Harga_beli  int    `validate:"required"`
	Harga_jual  int    `validate:"required"`
	Stok        int    `validate:"required"`
}
