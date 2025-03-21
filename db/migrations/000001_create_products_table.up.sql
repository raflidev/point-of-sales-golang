CREATE TABLE product
(
  id uuid DEFAULT gen_random_uuid(),
  kode_produk VARCHAR(255) NULL,
  nama_produk VARCHAR(255) NULL,
  merk VARCHAR(255) NULL,
  harga_beli VARCHAR(255) NULL,
  diskon VARCHAR(255) NULL,
  harga_jual VARCHAR(255) NULL,
  stok VARCHAR(255) NULL,
  PRIMARY KEY(id) 
)