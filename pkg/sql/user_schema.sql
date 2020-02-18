BEGIN TRANSACTION;
CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  email varchar(255),
  firstname varchar(255),
  secondname varchar(255),
  username varchar(255),
  password varchar(255)
);
END TRANSACTION;