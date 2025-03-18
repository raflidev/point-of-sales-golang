package web

type ProductCreateRequest struct {
	KodeProduk string
	NamaProduk string
	Merk       string
	HargaBeli  int
	HargaJual  int
	Stok       int
}
