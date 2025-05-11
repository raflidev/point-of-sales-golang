CREATE TABLE pembelian
(
  id CHAR(36) NOT NULL,
  supplier_id CHAR(36) NOT NULL,
  total_item INT NULL,
  total_harga INT NULL,
  diskon INT NULL,
  bayar INT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id),
  FOREIGN KEY (supplier_id) REFERENCES suppliers(id)
)