CREATE TABLE members
(
  id CHAR(36) NOT NULL,
  kode_member VARCHAR(255) NULL,
  nama VARCHAR(255) NULL,
  alamat Text NULL,
  telepon VARCHAR(255) NULL,
  keterangan VARCHAR(255) NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id)
)