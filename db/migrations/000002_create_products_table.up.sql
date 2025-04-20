CREATE TABLE product
(
  id CHAR(36) NOT NULL,
  kategori_id CHAR(36) NOT NULL,
  kode_produk VARCHAR(255) NULL,
  nama_produk VARCHAR(255) NULL,
  merk VARCHAR(255) NULL,
  harga_beli INT NULL,
  diskon INT NULL,
  harga_jual INT NULL,
  stok INT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id),
  FOREIGN KEY (kategori_id) REFERENCES Kategori(id)
)