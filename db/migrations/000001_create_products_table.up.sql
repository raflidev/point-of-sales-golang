CREATE TABLE product
(
  id uuid DEFAULT gen_random_uuid(),
  kode_produk VARCHAR(255) NULL,
  nama_produk VARCHAR(255) NULL,
  merk VARCHAR(255) NULL,
  harga_beli INT NULL,
  diskon INT NULL,
  harga_jual INT NULL,
  stok INT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id) 
)