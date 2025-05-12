CREATE TABLE pembelian_detail
(
  id CHAR(36) NOT NULL,
  pembelian_id CHAR(36) NOT NULL,
  produk_id CHAR(36) NOT NULL,
  harga_beli INT NULL,
  jumlah INT NULL,
  subtotal INT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id),
  FOREIGN KEY (pembelian_id) REFERENCES pembelian(id),
  FOREIGN KEY (produk_id) REFERENCES product(id)
)