CREATE TABLE kategori
(
  id CHAR(36) NOT NULL,
  nama_kategori TEXT NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id)
)

