CREATE TABLE users
(
  id CHAR(36) NOT NULL,
  nama VARCHAR(255) NULL,
  email VARCHAR(255) NULL,
  foto Text NULL,
  password VARCHAR(255) NULL,
  role VARCHAR(255) NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id)
)