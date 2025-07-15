CREATE TABLE tickets (
  id TEXT PRIMARY KEY,
  type TEXT,
  price REAL,
  quantity INTEGER
);

CREATE TABLE sales (
  id TEXT PRIMARY KEY,
  ticket_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  payment_type TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
