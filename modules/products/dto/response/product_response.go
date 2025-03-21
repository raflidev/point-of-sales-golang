package response

type ProductResponse struct {
	Id          int    `json:"id"`
	Kode_produk string `json:"kode_produk"`
	Nama_produk string `json:"nama_produk"`
	Merk        string `json:"merk"`
	Harga_beli  int    `json:"harga_beli"`
	Harga_jual  int    `json:"harga_jual"`
	Stok        int    `json:"stok"`
}
