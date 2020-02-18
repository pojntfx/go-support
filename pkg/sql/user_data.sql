BEGIN TRANSACTION;
INSERT INTO users
VALUES
  (
    NULL,
    "max@mustermann.com",
    "Max",
    "Mustermann",
    "maxm",
    "password"
  );
END TRANSACTION;