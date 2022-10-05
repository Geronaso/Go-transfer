DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS transfers;

CREATE TABLE account (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user TEXT UNIQUE NOT NULL,
  date_joined DATE NOT NULL,
  cpf VARCHAR(11) UNIQUE NOT NULL,
  balance REAL NOT NULL,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE transfers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  account_origin_id TEXT NOT NULL,
  account_destination_id TEXT NOT NULL,
  amount REAL NOT NULL,
  created_at DATE NOT NULL,
  FOREIGN KEY (account_origin_id) REFERENCES account (id),
  FOREIGN KEY (account_destination_id) REFERENCES account (id)
);