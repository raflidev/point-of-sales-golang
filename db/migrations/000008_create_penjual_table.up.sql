CREATE TABLE penjualan
(
  id CHAR(36) NOT NULL,
  member_id CHAR(36) NOT NULL,
  user_id CHAR(36) NOT NULL,
  total_item INT NULL,
  total_harga INT NULL,
  diskon INT NULL,
  bayar INT NULL,
  diterima INT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id),
  FOREIGN KEY (member_id) REFERENCES members(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
)