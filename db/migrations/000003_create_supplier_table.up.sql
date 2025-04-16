CREATE TABLE suppliers
(
  id uuid DEFAULT gen_random_uuid(),
  nama VARCHAR(255) NULL,
  alamat Text NULL,
  telepon VARCHAR(255) NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id) 
)