CREATE TABLE productLogs
(
  id uuid DEFAULT gen_random_uuid(),
  product_id gen_random_uuid() NULL,
  action VARCHAR(255) NULL,
  create_at TIMESTAMP DEFAULT now(),
  update_at TIMESTAMP DEFAULT now(),
  PRIMARY KEY(id) 
)