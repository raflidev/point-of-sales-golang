CREATE TABLE product
(
  id INT NOT NULL,
  kode_produk VARCHAR(255) NOT NULL,
  nama_produk VARCHAR(255) NOT NULL,
  merk VARCHAR(255) NOT NULL,
  harga_beli VARCHAR(255) NOT NULL,
  diskon VARCHAR(255) NOT NULL,
  harga_jual VARCHAR(255) NOT NULL,
  stok VARCHAR(255) NOT NULL,
  PRIMARY KEY(id) 
)